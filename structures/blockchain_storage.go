package structures

import (
	"encoding/json"
	"github.com/Suad0/blockchainGO/core"
	"log"
	"os"
)

// SaveBlockchain saves the blockchain to a file
func SaveBlockchain(bc *core.Blockchain, filename string) error {
	data, err := json.MarshalIndent(bc, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// LoadBlockchain loads the blockchain from a file, or creates a new one if file doesn't exist
func LoadBlockchain(filename string) *core.Blockchain {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Println("Blockchain file not found, creating a new one...")
		return core.NewBlockchain() // Create a new blockchain with genesis block if no file exists
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error reading blockchain file:", err)
	}

	var bc core.Blockchain
	err = json.Unmarshal(data, &bc)
	if err != nil {
		log.Fatal("Error unmarshaling blockchain:", err)
	}

	return &bc
}
