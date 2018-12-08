package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"log"
)

// UTXO Value 只能花费一次 余额等于未花费相加
type Transaction struct {
	TxHash []byte      //交易Hash,拼接所有交易的hash值
	Vins   []*TXInput  //输入
	Vouts  []*TXOutput //输出
}

//Transaction创建分为两种情况
//创世区块创建产生Transaction
func NewCoinbaseTransaction(address string) *Transaction {
	txInput := &TXInput{[]byte{}, -1, "Genesis Block"}
	txOutput := &TXOutput{10, address}
	txCoinbase := &Transaction{[]byte{}, []*TXInput{txInput}, []*TXOutput{txOutput}}
	txCoinbase.HashTransaction()
	return txCoinbase
}

func (tx *Transaction) IsCoinBaseTransaction() bool {
	return len(tx.Vins[0].TxHash) == 0 && tx.Vins[0].Vout == -1
}

//转帐时产生Transaction
func NewSimpleTransaction(from string, to string, amount int64, blockchain *BlockChain) *Transaction {

	//返回from的所有的未花费交易输出所对应的Transaction
	//unSpentTx := blockchain.UnUTXOsWithAddress(from)
	//返回适合转帐的花费，以及TXhash与Vout在Transaction的索引的统一字典 money,hash,vout索引 即money,{hash1:[0,1,2,3],hash2:[0,1,2]}
	money, spendableUTXOdic := blockchain.FindSpendableUTXOs(from, amount)

	var txInputs []*TXInput
	var txOutputs []*TXOutput
	//消费
	//bytes, _ := hex.DecodeString("5ae85331db2061580b5e404f47b0fe4b5aa48b96b363e4b3b19d34e481bb5d95")
	//txInput := &TXInput{bytes, 0, from}
	//txInputs = append(txInputs, txInput)
	for txHash, indexArray := range spendableUTXOdic {
		txHashBytes, _ := hex.DecodeString(txHash)
		for _, index := range indexArray {
			txInput := &TXInput{txHashBytes, (int64)(index), from}
			txInputs = append(txInputs, txInput)
		}
	}
	//转帐
	txOutput := &TXOutput{amount, to}
	txOutputs = append(txOutputs, txOutput)

	//找零
	txOutput = &TXOutput{money - amount, from}
	txOutputs = append(txOutputs, txOutput)

	txTransaction := &Transaction{[]byte{}, txInputs, txOutputs}
	txTransaction.HashTransaction()

	return txTransaction
}

//生成交易哈希
func (tx *Transaction) HashTransaction() {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash := sha256.Sum256(result.Bytes())
	tx.TxHash = hash[:]

}
