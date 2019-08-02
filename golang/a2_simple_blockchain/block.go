package main

import (
	"crypto/sha256"
)


// 0.定义一个区块
type Block struct {
	//1.前一个区块链的哈希
	PrevHash []byte
	//2.当前区块的哈希
	Hash []byte
	//3.数据
	Data []byte
}


// 2.创建一个区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevHash: prevBlockHash,
		Hash : []byte{},
		Data : []byte(data),
	}

	block.SetHash()

	return &block
}

// 3.生成哈希
func (block *Block) SetHash() {
	// 1.拼接数据
	blockinfo := append(block.PrevHash, block.Data...)

	// 2.sha256
	hash := sha256.Sum256(blockinfo)
	block.Hash = hash[:]
}