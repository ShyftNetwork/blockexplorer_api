package events

import (
	"math/big"
	"github.com/ShyftNetwork/go-empyrean/ethclient"
	"github.com/ShyftNetwork/go-empyrean"
	"log"
	"strings"
	"github.com/ShyftNetwork/go-empyrean/erc"
	"fmt"
	"github.com/ShyftNetwork/go-empyrean/common"
	"github.com/ShyftNetwork/go-empyrean/accounts/abi/bind"
	"context"
	"github.com/ShyftNetwork/go-empyrean/accounts/abi"
	"github.com/ShyftNetwork/go-empyrean/crypto"
)


func main() {
	client, err := ethclient.Dial("ws://localhost:8546")
	if err != nil {
	log.Fatal(err)
	}

	tokenAddress := common.HexToAddress("0x3e677837e3b69f80f2fee84a41a14e3be12a9123")
	instance, err := erc.NewErc(tokenAddress, client)
	if err != nil {
	log.Fatal(err)
	}
	LoggingERC20Events(tokenAddress, client)
	const key = `{"address":"b276840e8b89c64b517629de60de861e85f539ca","crypto":{"cipher":"aes-128-ctr","ciphertext":"9d042f1cc358064ff69ae86cec9db5756437c63468ba86c350c38dab567eb428","cipherparams":{"iv":"0a669aaefabd4540b44d6e43a20318ff"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"2971272a3cfe7e43a836d1e07684175290259a137a1443dcc28dbc2125f90768"},"mac":"b3bdd77d0898712de0efd359319d6c5af505d8d7e02ebb1c015752639f39eaa4"},"id":"afddee02-128c-4d02-bf99-1d2def8acf73","version":3}`
	auth, err := bind.NewTransactor(strings.NewReader(key), "123456")
	if err != nil {
	log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	address := common.HexToAddress("0x3e677837e3b69f80f2fee84a41a14e3be12a9123")
	tx, err := instance.Transfer(auth, address, big.NewInt(100000000))
	if err != nil {
	log.Fatal(err)
	}

	fmt.Println("TRANSFERHASH:::::", tx.Hash().String())
	fmt.Println("T0 ADDR     :::::", tx.To().String())
	fmt.Println("FROM ADDR   :::::", tx.From().String())
	fmt.Println("CONTRACT ADDR::::", tokenAddress.Hex())
}

type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

func LoggingERC20Events(addr common.Address, client *ethclient.Client) {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(360),
		ToBlock:   big.NewInt(700),
		Addresses: []common.Address{
			addr,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(erc.ErcABI)))
	if err != nil {
		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")

			var transferEvent LogTransfer

			err := contractAbi.Unpack(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())

		case logApprovalSigHash.Hex():
			fmt.Printf("Log Name: Approval\n")

			var approvalEvent LogApproval

			err := contractAbi.Unpack(&approvalEvent, "Approval", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
			approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
			fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
			fmt.Printf("Tokens: %s\n", approvalEvent.Tokens.String())
		}

		fmt.Printf("\n\n")
	}
}
