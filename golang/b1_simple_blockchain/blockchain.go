package main

import (
	"log"

	"./bolt"
)

// 4.引入区块链
type BlockChain struct {
	// // 定义一个区块链数组
	// blocks []*Block

	db *bolt.DB
	tail []byte
}

const blockChainDb = "blockChain.db"
const blockBucket = "blockBucket"

// 5.定义一个区块链
func NewBlockChain() *BlockChain {
	// 获取最后一个hash，从数据库里面读出来的

	var lastHash []byte

	// 1.打开数据库
	db, err := bolt.Open(blockChainDb, 0600, nil)

	defer db.Close()

	if err != nil {
		log.Panic("打开数据库失败！！")
	}

	// 更改数据库内容
	db.Update(func(tx *bolt.Tx) error {
		// 2.找到抽屉bucket(如果没有就创建)
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			//没有，就创建
			bucket, err = tx.CreateBucket([]byte(blockBucket))

			if err != nil {
				log.Panic("创建b1数据库失败")
			}

			//创建一个创世区块
			genesisBlock := GenesisBlock()

			// 3.写数据
			bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			bucket.Put([]byte("LastHashKey"), genesisBlock.Hash)
			lastHash = genesisBlock.Hash
		} else {
			lastHash = bucket.Get([]byte("LastHashKey"))
		}

		return nil
	})

	return &BlockChain{db, lastHash}
}

// 定义一个创世区块
func GenesisBlock() *Block {
	return NewBlock("第一个区块，创世区块，老牛逼了", []byte{})
}

// 6.添加一个区块到区块链中
func (bc *BlockChain) AddBlock(data string) {
	// 如果获取前一个区块的hash

	db := bc.db
	lastHash := bc.tail

	
	db.Update(func(tx *bolt.Tx) error {
		// bucket一定已经存在
		bucket := tx.Bucket([]byte(blockBucket))

		if bucket == nil {
			log.Panic("bucket 不应该为空，请检查")
		}

		// a.创建一个区块
		block := NewBlock(data, lastHash)
		// b.把区块写入区块链中
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte("LastHashKey"), block.Hash)
		lastHash = block.Hash

		return nil
	})
}
