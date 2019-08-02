package main

import (
	"fmt"
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

// 4.引入区块链
type BlockChain struct {
	// 定义一个区块链数组
	blocks []*Block
}

// 5.定义一个区块链
func NewBlockChain() *BlockChain {
	// 创建一个创世块并添加到区块链中
	genesisBlock := GenesisBlock()

	return &BlockChain{
		blocks : []*Block{genesisBlock},
	}
}

// 定义一个创世区块
func GenesisBlock() *Block {
	return NewBlock("第一个区块，创世区块，老牛逼了", []byte{})
}

// 6.添加一个区块到区块链中
func (bc *BlockChain) AddBlock(data string) {
	// 获取区块链最后一个区块
	lastBLock := bc.blocks[len(bc.blocks) -1]
	prevHash := lastBLock.Hash

	// a.生成一个新区块
	block := NewBlock(data, prevHash)

	// b.添加区块到区块链中
	bc.blocks = append(bc.blocks, block)
}

func main() {
	bc := NewBlockChain()

	bc.AddBlock("班长向班花转了50枚比特币")
	bc.AddBlock("班长又向班花转了50枚比特币")

	for i, block := range bc.blocks {
		fmt.Printf("==================   当前区块高度：%v   ==================\n", i)
		fmt.Printf("前一个区块的哈希： %x \n", block.PrevHash)
		fmt.Printf("当前区块的哈希： %x \n", block.Hash)
		fmt.Printf("区块的数据： %s \n", block.Data)
	}
}
