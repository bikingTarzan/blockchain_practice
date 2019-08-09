package main

import (
	"fmt"
	"bytes"
	"math/big"
	"crypto/sha256"
)

type ProofOfWork struct {
	block *Block

	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork {
		block : block,
	}

	targetStr := "00001000000000000000000000000000000000000000000000000000000000000"

	tmpInt := big.Int{}
	tmpInt.SetString(targetStr, 16)
	pow.target = &tmpInt
	return &pow
}

func (pow *ProofOfWork) Run() ([]byte, uint64) {
	// 1.拼装数据（就是区块链的交易数据）
	// 2.做哈希计算
	// 3.与pow中的target进行比较
	// 4.找到，退出
	// 5.没找到，继续，nonce+1

	var nonce uint64
	fmt.Println("刚刚声明的nonce：", nonce)
	block := pow.block
	var hash [32]byte


	for {
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.Data,
		}

		blockInfo := bytes.Join(tmp, []byte{})

		hash = sha256.Sum256(blockInfo)
		tmpInt := big.Int{}
		tmpInt.SetBytes(hash[:])
		if tmpInt.Cmp(pow.target) == -1 {
			fmt.Printf("挖矿成功！ hash: %x, nonce: %d\n", hash, nonce)
			return hash[:], nonce
		} else {
			nonce++
		}
	}

	return []byte("helloworld"), 10
}
