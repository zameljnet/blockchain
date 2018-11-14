package BLC

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
