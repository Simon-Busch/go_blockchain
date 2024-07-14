package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Now let's create a method for generating a hash of the block
// We will just concatenate all the data and hash it to obtain the block hash
func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))                                  // get the time and convert it into a unique series of digits
	headers := bytes.Join([][]byte{timestamp, block.PreviousBlockHash, block.AllData}, []byte{}) // concatenate all the block data
	hash := sha256.Sum256(headers)                                                               // hash the whole thing
	block.MyBlockHash = hash[:]                                                                  // now set the hash of the block
}

// Create a function for new block generation and return that block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), prevBlockHash, []byte{}, []byte(data)} // the block is received
	block.SetHash()                                                           // the block is hashed
	return block                                                              // the block is returned with all the information in it
}

/* let's now create the genesis block function that will return the first block. The genesis block is the first block on the chain */
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{}) // the genesis block is made with some data in it
}
