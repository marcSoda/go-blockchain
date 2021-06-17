package main

import (
	"crypto/sha256"
	"bytes"
	"fmt"
)

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	ThisHash []byte
	Data     InsuranceSummary
	LastHash []byte
}

type InsuranceSummary struct {
	Month           string
	Summary         string
	PremiumIncrease string
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.LastHash,
				    []byte(b.Data.Month),
				    []byte(b.Data.Summary),
				    []byte(b.Data.PremiumIncrease)},
				    []byte{})
	hash := sha256.Sum256(info)
	b.ThisHash = hash[:]
}

func MakeBlock(data InsuranceSummary, lastHash []byte) *Block {
	block := &Block{[]byte{}, data, lastHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) InsertBlock(data *InsuranceSummary) {
	lastBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := MakeBlock(*data, lastBlock.ThisHash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

func Init() *BlockChain {
	genesis := MakeBlock(InsuranceSummary{"N/A", "Initial Block", "N/A"}, []byte{})
	return &BlockChain{[]*Block{genesis}}
}

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
	    fmt.Printf("Data: | %s | %s | %s |\n", block.Data.Month, block.Data.Summary, block.Data.PremiumIncrease)
	    fmt.Printf("This Hash: %x\n\n", block.ThisHash)
	}
}
