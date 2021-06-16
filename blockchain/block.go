package blockchain

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
	month           string
	summary         string
	premiumIncrease string
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.LastHash,
				    []byte(b.Data.month),
				    []byte(b.Data.summary),
				    []byte(b.Data.premiumIncrease)},
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
