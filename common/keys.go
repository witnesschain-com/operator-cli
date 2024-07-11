package wc_common

import (
	"crypto/ecdsa"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

var m_useEncryptedKeys bool = false
var m_isFullPath bool = false
var m_retryMounting bool = false
var m_encryptedDir string = EncryptedDirName
var m_decryptedDir string = DecryptedDirName
var m_goCryptFSConfig string = GoCryptFSConfigName

func KeysCmd() *cli.Command {
	var keysCmd = &cli.Command{
		Name:  "keys",
		Usage: "Manage the operator's keys",
		Subcommands: []*cli.Command{
			InitCmd(),
			CreateCmd(),
			DeleteCmd(),
			ListCmd(),
		},
	}
	return keysCmd
}

func InitCmd() *cli.Command {
	var initCmd = &cli.Command{
		Name:      "init",
		Usage:     "init local keystore",
		UsageText: "init",
		Flags: []cli.Flag{
			&InsecureFlag,
		},
		Action: func(cCtx *cli.Context) error {
			CheckIfGocryptfsIsInstalled()
			InitKeyStore(cCtx)
			return nil
		},
	}
	return initCmd
}

func CreateCmd() *cli.Command {
	var createCmd = &cli.Command{
		Name:      "create",
		Usage:     "create encrypted key in local keystore",
		UsageText: "create <keyName>",
		Flags: []cli.Flag{
			&KeyNameFlag,
		},
		Action: func(cCtx *cli.Context) error {
			ValidateAndMount()
			CreateKeyCmd(cCtx)
			return nil
		},
	}
	return createCmd
}

func DeleteCmd() *cli.Command {
	var deleteCmd = &cli.Command{
		Name:      "delete",
		Usage:     "delete encrypted key from local keystore",
		UsageText: "delete <keyName>",
		Flags: []cli.Flag{
			&KeyNameFlag,
		},
		Action: func(cCtx *cli.Context) error {
			ValidateAndMount()
			DeleteKeyCmd(cCtx)
			return nil
		},
	}
	return deleteCmd
}

func ListCmd() *cli.Command {
	var listCmd = &cli.Command{
		Name:      "list",
		Usage:     "list all encrypted keys from local keystore",
		UsageText: "list",
		Action: func(cCtx *cli.Context) error {
			ValidateAndMount()
			ListKeyCmd()
			return nil
		},
	}
	return listCmd
}

func InitKeyStore(cCtx *cli.Context) {
	insecure := cCtx.Bool("insecure")

	if !DirectoryExists(m_encryptedDir) {
		CreateDirectory(m_encryptedDir)
	}

	if !DirectoryExists(m_decryptedDir) {
		CreateDirectory(m_decryptedDir)
	}

	InitGocryptfs(insecure)
}

func CreateKeyCmd(cCtx *cli.Context) {
	keyName := cCtx.String("key-name")
	err := ValidateKeyName(keyName)
	CheckError(err, "Error validating key name")

	CreateKey(keyName)
}

func DeleteKeyCmd(cCtx *cli.Context) {
	keyName := cCtx.String("key-name")
	err := ValidateKeyName(keyName)
	CheckError(err, "Error validating key name")

	DeleteKey(keyName)
}

func ListKeyCmd() {
	dir, err := os.Open(m_decryptedDir)
	CheckError(err, "Error opening directory")
	defer dir.Close()

	files, err := dir.Readdir(-1)
	CheckError(err, "Error reading directory")

	fmt.Printf("   " + strings.Repeat("-", 55) + "\n")
	fmt.Printf("   %-30s %-25s\n", "Name", "Created")
	fmt.Printf("   " + strings.Repeat("-", 55) + "\n")

	for _, file := range files {
		createdTime := file.ModTime().Format("02-01-2006 15:04:05")
		fmt.Printf("   %-30s %-25s\n", file.Name(), createdTime)
	}

	fmt.Printf("   " + strings.Repeat("-", 55) + "\n")
}

func InitGocryptfs(insecure bool) {
	initCmd := exec.Command("gocryptfs", "-init", "-plaintextnames", m_encryptedDir)

	RunCommandWithPassword(initCmd, "init", insecure)
}

func ValidateKeyName(keyName string) error {
	if len(keyName) == 0 {
		return ErrEmptyKeyName
	}

	if match, _ := regexp.MatchString("\\s", keyName); match {
		return ErrKeyContainsWhitespaces
	}
	return nil
}

func CreateKey(keyName string) {
	keyFile := m_decryptedDir + "/" + keyName

	_, err := os.Stat(keyFile)
	if !os.IsNotExist(err) {
		fmt.Printf("Key already exists, do you want to overwrite? (y/n): ")
		var response string
		fmt.Scanln(&response)

		if strings.ToLower(response) != "y" {
			return
		}
	}

	privateKey := GetPrivateKeyFromUser()
	CreateKeyFileAndStoreKey(keyFile, privateKey)

	fmt.Printf("Created key: %s\n", keyName)
}

func DeleteKey(keyName string) {
	keyFile := m_decryptedDir + "/" + keyName
	err := os.Remove(keyFile)
	CheckError(err, "Error deleting key\n")

	fmt.Printf("Deleted key: %s\n", keyName)
}

func GetPrivateKeyFromUser() string {
	fmt.Print("Enter private key: ")
	return ReadHiddenInput()
}

func CreateKeyFileAndStoreKey(keyFile string, privateKey string) {
	file, err := os.Create(keyFile)
	CheckError(err, "Error creating file")
	defer file.Close()

	_, err = file.WriteString(privateKey)
	CheckError(err, "Error writing to file")
}

func ValidEncryptedDir() bool {
	_, err := os.Stat(m_goCryptFSConfig)

	return !os.IsNotExist(err)
}

func GetPrivateKeyFromFile(keyName string) string {
	keyFile := m_decryptedDir + "/" + keyName
	data, err := os.ReadFile(keyFile)
	CheckError(err, "Error reading key file" + keyFile)
	return string(data)
}

func UseEncryptedKeys() {
	if m_useEncryptedKeys == true {
		return
	}
	m_useEncryptedKeys = true
	ValidateAndMount()
}

func RetryMounting() {
	m_retryMounting = true
}

func ProcessConfigKeyPath(keyPath string) {
	dir, file := filepath.Split(keyPath)

	if file == keyPath && dir == "." {
		fmt.Printf("Using the default key path : %s\n", m_encryptedDir)
		// this means they have given only the key name only,
		// do nothing and use default path
		return
	}

	// go to the grand parent directory of the key path to get the .encrypted_keys path
	parentPath := filepath.Dir(filepath.Dir(dir))

	m_encryptedDir = filepath.Join(parentPath, EncryptedDirName)
	m_decryptedDir = filepath.Join(parentPath, DecryptedDirName)
	m_goCryptFSConfig = filepath.Join(parentPath, GoCryptFSConfigName)
	m_isFullPath = true

	fmt.Printf("Using the key path : %s\n", m_encryptedDir)
}

func GetPrivateKey(key string) string {
	if m_useEncryptedKeys {
		keyName := key
		if m_isFullPath {
			_, keyName = filepath.Split(key)
		}
		return GetPrivateKeyFromFile(keyName)
	}
	return key
}

func LoadPrivateKey(path string) (*ecdsa.PrivateKey, error) {
	ProcessConfigKeyPath(path)

	fmt.Println("loading " + path)
	dir := filepath.Dir(path)
	KeyfileName := filepath.Base(path)
	m_encryptedDir = dir

	Mount()
	data, err := os.ReadFile(m_decryptedDir + "/" + KeyfileName)
	CheckError(err, "Error reading key file" + path)
	Unmount()

	priv, err := crypto.HexToECDSA(string(data))
	if err != nil {
		return nil, err
	}
	return priv, nil
}
