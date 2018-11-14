package BLC

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

const dbName = "blockchain.db"
const blockTableName = "blocks"

type BlockChain struct {

	//Blocks []*Block
	Top []byte //最新的区块Hash
	DB  *bolt.DB
}

func dbExists() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}

	return true
}

//创建带有创世区块的区块链
func CreateBlockChainWithGenesisBlock(data string) {
	if dbExists() {
		fmt.Println("创世区块已经存在")
		os.Exit(1)
	}

	fmt.Println("创建创世区块")

	//链接数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	//只写
	err = db.Update(func(tx *bolt.Tx) error {

		bucket, err := tx.CreateBucket([]byte(blockTableName))
		if err != nil {
			log.Panic("Blocks table create failed: ", err)
		}

		if bucket != nil {
			//创建创世区块
			genesisBlock := CreateGenesisBlock(data)
			err := bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			err = bucket.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			//更新区块Hash

		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	//defer db.Close()
}

//创建BlockChain对象
func GetBlockChain() *BlockChain {
	if dbExists() == false {
		fmt.Println("数据库不存在")
		os.Exit(1)
	}
	blockdb, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	var blockchain *BlockChain

	err = blockdb.View(func(tx *bolt.Tx) error {
		//获取表对象
		bucket := tx.Bucket([]byte(blockTableName))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(blockTableName))
			if err != nil {
				log.Panic("Blocks table create failed: ", err)
			}
		}
		//读取最新区块的hash
		hash := bucket.Get([]byte("l"))
		blockchain = &BlockChain{hash, blockdb}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return blockchain
}

//链接区块
func (blockchain *BlockChain) AddBlockToBlockChain(data string) {

	//只写
	err := blockchain.DB.Update(func(tx *bolt.Tx) error {
		//创建表对象
		bucket := tx.Bucket([]byte(blockTableName))
		if bucket == nil {
			log.Panic("get blockTable failed")
		}
		if bucket != nil {
			//获取最新区块
			latestBlockBytes := bucket.Get(blockchain.Top)
			latestBlock := DeserializeBlock(latestBlockBytes)
			//创建新区块
			newBlock := NewBlock(data, latestBlock.Height+1, latestBlock.Hash)
			err := bucket.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			err = bucket.Put([]byte("l"), newBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			//更新最新区块的Hash
			blockchain.Top = newBlock.Hash
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

}

//遍历输出区块链信息
func (blockchain *BlockChain) PrintBlockChain() {
	var latestBlock *Block
	iterator := blockchain.Iterator()
	for {
		latestBlock = iterator.Next()
		fmt.Printf("Height: %d\n", latestBlock.Height)
		fmt.Printf("Hash: %x\n", latestBlock.Hash)
		fmt.Printf("PrevHash: %x\n", latestBlock.PrevHash)
		fmt.Printf("Timestamp: %s\n", time.Unix(latestBlock.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Data: %s\n", latestBlock.Data)
		fmt.Printf("Nonce: %d\n", latestBlock.Nonce)
		var hashInt big.Int
		hashInt.SetBytes(latestBlock.PrevHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
	}

}

/*
//遍历输出区块链信息
//for循环版本
func (blockchain *BlockChain) PrintBlockChain() {
	var latestBlock *Block
	var CurrentHash []byte = blockchain.Top

	for {
		err := blockchain.DB.View(func(tx *bolt.Tx) error {
			//获取表对象
			bucket := tx.Bucket([]byte(blockTableName))
			if bucket == nil {
				log.Panic("get blockTable failed")
			}
			if bucket != nil {

				latestBlockBytes := bucket.Get(CurrentHash)
				latestBlock = DeserializeBlock(latestBlockBytes)

				fmt.Printf("Height: %d\n", latestBlock.Height)
				fmt.Printf("Hash: %x\n", latestBlock.Hash)
				fmt.Printf("PrevHash: %x\n", latestBlock.PrevHash)
				fmt.Printf("Timestamp: %s\n", time.Unix(latestBlock.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
				fmt.Printf("Data: %s\n", latestBlock.Data)
				fmt.Printf("Nonce: %d\n", latestBlock.Nonce)

			}
			return nil
		})
		if err != nil {
			log.Panic(err)
		}
		var hashInt big.Int
		hashInt.SetBytes(latestBlock.PrevHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
		CurrentHash = latestBlock.PrevHash
	}
}
*/
