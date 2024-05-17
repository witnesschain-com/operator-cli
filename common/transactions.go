package wc_common

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetLatestNonce(client *ethclient.Client, privateKey *ecdsa.PrivateKey) *big.Int {
	senderAddress := GetPublicAddressFromPrivateKey(privateKey)
	nonce, err := client.PendingNonceAt(context.Background(), senderAddress)
	CheckError(err, "Pending nonce calculation failed")
	nonceNew := big.NewInt(int64(nonce))
	return nonceNew
}

func PrepareTransactionOptions(client *ethclient.Client, chainId big.Int, gasLimit uint64, privateKey *ecdsa.PrivateKey) *bind.TransactOpts {

	gasPrice, err := client.SuggestGasPrice(context.Background())
	CheckError(err, "Gas price calculation failed")

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, &chainId)
	CheckError(err, "Transactor creation failed")


	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = gasLimit   // in units
	auth.GasPrice = gasPrice

	return auth
}

func WaitForTransactionReceipt(client *ethclient.Client, txn *types.Transaction, timeout int64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*(time.Second))

	defer cancel()

	receipt, err := bind.WaitMined(ctx, client, txn)
	if err != nil {
		CheckError(err, "Transaction failed")
	} else if receipt.Status == 1 {
		fmt.Println("Transaction executed successfully, logs are ...")
		fmt.Println(receipt.Logs)
	} else if receipt.Status == 0 {
		FatalError("Transaction submitted successfully but failed to execute!")
	}
}
