package main

import (
	"fmt"
)

// 存储方式由slice更换为blkot数据库（kv库）
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
