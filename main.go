package main

import (
	"github.com/Suad0/blockchainGO/network"
	"github.com/Suad0/blockchainGO/structures"
)

const blockchainFile = "blockchain.json"

func main() {

	/*

		blockchain := structures.LoadBlockchain(blockchainFile)

		// Add a new block
		blockchain.AddBlock([]core.Transaction{
			core.NewTransaction("Alice", "Bob", 10),
			core.NewTransaction("Bob", "Charlie", 5),
		})

		// Save the blockchain to file
		err := structures.SaveBlockchain(blockchain, blockchainFile)
		if err != nil {
			log.Fatal("Error saving blockchain:", err)
		}

		// Print blockchain
		for _, block := range blockchain.Blocks {
			fmt.Printf("Prev Hash: %s\n", block.PrevHash)
			fmt.Printf("Hash: %s\n", block.Hash)
			fmt.Println()
		}


	*/

	// Load the blockchain from file
	blockchain := structures.LoadBlockchain(blockchainFile)

	// Create a new node
	node := network.Node{
		Address:    "localhost:3000", // Local node's address
		Blockchain: blockchain,
	}

	// Start the node and listen for incoming connections
	go node.StartNode()

	// Simulate connecting to a peer (optional)
	node.ConnectToPeer("localhost:3001")

	// Save the blockchain periodically or after every block addition
	err := structures.SaveBlockchain(node.Blockchain, blockchainFile)
	if err != nil {
		return
	}

	// Keep the main function running
	select {}

}
