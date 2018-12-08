package BLC

type UTXO struct {
	TxHash   []byte
	index    int
	txOutput *TXOutput
}
