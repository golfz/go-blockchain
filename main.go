package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Coordinate struct {
	x int
	y int
}

type Block struct {
	Index          int
	Timestamp      string
	MoveCoordinate Coordinate
	Hash           string
	PrevHash       string
}

var Blockchain []Block

func calculateHash(block Block) string {
	record := string(block.Index) +
		block.Timestamp +
		string(block.MoveCoordinate.x) +
		string(block.MoveCoordinate.y) +
		block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, moveCoordinate Coordinate) (Block, error) {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.MoveCoordinate = moveCoordinate
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

func main() {

}
