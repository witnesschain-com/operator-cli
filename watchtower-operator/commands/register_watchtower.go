package operator_commands

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

	operatorRegistry, err := OperatorRegistry.NewOperatorRegistry(wc_common.NetworkConfig[chainID.String()].OperatorRegistryAddress, client)
	wc_common.CheckError(err, "Instantiating OperatorRegistry contract failed")


	// operatorVault
	var operatorPrivateKey *ecdsa.PrivateKey
	if len(config.OperatorPrivateKey) != 0 {
		operatorPrivateKey, err = crypto.HexToECDSA(config.OperatorPrivateKey)
		wc_common.CheckError(err, "unable to import operator privateKey")
	}

	fmt.Println("keystore args:", config.OperatorAddress, chainID, operatorPrivateKey, config.Endpoint)
	
	operatorVault, err := keystore.SetupVault(config.OperatorAddress, chainID, operatorPrivateKey, config.Endpoint)
	if err != nil {
		wc_common.CheckError(err, "unable to setup vault")
	}

	if !wc_common.IsOperatorWhitelisted(config.OperatorAddress, operatorRegistry) {
		fmt.Printf("Operator %s is not whitelisted\n", config.OperatorAddress.Hex())
		return
	}

	// regTransactOpts := wc_common.PrepareTransactionOptions(client, config.ChainId, config.GasLimit, operatorPrivateKey)
	transactOpts := operatorVault.NewTransactOpts(chainID)

	expiry := wc_common.CalculateExpiry(client, config.ExpiryInDays)

	for i, watchtowerAddressHex := range config.WatchtowerAddresses {
		watchtowerAddress := common.HexToAddress(watchtowerAddressHex)

		// watchtowerVault
		// watchtowerPrivateKey, watchtowerAddress := wc_common.GetECDSAPrivateAndPublicKey(wc_common.GetPrivateKey(watchTowerPkName))
		var privKey *ecdsa.PrivateKey
		if len(config.WatchtowerPrivateKeys) != 0 {
			privKey, err = crypto.HexToECDSA(config.WatchtowerPrivateKeys[i])
			wc_common.CheckError(err, "unable to import PrivateKey")
		}
		watchtowerVault, err := keystore.SetupVault(watchtowerAddress, chainID, privKey, config.Endpoint)
		wc_common.CheckError(err, "unable to setup watchtower vault")

		if wc_common.IsWatchtowerRegistered(watchtowerAddress, operatorRegistry) {
			fmt.Printf("Watchtower %s is already registered\n", watchtowerAddress.Hex())
			continue
		}

		signedMessage := SignOperatorAddress(client, watchtowerVault, config.OperatorAddress, expiry)

		regTx, err := operatorRegistry.RegisterWatchtowerAsOperator(transactOpts, watchtowerAddress, expiry, signedMessage)
		wc_common.CheckError(err, "Registering watchtower as operator failed")
		fmt.Printf("Tx sent: %s\n", regTx.Hash().Hex())
		wc_common.WaitForTransactionReceipt(client, regTx, config.TxReceiptTimeout)
	}
}
