package operator_commands

import (
	"math/big"

	"github.com/witnesschain-com/diligencewatchtower-client/keystore"
	wc_common "github.com/witnesschain-com/operator-cli/common"
	"github.com/witnesschain-com/operator-cli/common/bindings/AvsDirectory"
	"github.com/witnesschain-com/operator-cli/common/bindings/OperatorRegistry"
	"github.com/witnesschain-com/operator-cli/common/bindings/WitnessHub"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func GetOpertorSignature(client *ethclient.Client, avsDirectory *AvsDirectory.AvsDirectory, witnessHubAddress common.Address, vault *keystore.Vault, operatorAddress common.Address, expiry *big.Int) WitnessHub.ISignatureUtilsSignatureWithSaltAndExpiry {
	salt := wc_common.GenerateSalt()

	//ON AVS DIRECTORY
	digestHash, err := avsDirectory.CalculateOperatorAVSRegistrationDigestHash(&bind.CallOpts{}, operatorAddress, witnessHubAddress, salt, expiry)
	wc_common.CheckError(err, "Digest hash calculation failed")

	signature, err := vault.SignData(digestHash[:], apitypes.DataTyped.Mime)

	wc_common.CheckError(err, "Signing the digest hash failed")

	operatorSignature := WitnessHub.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: signature,
		Salt:      salt,
		Expiry:    expiry,
	}

	return operatorSignature
}

func SignOperatorAddress(client *ethclient.Client, operatorRegistry *OperatorRegistry.OperatorRegistry, vault *keystore.Vault, OperatorAddress common.Address, salt [32]byte, expiry *big.Int) []byte {
	digestHash, err := operatorRegistry.CalculateWatchtowerRegistrationMessageHash(&bind.CallOpts{}, OperatorAddress, salt, expiry)
	fullSignature, err := vault.SignData(digestHash[:], apitypes.DataTyped.Mime)
	wc_common.CheckError(err, "unable to sign operator address")
	return fullSignature
}
