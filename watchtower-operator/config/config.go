package operator_config

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	wc_common "github.com/witnesschain-com/operator-cli/common"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

type OperatorConfig struct {
	WatchtowerPrivateKeys   []string       `json:"watchtower_private_keys"`
	WatchtowerAddresses     []string       `json:"watchtower_addresses"`
	WatchtowerEncryptedKeys []string       `json:"watchtower_encrypted_keys"`
	OperatorPrivateKey      string         `json:"operator_private_key"`
	OperatorAddress         common.Address `json:"operator_address"`
	OperatorEncryptedKey    string         `json:"operator_encrypted_key"`
	OperatorRegistryAddress common.Address `json:"operator_registry_address"`
	WitnessHubAddress       common.Address `json:"witnesshub_address"`
	AvsDirectoryAddress     common.Address `json:"avs_directory_address"`
	EthRPCUrl               string         `json:"eth_rpc_url"`
	ChainId                 big.Int        `json:"chain_id"`
	GasLimit                uint64         `json:"gas_limit"`
	TxReceiptTimeout        int64          `json:"tx_receipt_timeout"`
	ExpiryInDays            int64          `json:"expiry_in_days"`
	UseEncryptedKeys        bool           `json:"use_encrypted_keys"`
	Endpoint                string         `json:"external_signer_endpoint"`
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

	if len(config.OperatorPrivateKey) != 0 {
		priv, err := crypto.HexToECDSA(config.OperatorPrivateKey)
		wc_common.CheckError(err, "unable to convert privateKey")
		config.OperatorAddress =  crypto.PubkeyToAddress(priv.PublicKey)
	}

	if len(config.WatchtowerPrivateKeys) != 0 {
		fmt.Println(config.WatchtowerAddresses)
		for _, privKey := range config.WatchtowerPrivateKeys{
			fmt.Println(privKey)
			key, err := crypto.HexToECDSA(privKey)
			wc_common.CheckError(err, "unable to convert watchtower privatekey")
			config.WatchtowerAddresses = append(config.WatchtowerAddresses, crypto.PubkeyToAddress(key.PublicKey).Hex())
		}
	}

	return &config
}
