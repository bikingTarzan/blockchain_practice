package main

import (
	"fmt"
	"time"
)

/*
func (cli *CLI) AddBlock(txs []*Transaction) {
	cli.bc.AddBlock(txs)
	fmt.Printf("添加区块成功！\n")
}
*/

//正向打印
func (cli *CLI) PrinBlockChain() {
	cli.bc.Printchain()
	fmt.Printf("打印区块链完成\n")
}

func (cli *CLI) PrinBlockChainReverse() {
	bc := cli.bc
	//创建迭代器
	it := bc.NewIterator()

	//调用迭代器，返回我们的每一个区块数据
	for {
		//返回区块，左移
		block := it.Next()

		fmt.Printf("===========================\n\n")
		fmt.Printf("版本号: %d\n", block.Version)
		fmt.Printf("前区块哈希值: %x\n", block.PrevHash)
		fmt.Printf("梅克尔根: %x\n", block.MerkelRoot)
		timeFormat := time.Unix(int64(block.TimeStamp), 0).Format("2006-01-02 15:04:05")
		fmt.Printf("时间戳: %s\n", timeFormat)
		fmt.Printf("难度值(随便写的）: %d\n", block.Difficulty)
		fmt.Printf("随机数 : %d\n", block.Nonce)
		fmt.Printf("当前区块哈希值: %x\n", block.Hash)
		fmt.Printf("区块数据 :%s\n", block.Transactions[0].TXInputs[0].Sig)

		if len(block.PrevHash) == 0 {
			fmt.Printf("区块链遍历结束！")
			break
		} 
	}
}


func (cli *CLI) GetBalance(address string) {
	utxos := cli.bc.FindUTXOs(address)
	
	total :=0.0

	for _, utxo := range utxos {
		total += utxo.Value
	}

	fmt.Printf("\"%s\"的余额位：%f\n", address, total)
}

func (cli *CLI) Send(from, to string, amount float64, miner, data string) {


	//1. 创建挖矿交易
	coinbase := NewCoinbaseTx(miner, data)
	//2. 创建一个普通交易
	tx := NewTransaction(from, to, amount, cli.bc)
	if tx == nil {
		//fmt.Printf("无效的交易")
		return
	}
	//3. 添加到区块

	cli.bc.AddBlock([]*Transaction{coinbase, tx})
	fmt.Printf("转账成功！\n")
}
