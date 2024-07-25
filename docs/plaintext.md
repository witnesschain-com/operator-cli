### Config file format for plaintext private keys in configuration file

Warning: storing keys in plaintext provide least security, therefore use 
this format at your own risk.

The below example shows how you can use the key names which will be taken from the default path
```
{
  "watchtower_private_keys": [
    "<raw-watchtower-private-key e.g. 0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef>"
  ],
  "operator_private_key": "<raw-operator-private-key e.g. abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789>",
  "eth_rpc_url": "<Mainnet RPC URL>",
}
```
