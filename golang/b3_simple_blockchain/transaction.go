package main


import (
	"log"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
)

const reward = 50.0
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
	Value float64
	// 锁定脚本，我们用地址模拟
	PubkeyHash string
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

// 实现一个函数，判断当前的交易是否为挖矿交易
func (tx *Transaction) IsCoinbase() bool {
	// if len(tx.TXInputs) == 1 {
	// 	input := tx.TXInputs[0]

	// 	if !bytes.Equal(input.TXid, []byte{}) || input.Index != -1 {
	// 		return false
	// 	}
	// }

	if len(tx.TXInputs) == 1 && len(tx.TXInputs[0].TXid) == 0 && tx.TXInputs[0].Index == -1 {
		return true
	}

	return false
}

// 2.创建挖矿交易
func NewCoinbaseTx(address string, data string) *Transaction {
	input := TXInput{[]byte{}, -1, data}
	output := TXOutput{reward, address}
	
	tx := Transaction{[]byte{}, []TXInput{input}, []TXOutput{output}}
	tx.SetHash()

	return &tx
}


// 创建普通的转账交易
// 创建output
// 如果有零钱，要找零

func NewTransaction(from, to string, amount float64, bc *BlockChain) *Transaction{
	utxos, resValue := bc.FindNeedUTXOs(from, amount)

	if resValue < amount {
		return nil
	}

	var inputs []TXInput
	var outputs []TXOutput

	//2.创建交易输入，将这些UTXO逐一转成inputs
	for id, indexArray := range utxos {
		for _, i := range indexArray {
			input := TXInput{[]byte(id), int64(i), from}
			inputs = append(inputs, input)
		}
	}

	// 创建交易输出
	output := TXOutput{amount, to}
	outputs = append(outputs, output)

	// 找零
	if resValue > amount {
		outputs = append(outputs, TXOutput{resValue-amount, from})
	}

	tx := Transaction{[]byte{}, inputs, outputs}
	tx.SetHash()

	return &tx

}