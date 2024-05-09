package wc_common

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"os/exec"

	"github.com/witnesschain-com/operator-cli/common/bindings/AvsDirectory"
	"github.com/witnesschain-com/operator-cli/common/bindings/OperatorRegistry"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var isMounted bool = false

func ConnectToUrl(url string) *ethclient.Client {
	client, err := ethclient.Dial(url)
	CheckError(err, "Connection to RPC failed")

	id, _ := client.ChainID(context.Background())
	fmt.Println("Connection successful : ", id)

	return client
}

func GetECDSAPrivateKey(privateKeyString string) *ecdsa.PrivateKey {
	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKeyString)
	CheckError(err, "Converting private key to ECDSA format failed")
	return ecdsaPrivateKey
}

func GetPublicAddressFromPrivateKey(privateKey *ecdsa.PrivateKey) common.Address {
	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		FatalError("Error casting public key to ECDSA")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA)
}

func GetECDSAPrivateAndPublicKey(privateKeyString string) (*ecdsa.PrivateKey, common.Address) {
	ecdsaPrivateKey := GetECDSAPrivateKey(privateKeyString)
	ecdsaPublicKey := GetPublicAddressFromPrivateKey(ecdsaPrivateKey)
	return ecdsaPrivateKey, ecdsaPublicKey
}

func GetPaddedValue(value []byte) [32]byte {
	var paddedValue [32]byte
	startIndex := len(paddedValue) - len(value)
	copy(paddedValue[startIndex:], value)

	return paddedValue
}

func CalculateExpiry(client *ethclient.Client, expectedExpiryDays int64) *big.Int {
	// Get the latest block header
	header, err := client.HeaderByNumber(context.Background(), nil)
	CheckError(err, "Could not get HeaderByNumber")

	// Get the current timestamp from the latest block header
	currentTimestamp := big.NewInt(int64(header.Time))

	expiryInSeconds := expectedExpiryDays * 24 * 60 * 60
	timeToElapse := big.NewInt(expiryInSeconds)

	expiry := new(big.Int).Add(currentTimestamp, timeToElapse)
	return expiry
}

func GenerateSalt() [32]byte {
	var salt [32]byte

	// Generate random bytes
	_, err := rand.Read(salt[:])
	CheckError(err, "Generating salt failed")

	return salt
}

func ValidateAndMount() {
	if !ValidEncryptedDir() {
		FatalErrorWithoutUnmount(fmt.Sprintf("%v: %s\n", ErrInvalidEncryptedDirectory,
			" : check if "+GoCryptFSConfig+" and "+GoCryptFSDiriv+" exist. Or try initiating again after deleting those directories"))
	}

	Mount()
}

func Mount() {
	mountCmd := exec.Command("gocryptfs", EncryptedDir, DecryptedDir)
	RunCommandWithPassword(mountCmd, "mount", true)

	isMounted = true
}

func Unmount() {
	if !isMounted {
		return
	}

	umountCmd := exec.Command("fusermount", "-u", DecryptedDir)
	err := umountCmd.Run()
	if err != nil {
		CheckErrorWithoutUnmount(err, "Error unmounting GoCryptFS filesystem")
	}

	isMounted = false
}

func DirectoryExists(path string) bool {
	fileInfo, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	CheckError(err, "Error checking directory")

	if fileInfo.Mode().IsRegular() {
		CheckError(ErrNotADirectory, path)
	}

	return true
}

func CreateDirectory(path string) {
	fmt.Println("Creating directory: ", path)
	err := os.Mkdir(path, 0755)
	CheckError(err, "Error creating directory")
}

func RunCommandWithPassword(cmd *exec.Cmd, desc string, insecure bool) {
	fmt.Printf("Enter password to %s: ", desc)
	password := ReadHiddenInput()

	if !insecure {
		ValidatePassword(password)
	}

	cmdStdin, err := cmd.StdinPipe()
	CheckError(err, "Error creating stdin pipe for "+desc)

	err = cmd.Start()
	CheckError(err, "Error starting command for "+desc)

	_, err = cmdStdin.Write([]byte(password))
	CheckError(err, "Error writing to command stdin for "+desc)

	err = cmdStdin.Close()
	CheckError(err, "Error closing command stdin for "+desc)

	err = cmd.Wait()
	CheckError(err, "Command failed for "+desc)
}

func IsWatchtowerRegistered(watchtower common.Address, operatorRegistry *OperatorRegistry.OperatorRegistry) bool {
	registered, err := operatorRegistry.IsValidWatchtower(&bind.CallOpts{}, watchtower)
	CheckError(err, "Error checking if watchtower is already registered")
	return registered
}

func IsOperatorWhitelisted(operator common.Address, operatorRegistry *OperatorRegistry.OperatorRegistry) bool {
	active, err := operatorRegistry.IsActiveOperator(&bind.CallOpts{}, operator)
	CheckError(err, "Error checking if operator is whitelisted")
	return active
}

func IsOperatorRegistered(witnessHubAddress common.Address, operator common.Address, avsDirectory *AvsDirectory.AvsDirectory) bool {
	status, err := avsDirectory.AvsOperatorStatus(&bind.CallOpts{}, witnessHubAddress, operator)
	CheckError(err, "Checking operator status failed")
	return status != 0
}
