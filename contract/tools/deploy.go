package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/skjune12/grpc-eth/contract"
)

const GasLimit uint = 4712388

// const GasLimit uint = 2100000000

func main() {
	// load .env
	err := godotenv.Load()

	if err != nil {
		log.Println(".env is not specified.")
		if os.Getenv("ETH_SECRET_KEY") == "" {
			log.Fatalln("ETH_SECRET_KEY is not specified.")
		}
	}

	// client, err := ethclient.Dial("http://localhost:7545")
	client, err := ethclient.Dial("http://localhost:8000")
	if err != nil {
		log.Fatalln("ethclient.Dial:", err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("ETH_SECRET_KEY"))
	if err != nil {
		log.Fatalln("crypto.HexToECDSA:", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalln("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalln("client.PendingNonceAt:", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalln("client.SuggestGasPrice:", err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	// auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(GasLimit)
	auth.GasPrice = gasPrice

	address, tx, _, err := contract.DeployExampleContract(auth, client)
	if err != nil {
		log.Fatalln("contract.DeployExampleContract:", err)
	}

	fmt.Println("Contract Address:", address.Hex())
	fmt.Println("Transaction:", tx.Hash().Hex())
}
