package core

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index         int64
	Timestamp     int64
	PrevBlockHash string
	Hash          string
	Data          string
}

func (this Block) Print() {
	println("index:", this.Index)
	println("time:", this.Timestamp)
	println("preHash:", this.PrevBlockHash)
	println("hash:", this.Hash)
	println("data:", this.Data)
	println("--------------------------------------")
}

func GenerateBlock(prevBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.PrevBlockHash = prevBlock.Hash
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

func GenerateGenesisBlock() Block {
	prevBlock := Block{Index: -1, Hash: ""}
	return GenerateBlock(prevBlock, "this is genesis block")
}

func calculateHash(block Block) string {
	data := strconv.FormatInt(block.Index, 10) + strconv.FormatInt(block.Timestamp, 10) + block.PrevBlockHash + block.Data
	hashData := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hashData[:])
}

func (this Block) IsValid(prevBlock Block) bool {
	if (this.Index-1 != prevBlock.Index) {
		return false
	}

	if (this.PrevBlockHash != prevBlock.Hash) {
		return false
	}

	if (this.Hash != calculateHash(this)) {
		return false
	}
	return true
}
