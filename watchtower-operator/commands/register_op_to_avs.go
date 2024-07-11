package operator_commands

import (
	"fmt"

	"github.com/witnesschain-com/diligencewatchtower-client/keystore"
	wc_common "github.com/witnesschain-com/operator-cli/common"
	"github.com/witnesschain-com/operator-cli/common/bindings/AvsDirectory"
	"github.com/witnesschain-com/operator-cli/common/bindings/OperatorRegistry"
	"github.com/witnesschain-com/operator-cli/common/bindings/WitnessHub"
	operator_config "github.com/witnesschain-com/operator-cli/watchtower-operator/config"

	"github.com/urfave/cli/v2"
)

func RegisterOperatorToAVSCmd() *cli.Command {
	wc_common.ConfigPathFlag.Value = wc_common.DefaultOpL1Config
	var registerOperatorToAVSCmd = &cli.Command{
		Name:  "registerOperatorToAVS",
		Usage: "Register the operator to AVS",
		Flags: []cli.Flag{
			&wc_common.ConfigPathFlag,
		},
		Action: func(cCtx *cli.Context) error {
			config := operator_config.GetConfigFromContext(cCtx)
			RegisterOperatorToAVS(config)
			return nil
		},
	}
	return registerOperatorToAVSCmd
}

func RegisterOperatorToAVS(config *operator_config.OperatorConfig) {
	client := wc_common.ConnectToUrl(config.EthRPCUrl)

	operatorRegistry, err := OperatorRegistry.NewOperatorRegistry(wc_common.NetworkConfig[config.ChainID.String()].OperatorRegistryAddress, client)
	wc_common.CheckError(err, "Instantiating OperatorRegistry contract failed")


	if !wc_common.IsOperatorWhitelisted(config.OperatorAddress, operatorRegistry) {
		fmt.Printf("Operator %s is not whitelisted\n", config.OperatorAddress.Hex())
		return
	}

	avsDirectory, err := AvsDirectory.NewAvsDirectory(wc_common.NetworkConfig[config.ChainID.String()].AVSDirectoryAddress, client)
	wc_common.CheckError(err, "Instantiating AvsDirectory contract failed")

	if wc_common.IsOperatorRegistered(wc_common.NetworkConfig[config.ChainID.String()].WitnessHubAddress, config.OperatorAddress, avsDirectory) {
		fmt.Printf("Operator %s is already registered\n", config.OperatorAddress.Hex())
		return
	}

	witnessHub, err := WitnessHub.NewWitnessHub(wc_common.NetworkConfig[config.ChainID.String()].WitnessHubAddress, client)
	wc_common.CheckError(err, "Instantiating WitnessHub contract failed")

	expiry := wc_common.CalculateExpiry(client, config.ExpiryInDays)
	vc := &keystore.VaultConfig{Address: config.OperatorAddress, PrivateKey: config.OperatorPrivateKey, GocryptfsKey: config.OperatorEncryptedKey, Endpoint: config.Endpoint}
	operatorVault, err := keystore.SetupVault(vc)
	wc_common.CheckError(err, "unable to setup operator Vault: " + vc.Address.Hex())
	operatorSignature := GetOpertorSignature(client, avsDirectory, wc_common.NetworkConfig[config.ChainID.String()].WitnessHubAddress, operatorVault, config.OperatorAddress, expiry)

	transactOpts := operatorVault.NewTransactOpts(config.ChainID)

	tx, err := witnessHub.RegisterOperatorToAVS(transactOpts, config.OperatorAddress, operatorSignature)
	wc_common.CheckError(err, "Registering operator to AVS failed")

	fmt.Printf("Tx sent: %s\n", tx.Hash().Hex())

	wc_common.WaitForTransactionReceipt(client, tx, config.TxReceiptTimeout)
}
