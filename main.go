package main

import (
	"fmt"

	"gitbhub.com/bhavyamanocha849/golang-blockchain/blockchain"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	chain := blockchain.InitBlockChain()

	chain.AddBlock("Firxdccdcdcdcdst After Geneisis")
	chain.AddBlock("Secondcdcdcdc After GEnesisis")

	for _, val := range chain.Blocks {
		fmt.Printf("prevHash: %x\n", val.PrevHash)
		fmt.Printf("Data: %s\n", val.Data)
		fmt.Printf("Hash: %x\n", val.Hash)
	}
}
