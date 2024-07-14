package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))                                  // get the time and convert it into a unique series of digits
	headers := bytes.Join([][]byte{timestamp, block.PreviousBlockHash, block.AllData}, []byte{}) // concatenate all the block data
	hash := sha256.Sum256(headers)                                                               // hash the whole thing
	block.MyBlockHash = hash[:]                                                                  // now set the hash of the block
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), prevBlockHash, []byte{}, []byte(data)} // the block is received
	block.SetHash()                                                           // the block is hashed
	return block                                                              // the block is returned with all the information in it
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{}) // the genesis block is made with some data in it
}
