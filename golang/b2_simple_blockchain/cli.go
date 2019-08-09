package main

import (
	"os"
	"fmt"
)


type CLI struct {
	bc *BlockChain
}


const Usage = `
	addBlock --data DATA     "add data to blockchain"
	printChain               "print all blockchain data" 
	getBalance --address ADDRESS "获取指定地址的余额"
`

func (cli *CLI) Run() {

	//./block printChain
	//./block addBlock --data "HelloWorld"
	//1. 得到所有的命令
	args := os.Args
	if len(args) < 2 {
		fmt.Printf(Usage)
		return
	}

	// 2.分析指令
	cmd := args[1]

	switch cmd {
	case "addBlock":
		// 3.添加区块
		fmt.Printf("添加区块\b")

		if len(args) == 4 && args[2] == "--data" {
			// a.获取数据
			// data := args[3]
			// b.添加数据
			// cli.AddBlock(data)
		} else {
			fmt.Printf("添加区块参数使用不当，请检查")
			fmt.Printf(Usage)
		}
	case "printChain":
		fmt.Printf("打印区块\n")
		cli.PrinBlockChain()
	case "getBalance":
		fmt.Printf("获取余额\n")
		if len(args) == 4 && args[2] == "--address" {
			address := args[3]
			cli.GetBalance(address)
		}
	default:
		fmt.Printf("无效的命令，请检查!\n")
		fmt.Printf(Usage)
	}
}