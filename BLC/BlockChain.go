package BLC

<<<<<<< HEAD
import (
	"encoding/hex"
=======
<<<<<<< HEAD
import (
>>>>>>> 6fb89150a277ef0b9a9dcaa6e5acd74d7d5e8b8a
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
<<<<<<< HEAD
func CreateBlockChainWithGenesisBlock(address string) *BlockChain {
=======
func CreateBlockChainWithGenesisBlock(data string) {
>>>>>>> 6fb89150a277ef0b9a9dcaa6e5acd74d7d5e8b8a
	if dbExists() {
		fmt.Println("创世区块已经存在")
		os.Exit(1)
	}
<<<<<<< HEAD
	var genesisBlock *Block
=======
>>>>>>> 6fb89150a277ef0b9a9dcaa6e5acd74d7d5e8b8a

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
<<<<<<< HEAD
			//创建Coinbase Transaction
			txCoinbase := NewCoinbaseTransaction(address)
			genesisBlock = CreateGenesisBlock([]*Transaction{txCoinbase})
=======
			genesisBlock := CreateGenesisBlock(data)
>>>>>>> 6fb89150a277ef0b9a9dcaa6e5acd74d7d5e8b8a
			err := bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			err = bucket.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
<<<<<<< HEAD
=======
			//更新区块Hash

>>>>>>> 6fb89150a277ef0b9a9dcaa6e5acd74d7d5e8b8a
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	//defer db.Close()
<<<<<<< HEAD
	return &BlockChain{genesisBlock.Hash, db}
=======
>>>>>>> 6fb89150a277ef0b9a9dcaa6e5acd74d7d5e8b8a
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

<<<<<<< HEAD
//转帐挖区块
func (blockchain *BlockChain) MineNewBlock(from []string, to []string, amount []string) {
	//建立Transaction数组,打包交易
	//建立新的区块
}

//链接区块
func (blockchain *BlockChain) AddBlockToBlockChain(txs []*Transaction) {
=======
//链接区块
func (blockchain *BlockChain) AddBlockToBlockChain(data string) {
>>>>>>> 6fb89150a277ef0b9a9dcaa6e5acd74d7d5e8b8a

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
<<<<<<< HEAD
			newBlock := NewBlock(txs, latestBlock.Height+1, latestBlock.Hash)
=======
			newBlock := NewBlock(data, latestBlock.Height+1, latestBlock.Hash)
>>>>>>> 6fb89150a277ef0b9a9dcaa6e5acd74d7d5e8b8a
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
<<<<<<< HEAD
		fmt.Printf("Nonce: %d\n", latestBlock.Nonce)
		//fmt.Printf("Txs: %v\n", latestBlock.Txs)
		fmt.Println("Transactions:")
		fmt.Println("{")
		for _, tx := range latestBlock.Txs {
			fmt.Printf("  TxHash: %x\n", tx.TxHash)
			fmt.Println("  Transaction Vins")
			for _, in := range tx.Vins {
				fmt.Printf("    Vins TxHash: %x\n", in.TxHash)
				fmt.Printf("    Vins Vout: %d\n", in.Vout)
				fmt.Printf("    Vins ScriptSig: %s\n", in.ScriptSig)
			}
			fmt.Println("  Transaction Vouts")
			for _, out := range tx.Vouts {
				fmt.Printf("    Vouts Value: %d\n", out.Value)
				fmt.Printf("    Vouts ScriptPubKey: %s\n", out.ScriptPubKey)
			}
		}
		fmt.Println("}")
		fmt.Println("----------------------------------")
=======
		fmt.Printf("Data: %s\n", latestBlock.Data)
		fmt.Printf("Nonce: %d\n", latestBlock.Nonce)
>>>>>>> 6fb89150a277ef0b9a9dcaa6e5acd74d7d5e8b8a
		var hashInt big.Int
		hashInt.SetBytes(latestBlock.PrevHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
	}

}

<<<<<<< HEAD
//转帐时查找可用的UTXO
func (blockchain *BlockChain) FindSpendableUTXOs(from string, amount int64) (int64, map[string][]int) {
	//获取from的所有的未花费交易输出所对应的Transaction
	utxos := blockchain.UnUTXOsWithAddress(from)
	spendableUTXOdic := make(map[string][]int)
	//遍历查找可用UTXO
	var value int64
	for _, utxo := range utxos {
		value = value + utxo.txOutput.Value
		hash := hex.EncodeToString(utxo.TxHash)
		spendableUTXOdic[hash] = append(spendableUTXOdic[hash], utxo.index)
		if value >= amount {
			break
		}

	}
	if value < amount {
		fmt.Printf("%s's fund is not enough\n", from)
		os.Exit(1)
	}
	return value, spendableUTXOdic
}

//返回address的所有未花费交易输出所对应的Transaction
func (blockchain *BlockChain) UnUTXOsWithAddress(address string) []*UTXO {

	var unUTXOs []*UTXO                        //所有未花费交易输出
	spentTxOutputs := make(map[string][]int64) //已花费{hash:[0,1,2]}

	var latestBlock *Block
	iterator := blockchain.Iterator()
	for {
		latestBlock = iterator.Next()

		for _, tx := range latestBlock.Txs {
			if tx.IsCoinBaseTransaction() == false {
				for _, in := range tx.Vins {
					//是否能够解锁
					if in.UnLockWithAddress(address) {
						key := hex.EncodeToString(in.TxHash)
						spentTxOutputs[key] = append(spentTxOutputs[key], in.Vout)
					}
				}
			}
		work:
			for index, out := range tx.Vouts {
				if out.UnLockWithAddress(address) {
					//判断是否消费 hash相等 index与Vout索引相等
					//fmt.Println(out)
					if spentTxOutputs != nil {
						if len(spentTxOutputs) != 0 {
							var isSpentUTXO bool
							for txHash, indexArray := range spentTxOutputs {
								for _, i := range indexArray {
									if (int64)(index) == i && txHash == hex.EncodeToString(tx.TxHash) {
										isSpentUTXO = true
										continue work
									}
								}
							}
							if isSpentUTXO == false {
								utxo := &UTXO{tx.TxHash, index, out}
								unUTXOs = append(unUTXOs, utxo)
							}

						} else {
							utxo := &UTXO{tx.TxHash, index, out}
							unUTXOs = append(unUTXOs, utxo)
						}

					}
				}
			}
		}
		//fmt.Println(spentTxOutputs)
		var hashInt big.Int
		hashInt.SetBytes(latestBlock.PrevHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
	}
	return unUTXOs
}
func (blockchain *BlockChain) GetBalance(address string) int64 {
	utxos := blockchain.UnUTXOsWithAddress(address)
	var amount int64
	for _, utxo := range utxos {
		amount = amount + utxo.txOutput.Value
	}
	return amount
}

=======
>>>>>>> 6fb89150a277ef0b9a9dcaa6e5acd74d7d5e8b8a
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
<<<<<<< HEAD
=======

>>>>>>> 6fb89150a277ef0b9a9dcaa6e5acd74d7d5e8b8a
				fmt.Printf("Height: %d\n", latestBlock.Height)
				fmt.Printf("Hash: %x\n", latestBlock.Hash)
				fmt.Printf("PrevHash: %x\n", latestBlock.PrevHash)
				fmt.Printf("Timestamp: %s\n", time.Unix(latestBlock.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
				fmt.Printf("Data: %s\n", latestBlock.Data)
				fmt.Printf("Nonce: %d\n", latestBlock.Nonce)
<<<<<<< HEAD
=======

>>>>>>> 6fb89150a277ef0b9a9dcaa6e5acd74d7d5e8b8a
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
<<<<<<< HEAD
=======
=======
type BlockChain struct {
	Blocks []*Block
}

//创建带有创世区块的区块链
func CreateBlockChainWithGenesisBlock() *BlockChain {
	genesisBlock := CreateGenesisBlock("Genesis Block")
	return &BlockChain{[]*Block{genesisBlock}}
}

//链接区块
func (blockchain *BlockChain) AddBlockToBlockChain(data string, height int64, preHash []byte) {
	newBlock := NewBlock(data, height, preHash)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)

}
>>>>>>> 6f0bdfe282c989920a0911c550946d37c54924de
>>>>>>> 6fb89150a277ef0b9a9dcaa6e5acd74d7d5e8b8a
