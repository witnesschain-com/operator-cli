package operator_commands

import (
	"fmt"

	wc_common "github.com/witnesschain-com/operator-cli/common"
	"github.com/witnesschain-com/operator-cli/common/bindings/AvsDirectory"
	"github.com/witnesschain-com/operator-cli/common/bindings/WitnessHub"
	operator_config "github.com/witnesschain-com/operator-cli/watchtower-operator/config"

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

	operatorPrivateKey, operatorAddress := wc_common.GetECDSAPrivateAndPublicKey(wc_common.GetPrivateKey(config.OperatorPrivateKey))
	avsDirectory, err := AvsDirectory.NewAvsDirectory(config.AvsDirectoryAddress, client)
	wc_common.CheckError(err, "Instantiating AvsDirectory contract failed")

	if !wc_common.IsOperatorRegistered(config.WitnessHubAddress, operatorAddress, avsDirectory) {
		fmt.Printf("Operator %s is not registered\n", operatorAddress.Hex())
		return
	}

	witnessHub, err := WitnessHub.NewWitnessHub(config.WitnessHubAddress, client)
	wc_common.CheckError(err, "Instantiating WitnessHub contract failed")

	avsRegtransactOpts := wc_common.PrepareTransactionOptions(client, config.ChainId, config.GasLimit, operatorPrivateKey)
	avsRegtransactOpts.Nonce = wc_common.GetLatestNonce(client, operatorPrivateKey)
	tx, err := witnessHub.DeregisterOperatorFromAVS(avsRegtransactOpts, operatorAddress)
	wc_common.CheckError(err, "Deregistering operator from AVS failed")

	fmt.Printf("Tx sent: %s\n", tx.Hash().Hex())

	wc_common.WaitForTransactionReceipt(client, tx, config.TxReceiptTimeout)
}
