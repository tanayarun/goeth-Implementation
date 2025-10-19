package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://mainnet.infura.io/v3/a54d875e40a344ef9da422b50ca4d226"
var address = common.HexToAddress("0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97")

func etherConversion(balance *big.Int) *big.Float {
	fBalance := new(big.Float)
	fBalance.SetInt(balance)
	EtherValue := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	return EtherValue
}

func main() {
	client, err := ethclient.DialContext(context.Background(), infuraURL)
	if err != nil {
		log.Fatalf("Error creating eth client: %v", err)
	}
	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error getting block number: %v", err)
	}

	fmt.Println("Latest block number:", block.Number())

	balance, err := client.BalanceAt(context.Background(), address , nil)
	if err != nil {
		log.Fatalf("Error getting balance: %v", err)
	}
	fmt.Printf("Balance: %.4f ETH ", etherConversion(balance))
}
