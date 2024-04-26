package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli"

	"watchtower-operator/bindings/AvsDirectory"
	"watchtower-operator/bindings/OperatorRegistry"
	"watchtower-operator/bindings/WitnessHub"
)

type Config struct {
	WatchtowerPrivateKeys   []string       `json:"watchtower_private_keys"`
	OperatorPrivateKey      string         `json:"operator_private_key"`
	OperatorRegistryAddress common.Address `json:"operator_registry_address"`
	WitnessHubAddress       common.Address `json:"witnesshub_address"`
	AvsDirectoryAddress     common.Address `json:"avs_directory_address"`
	EthRPCUrl               string         `json:"eth_rpc_url"`
	ChainId                 big.Int        `json:"chain_id"`
	GasLimit                uint64         `json:"gas_limit"`
	TxReceiptTimeout        int64          `json:"tx_receipt_timeout"`
	ExpiryInDays            int64          `json:"expiry_in_days"`
}

var VERSION string = "UNKNOWN"

func main() {
	cli.AppHelpTemplate = fmt.Sprintf(`
	__        ___ _                        ____ _           _       
	\ \      / (_) |_ _ __   ___  ___ ___ / ___| |__   __ _(_)_ __  
	 \ \ /\ / /| | __|  _ \ / _ \/ __/ __| |   |  _ \ / _  | |  _ \ 
	  \ V  V / | | |_| | | |  __/\__ \__ \ |___| | | | (_| | | | | |
	   \_/\_/  |_|\__|_| |_|\___||___/___/\____|_| |_|\__,_|_|_| |_|

	%s`, cli.AppHelpTemplate)
	app := cli.NewApp()

	app.Name = "WitnessChain"
	app.Usage = "WitnessChain Operator CLI"
	app.Version = VERSION
	app.Copyright = "(c) 2024 WitnessChain"

	app.Commands = []cli.Command{
		{
			Name:  "registerWatchtower",
			Usage: "registerWatchtower --config-file ~/path/to/operator_config.json",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "config-file",
					Value: "config/l2-operator-config.json",
					Usage: "Path to the operator config file",
				},
			},
			Action: func(cCtx *cli.Context) error {
				config := GetConfigFromContext(cCtx)
				RegisterWatchtower(config)
				return nil
			},
		},
		{
			Name:  "deRegisterWatchtower",
			Usage: "deRegisterWatchtower --config-file ~/path/to/operator_config.json",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "config-file",
					Value: "config/l2-operator-config.json",
					Usage: "Path to the operator config file",
				},
			},
			Action: func(cCtx *cli.Context) error {
				config := GetConfigFromContext(cCtx)
				DeRegisterWatchtower(config)
				return nil
			},
		},
		{
			Name:  "registerOperatorToAVS",
			Usage: "registerOperatorToAVS --config-file ~/path/to/operator_config.json",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "config-file",
					Value: "config/l1-operator-config.json",
					Usage: "Path to the operator config file",
				},
			},
			Action: func(cCtx *cli.Context) error {
				config := GetConfigFromContext(cCtx)
				RegisterOperatorToAVS(config)
				return nil
			},
		},
		{
			Name:  "deRegisterOperatorFromAVS",
			Usage: "deRegisterOperatorFromAVS --config-file ~/path/to/operator_config.json",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "config-file",
					Value: "config/l1-operator-config.json",
					Usage: "Path to the operator config file",
				},
			},
			Action: func(cCtx *cli.Context) error {
				config := GetConfigFromContext(cCtx)
				DeregisterOperatorFromAVS(config)
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

func GetConfigFromContext(cCtx *cli.Context) *Config {
	configFilePath := cCtx.String("config-file")
	fmt.Printf("Using config file path : %s\n", configFilePath)

	data, err := os.ReadFile(configFilePath)
	CheckError(err, "Error reading json file")

	// Parse the json data into a struct
	var config Config
	err = json.Unmarshal(data, &config)
	CheckError(err, "Error unmarshaling json data")

	SetDefaultConfigValues(&config)

	return &config
}

func SetDefaultConfigValues(config *Config) {
	if config.ExpiryInDays == 0 {
		config.ExpiryInDays = 1 // 1 day
	}

	if config.TxReceiptTimeout == 0 {
		config.TxReceiptTimeout = 5 * 60 // 5 minutes
	}

	if config.GasLimit == 0 {
		config.GasLimit = 500000
	}
}

func RegisterWatchtower(config *Config) {
	client, err := ethclient.Dial(config.EthRPCUrl)
	CheckError(err, "Connection to RPC failed")

	id, _ := client.ChainID(context.Background())
	fmt.Println("Connection successful : ", id)

	operatorRegistry, err := OperatorRegistry.NewOperatorRegistry(config.OperatorRegistryAddress, client)
	CheckError(err, "Instantiating OperatorRegistry contract failed")

	operatorPrivateKey, err := crypto.HexToECDSA(config.OperatorPrivateKey)
	CheckError(err, "Converting private key to ECDSA format failed")

	operatorAddress := GetPublicAddressFromPrivateKey(operatorPrivateKey)
	regTransactOpts := PrepareTransactionOptions(client, config.ChainId, config.GasLimit, operatorPrivateKey)
	expiry := CalculateExpiry(client, config.ExpiryInDays)

	for _, watchTowerPkString := range config.WatchtowerPrivateKeys {
		watchtowerPrivateKey, err := crypto.HexToECDSA(watchTowerPkString)
		CheckError(err, "Converting private key to ECDSA format failed")

		signedMessage := SignOperatorAddress(client, config, watchtowerPrivateKey, operatorAddress, *expiry)

		watchtowerAddress := GetPublicAddressFromPrivateKey(watchtowerPrivateKey)

		regTx, err := operatorRegistry.RegisterWatchtowerAsOperator(regTransactOpts, watchtowerAddress, expiry, signedMessage)
		CheckError(err, "Registering watchtower as operator failed")
		fmt.Printf("Tx sent: %s\n", regTx.Hash().Hex())
		WaitForTransactionReceipt(client, regTx, config.TxReceiptTimeout)
	}
}

func DeRegisterWatchtower(config *Config) {
	client, err := ethclient.Dial(config.EthRPCUrl)
	CheckError(err, "Connection to RPC failed")

	id, _ := client.ChainID(context.Background())
	fmt.Println("Connection successful : ", id)

	operatorRegistry, err := OperatorRegistry.NewOperatorRegistry(config.OperatorRegistryAddress, client)
	CheckError(err, "Instantiating OperatorRegistry contract failed")

	operatorPrivateKey, err := crypto.HexToECDSA(config.OperatorPrivateKey)
	CheckError(err, "Converting private key to ECDSA format failed")

	deRegTransactOpts := PrepareTransactionOptions(client, config.ChainId, config.GasLimit, operatorPrivateKey)

	for _, watchTowerPkString := range config.WatchtowerPrivateKeys {
		watchtowerPrivateKey, err := crypto.HexToECDSA(watchTowerPkString)
		CheckError(err, "Converting private key to ECDSA format failed")

		watchtowerAddress := GetPublicAddressFromPrivateKey(watchtowerPrivateKey)

		deRegTx, err := operatorRegistry.DeRegister(deRegTransactOpts, watchtowerAddress)
		CheckError(err, "Deregister failed")
		fmt.Printf("Tx sent: %s\n", deRegTx.Hash().Hex())
		WaitForTransactionReceipt(client, deRegTx, config.TxReceiptTimeout)
	}
}

func RegisterOperatorToAVS(config *Config) {
	client, err := ethclient.Dial(config.EthRPCUrl)
	CheckError(err, "Connection to RPC failed")

	id, _ := client.ChainID(context.Background())
	fmt.Println("Connection successful : ", id)

	avsDirectory, err := AvsDirectory.NewAvsDirectory(config.AvsDirectoryAddress, client)
	CheckError(err, "Instantiating AvsDirectory contract failed")

	witnessHub, err := WitnessHub.NewWitnessHub(config.WitnessHubAddress, client)
	CheckError(err, "Instantiating WitnessHub contract failed")

	operatorPrivateKey, err := crypto.HexToECDSA(config.OperatorPrivateKey)
	CheckError(err, "Converting private key to ECDSA format failed")
	operatorAddress := GetPublicAddressFromPrivateKey(operatorPrivateKey)

	operatorSignature := GetOpertorSignature(client, avsDirectory, config, operatorPrivateKey, operatorAddress)

	avsRegtransactOpts := PrepareTransactionOptions(client, config.ChainId, config.GasLimit, operatorPrivateKey)

	tx, err := witnessHub.RegisterOperatorToAVS(avsRegtransactOpts, operatorAddress, operatorSignature)
	CheckError(err, "Registering operator to AVS failed")

	fmt.Printf("Tx sent: %s\n", tx.Hash().Hex())
	WaitForTransactionReceipt(client, tx, config.TxReceiptTimeout)
}

func DeregisterOperatorFromAVS(config *Config) {
	client, err := ethclient.Dial(config.EthRPCUrl)
	CheckError(err, "Connection to RPC failed")

	id, _ := client.ChainID(context.Background())
	fmt.Println("Connection successful : ", id)

	witnessHub, err := WitnessHub.NewWitnessHub(config.WitnessHubAddress, client)
	CheckError(err, "Instantiating WitnessHub contract failed")

	operatorPrivateKey, err := crypto.HexToECDSA(config.OperatorPrivateKey)
	CheckError(err, "Converting private key to ECDSA format failed")
	operatorAddress := GetPublicAddressFromPrivateKey(operatorPrivateKey)

	avsRegtransactOpts := PrepareTransactionOptions(client, config.ChainId, config.GasLimit, operatorPrivateKey)

	tx, err := witnessHub.DeregisterOperatorFromAVS(avsRegtransactOpts, operatorAddress)
	CheckError(err, "Dregistering operator from AVS failed")

	fmt.Printf("Tx sent: %s\n", tx.Hash().Hex())
	WaitForTransactionReceipt(client, tx, config.TxReceiptTimeout)
}

func CalculateExpiry(client *ethclient.Client, expectedExpiryDays int64) *big.Int {
	// Get the latest block header
	header, err := client.HeaderByNumber(context.Background(), nil)
	CheckError(err, "Could not get HeaderByNumber")

	// Get the current timestamp from the latest block header
	currentTimestamp := big.NewInt(int64(header.Time))

	expiryInSeconds := expectedExpiryDays * 24 * 60 * 60
	timeToElapse := big.NewInt(expiryInSeconds)

	expiry := new(big.Int).Add(currentTimestamp, timeToElapse)
	return expiry
}

func GenerateSalt() [32]byte {
	var salt [32]byte

	// Generate random bytes
	_, err := rand.Read(salt[:])
	CheckError(err, "Generating salt failed")

	return salt
}

func GetOpertorSignature(client *ethclient.Client, contractInstance *AvsDirectory.AvsDirectory, config *Config, operatorPrivateKey *ecdsa.PrivateKey, operatorAddress common.Address) WitnessHub.ISignatureUtilsSignatureWithSaltAndExpiry {
	salt := GenerateSalt()
	expiry := CalculateExpiry(client, config.ExpiryInDays)

	//ON AVS DIRECTORY
	digestHash, err := contractInstance.CalculateOperatorAVSRegistrationDigestHash(&bind.CallOpts{}, operatorAddress, config.WitnessHubAddress, salt, expiry)
	CheckError(err, "Digest hash calculation failed")

	signature, err := crypto.Sign(digestHash[:], operatorPrivateKey)
	CheckError(err, "Signing the digest hash failed")

	v := new(big.Int).SetBytes(signature[64:])
	v.Add(v, big.NewInt(27))

	// Construct the full signature (r, s, v)
	fullSignature := append(signature[:64], v.Bytes()...)
	operatorSignature := WitnessHub.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: fullSignature,
		Salt:      salt,
		Expiry:    expiry,
	}

	return operatorSignature
}

func SignOperatorAddress(client *ethclient.Client, config *Config, privateKey *ecdsa.PrivateKey, OperatorAddress common.Address, expiry big.Int) []byte {
	paddedAddr := GetPaddedValue(OperatorAddress.Bytes())
	paddedExpiry := GetPaddedValue(expiry.Bytes())

	encodedData := append(paddedAddr[:], paddedExpiry[:]...)
	// fmt.Printf("Encoded data : 0x%x\n", encodedData)

	hashedMessage := crypto.Keccak256(encodedData)
	// fmt.Printf("Hashed message : 0x%x\n", hashedMessage)

	hashedEthMessage := crypto.Keccak256(append([]byte("\x19Ethereum Signed Message:\n32"), hashedMessage...))
	// fmt.Printf("Hashed Eth message : 0x%x\n", hashedEthMessage)

	signature, err := crypto.Sign(hashedEthMessage, privateKey)
	CheckError(err, "Signing operator address failed")

	// fmt.Printf("Signature : 0x%x\n", signature)
	v := new(big.Int).SetBytes(signature[64:])
	v.Add(v, big.NewInt(27))

	// Construct the full signature (r, s, v)
	fullSignature := append(signature[:64], v.Bytes()...)

	// fmt.Printf("Full Signature hex : 0x%x\n", fullSignature)

	return fullSignature
}

func GetPaddedValue(value []byte) [32]byte {
	var paddedValue [32]byte
	startIndex := len(paddedValue) - len(value)
	copy(paddedValue[startIndex:], value)

	return paddedValue
}

func GetPublicAddressFromPrivateKey(privateKey *ecdsa.PrivateKey) common.Address {
	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA)
}

func PrepareTransactionOptions(client *ethclient.Client, chainId big.Int, gasLimit uint64, privateKey *ecdsa.PrivateKey) *bind.TransactOpts {

	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}

	senderAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), senderAddress)
	CheckError(err, "Pending nonce calculation failed")

	gasPrice, err := client.SuggestGasPrice(context.Background())
	CheckError(err, "Gas price calculation failed")

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, &chainId)
	CheckError(err, "Transactor creation failed")

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = gasLimit   // in units
	auth.GasPrice = gasPrice

	return auth
}

func CheckError(err error, description string) {
	if err != nil {
		log.Fatalln(description, ": ", err)
	}
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
		log.Fatalln("Transaction submitted successfully but failed to execute!")
	}
}
