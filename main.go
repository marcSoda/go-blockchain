package main

import (
	"fmt"
	"strconv"
	"blockchain"
)

func main() {
	chain := Init()
	chain.InsertBlock(&InsuranceSummary{"Jan", "Bob was a safe driver", "0"})
	chain.InsertBlock(&InsuranceSummary{"Feb", "Bob was a safe driver", "0"})
	chain.InsertBlock(&InsuranceSummary{"Mar", "Bob was a safe dirver", "0"})
	chain.InsertBlock(&InsuranceSummary{"Apr", "Bob was a safe driver", "0"})
	chain.InsertBlock(&InsuranceSummary{"May", "Bob was a safe driver", "0"})

	for i, block := range chain.Blocks {
	    fmt.Printf("Block: %d\n", i)
	    fmt.Printf("Last Hash: %x\n", block.LastHash)
	    fmt.Printf("Data: | %s | %s | %s |\n", block.Data.month, block.Data.summary, block.Data.premiumIncrease)
	    fmt.Printf("This Hash: %x\n\n", block.ThisHash)

	    pow := blockchain.NewProof(block)
	    fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	    fmt.Println()
	}
}
