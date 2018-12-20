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
			log.Fatal("ETH_SECRET_KEY is not specified.")
		}
	}

	// setup ethereum
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("ETH_SECRET_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("error casting public key to ECDSA.")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)

	auth.Nonce = big.NewInt(int64(nonce))
	// auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(GasLimit)
	auth.GasPrice = gasPrice

	// contract address (string)
	address := common.HexToAddress(os.Getenv("CONTRACT_ADDR"))
	instance, err := contract.NewExampleContract(address, client)
	if err != nil {
		log.Fatal(err)
	}

	switch in.Method {
	case GET:
		value, err := instance.Number(nil)
		if err != nil {
			log.Fatal("get", err)
		}

		log.Printf("GET from %s\n", fromAddress.String())
		return &ReturnMsg{Msg: fmt.Sprint(value)}, nil

	case ADD:
		tx, err := instance.SetNumber(auth, uint32(in.Value))

		if err != nil {
			log.Fatal("add", err)
		}

		log.Printf("ADD from %s (value = %d)\n", fromAddress.String(), in.Value)
		log.Printf("tx sent: %s\n", tx.Hash().Hex())
		return &ReturnMsg{Msg: fmt.Sprintf(tx.Hash().Hex())}, nil
	}

	return &ReturnMsg{Msg: "fail"}, nil
}
