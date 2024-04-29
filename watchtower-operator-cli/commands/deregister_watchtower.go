package operator_commands

import (
	"fmt"
	wc_common "operator-cli/common"
	"operator-cli/common/bindings/OperatorRegistry"
	operator_config "operator-cli/watchtower-operator-cli/config"

	"github.com/urfave/cli/v2"
)

func DeRegisterWatchtowerCmd() *cli.Command {
	wc_common.ConfigPathFlag.Value = wc_common.DefaultOpL2Config
	var deregisterWatchtowerCmd = &cli.Command{
		Name:  "deRegisterWatchtower",
		Usage: "De-register the watchtower",
		Flags: []cli.Flag{
			&wc_common.ConfigPathFlag,
		},
		Action: func(cCtx *cli.Context) error {
			config := operator_config.GetConfigFromContext(cCtx)
			DeRegisterWatchtower(config)
			return nil
		},
	}
	return deregisterWatchtowerCmd
}

func DeRegisterWatchtower(config *operator_config.OperatorConfig) {
	client := wc_common.ConnectToUrl(config.EthRPCUrl)

	operatorRegistry, err := OperatorRegistry.NewOperatorRegistry(config.OperatorRegistryAddress, client)
	wc_common.CheckError(err, "Instantiating OperatorRegistry contract failed")

	operatorPrivateKey, operatorAddress := wc_common.GetECDSAPrivateAndPublicKey(wc_common.GetPrivateKey(config.OperatorPrivateKey))

	if !wc_common.IsOperatorWhitelisted(operatorAddress, operatorRegistry) {
		fmt.Printf("Operator %s is not whitelisted\n", operatorAddress.Hex())
		return
	}

	deRegTransactOpts := wc_common.PrepareTransactionOptions(client, config.ChainId, config.GasLimit, operatorPrivateKey)

	for _, watchTowerPkName := range config.WatchtowerPrivateKeys {

		_, watchtowerAddress := wc_common.GetECDSAPrivateAndPublicKey(wc_common.GetPrivateKey(watchTowerPkName))

		if !wc_common.IsWatchtowerRegistered(watchtowerAddress, operatorRegistry) {
			fmt.Printf("Watchtower %s is not registered\n", watchtowerAddress.Hex())
			continue
		}

		deRegTx, err := operatorRegistry.DeRegister(deRegTransactOpts, watchtowerAddress)
		wc_common.CheckError(err, "Deregister failed")
		fmt.Printf("Tx sent: %s\n", deRegTx.Hash().Hex())
		wc_common.WaitForTransactionReceipt(client, deRegTx, config.TxReceiptTimeout)
	}
}
