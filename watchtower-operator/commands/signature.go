package operator_commands

import (
	"crypto/ecdsa"
	"math/big"

	wc_common "github.com/witnesschain-com/operator-cli/common"
	"github.com/witnesschain-com/operator-cli/common/bindings/AvsDirectory"
	"github.com/witnesschain-com/operator-cli/common/bindings/WitnessHub"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetOpertorSignature(client *ethclient.Client, avsDirectory *AvsDirectory.AvsDirectory, witnessHubAddress common.Address, operatorPrivateKey *ecdsa.PrivateKey, operatorAddress common.Address, expiry *big.Int) WitnessHub.ISignatureUtilsSignatureWithSaltAndExpiry {
	salt := wc_common.GenerateSalt()

	//ON AVS DIRECTORY
	digestHash, err := avsDirectory.CalculateOperatorAVSRegistrationDigestHash(&bind.CallOpts{}, operatorAddress, witnessHubAddress, salt, expiry)
	wc_common.CheckError(err, "Digest hash calculation failed")

	signature, err := crypto.Sign(digestHash[:], operatorPrivateKey)
	wc_common.CheckError(err, "Signing the digest hash failed")

	v := new(big.Int).SetBytes(signature[64:])
	v.Add(v, big.NewInt(27))

	// Construct the full signature (r, s, v)
	fullSignature := append(signature[:64], v.Bytes()...)
	operatorSignature := WitnessHub.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: fullSignature,
		Salt:      salt,
		Expiry:    expiry,
	}

	return operatorSignature
}

func SignOperatorAddress(client *ethclient.Client, privateKey *ecdsa.PrivateKey, OperatorAddress common.Address, expiry big.Int) []byte {
	paddedAddr := wc_common.GetPaddedValue(OperatorAddress.Bytes())
	paddedExpiry := wc_common.GetPaddedValue(expiry.Bytes())

	encodedData := append(paddedAddr[:], paddedExpiry[:]...)
	hashedMessage := crypto.Keccak256(encodedData)
	hashedEthMessage := crypto.Keccak256(append([]byte("\x19Ethereum Signed Message:\n32"), hashedMessage...))

	signature, err := crypto.Sign(hashedEthMessage, privateKey)
	wc_common.CheckError(err, "Signing operator address failed")

	v := new(big.Int).SetBytes(signature[64:])
	v.Add(v, big.NewInt(27))

	// Construct the full signature (r, s, v)
	fullSignature := append(signature[:64], v.Bytes()...)
	return fullSignature
}
