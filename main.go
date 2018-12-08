package main

import (
<<<<<<< HEAD
	"flag"
	"fmt"
	"log"
	"os"
	"publicChain/BLC"
=======
	"fmt"
	"log"
	"publicChain/BLC"

	"github.com/boltdb/bolt"
>>>>>>> 6f0bdfe282c989920a0911c550946d37c54924de
)

func main() {
	//block := BLC.NewBlock("Genesis Block", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	//block := BLC.CreateGenesisBlock("Genesis Block")
	//fmt.Println(blockchain.Blocks)
	//fmt.Println(blockchain.Blocks[0])
	// block := BLC.NewBlock("Test", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	// pow :=BLC.NewProofOfWork(block)
	// fmt.Printf("%v\n", pow.IsVaild())
	// blockBytes := block.Serialize()
	// fmt.Println(blockBytes)
	// block1 := BLC.DeserializeBlock(blockBytes)
	// fmt.Println(block1)
<<<<<<< HEAD

	cli := &CLI{}
	cli.Run()
	//CLI命令行输入的字符串---Args

	/*
		blockchain := BLC.CreateBlockChainWithGenesisBlock()
		blockchain.AddBlockToBlockChain("Send 10 to B")
		time.Sleep(1e10)
		blockchain.AddBlockToBlockChain("Send 10 to C")
		time.Sleep(1e10)
		blockchain.AddBlockToBlockChain("Send 10 to D")
		time.Sleep(1e10)
		blockchain.AddBlockToBlockChain("Send 10 to E")
		fmt.Println("\nBlockChain")
		blockchain.PrintBlockChain()
		defer blockchain.DB.Close()
	*/
	/*
		block := BLC.NewBlock("Test", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
		//打开数据库
		// It will be created if it doesn't exist.mode代表权限 ，最大为777,可读1,可执行2,可写4
		db, err := bolt.Open("block.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		//只写
		err = db.Update(func(tx *bolt.Tx) error {
			//获取表对象
			bucket := tx.Bucket([]byte("blocks"))
			if bucket == nil {
				bucket, err = tx.CreateBucket([]byte("blocks"))
				if err != nil {
					log.Panic("Blocks table create failed: ", err)
				}
			}
			err := bucket.Put([]byte("l"), block.Serialize())
			if err != nil {
				log.Panic(err)
			}

			return nil
		})
=======
	/*
		blockchain := BLC.CreateBlockChainWithGenesisBlock()
		blockchain.AddBlockToBlockChain("Send 10 to B", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
		blockchain.AddBlockToBlockChain("Send 10 to C", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
		blockchain.AddBlockToBlockChain("Send 10 to D", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
		blockchain.AddBlockToBlockChain("Send 10 to E", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	*/

	block := BLC.NewBlock("Test", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	//打开数据库
	// It will be created if it doesn't exist.mode代表权限 ，最大为777,可读1,可执行2,可写4
	db, err := bolt.Open("block.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//只写
	err = db.Update(func(tx *bolt.Tx) error {
		//获取表对象
		bucket := tx.Bucket([]byte("blocks"))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte("blocks"))
			if err != nil {
				log.Panic("Blocks table create failed: ", err)
			}
		}
		err := bucket.Put([]byte("l"), block.Serialize())
>>>>>>> 6f0bdfe282c989920a0911c550946d37c54924de
		if err != nil {
			log.Panic(err)
		}

<<<<<<< HEAD
		//只读
		err = db.View(func(tx *bolt.Tx) error {
			//获取表对象
			bucket := tx.Bucket([]byte("blocks1"))
			if bucket == nil {
				log.Panic(err)
			}
			blockdata := bucket.Get([]byte("l"))
			fmt.Println(BLC.DeserializeBlock(blockdata))
			return nil
		})
		if err != nil {
			log.Panic(err)
		}
	*/
}

/* flag用法 -printchain flagPrintChainCmd
flagPringtChainCmd := flag.String("printchain", "", "输出所有的区块信息")
	flagInt := flag.Int("number", 6, "输出一个整数")
	flagBool := flag.Bool("open", false, "判断真假")
	flag.Parse()

	fmt.Printf("%s\n", *flagPringtChainCmd)
	fmt.Printf("%d\n", *flagInt)
	fmt.Printf("%v\n", *flagBool)
	if *flagPringtChainCmd == "printchain" {

	}
*/
=======
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	//只读
	err = db.View(func(tx *bolt.Tx) error {
		//获取表对象
		bucket := tx.Bucket([]byte("blocks1"))
		if bucket == nil {
			log.Panic(err)
		}
		blockdata := bucket.Get([]byte("l"))
		fmt.Println(BLC.DeserializeBlock(blockdata))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

}

>>>>>>> 6f0bdfe282c989920a0911c550946d37c54924de
/*
//打开数据库
db, err := bolt.Open("block.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
//数据库写
err = db.Update(func(tx *bolt.Tx) error {

})
if err != nil {
	log.Panic(err)
}
// 创建表
bucket, err := tx.CreateBucket([]byte("blockBucket"))
if err != nil {
	return fmt.Errorf("create bucket: %s", err)
}
// 往表里存数据
if bucket != nil {
	err := bucket.Put([]byte("l"), []byte("Send 10 to B"))
	if err != nil {
		log.Panic("数据存储失败")
	}
}
// 从表里读取数据
if bucket != nil {
	data := bucket.Get([]byte("l"))
	fmt.Printf("%s\n", data)
}
*/
<<<<<<< HEAD

//CLI
type CLI struct {
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreateBlockChain -data Data--交易数据")
	fmt.Println("\taddBlock -data Data --交易数据")
	fmt.Println("\tprintChain --输出区块信息")
}
func isVaildArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) createBlockChain(data string) {
	//cli.createBlockChainCmd(data)
	BLC.CreateBlockChainWithGenesisBlock(data)
}

func (cli *CLI) addBlock(data string) {
	blockchain := BLC.GetBlockChain()
	defer blockchain.DB.Close()
	blockchain.AddBlockToBlockChain(data)
}

func (cli *CLI) printChain() {
	blockchain := BLC.GetBlockChain()
	defer blockchain.DB.Close()
	blockchain.PrintBlockChain()
}

func (cli *CLI) Run() {
	isVaildArgs()

	//args := os.Args
	//fmt.Printf("arg: %v\n", args)
	createBlockChainCmd := flag.NewFlagSet("createBlockChain", flag.ExitOnError)
	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	flagAddBlockData := addBlockCmd.String("data", "Send 10 to A", "交易数据")
	flagCreateBlockChainData := createBlockChainCmd.String("data", "Genesis Block...", "创世区块交易数据")
	switch os.Args[1] {
	case "createBlockChain":
		err := createBlockChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printChain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)
	}
	if createBlockChainCmd.Parsed() {
		if *flagCreateBlockChainData == "" {
			printUsage()
			os.Exit(1)
		}
		cli.createBlockChain(*flagCreateBlockChainData)
	}

	if addBlockCmd.Parsed() {
		if *flagAddBlockData == "" {
			printUsage()
			os.Exit(1)
		}
		cli.addBlock(*flagAddBlockData)
	}

	if printChainCmd.Parsed() {

		cli.printChain()
	}

}
=======
>>>>>>> 6f0bdfe282c989920a0911c550946d37c54924de
