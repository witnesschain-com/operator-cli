package operator_commands

import (
	"fmt"

	wc_common "github.com/witnesschain-com/operator-cli/common"
	"github.com/witnesschain-com/operator-cli/common/bindings/OperatorRegistry"
	operator_config "github.com/witnesschain-com/operator-cli/watchtower-operator/config"

	"github.com/urfave/cli/v2"
)

func RegisterWatchtowerCmd() *cli.Command {
	wc_common.ConfigPathFlag.Value = wc_common.DefaultOpL2Config
	var registerWatchtowerCmd = &cli.Command{
		Name:  "registerWatchtower",
		Usage: "Register a watchtower",
		Flags: []cli.Flag{
			&wc_common.ConfigPathFlag,
		},
		Action: func(cCtx *cli.Context) error {
			config := operator_config.GetConfigFromContext(cCtx)
			RegisterWatchtower(config)
			return nil
		},
	}
	return registerWatchtowerCmd
}

func RegisterWatchtower(config *operator_config.OperatorConfig) {
	client := wc_common.ConnectToUrl(config.EthRPCUrl)

	operatorRegistry, err := OperatorRegistry.NewOperatorRegistry(config.OperatorRegistryAddress, client)
	wc_common.CheckError(err, "Instantiating OperatorRegistry contract failed")

	operatorPrivateKey, operatorAddress := wc_common.GetECDSAPrivateAndPublicKey(wc_common.GetPrivateKey(config.OperatorPrivateKey))

	if !wc_common.IsOperatorWhitelisted(operatorAddress, operatorRegistry) {
		fmt.Printf("Operator %s is not whitelisted\n", operatorAddress.Hex())
		return
	}

	regTransactOpts := wc_common.PrepareTransactionOptions(client, config.ChainId, config.GasLimit, operatorPrivateKey)
	expiry := wc_common.CalculateExpiry(client, config.ExpiryInDays)

	for _, watchTowerPkName := range config.WatchtowerPrivateKeys {

		watchtowerPrivateKey, watchtowerAddress := wc_common.GetECDSAPrivateAndPublicKey(wc_common.GetPrivateKey(watchTowerPkName))

		if wc_common.IsWatchtowerRegistered(watchtowerAddress, operatorRegistry) {
			fmt.Printf("Watchtower %s is already registered\n", watchtowerAddress.Hex())
			continue
		}

		signedMessage := SignOperatorAddress(client, watchtowerPrivateKey, operatorAddress, *expiry)

		regTx, err := operatorRegistry.RegisterWatchtowerAsOperator(regTransactOpts, watchtowerAddress, expiry, signedMessage)
		wc_common.CheckError(err, "Registering watchtower as operator failed")
		fmt.Printf("Tx sent: %s\n", regTx.Hash().Hex())
		wc_common.WaitForTransactionReceipt(client, regTx, config.TxReceiptTimeout)
	}
}
