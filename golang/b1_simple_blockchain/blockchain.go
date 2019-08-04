package main

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
