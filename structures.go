package main

type Block struct {
	Timestamp 	 			int64  // when the block was created
	PreviousBlockHash []byte // hash of the previous block
	MyBlockHash 			[]byte // hash of the block
	AllData 			  	[]byte // data of the block
}

type Blockchain struct {
	Blocks []*Block
}
