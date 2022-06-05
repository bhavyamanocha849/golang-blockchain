package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

type Proof struct {
	Block  *Block
	target *big.Int
}

const Difficulty = 10

func NewProof(b *Block) *Proof {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &Proof{b, target}

	return pow
}

func (pow *Proof) InitialiseBlock(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.Data,
			pow.Block.PrevHash,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

func (pow *Proof) Run() (int, []byte) {
	fmt.Println("Running.....")
	var intHash big.Int
	var hash [32]byte
	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitialiseBlock(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("Hash : %x\n", hash)

		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash[:]
}

func (pow *Proof) Validate() bool {
	var intHash big.Int

	data := pow.InitialiseBlock(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.target) == -1
}

//helper
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)

	}

	// fmt.Printf("buff data: %x\n", buff)
	return buff.Bytes()
}
