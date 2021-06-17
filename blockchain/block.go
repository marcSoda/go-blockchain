package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
)

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	ThisHash []byte
	Data     InsuranceSummary
	LastHash []byte
	Nonce    int
}

type InsuranceSummary struct {
	Month           string
	Summary         string
	PremiumIncrease string
}

func (b *Block) DeriveHash(difficulty int) [32]byte {
	data := bytes.Join(
		[][]byte{
			b.LastHash,
			[]byte(b.Data.Month),
			[]byte(b.Data.Summary),
			[]byte(b.Data.PremiumIncrease),
			IntToHex(int64(b.Nonce)),
			IntToHex(int64(difficulty)),
		},
		[]byte{},)
	hash := sha256.Sum256(data)
	return hash
}

func MakeBlock(data InsuranceSummary, lastHash []byte) *Block {
	block := &Block{[]byte{}, data, lastHash, 0}
	pow := NewPOW(block)
	hash := pow.Generate()
	block.ThisHash = hash[:]
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

func IntToHex(n int64) []byte {
	buff := new(bytes.Buffer)
	binary.Write(buff, binary.BigEndian, n)
	return buff.Bytes()
}
