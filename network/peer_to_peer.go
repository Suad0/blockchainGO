package network

import (
	"encoding/json"
	"fmt"
	"github.com/Suad0/blockchainGO/core"
	"log"
	"net"
)

type Node struct {
	Address    string
	Peers      []string
	Blockchain *core.Blockchain
}

// StartNode starts a new node and listens for incoming connections
func (n *Node) StartNode() {
	listener, err := net.Listen("tcp", n.Address)
	if err != nil {
		log.Fatal("Error starting node:", err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {

		}
	}(listener)

	log.Println("Node started at", n.Address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go n.handleConnection(conn)
	}
}

// handleConnection handles incoming requests from peers
func (n *Node) handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read incoming message
	var request string
	_, err := fmt.Fscan(conn, &request)
	if err != nil {
		log.Println("Error reading request:", err)
		return
	}

	switch request {
	case "GET_BLOCKCHAIN":
		n.sendBlockchain(conn)
	case "NEW_BLOCK":
		n.receiveNewBlock(conn)
	default:
		log.Println("Unknown request:", request)
	}
}

// sendBlockchain sends the blockchain to the requesting node
func (n *Node) sendBlockchain(conn net.Conn) {
	blockchainData, err := json.Marshal(n.Blockchain)
	if err != nil {
		log.Println("Error marshaling blockchain:", err)
		return
	}
	_, err = conn.Write(blockchainData)
	if err != nil {
		return
	}
}

// receiveNewBlock receives a new block from a peer and adds it to the blockchain
func (n *Node) receiveNewBlock(conn net.Conn) {
	var newBlock core.Block
	decoder := json.NewDecoder(conn)
	err := decoder.Decode(&newBlock)
	if err != nil {
		log.Println("Error decoding block:", err)
		return
	}
	n.Blockchain.AddBlock(newBlock.Transactions)
	log.Println("New block added from peer")
}

// ConnectToPeer connects to another node and requests the blockchain
func (n *Node) ConnectToPeer(peerAddress string) {
	conn, err := net.Dial("tcp", peerAddress)
	if err != nil {
		log.Println("Error connecting to peer:", err)
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// Request the blockchain from the peer
	_, err = fmt.Fprintln(conn, "GET_BLOCKCHAIN")
	if err != nil {
		return
	}

	// Read the peer's blockchain
	var peerBlockchain core.Blockchain
	decoder := json.NewDecoder(conn)
	err = decoder.Decode(&peerBlockchain)
	if err != nil {
		log.Println("Error decoding blockchain from peer:", err)
		return
	}

	// Replace the local blockchain if the peer's is longer
	if len(peerBlockchain.Blocks) > len(n.Blockchain.Blocks) {
		n.Blockchain = &peerBlockchain
		log.Println("Replaced local blockchain with peer's blockchain")
	}
}
