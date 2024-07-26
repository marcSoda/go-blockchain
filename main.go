package main

import (
	"fmt"
	"strconv"
	"github.com/marcSoda/go-blockchain/blockchain"
)

func main() {
	chain := blockchain.Init()
	chain.InsertBlock(&blockchain.InsuranceSummary{"Jan", "Bob was a safe driver", "0"})
	chain.InsertBlock(&blockchain.InsuranceSummary{"Feb", "Bob was a safe driver", "0"})
	chain.InsertBlock(&blockchain.InsuranceSummary{"Mar", "Bob was a safe dirver", "0"})
	chain.InsertBlock(&blockchain.InsuranceSummary{"Apr", "Bob was a safe driver", "0"})
	chain.InsertBlock(&blockchain.InsuranceSummary{"May", "Bob was a safe driver", "0"})

	for i, block := range chain.Blocks {
	    fmt.Printf("Block: %d\n", i)
	    fmt.Printf("Last Hash: %x\n", block.LastHash)
	    fmt.Printf("Data: | %s | %s | %s |\n", block.Data.Month, block.Data.Summary, block.Data.PremiumIncrease)
	    fmt.Printf("This Hash: %x\n", block.ThisHash)
	    pow := blockchain.NewPOW(block)
	    fmt.Printf("POW Valid: %s\n", strconv.FormatBool(pow.Validate()))
	}

	fmt.Println("\nBinary representation of hash of second block:")
	for _, n := range(chain.Blocks[1].ThisHash) {
	    fmt.Printf("%08b", n)
	}
	fmt.Println()
}
