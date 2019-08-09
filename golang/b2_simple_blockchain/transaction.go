package main


import (
	"log"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
)

const reward = 12.5
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

// 对交易序列化
func (tx *Transaction) SetHash() {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	data := buffer.Bytes()
	hash := sha256.Sum256(data)
	tx.TXID = hash[:]
}

// 2.创建挖矿交易
func NewCoinbaseTx(address string, data string) *Transaction {
	input := TXInput{[]byte{}, -1, data}
	output := TXOutput{reward, address}
	
	tx := Transaction{[]byte{}, []TXInput{input}, []TXOutput{output}}
	tx.SetHash()

	return &tx
}