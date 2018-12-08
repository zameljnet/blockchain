package main

import "publicChain/BLC"

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

	cli := &BLC.CLI{}
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
