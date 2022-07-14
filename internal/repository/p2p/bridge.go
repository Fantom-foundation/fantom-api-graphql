// Package p2p implements basic P2P communication module to aid connecting Opera nodes
// and extracting basic status information from them.
package p2p

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/rlpx"
	"github.com/ethereum/go-ethereum/rlp"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	// peerConnectionTimeout is the timeout enforced on a new connection to remote p2p peer.
	peerConnectionTimeout = 2 * time.Second

	// peerChatDeadline specifies the max aount of time we are willing to chat with a remote peer.
	peerChatDeadline = 5 * time.Second
)

const (
	// p2p chat stages used to communicate with a peer and to get the info we need
	chatStageHandshake = iota
	chatStageSendHello
	chatStageReceiveInfo
	chatStageGoodbye
	chatStageDone

	// syncedNodeFlagBlockDiff is the max number of blocks below current known head we tolerate
	// to consider peer to be synced for the purpose of this detailed peer node check.
	syncedNodeFlagBlockDiff = 25
)

// BlockHeightProvider is an interface capable of providing the current known block height.
type BlockHeightProvider interface {
	BlockHeight() uint64
}

// cfg contains instance of the configuration to be used by the p2p module.
var cfg *config.Config

// log contains instance of the app logger.
var log logger.Logger

// ErrNonOperaPeer defines error sent if non-opera node has been found.
var ErrNonOperaPeer = fmt.Errorf("useless non-opera node found")

// SetConfig configures default configuration.
func SetConfig(c *config.Config) {
	cfg = c
}

// SetLogger configures default logger to be used by the module.
func SetLogger(l logger.Logger) {
	log = l
}

// PeerInformation returns detailed information of the given peer, if it can be obtained.
func PeerInformation(node *enode.Node, bhp BlockHeightProvider) (*types.OperaNodeInformation, error) {
	// make sure we have a way to sign
	if cfg == nil || cfg.Signature.PrivateKey == nil {
		return nil, fmt.Errorf("p2p key configuration is missing")
	}

	// establish TCP connection to the remote peer
	addr := fmt.Sprintf("%s:%d", node.IP().String(), node.TCP())
	log.Debugf("p2p connecting to %s; signing as %s", addr, crypto.PubkeyToAddress(cfg.Signature.PrivateKey.PublicKey).String())

	c, err := net.DialTimeout("tcp", addr, peerConnectionTimeout)
	if err != nil {
		log.Debugf("no connection to peer %s; %s", node.URLv4(), err.Error())
		return nil, err
	}

	con := rlpx.NewConn(c, node.Pubkey())
	defer func() {
		if e := con.Close(); e != nil {
			log.Debugf("p2p could not close connection to %s; %s", addr, e.Error())
		}
	}()

	// make sure the com is finished in the specified time
	if err := con.SetDeadline(time.Now().Add(peerChatDeadline)); err != nil {
		log.Errorf("can not set p2p deadline; %s", err.Error())
	}

	return chat(con, bhp)
}

// chat with the connected peer to get the node information we need.
func chat(con *rlpx.Conn, bhp BlockHeightProvider) (*types.OperaNodeInformation, error) {
	var stage int
	var err error

	// prep the info container
	var info = types.OperaNodeInformation{
		Updated: time.Now().UTC(),
	}

	for stage < chatStageDone {
		switch stage {
		case chatStageHandshake:
			_, err = con.Handshake(cfg.Signature.PrivateKey)
			if err != nil {
				log.Debugf("p2p handshake failed; %s", err.Error())
				stage = chatStageDone
				continue
			}

			stage = chatStageSendHello

		case chatStageSendHello:
			_, err = con.Write(msgTypeHello, myHello())
			if err != nil {
				log.Debugf("p2p hello failed; %s", err.Error())
				stage = chatStageDone
				continue
			}

			stage = chatStageReceiveInfo

		case chatStageReceiveInfo:
			// extract the peer information from received messages coming from them
			stage, err = readNext(con, &info, bhp)

		case chatStageGoodbye:
			// an error here does not need to propagate; we are just saying goodbye
			_, e := con.Write(msgTypeDisconnect, myDisconnect())
			if e != nil {
				log.Debugf("p2p graceful disconnect failed; %s", err.Error())
			}
			stage = chatStageDone
		}
	}

	return &info, err
}

// readNext reads next message and updates peer information with the data received.
// The call returns the next chat state to be executed.
func readNext(con *rlpx.Conn, info *types.OperaNodeInformation, bhp BlockHeightProvider) (int, error) {
	mt, msg, err := receive(con)
	if err != nil {
		log.Warningf("p2p receiver failed; %s", err.Error())
		return chatStageGoodbye, err
	}

	// update info
	switch mt {
	case msgTypeDisconnect:
		log.Debugf("peer sent disconnect, %s", msg.(*msgDisconnect).Reason.String())
		if msg.(*msgDisconnect).Reason == p2p.DiscUselessPeer {
			return chatStageDone, ErrNonOperaPeer
		}
		return chatStageDone, nil

	case msgTypeHandshake:
		log.Debugf("peer network is #%d, with genesis %s", msg.(*msgHandshake).NetworkID, msg.(*msgHandshake).Genesis.String())

	case msgTypeHello:
		info.Name = msg.(*msgHello).Name
		info.Version = strconv.FormatUint(msg.(*msgHello).Version, 16)

		// validate protocols selection
		info.IsOperaConfirmed, info.Protocols = hasOperaProtocol(msg.(*msgHello).Caps)
		if !info.IsOperaConfirmed {
			log.Errorf("useless peer is %s, ver.%s with caps [%s]", info.Name, info.Version, info.Protocols)
			return chatStageGoodbye, ErrNonOperaPeer
		}
		log.Infof("p2p peer is %s, ver.%s, [%s]", info.Name, info.Version, info.Protocols)

	case msgTypeProgress:
		bh := int64(bhp.BlockHeight())

		info.Epoch = int64(msg.(*msgPeerProgress).Epoch)
		info.BlockHeight = int64(msg.(*msgPeerProgress).LastBlock)
		info.IsSynced = (bh - info.BlockHeight) <= syncedNodeFlagBlockDiff

		log.Debugf("peer epoch is #%d, block #%d (head is #%d)", info.Epoch, info.BlockHeight, bh)

		// we hang up the connection with an excuse after the progress
		return chatStageGoodbye, nil
	}

	return chatStageReceiveInfo, nil
}

// hasOperaProtocol checks if the set of p2p protocols contains Opera protocol.
func hasOperaProtocol(caps []p2p.Cap) (bool, string) {
	if caps == nil {
		return false, "none/0"
	}

	var sb strings.Builder
	var ok bool

	for i, c := range caps {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(c.String())

		if c.Name == "opera" {
			ok = true
		}
	}

	return ok, sb.String()
}

// receive a message from the connected remote party.
func receive(con *rlpx.Conn) (mt uint64, target interface{}, err error) {
	var rawData []byte
	var data []byte

	// get the next RLP message from the connection
	mt, rawData, _, err = con.Read()
	if err != nil {
		return 0, nil, err
	}

	// prep the target interface based on message type
	switch mt {
	case msgTypeHandshake:
		var hs msgHandshake
		target = &hs
		data = rawData[2:]
	case msgTypeHello:
		var he msgHello
		data = rawData[:]
		target = &he
	case msgTypeProgress:
		var pp msgPeerProgress
		target = &pp
		data = rawData[2:]
	case msgTypeDisconnect:
		var di msgDisconnect
		target = &di
		if len(rawData) > 2 {
			data = rawData[2:]
		}
	}

	// decode received data block to the target structure
	// Check your RLP here, if needed https://toolkit.abdk.consulting/ethereum#rlp
	err = rlp.DecodeBytes(data[:], target)
	if err != nil {
		log.Errorf("p2p data block %s; %s", hexutil.Encode(rawData), err.Error())
		return 0, nil, err
	}

	return mt, target, nil
}
