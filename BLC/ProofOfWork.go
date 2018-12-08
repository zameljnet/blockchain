package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//sha256 0000 0000 0000 0000 1001 0001 0000 .... 0001

//意为256位Hash里前面至少要有16个0
const targetBit = 16

type ProofOfWork struct {
	Block  *Block   //需要进行验证的区块
	Target *big.Int //难度大数列 有几个0
}

//多行注释ctrl+k,ctrl+c，块注释ctrl+shift+A
/* 创建工作量证明对象，对区块进行验证
*假设8位Hash值，实现前面有两个0
*0010 0000  最大为0011 1111 即为63
*那么难度值应该设置为0100 0000 即64，只要找到hash值比难度值小，即满足验证条件
*显然难度是通过0000 0001 左移(8-2)得到的
 */
func NewProofOfWork(block *Block) *ProofOfWork {

	//创建初始值为1的target
	target := big.NewInt(1)
	//左移256-targetBit,设置难度
	target = target.Lsh(target, 256-targetBit)

	return &ProofOfWork{Block: block, Target: target}
}

func (proofOfWork *ProofOfWork) IsVaild() bool {
	var hashInt big.Int
	hashInt.SetBytes(proofOfWork.Block.Hash[:])
	if proofOfWork.Target.Cmp(&hashInt) == 1 {
		return true
	}
	return false
}

//工作量证明方法,判断生成hash的有效性
func (proofOfWork *ProofOfWork) Run() ([]byte, int64) {
	nonce := 0          //随机数
	var hashInt big.Int //存储新生成的hash值
	var hash [32]byte

	for {
		dataBytes := proofOfWork.prepareData(nonce)
		hash = sha256.Sum256(dataBytes)
<<<<<<< HEAD
		//fmt.Printf("\r%x", hash)
		fmt.Printf("mining---hash: %x\n", hash)
=======
		fmt.Printf("\r%x", hash)
>>>>>>> 6f0bdfe282c989920a0911c550946d37c54924de
		//将hash存储到hashInt中
		hashInt.SetBytes(hash[:])
		//找到比难度小的hashInt，cmp大于=1,小于=-1
		if proofOfWork.Target.Cmp(&hashInt) == 1 {
<<<<<<< HEAD
			fmt.Printf("mining success---hash: %x ", hash)
=======
>>>>>>> 6f0bdfe282c989920a0911c550946d37c54924de
			break
		}
		nonce = nonce + 1

	}
	return hash[:], int64(nonce)

}

//拼接字符数组
func (proofOfWork *ProofOfWork) prepareData(nonce int) []byte {
	block := proofOfWork.Block
	heightBytes := IntToHex(block.Height)
	timeBytes := IntToHex(block.Timestamp)
	nonceBytes := IntToHex(int64(nonce))
	targetBytes := IntToHex(int64(targetBit)) //加不加都可以
	blockBytes := bytes.Join(
		[][]byte{
			block.PrevHash,
			block.Data,
			timeBytes,
			targetBytes,
			nonceBytes,
			heightBytes,
		},
		[]byte{},
	) //拼接成二维字符数组

	return blockBytes

}
