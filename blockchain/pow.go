package blockchain

import (
	"fmt"
	"math"
	"math/big"
)

const Difficulty = 16

type POW struct {
	Block *Block
	Target *big.Int
}

func NewPOW(block *Block) *POW {
	targ := big.NewInt(1)
	targ.Lsh(targ, uint(256 - Difficulty))
	return &POW{block, targ}
}

func (pow *POW) Generate() []byte {
	var intHash big.Int
	var hash [32]byte
	pow.Block.Nonce = 0
	for pow.Block.Nonce < math.MaxInt64 {
		hash = pow.Block.DeriveHash(Difficulty)
		fmt.Printf("\r| Nonce: %d: | Hash: %x |", pow.Block.Nonce, hash)
		intHash.SetBytes(hash[:])

		if (intHash.Cmp(pow.Target) == -1) {
			break
		} else {
			pow.Block.Nonce++
		}
	}
	fmt.Println()
	return hash[:]
}

func (pow *POW) Validate() bool {
	var intHash big.Int
	hash := pow.Block.DeriveHash(Difficulty)
	intHash.SetBytes(hash[:])
	return intHash.Cmp(pow.Target) == -1
}
