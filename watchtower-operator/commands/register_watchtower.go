package operator_commands

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"github.com/witnesschain-com/diligencewatchtower-client/keystore"
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
	chainID, err := client.ChainID(context.Background())
	wc_common.CheckError(err, "failed to retrive chain ID")
	fmt.Println("chainID: " + chainID.String())

	operatorRegistry, err := OperatorRegistry.NewOperatorRegistry(wc_common.NetworkConfig[chainID.String()].OperatorRegistryAddress, client)
	wc_common.CheckError(err, "Instantiating OperatorRegistry contract failed")

	vc := &keystore.VaultConfig{Address: config.OperatorAddress, ChainID: chainID, PrivateKey: config.OperatorPrivateKey, Endpoint: config.Endpoint}
	operatorVault, err := keystore.SetupVault(vc)
	if err != nil {
		wc_common.CheckError(err, "unable to setup vault")
	}

	if !wc_common.IsOperatorWhitelisted(config.OperatorAddress, operatorRegistry) {
		fmt.Printf("Operator %s is not whitelisted\n", config.OperatorAddress.Hex())
		return
	}

	transactOpts := operatorVault.NewTransactOpts(chainID)

	expiry := wc_common.CalculateExpiry(client, config.ExpiryInDays)

	for i, watchtowerAddress := range config.WatchtowerAddresses {
		fmt.Println("watchtowerAddress: " + watchtowerAddress.Hex())

		var watchtowerPrivateKey *ecdsa.PrivateKey
		if len(config.WatchtowerPrivateKeys) != 0{
			watchtowerPrivateKey = config.WatchtowerPrivateKeys[i]
		}

		var gocryptfsKey string
		if len(config.WatchtowerEncryptedKeys) != 0 {
			gocryptfsKey = config.WatchtowerEncryptedKeys[i]
		}

		vc := &keystore.VaultConfig{Address: watchtowerAddress, ChainID: chainID, PrivateKey: watchtowerPrivateKey, Endpoint: config.Endpoint, GocryptfsKey: gocryptfsKey}
		watchtowerVault, err := keystore.SetupVault(vc)
		wc_common.CheckError(err, "unable to setup watchtower vault")

		if wc_common.IsWatchtowerRegistered(watchtowerAddress, operatorRegistry) {
			fmt.Printf("Watchtower %s is already registered\n", watchtowerAddress.Hex())
			continue
		}

		salt := wc_common.GenerateSalt()
		signedMessage := SignOperatorAddress(client, operatorRegistry, watchtowerVault, config.OperatorAddress, salt, expiry)
		regTx, err := operatorRegistry.RegisterWatchtowerAsOperator(transactOpts, watchtowerAddress, salt, expiry, signedMessage)
		wc_common.CheckError(err, "Registering watchtower as operator failed")
		fmt.Printf("Tx sent: %s/tx/%s\n", wc_common.NetworkConfig[chainID.String()].BlockExplorer, regTx.Hash().Hex())
		wc_common.WaitForTransactionReceipt(client, regTx, config.TxReceiptTimeout)
	}
}
