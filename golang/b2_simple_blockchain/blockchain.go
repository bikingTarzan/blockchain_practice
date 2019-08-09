package main

import (
	"log"
	"fmt"
	"bytes"
	
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
func NewBlockChain(address string) *BlockChain {
	// 获取最后一个hash，从数据库里面读出来的

	var lastHash []byte

	// 1.打开数据库
	db, err := bolt.Open(blockChainDb, 0600, nil)

	//defer db.Close()

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
			genesisBlock := GenesisBlock(address)

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
func GenesisBlock(address string) *Block {
	coinbase := NewCoinbaseTx(address,"第一个区块，创世区块，老牛逼了")
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

// 6.添加一个区块到区块链中
func (bc *BlockChain) AddBlock(txs []*Transaction) {
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
		block := NewBlock(txs, lastHash)
		// b.把区块写入区块链中
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte("LastHashKey"), block.Hash)
		lastHash = block.Hash

		return nil
	})
}

func (bc *BlockChain) Printchain() {

	blockHeight := 0
	bc.db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("blockBucket"))

		//从第一个key-> value 进行遍历，到最后一个固定的key时直接返回
		b.ForEach(func(k, v []byte) error {
			if bytes.Equal(k, []byte("LastHashKey")) {
				return nil
			}

			block := Deserialize(v)
			//fmt.Printf("key=%x, value=%s\n", k, v)
			fmt.Printf("=============== 区块高度: %d ==============\n", blockHeight)
			blockHeight++
			fmt.Printf("版本号: %d\n", block.Version)
			fmt.Printf("前区块哈希值: %x\n", block.PrevHash)
			fmt.Printf("梅克尔根: %x\n", block.MerkelRoot)
			fmt.Printf("时间戳: %d\n", block.TimeStamp)
			fmt.Printf("难度值(随便写的）: %d\n", block.Difficulty)
			fmt.Printf("随机数 : %d\n", block.Nonce)
			fmt.Printf("当前区块哈希值: %x\n", block.Hash)
			fmt.Printf("区块数据 :%s\n", block.Transactions[0].TXInputs[0].Sig)
			return nil
		})
		return nil
	})
}


// 寻找指定地址的所有UTXO
func (bc *BlockChain) FindUTXOs(address string) []TXOutput {
	var UTXO []TXOutput

	spentOutputs := make(map[string][]int64)


	it := bc.NewIterator()

	for {
		// 1.遍历区块
		block := it.Next()

		// 2.遍历交易
		for _, tx := range block.Transactions {
			fmt.Printf("current txid : %x\n", tx.TXID)

			// 遍历output
			for i, output := range tx.TXOutputs {
				fmt.Printf("current index : %d\n", i)

				if output.PubkeyHash == address {
					UTXO = append(UTXO, output)
				}
			}

			// 遍历input
			for _, input := range tx.TXInputs {
				if input.Sig == address {
					indexArray := spentOutputs[string(input.TXid)]
					indexArray = append(indexArray, input.Index)
				}
			}
		}

		if len(block.PrevHash) == 0 {
			break
			fmt.Printf("区块遍历完成退出")
		}
	}

	return UTXO
}
