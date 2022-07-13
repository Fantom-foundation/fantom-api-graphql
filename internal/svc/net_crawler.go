// Package svc implements blockchain data processing services.
package svc

import (
	"bytes"
	"fantom-api-graphql/internal/repository/db"
	"fmt"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"net"
	"time"
)

// discoveryNodeEmptyID represents an undefined ID of a discovery node.
var discoveryNodeEmptyID = [32]byte{}

// netCrawler implements Opera network discovery service used to detect and maintain
// a list of running network nodes to allow network health and layout analysis
type netCrawler struct {
	client *discover.UDPv5
	service
}

// name returns the name of the service used by orchestrator.
func (nc *netCrawler) name() string {
	return "discovery crawler"
}

// init prepares the discovery analyzer.
func (nc *netCrawler) init() {
	nc.sigStop = make(chan struct{})

	// initialise discovery protocol
	var err error
	nc.client, err = newDiscovery()
	if err != nil {
		log.Criticalf("network discovery is not available; %s", err.Error())
		return
	}

	log.Notice("discovery v5 started")
}

// run starts the network discovery scanner
func (nc *netCrawler) run() {
	// make sure we are orchestrated
	if nc.mgr == nil {
		panic(fmt.Errorf("no svc manager set on %s", nc.name()))
	}

	// if we can not communicate, terminate the crawler right away
	if nc.client == nil {
		log.Warningf("client not open, network discovery disabled")
		return
	}

	// signal orchestrator we started and go
	nc.mgr.started(nc)
	go nc.execute()
}

// run starts the network discovery scanner
func (nc *netCrawler) execute() {
	// keep track of active iterators and prep their signaling channels
	iterators := make(map[enode.Iterator]string)
	iteratorDone := make(chan enode.Iterator, 5)
	nodeFeed := make(chan *enode.Node)

	defer func() {
		// close active iterators and clean-up their closing signals to make sure they are done
		log.Infof("closing iterators")
		for i := range iterators {
			i.Close()
			<-iteratorDone
		}

		// the crawler done
		close(nodeFeed)
		nc.mgr.finished(nc)
	}()

	// run the iterator handler
	nc.handle(iterators, nodeFeed, iteratorDone)
}

// handle running iterators collecting and/or updating network nodes.
func (nc *netCrawler) handle(iterators map[enode.Iterator]string, inNode chan *enode.Node, inDone chan enode.Iterator) {
	// monitor update batch status periodically
	updateTick := time.NewTicker(30 * time.Second)
	defer updateTick.Stop()

	// start initial batches
	nc.scheduleDefaultBatch(iterators, inNode, inDone)
	updateIterator := nc.scheduleUpdateBatch(iterators, inNode, inDone)

	for {
		select {
		case <-nc.sigStop:
			return

		case done := <-inDone:
			name, ok := iterators[done]
			if ok {
				delete(iterators, done)
				log.Infof("discovery iterator %s done", name)
			}

		case node, ok := <-inNode:
			if !ok {
				return
			}
			nc.update(node)

		case <-updateTick.C:
			_, ok := iterators[updateIterator]
			if !ok {
				updateIterator = nc.scheduleUpdateBatch(iterators, inNode, inDone)
			}
		}
	}
}

// scheduleDefaultBatch schedules the default discovery iterator
// this is the one providing newly discovered nodes from the interaction with known nodes
func (nc *netCrawler) scheduleDefaultBatch(iterators map[enode.Iterator]string, inNode chan *enode.Node, inDone chan enode.Iterator) {
	di := nc.client.RandomNodes()
	iterators[di] = "default"
	go nc.traverse(di, inNode, inDone)
}

// scheduleUpdateBatch runs a new update batch iterator through the handler.
func (nc *netCrawler) scheduleUpdateBatch(iterators map[enode.Iterator]string, inNode chan *enode.Node, inDone chan enode.Iterator) enode.Iterator {
	// load update batch from repository; make sure we have some
	nodes, err := repo.NetworkNodeUpdateBatch()
	if err != nil {
		log.Criticalf("can not pull network nodes update batch; %s", err.Error())
		nodes = make([]*enode.Node, 0)
	}

	// make the iterator and run it
	updateIterator := enode.IterNodes(nodes)
	iterators[updateIterator] = fmt.Sprintf("batch %s (%d nodes)", time.Now().Format("15:04:05"), len(nodes))

	go nc.traverse(updateIterator, inNode, inDone)
	return updateIterator
}

// traverse given node iterator extracting nodes in it and sending them into the feed channel for processing.
func (nc *netCrawler) traverse(iter enode.Iterator, nodeFeed chan *enode.Node, done chan enode.Iterator) {
	defer func() { done <- iter }()

	for iter.Next() {
		select {
		case nodeFeed <- iter.Node():
		case <-nc.sigStop:
			return
		}
	}
}

// update the given Node information, if possible.
func (nc *netCrawler) update(node *enode.Node) {
	defer func() {
		if re := recover(); re != nil {
			log.Criticalf("network discovery panicked at %s, %s", node.ID(), node.URLv4())
		}
	}()

	// is the node ID valid?
	if bytes.Compare(node.ID().Bytes(), discoveryNodeEmptyID[:]) == 0 {
		log.Errorf("empty discovery node ID skipped")
		return
	}

	// request node ENR record
	log.Debugf("checking peer %s at %s", node.ID(), node.URLv4())
	enr, err := nc.client.RequestENR(node)
	if err != nil {
		nc.fail(node)
		return
	}

	nc.confirm(enr)
}

// confirm the given node was contacted, is live and available.
func (nc *netCrawler) confirm(node *enode.Node) {
	log.Debugf("node %s confirmed at %s", node.IP().String(), node.URLv4())

	_, err := repo.NetworkNodeConfirmCheck(node, nc.mgr)
	if err != nil {
		log.Errorf("could not confirm node %s; %s", node.ID().String(), err.Error())
	}
}

// fail the given node check - e.g. register a failure in verifying node status.
func (nc *netCrawler) fail(node *enode.Node) {
	log.Debugf("node %s failed at %s [#%s]", node.IP().String(), node.URLv4(), node.ID())

	err := repo.NetworkNodeFailCheck(node)
	if err != nil && err != db.ErrUnknownNetworkNode {
		log.Errorf("could not store node %s check failure; %s", node.ID().String(), err.Error())
	}
}

// newDiscovery creates a new configured local instance of the discovery v5 protocol client.
func newDiscovery() (*discover.UDPv5, error) {
	// combine pre-configured list of nodes taken from Opera
	// with a random set of servers we already discovered and verified to get the bootstrap set
	boot := repo.NetworkNodeBootstrapSet()
	boot = append(boot, cfg.OperaNetwork.BootstrapNodes...)

	// prepare configuration using global config
	conf := discover.Config{
		PrivateKey: cfg.Signature.PrivateKey,
		Bootnodes:  boot,
	}

	// make in-memory key-value database; it's used by the client to track neighbors' buckets
	// it always starts with an empty database and the client builds the network topology schema from scratch
	// we provide a random set of servers collected before along with preconfigured bootstraps
	// to kick off the initial search and discovery
	buckets, err := enode.OpenDB("")
	if err != nil {
		log.Criticalf("can not initialize discovery database; %s", err.Error())
		return nil, err
	}

	// construct local node; it's going to be a full-featured participant of the discovery protocol
	ln := enode.NewLocalNode(buckets, conf.PrivateKey)

	// get UDP socket listener
	socket, err := attachUDPSocket(ln)
	if err != nil {
		log.Criticalf("can not connect discovery to network; %s", err.Error())
	}

	return discover.ListenV5(socket, ln, conf)
}

// attachUDPSocket creates UDP socket used by the discovery client
// to communicate with other nodes over the network
func attachUDPSocket(ln *enode.LocalNode) (*net.UDPConn, error) {
	// we try to attach to all network interfaces;
	// port is not defined and is assigned automatically by the networking layer
	addr := cfg.OperaNetwork.DiscoveryUDP

	socket, err := net.ListenPacket("udp4", addr)
	if err != nil {
		log.Criticalf("could not attach UDP listener; %s", err.Error())
		return nil, err
	}

	// do we have an IP address? use localhost as a fallback if we don't
	ipAddress := socket.LocalAddr().(*net.UDPAddr).IP
	if ipAddress.IsUnspecified() {
		ipAddress = net.IP{127, 0, 0, 1}
	}

	// update local node networking fallback
	ln.SetFallbackIP(ipAddress)
	ln.SetFallbackUDP(socket.LocalAddr().(*net.UDPAddr).Port)

	return socket.(*net.UDPConn), nil
}
