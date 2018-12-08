package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

//CLI
type CLI struct {
}

//    ./main send -from '["a","b"]' -to '["c","d"]' -amount '["2","3"]'
func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreateBlockChain -address --交易地址")
	fmt.Println("\tsend -from From -to To -amount Amount --交易明细")
	fmt.Println("\tprintChain --输出区块信息")
	fmt.Println("\tgetBalance -address --查看余额")
}
func isVaildArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

//创建创世区块
func (cli *CLI) createBlockChain(address string) {
	//cli.createBlockChainCmd(data)
	blockchain := CreateBlockChainWithGenesisBlock(address)
	blockchain.DB.Close()
}

//转帐
//blockchain.MineNewBlock(from, to, amount)
func (cli *CLI) send(from []string, to []string, amount []string) {
	blockchain := GetBlockChain()
	defer blockchain.DB.Close()
	//建立Transaction数组,打包交易
	//send -from '["zameljnet"]' -to '["a"]' -amount '["4"]'
	//send -from '["b"]' -to '["c"]' -amount '["2"]'
	//send -from '["zameljnet"]' -to '["c"]' -amount '["2"]'
	value, _ := strconv.Atoi(amount[0])

	tx := NewSimpleTransaction(from[0], to[0], (int64)(value), blockchain)

	var txs []*Transaction
	txs = append(txs, tx)

	//建立新的区块
	blockchain.AddBlockToBlockChain(txs)

}
func (cli *CLI) addBlock(txs []*Transaction) {
	blockchain := GetBlockChain()
	defer blockchain.DB.Close()
	blockchain.AddBlockToBlockChain(txs)
}

func (cli *CLI) printChain() {
	blockchain := GetBlockChain()
	defer blockchain.DB.Close()
	blockchain.PrintBlockChain()
}

//查询账户余额
func (cli *CLI) getBalance(address string) {
	blockchain := GetBlockChain()
	defer blockchain.DB.Close()
	amount := blockchain.GetBalance(address)
	fmt.Printf("%s的余额为: %d个Token\n", address, amount)
}

func (cli *CLI) Run() {
	isVaildArgs()

	//args := os.Args
	//fmt.Printf("arg: %v\n", args)
	createBlockChainCmd := flag.NewFlagSet("createBlockChain", flag.ExitOnError)
	sendBlockCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getBalance", flag.ExitOnError)
	flagFrom := sendBlockCmd.String("from", "", "转帐源地址")
	flagTo := sendBlockCmd.String("to", "", "转帐目的地址")
	flagAmount := sendBlockCmd.String("amount", "", "转帐金额")
	flagCreateBlockChainWithAddress := createBlockChainCmd.String("address", "Genesis Block...", "创建创世区块的地址")
	flagGetBalanceWithAddress := getBalanceCmd.String("address", "", "查询账户余额")
	switch os.Args[1] {
	case "createBlockChain":
		err := createBlockChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printChain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getBalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)
	}
	if createBlockChainCmd.Parsed() {
		if *flagCreateBlockChainWithAddress == "" {
			fmt.Println("地址不能为空")
			printUsage()
			os.Exit(1)
		}
		cli.createBlockChain(*flagCreateBlockChainWithAddress)
	}

	if sendBlockCmd.Parsed() {
		if *flagFrom == "" || *flagTo == "" || *flagAmount == "" {
			printUsage()
			os.Exit(1)
		}
		//判断符合JSON格式，及参数个数相同
		//cli.addBlock([]*Transaction{})
		from := JSONToArray(*flagFrom)
		to := JSONToArray(*flagTo)
		amoumt := JSONToArray(*flagAmount)
		cli.send(from, to, amoumt)
	}

	if printChainCmd.Parsed() {

		cli.printChain()
	}

	if getBalanceCmd.Parsed() {
		if *flagGetBalanceWithAddress == "" {
			fmt.Println("地址不能为空")
			printUsage()
			os.Exit(1)
		}
		cli.getBalance(*flagGetBalanceWithAddress)
	}

}
