package wc_common

const (
	DefaultAdminConfig string  = "config/admin-config.json"
	DefaultOpL1Config  string  = "config/l1-operator-config.json"
	DefaultOpL2Config  string  = "config/l2-operator-config.json"
	EncryptedDir       string  = ".encrypted_keys"
	DecryptedDir       string  = ".decrypted_keys"
	GoCryptFSConfig    string  = EncryptedDir + "/gocryptfs.conf"
	GoCryptFSDiriv     string  = EncryptedDir + "/gocryptfs.diriv"
	MinEntropyBits     float64 = 50
)
