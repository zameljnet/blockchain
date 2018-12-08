package BLC

type TXOutput struct {
	Value        int64  //Value
	ScriptPubKey string //用户名
}

//判断当前这笔消费是哪个账户的
func (txOutput *TXOutput) UnLockWithAddress(address string) bool {
	return txOutput.ScriptPubKey == address
}
