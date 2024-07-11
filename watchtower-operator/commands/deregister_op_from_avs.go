package operator_commands

import (
	"fmt"

	"github.com/witnesschain-com/diligencewatchtower-client/keystore"
	wc_common "github.com/witnesschain-com/operator-cli/common"
	"github.com/witnesschain-com/operator-cli/common/bindings/AvsDirectory"
	"github.com/witnesschain-com/operator-cli/common/bindings/WitnessHub"
	operator_config "github.com/witnesschain-com/operator-cli/watchtower-operator/config"
	"github.com/witnesschain-com/operator-cli/common/bindings/OperatorRegistry"

	"github.com/urfave/cli/v2"
)

func DeRegisterOperatorFromAVSCmd() *cli.Command {
	wc_common.ConfigPathFlag.Value = wc_common.DefaultOpL1Config
	var deregisterOperatorFromAVSCmd = &cli.Command{
		Name:  "deRegisterOperatorFromAVS",
		Usage: "De-register the operator from AVS",
		Flags: []cli.Flag{
			&wc_common.ConfigPathFlag,
		},
		Action: func(cCtx *cli.Context) error {
			config := operator_config.GetConfigFromContext(cCtx)
			DeRegisterOperatorFromAVS(config)
			return nil
		},
	}
	return deregisterOperatorFromAVSCmd
}

func DeRegisterOperatorFromAVS(config *operator_config.OperatorConfig) {
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

	vc := &keystore.VaultConfig{Address: config.OperatorAddress, PrivateKey: config.OperatorPrivateKey, GocryptfsKey: config.OperatorEncryptedKey, Endpoint: config.Endpoint}
	operatorVault, err := keystore.SetupVault(vc)
	wc_common.CheckError(err, "unable to setup operator Vault: " + vc.Address.Hex())

	transactOpts := operatorVault.NewTransactOpts(config.ChainID)

	tx, err := witnessHub.DeregisterOperatorFromAVS(transactOpts, config.OperatorAddress)
	wc_common.CheckError(err, "Registering operator to AVS failed")

	fmt.Printf("Tx sent: %s\n", tx.Hash().Hex())

	wc_common.WaitForTransactionReceipt(client, tx, config.TxReceiptTimeout)
}
