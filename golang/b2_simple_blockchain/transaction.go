package main


// 1.定义交易结构
type Transaction struct {
	TXID		[]byte
	TXInputs 	[]TXInput
	TXOutputs	[]TXOutput
}

// 定义交易输入
type TXInput struct {
	// 交易id
	TXid []byte
	// 引入的output的索引值
	Index int64
	// 解锁脚本
	Sig string
}

type TXOutput struct {
	// 转账金额
	value float64
	// 锁定脚本，我们用地址模拟
	PubKeyHash string
}