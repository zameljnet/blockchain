package BLC

import (
	"log"

	"github.com/boltdb/bolt"
)

//迭代器
type BlockChainIterator struct {
	CurrentHash []byte
	DB          *bolt.DB
}

func (blockchain *BlockChain) Iterator() *BlockChainIterator {
	//返回迭代器对象
	return &BlockChainIterator{blockchain.Top, blockchain.DB}
}

func (blockChainIterator *BlockChainIterator) Next() *Block {
	var latestBlock *Block

	err := blockChainIterator.DB.View(func(tx *bolt.Tx) error {
		//获取表对象
		bucket := tx.Bucket([]byte(blockTableName))
		if bucket == nil {
			log.Panic("get blockTable failed")
		}
		if bucket != nil {
			//获取当前迭代器里面的CurrentHash对应的区块
			latestBlockBytes := bucket.Get(blockChainIterator.CurrentHash)
			latestBlock = DeserializeBlock(latestBlockBytes)
			//更新迭代器，指向下一个区块
			blockChainIterator.CurrentHash = latestBlock.PrevHash

		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return latestBlock

}
