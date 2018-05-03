// This is an example of creating a blockchain and blocks
package main

import (
	"fmt"
	"math/big"
)

// Main function will be called first
func main() {
	// Create a new blockchain
	bc := NewBlockchain()

  // Add blocks with arbitrary data
	bc.AddBlock("Send 1 BTC to Dr. Lam")
	bc.AddBlock("Send 2 more BTC to Dr. Minh")

  // Print a block
	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
