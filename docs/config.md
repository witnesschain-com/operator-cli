### Configuration options that can be set in config file
| Field | Description |
|----------|----------|
|watchtower_private_keys | Private keys of the watchtowers (use this field if you want to enter raw keys)|
|watchtower_encrypted_keys | Encrypted private keys of the watchtowers (use this field if you want to enter encrypted key names)|
|operator_private_key | Private key of the operator(on which the actions will be performed) (use this field if you want to enter raw key)|
|operator_encrypted_key | Encrypted private key of the operator(on which the actions will be performed) (use this field if you want to enter raw key)|
|encrypted_key_type | The type of encryption used for the keys (valid values = w3secretkeys/gocryptfs) |
|eth_rpc_url | The RPC URL where you want to perform the transactions |
|gas_limit | The gas limit you want to set while sending the transactions (Default value = 1000000). No need to add in the config unless you want to overwrite the default values.  |
|tx_receipt_timeout| Timeout in seconds for waiting of tx receipts (Default value = 300). No need to add in the config unless you want to overwrite the default values. |
|expiry| Expiry in days after which the operator signature becomes invalid (Default value = 1). No need to add in the config unless you want to overwrite the default values. |

