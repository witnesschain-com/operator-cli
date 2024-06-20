package wc_common

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/urfave/cli/v2"
)

var useEncryptedKeys bool = false
var EncryptedDir string = EncryptedDirName
var DecryptedDir string = DecryptedDirName
var GoCryptFSConfig string = GoCryptFSConfigName

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

	if !DirectoryExists(EncryptedDir) {
		CreateDirectory(EncryptedDir)
	}

	if !DirectoryExists(DecryptedDir) {
		CreateDirectory(DecryptedDir)
	}

	InitGocryptfs(insecure)
}

func CreateKeyCmd(cCtx *cli.Context) {
	keyName := cCtx.String("key-name")
	err := ValidateKeyName(keyName)
	CheckError(err, "Error validating key name")

	CreateKey(keyName)
	fmt.Printf("Created key: %s\n", keyName)
}

func DeleteKeyCmd(cCtx *cli.Context) {
	keyName := cCtx.String("key-name")
	err := ValidateKeyName(keyName)
	CheckError(err, "Error validating key name")

	DeleteKey(keyName)

	fmt.Printf("Deleted key: %s\n", keyName)
}

func ListKeyCmd() {
	dir, err := os.Open(DecryptedDir)
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
	initCmd := exec.Command("gocryptfs", "-init", "-plaintextnames", EncryptedDir)

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
	keyFile := DecryptedDir + "/" + keyName
	privateKey := GetPrivateKeyFromUser()
	CreateKeyFileAndStoreKey(keyFile, privateKey)
}

func DeleteKey(keyName string) {
	keyFile := DecryptedDir + "/" + keyName
	err := os.Remove(keyFile)
	CheckError(err, "Error deleting key\n")
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
	_, err := os.Stat(GoCryptFSConfig)

	return !os.IsNotExist(err)
}

func GetPrivateKeyFromFile(keyName string) string {
	keyFile := DecryptedDir + "/" + keyName
	data, err := os.ReadFile(keyFile)
	CheckError(err, "Error reading key file")
	return string(data)
}

func UseEncryptedKeys() {
	useEncryptedKeys = true
	ValidateAndMount()
}

func SetKeysPath(keysPath string) {
	if len(keysPath) > 1 {
		lastChar := keysPath[len(keysPath)-1:]
		if lastChar == "/" {
			EncryptedDir = keysPath + EncryptedDirName
			DecryptedDir = keysPath + DecryptedDirName
			GoCryptFSConfig = keysPath + GoCryptFSConfigName
		} else {
			EncryptedDir = keysPath + "/" + EncryptedDirName
			DecryptedDir = keysPath + "/" + DecryptedDirName
			GoCryptFSConfig = keysPath + "/" + GoCryptFSConfigName
		}
	}
}

func GetPrivateKey(key string) string {
	if useEncryptedKeys {
		return GetPrivateKeyFromFile(key)
	}
	return key
}
