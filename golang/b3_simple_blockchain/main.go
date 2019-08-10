package main

import (
	// "fmt"
)


// 存储方式由slice更换为bolt数据库（kv库）
// 引入交易结构
func main() {
	bc := NewBlockChain("张三")

	cli := CLI{bc}
	cli.Run()



}
