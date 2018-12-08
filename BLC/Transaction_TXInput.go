package BLC

type TXInput struct {
	TxHash    []byte //交易Hash-id
	Vout      int64  //TXOut在当前数组的索引 即对应的要消费的钱或者TXOut
	ScriptSig string //用户签名--用户名 即对应的要消费的Value所属
}

//判断当前这笔消费是哪个账户的
func (txInput *TXInput) UnLockWithAddress(address string) bool {
	return txInput.ScriptSig == address
}
