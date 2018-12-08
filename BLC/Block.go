package BLC

import (
	"bytes"
<<<<<<< HEAD
	"crypto/sha256"
=======
>>>>>>> 6fb89150a277ef0b9a9dcaa6e5acd74d7d5e8b8a
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

<<<<<<< HEAD
	Nonce int64 //随机数
	Txs   []*Transaction
=======
	Data []byte //数据

	Nonce int64 //随机数
>>>>>>> 6fb89150a277ef0b9a9dcaa6e5acd74d7d5e8b8a
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

<<<<<<< HEAD
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

=======
>>>>>>> 6fb89150a277ef0b9a9dcaa6e5acd74d7d5e8b8a
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
<<<<<<< HEAD
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

=======
func NewBlock(data string, height int64, preBlockHash []byte) *Block {
	block := &Block{Height: height, Hash: nil, PrevHash: preBlockHash, Timestamp: time.Now().Unix(), Data: []byte(data), Nonce: 0}
	//block.SetHash()
	//调用工作量证明方法，返回有效的Hash和Nonce值
	pow := NewProofOfWork(block)
<<<<<<< HEAD
	//000000,符合hash的nonce,验证
	hash, nonce := pow.Run()

=======

	//000000,符合hash的nonce,验证
	hash, nonce := pow.Run()
>>>>>>> 6f0bdfe282c989920a0911c550946d37c54924de
	block.Hash = hash[:]
	block.Nonce = nonce
>>>>>>> 6fb89150a277ef0b9a9dcaa6e5acd74d7d5e8b8a
	fmt.Println()
	fmt.Println("Block:", block)
	return block
}

//创建创世区块
<<<<<<< HEAD
func CreateGenesisBlock(txs []*Transaction) *Block {
	block := NewBlock(txs, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
=======
func CreateGenesisBlock(data string) *Block {
	block := NewBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
>>>>>>> 6fb89150a277ef0b9a9dcaa6e5acd74d7d5e8b8a
	//block.SetHash()
	//调用工作量证明方法，返回有效的Hash和Nonce值
	return block
}
<<<<<<< HEAD
=======
<<<<<<< HEAD
>>>>>>> 6fb89150a277ef0b9a9dcaa6e5acd74d7d5e8b8a

func SetHash(block *Block) {
	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
}
<<<<<<< HEAD
=======
=======
>>>>>>> 6f0bdfe282c989920a0911c550946d37c54924de
>>>>>>> 6fb89150a277ef0b9a9dcaa6e5acd74d7d5e8b8a
