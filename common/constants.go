package wc_common

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	DefaultAdminConfig   string  = "config/admin-config.json"
	DefaultOpL1Config    string  = "config/l1-operator-config.json"
	DefaultOpL2Config    string  = "config/l2-operator-config.json"
	EncryptedDirName     string  = ".encrypted_keys"
	DecryptedDirName     string  = ".decrypted_keys"
	GoCryptFSConfigName  string  = EncryptedDirName + "/gocryptfs.conf"
	MinEntropyBits       float64 = 50
	MaxMountRetries      int     = 5
	RetryPeriodInSeconds uint    = 1
)

type ChainConfig struct{
	OperatorRegistryAddress common.Address
	WitnessHubAddress       common.Address
	AVSDirectoryAddress     common.Address
	ChainID                 big.Int
	BlockExplorer           string
}

var BlueOrangutan = ChainConfig {
	OperatorRegistryAddress: common.HexToAddress("0x26710e60A36Ace8A44e1C3D7B33dc8B80eAb6cb7"),
	ChainID: *big.NewInt(1237146866),
	BlockExplorer: "https://blue-orangutan-blockscout.eu-north-2.gateway.fm",
}

var Holesky = ChainConfig {
	OperatorRegistryAddress: common.HexToAddress("0x708CBDDdab358c1fa8efB82c75bB4a116F316Def"),
	WitnessHubAddress: common.HexToAddress("0xa987EC494b13b21A8a124F8Ac03c9F530648C87D"),
	AVSDirectoryAddress: common.HexToAddress("0x055733000064333CaDDbC92763c58BF0192fFeBf"),
	ChainID: *big.NewInt(17000),
	BlockExplorer: "https://holesky.etherscan.io",
}

var NetworkConfig = map[string] ChainConfig {
	"1237146866": BlueOrangutan,
	"17000": Holesky,
}



