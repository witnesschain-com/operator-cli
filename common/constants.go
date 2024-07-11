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
}

var BlueOrangutan = ChainConfig {
	OperatorRegistryAddress: common.HexToAddress("0x26710e60A36Ace8A44e1C3D7B33dc8B80eAb6cb7"),
	WitnessHubAddress: common.HexToAddress("0xD25c2c5802198CB8541987b73A8db4c9BCaE5cC7"),
	AVSDirectoryAddress: common.HexToAddress("0x135dda560e946695d6f155dacafc6f1f25c1f5af"),
	ChainID: *big.NewInt(1237146866),
}

var Holesky = ChainConfig {
	OperatorRegistryAddress: common.HexToAddress("0x708CBDDdab358c1fa8efB82c75bB4a116F316Def"),
	WitnessHubAddress: common.HexToAddress("0xa987EC494b13b21A8a124F8Ac03c9F530648C87D"),
	AVSDirectoryAddress: common.HexToAddress("0xc697aF7aF3C8c5Ce614d8BC0F252377233D9588a"),
	ChainID: *big.NewInt(17000),
}

var NetworkConfig = map[string] ChainConfig {
	"1237146866": BlueOrangutan,
	"17000": Holesky,
}



