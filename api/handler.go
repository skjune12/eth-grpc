package api

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/skjune12/grpc-eth/contract"
)

const GasLimit uint = 4712388

type Server struct {
}

// TODO: commonize const
const (
	ADD = iota
	GET = iota
)

func (s *Server) Exec(ctx context.Context, in *TestMsg) (*ReturnMsg, error) {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println(".env is not specified")

		if os.Getenv("ETH_SECRET_KEY") == "" {
			log.Fatalln("ETH_SECRET_KEY is not specified.")
		}
	}

	// setup ethereum client
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
		log.Fatalln("error casting public key to ECDSA.")
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
	auth.GasLimit = uint64(GasLimit)
	// auth.Value = big.NewInt(0)
	auth.GasPrice = gasPrice

	address := common.HexToAddress(os.Getenv("CONTRACT_ADDR"))
	instance, err := contract.NewExampleContract(address, client)
	if err != nil {
		log.Fatalln("contract.NewExampleContract:", err)
	}

	switch in.Method {
	case GET:
		value, err := instance.Number(nil)
		if err != nil {
			log.Fatalln("GET:", err)
		}

		log.Printf("GET from %s\n", fromAddress.String())
		return &ReturnMsg{Msg: fmt.Sprint(value)}, nil

	case ADD:
		tx, err := instance.SetNumber(auth, uint32(in.Value))

		if err != nil {
			log.Fatalln("ADD:", err)
		}

		log.Printf("ADD from %s (value = %d)\n", fromAddress.String(), in.Value)
		log.Printf("tx sent: %s\n", tx.Hash().Hex())

		fmt.Println("auth.GasLimit", auth.GasLimit)
		fmt.Println("auth.GasPrice", auth.GasPrice)
		return &ReturnMsg{Msg: fmt.Sprintf(tx.Hash().Hex())}, nil
	}

	return &ReturnMsg{Msg: "fail"}, nil
}
