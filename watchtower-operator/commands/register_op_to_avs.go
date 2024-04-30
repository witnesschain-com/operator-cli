package operator_commands

import (
	"fmt"
	wc_common "operator-cli/common"
	"operator-cli/common/bindings/AvsDirectory"
	"operator-cli/common/bindings/OperatorRegistry"
	"operator-cli/common/bindings/WitnessHub"
	operator_config "operator-cli/watchtower-operator/config"

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

	operatorRegistry, err := OperatorRegistry.NewOperatorRegistry(config.OperatorRegistryAddress, client)
	wc_common.CheckError(err, "Instantiating OperatorRegistry contract failed")

	operatorPrivateKey, operatorAddress := wc_common.GetECDSAPrivateAndPublicKey(wc_common.GetPrivateKey(config.OperatorPrivateKey))
	wc_common.CheckError(err, "Converting private key to ECDSA format failed")

	if !wc_common.IsOperatorWhitelisted(operatorAddress, operatorRegistry) {
		fmt.Printf("Operator %s is not whitelisted\n", operatorAddress.Hex())
		return
	}

	avsDirectory, err := AvsDirectory.NewAvsDirectory(config.AvsDirectoryAddress, client)
	wc_common.CheckError(err, "Instantiating AvsDirectory contract failed")

	witnessHub, err := WitnessHub.NewWitnessHub(config.WitnessHubAddress, client)
	wc_common.CheckError(err, "Instantiating WitnessHub contract failed")

	expiry := wc_common.CalculateExpiry(client, config.ExpiryInDays)
	operatorSignature := GetOpertorSignature(client, avsDirectory, config.WitnessHubAddress, operatorPrivateKey, operatorAddress, expiry)

	avsRegtransactOpts := wc_common.PrepareTransactionOptions(client, config.ChainId, config.GasLimit, operatorPrivateKey)

	tx, err := witnessHub.RegisterOperatorToAVS(avsRegtransactOpts, operatorAddress, operatorSignature)
	wc_common.CheckError(err, "Registering operator to AVS failed")

	fmt.Printf("Tx sent: %s\n", tx.Hash().Hex())

	wc_common.WaitForTransactionReceipt(client, tx, config.TxReceiptTimeout)
}
