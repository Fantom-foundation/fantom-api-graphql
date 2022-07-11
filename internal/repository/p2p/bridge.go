// Package p2p implements basic P2P communication module to aid connecting Opera nodes
// and extracting basic status information from them.
package p2p

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/rlpx"
	"github.com/ethereum/go-ethereum/rlp"
	"net"
	"strconv"
	"time"
)

// peerConnectionTimeout is the timeout enforced on a new connection to remote p2p peer.
const peerConnectionTimeout = 5 * time.Second

const (
	// p2p chat stages used to communicate with a peer and to get the info we need
	chatStageHandshake = iota
	chatStageSendHello
	chatStageReceiveInfo
	chatStageGoodbye
	chatStageDone
)

// cfg contains instance of the configuration to be used by the p2p module.
var cfg *config.Config

// log contains instance of the app logger.
var log logger.Logger

// SetConfig configures default configuration.
func SetConfig(c *config.Config) {
	cfg = c
}

// SetLogger configures default logger to be used by the module.
func SetLogger(l logger.Logger) {
	log = l
}

// PeerInformation returns detailed information of the given peer, if it can be obtained.
func PeerInformation(node *enode.Node) (*types.OperaNodeInformation, error) {
	// make sure we have a way to sign
	if cfg == nil || cfg.Signature.PrivateKey == nil {
		return nil, fmt.Errorf("p2p key configuration is missing")
	}

	// establish TCP connection to the remote peer
	addr := fmt.Sprintf("%s:%d", node.IP().String(), node.TCP())
	log.Debugf("p2p connecting to %s", addr)

	c, err := net.DialTimeout("tcp", addr, peerConnectionTimeout)
	if err != nil {
		log.Warningf("no connection to peer %s; %s", node.URLv4(), err.Error())
		return nil, err
	}

	con := rlpx.NewConn(c, &cfg.Signature.PrivateKey.PublicKey)
	defer func() {
		if e := con.Close(); e != nil {
			log.Warningf("p2p could not close connection to %s; %s", addr, e.Error())
		}
	}()

	log.Debug("p2p connected")
	return chat(con)
}

// chat with the connected peer to get the node information we need.
func chat(con *rlpx.Conn) (*types.OperaNodeInformation, error) {
	var info types.OperaNodeInformation
	var stage int
	var err error

	for stage < chatStageDone {
		switch stage {
		case chatStageHandshake:
			_, err = con.Handshake(cfg.Signature.PrivateKey)
			if err != nil {
				log.Warningf("p2p handshake failed; %s", err.Error())
				stage = chatStageDone
				continue
			}

			stage = chatStageSendHello

		case chatStageSendHello:
			_, err = con.Write(msgTypeHello, myHello())
			if err != nil {
				log.Warningf("p2p hello failed; %s", err.Error())
				stage = chatStageDone
				continue
			}

			stage = chatStageReceiveInfo

		case chatStageReceiveInfo:
			// extract the actual peer information, if they will not reject us
			stage, err = readNextInfo(con, &info)

		case chatStageGoodbye:
			// an error here does not need to propagate; we are just saying goodbye
			_, e := con.Write(msgTypeDisconnect, myDisconnect())
			if e != nil {
				log.Warningf("p2p graceful disconnect failed; %s", err.Error())
			}
			stage = chatStageDone
		}
	}

	return &info, err
}

// readNextInfo reads next message and updates peer information with the data received.
// The call returns the next chat state to be executed.
func readNextInfo(con *rlpx.Conn, info *types.OperaNodeInformation) (int, error) {
	mt, msg, err := receive(con)
	if err != nil {
		log.Warningf("p2p receiver failed; %s", err.Error())
		return chatStageGoodbye, err
	}

	// update info
	switch mt {
	case msgTypeDisconnect:
		log.Noticef("peer sent disconnect, %s", msg.(*msgDisconnect).Reason.String())
		return chatStageDone, nil

	case msgTypeHandshake:
		log.Noticef("peer network is #%d, with genesis %s", msg.(*msgHandshake).NetworkID, msg.(*msgHandshake).Genesis.String())

	case msgTypeHello:
		info.Name = msg.(*msgHello).Name
		info.Version = strconv.FormatUint(msg.(*msgHello).Version, 16)
		log.Noticef("peer is %s, version %s", info.Name, info.Version)

	case msgTypeProgress:
		info.Epoch = int64(msg.(*msgPeerProgress).Epoch)
		info.BlockHeight = int64(msg.(*msgPeerProgress).LastBlock)
		log.Noticef("peer epoch is #%d, block #%d", info.Epoch, info.BlockHeight)

		// we hang up the connection with an excuse after the progress
		return chatStageGoodbye, nil
	}

	return chatStageReceiveInfo, nil
}

// receive a message from the connected remote party.
func receive(con *rlpx.Conn) (mt uint64, target interface{}, err error) {
	// get a message from the connection
	var data []byte
	mt, data, _, err = con.Read()
	if err != nil {
		return 0, nil, err
	}

	log.Infof("p2p received #%d", mt)

	// prep the target interface based on message type
	switch mt {
	case msgTypeHandshake:
		var hs msgHandshake
		target = &hs
	case msgTypeHello:
		var he msgHello
		target = &he
	case msgTypeProgress:
		var pp msgPeerProgress
		target = &pp
	case msgTypeDisconnect:
		var di msgDisconnect
		target = &di
	}

	// decode received data block to the target structure
	// @todo why do we skip 2 bytes here?
	if err := rlp.DecodeBytes(data[2:], target); err != nil {
		return 0, nil, err
	}
	return mt, target, nil
}
