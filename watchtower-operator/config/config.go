package operator_config

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	wc_common "github.com/witnesschain-com/operator-cli/common"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

type OperatorConfig struct {
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
	UseEncryptedKeys        bool           `json:"use_encrypted_keys"`
}

func GetConfigFromContext(cCtx *cli.Context) *OperatorConfig {
	configFilePath := cCtx.String("config-file")
	fmt.Printf("Using config file path : %s\n", configFilePath)

	data, err := os.ReadFile(configFilePath)
	wc_common.CheckError(err, "Error reading json file")

	// Parse the json data into a struct
	var config OperatorConfig
	err = json.Unmarshal(data, &config)
	wc_common.CheckError(err, "Error unmarshaling json data")

	SetDefaultConfigValues(&config)

	if config.UseEncryptedKeys {
		wc_common.UseEncryptedKeys()
	}

	return &config
}

func SetDefaultConfigValues(config *OperatorConfig) {
	if config.ExpiryInDays == 0 {
		config.ExpiryInDays = 1 // 1 day
	}

	if config.TxReceiptTimeout == 0 {
		config.TxReceiptTimeout = 5 * 60 // 5 minutes
	}

	if config.GasLimit == 0 {
		config.GasLimit = 300000
	}
}
