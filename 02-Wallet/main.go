package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Cannot generate private key: %v", err)
	}
	privateData := crypto.FromECDSA(pvk)
	fmt.Println(hexutil.Encode(privateData))

	publicKey := crypto.FromECDSAPub(&pvk.PublicKey)
	fmt.Println(hexutil.Encode(publicKey))

	fmt.Println(crypto.PubkeyToAddress(pvk.PublicKey).Hex())
}
