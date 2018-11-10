package main

import "time"

type Block struct {
	Version       int64
	PrevBlockHash []byte
	MerkelRoot    []byte
	TimeStamp     int64
	Bits          int64
	Nonce         int64
	Data          []byte
}

func NewBlock(data string, preBlockHash []byte) *Block {
	var block Block

	block = Block{
		Version:       1,
		PrevBlockHash: preBlockHash,
		MerkelRoot:    []byte{},
		TimeStamp:     time.Now().Unix(),
		Bits:          1,
		Nonce:         1,
		Data:          []byte(data),
	}
	block.SetHash()

	return &block
}

func (block *Block) SetHash() {

}
