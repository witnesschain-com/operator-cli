package operator_config

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"

	wc_common "github.com/witnesschain-com/operator-cli/common"
)

type OperatorConfig struct {
	WatchtowerPrivateKeysHex   []string       `json:"watchtower_private_keys"`
	WatchtowerAddresses     []common.Address       `json:"watchtower_addresses"`
	WatchtowerEncryptedKeys []string       `json:"watchtower_encrypted_keys"`
	OperatorPrivateKeyHex      string         `json:"operator_private_key"`
	OperatorAddress         common.Address `json:"operator_address"`
	OperatorEncryptedKey    string         `json:"operator_encrypted_key"`
	EthRPCUrl               string         `json:"eth_rpc_url"`
	GasLimit                uint64         `json:"gas_limit"`
	TxReceiptTimeout        int64          `json:"tx_receipt_timeout"`
	ExpiryInDays            int64          `json:"expiry_in_days"`
	Endpoint                string         `json:"external_signer_endpoint"`
	WatchtowerPrivateKeys   []*ecdsa.PrivateKey
	OperatorPrivateKey       *ecdsa.PrivateKey
	ChainID                  *big.Int

}

func GetConfigFromContext(cCtx *cli.Context) *OperatorConfig {
	configFilePath := cCtx.String("config-file")
	fmt.Printf("Using config file path : %s\n", configFilePath)

	data, err := os.ReadFile(configFilePath)
	wc_common.CheckError(err, "Error reading json file")

	// Parse the json data into a struct
	var config OperatorConfig = OperatorConfig{ExpiryInDays: 1, TxReceiptTimeout: 300, GasLimit: 300000}
	err = json.Unmarshal(data, &config)
	wc_common.CheckError(err, "Error unmarshaling json data")

	if len(config.WatchtowerEncryptedKeys) != 0 {
		// get the path from the first key, as others should be same
		// will not work with different paths
		wc_common.RetryMounting()
		wc_common.ProcessConfigKeyPath(config.WatchtowerEncryptedKeys[0])
		wc_common.UseEncryptedKeys()
	}

	if len(config.WatchtowerPrivateKeysHex) != 0 {
		for _, privKey := range config.WatchtowerPrivateKeysHex{
			fmt.Println(privKey)
			key, err := crypto.HexToECDSA(privKey)
			wc_common.CheckError(err, "unable to convert watchtower privatekey")
			config.WatchtowerAddresses = append(config.WatchtowerAddresses, crypto.PubkeyToAddress(key.PublicKey))
			config.WatchtowerPrivateKeys = append(config.WatchtowerPrivateKeys, key)
		}
	}

	if len(config.WatchtowerEncryptedKeys) != 0 {
		for _, keyPath := range config.WatchtowerEncryptedKeys{
			privKey, err := wc_common.LoadPrivateKey(keyPath)
			wc_common.CheckError(err, "unable to load encrypted keys")

			config.WatchtowerPrivateKeys = append(config.WatchtowerPrivateKeys, privKey)
			config.WatchtowerAddresses = append(config.WatchtowerAddresses, crypto.PubkeyToAddress(privKey.PublicKey))
		}
	}

	if len(config.OperatorEncryptedKey) != 0 {
		priv, err := wc_common.LoadPrivateKey(config.OperatorEncryptedKey)
		if err != nil {
			log.Fatal("unable to retive operator privateKey")
		}
		config.OperatorAddress = crypto.PubkeyToAddress(priv.PublicKey)
		config.OperatorPrivateKey = priv
	}

	if len(config.OperatorPrivateKeyHex) != 0 {
		priv, err := crypto.HexToECDSA(config.OperatorPrivateKeyHex)
		wc_common.CheckError(err, "unable to convert privateKey")
		config.OperatorAddress =  crypto.PubkeyToAddress(priv.PublicKey)
		config.OperatorPrivateKey = priv
	}

	client := wc_common.ConnectToUrl(config.EthRPCUrl)
	wc_common.CheckError(err, "unable to connect to RPC")
	chainID, err := client.ChainID(context.Background())
	wc_common.CheckError(err, "unable to get chainID")
	config.ChainID = chainID

	if config.OperatorAddress.Cmp(common.Address{0}) == 0 {
		panic("operatorAddress is zero")
	}
	
	return &config
}
