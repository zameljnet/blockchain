package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

type Block struct {
	Height int64 //区块高度

	Hash []byte //区块HASH

	PrevHash []byte //上一个区块HASH

	Timestamp int64 //时间戳

	Nonce int64 //随机数
	Txs   []*Transaction
}

//设置区块HASH，需要通过block来调用
/*
func (block *Block) SetHash() {
	heightBytes := IntToHex(block.Height)
	timeString := strconv.FormatInt(block.Timestamp, 2) //转换成二进制，字符串
	timeBytes := []byte(timeString)
	blockBytes := bytes.Join([][]byte{heightBytes, block.Hash, block.PrevHash, timeBytes, block.Data}, []byte{}) //拼接成二维字符数组
	hash := sha256.Sum256(blockBytes)                                                                            //sha256 得到32个字节的HASH值
	block.Hash = hash[:]                                                                                         //进行截断，固定大小的字节数组
}
*/

//将交易转化成字节数组
func (block *Block) HashTransaction() []byte {
	var txHashes [][]byte
	var txHash [32]byte
	for _, tx := range block.Txs {
		txHashes = append(txHashes, tx.TxHash)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

//将区块序列化
func (block *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()

}

//将区块反序列化
func DeserializeBlock(blockBytes []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}

//创建区块 引入包直接可以NewBlock
func NewBlock(txs []*Transaction, height int64, preBlockHash []byte) *Block {

	block := &Block{Height: height, Hash: nil, PrevHash: preBlockHash, Timestamp: time.Now().Unix(), Txs: txs, Nonce: 0}

	//block.SetHash()
	//调用工作量证明方法，返回有效的Hash和Nonce值
	pow := NewProofOfWork(block)
	//000000,符合hash的nonce,验证
	hash, nonce := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	//go SetHash(block)

	fmt.Println()
	fmt.Println("Block:", block)
	return block
}

//创建创世区块
func CreateGenesisBlock(txs []*Transaction) *Block {
	block := NewBlock(txs, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	//block.SetHash()
	//调用工作量证明方法，返回有效的Hash和Nonce值
	return block
}

func SetHash(block *Block) {
	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
}