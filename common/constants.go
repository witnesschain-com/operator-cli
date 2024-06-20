package wc_common

const (
	DefaultAdminConfig  string  = "config/admin-config.json"
	DefaultOpL1Config   string  = "config/l1-operator-config.json"
	DefaultOpL2Config   string  = "config/l2-operator-config.json"
	EncryptedDirName    string  = ".encrypted_keys"
	DecryptedDirName    string  = ".decrypted_keys"
	GoCryptFSConfigName string  = EncryptedDirName + "/gocryptfs.conf"
	MinEntropyBits      float64 = 50
)
