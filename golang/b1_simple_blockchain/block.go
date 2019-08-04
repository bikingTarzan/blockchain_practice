package main

import (
	"log"
	"time"
	"bytes"
	"crypto/sha256"
	"encoding/binary"
)


// 0.定义一个区块
type Block struct {
	// 1.版本号
	Version uint64
	// 2.前一个区块链的哈希
	PrevHash []byte
	// 3.Merkel根
	MerkelRoot []byte
	// 4.时间戳
	TimeStamp uint64
	// 5.难度值
	Difficult uint64
	// 6.随机数，挖矿时要用的
	Nonce uint64

	Hash []byte
	//3.数据
	Data []byte
}

// 实现一个辅助函数，把unit64转成byte
func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer

	err := binary.Write(&buffer, binary.BigEndian, num)

	if err != nil{
		log.Panic(err)
	}

	return buffer.Bytes()
}

// 2.创建一个区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:	00,
		PrevHash:	prevBlockHash,
		MerkelRoot: []byte{},
		TimeStamp:	uint64(time.Now().Unix()),
		Difficult:	0,
		Nonce:		0,
		Hash:		[]byte{},
		Data:		[]byte(data),
	}

	// block.SetHash()

	//创建一个pow对象
	pow := NewProofOfWork(&block)
	hash, nonce := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return &block
}

// 3.生成哈希
func (block *Block) SetHash() {
	// 1.拼接数据
	// blockinfo := append(block.PrevHash, block.Data...)

	//1. 拼装数据
	/*
	blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
	blockInfo = append(blockInfo, block.PrevHash...)
	blockInfo = append(blockInfo, block.MerkelRoot...)
	blockInfo = append(blockInfo, Uint64ToByte(block.TimeStamp)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
	blockInfo = append(blockInfo, block.Data...)
	*/

	//1. 拼装数据,join
	tmp := [][]byte{
		Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficult),
		Uint64ToByte(block.Nonce),
		block.Data,
	}

	blockinfo := bytes.Join(tmp, []byte{})

	// 2.sha256
	hash := sha256.Sum256(blockinfo)
	block.Hash = hash[:]
}