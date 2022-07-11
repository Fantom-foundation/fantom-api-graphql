// Package p2p implements basic P2P communication module to aid connecting Opera nodes
// and extracting basic status information from them.
package p2p

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	ep2p "github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rlp"
)

const (
	// message protocol params and message types
	baseProtocolLength uint64 = 16
	msgTypeHello       uint64 = 0x0
	msgTypeDisconnect  uint64 = 0x1
	msgTypeHandshake          = baseProtocolLength + 0
	msgTypeProgress           = baseProtocolLength + 1
)

// msgHandshake represents a message used for network negotiation.
type msgHandshake struct {
	ProtocolVersion uint32
	NetworkID       uint64
	Genesis         common.Hash
}

// msgHello represents a message sent to a connected peer to start communication.
type msgHello struct {
	Version    uint64
	Name       string
	Caps       []ep2p.Cap
	ListenPort uint64
	ID         []byte // secp256k1 public key

	// Ignore additional fields (for forward compatibility).
	Rest []rlp.RawValue `rlp:"tail"`
}

// msgPeerProgress contains structures identifying a remote peer state progress.
type msgPeerProgress struct {
	Epoch            uint32
	LastBlock        uint64
	LastBlockAtropos common.Hash
	HighestLamport   uint32 // unused
}

// msgDisconnect represents a message indicating a graceful connection termination.
type msgDisconnect struct {
	Reason ep2p.DiscReason
}

// defaultHelloMsg is a default RLP encoded instance of p2p hello message we send to contacted peers.
var defaultHelloMsg []byte

// defaultDisconnectMsg is the message we send to indicate graceful disconnection to remote peer.
var defaultDisconnectMsg []byte

// myHello provides an encoded p2p hello packet
// to be sent to a contacted peer as our opening message.
func myHello() []byte {
	// make the message on the first use
	if defaultHelloMsg == nil {
		h := msgHello{
			Version: 5,
			Name:    cfg.AppName,
			Caps: []ep2p.Cap{
				{Name: "opera", Version: 62},
				{Name: "opera", Version: 63},
			},
			ID: crypto.FromECDSAPub(&cfg.Signature.PrivateKey.PublicKey)[1:],
		}

		var err error
		defaultHelloMsg, err = rlp.EncodeToBytes(&h)
		if err != nil {
			log.Criticalf("could not encode p2p hello message", err.Error())
			return nil
		}
	}
	return defaultHelloMsg
}

// myDisconnect provides and rlp encoded message for graceful disconnection from a remote peer.
func myDisconnect() []byte {
	if defaultDisconnectMsg == nil {
		// come up with a reason why we terminate the connection early
		d := msgDisconnect{Reason: ep2p.DiscRequested}

		var err error
		defaultDisconnectMsg, err = rlp.EncodeToBytes(&d)
		if err != nil {
			log.Criticalf("could not encode p2p disconnect message", err.Error())
			return nil
		}
	}

	return defaultDisconnectMsg
}
