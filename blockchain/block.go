package blockchain

import "fmt"

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	PrevHash []byte
	Data     []byte
	Hash     []byte
	Nonce    int
}

// func (b *Block) DeriveHash() {
// 	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})

// 	hash := sha256.Sum256(info)
// 	b.Hash = hash[:]
// }

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{Hash: []byte{}, Data: []byte(data), PrevHash: prevHash, Nonce: 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	// block.PrevHash = prevHash
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]

	newBlock := CreateBlock(data, prevBlock.Hash)
	fmt.Printf("prev Hash: %x\n", prevBlock.Hash)
	fmt.Printf("new block ka prevhash: %x\n", newBlock.PrevHash)
	fmt.Printf("new block ka hash: %x\n", newBlock.Hash)

	chain.Blocks = append(chain.Blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
