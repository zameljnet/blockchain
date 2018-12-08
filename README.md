# blockchain
v1.2 blockchain demo using go with UTXO Transaction model and CLI
Key-Value database: boltdb
# Quick Start
The following command will run the entire process (code build and run ).

```sh
$ go build main.go
$ ./main
```
# Usage
createBlockChain -address --交易地址
send -from From -to To -amount Amount --交易明细
printChain --输出区块信息
getBalance -address --查看余额
# Test
The following command will run the CLI.

create GenesisBlock
```sh
$ ./main createBlockChain -address
```

transaction
./main send -from '["a","b"]' -to '["c","d"]' -amount '["2","3"]'
```sh
$ ./send -from From -to To -amount Amount
example From [Array] to [Array] Amount [Array]
$ ./main send -from '["addressA","addressB"]' -to '["addressC","addressD"]' -amount '["amountA","amountB"]' 
```

print the whole blockchain
```sh
$ ./main printChain
```

Query Balance
```sh
$ ./main getBalance -address
```
