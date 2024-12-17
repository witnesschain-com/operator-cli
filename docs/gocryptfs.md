# Gocryptfs keystore

One of the methods to store private keys of watchtowers and operators is 
gocryptfs. It store keys in encrypted format when the key is not in use 
(data at rest).

## Gocryptfs: store keys in encrypted format
Gocryptfs provide an encrypted file system. It uses FUSE (Filesystem in 
userspace) to mount encrypted directory which are decrypted when the 
files system is mounted.

### Pre-requisite
Install gocryptfs

```bash
sudo apt install gocryptfs
```

### Gocryptfs key management

```
watchtower-operator keys init -t gocryptfs
```
After this command, two directories `.encrypted_keys` and 
`.decrypted_keys` are created inside a directory `.gocryptfs`. The names 
indicate their functions. Once this is done, we don't need to do it 
again, unless the `.encrypted_keys` or `.decrypted_keys` are tampered 
with. Once the command is successfully run, all other actions to 
create/import/export/delete `gocryptfs` type will need this password.

The usage of `import`, `create`, `list` is similar to [web3 secret 
storage](../README.md). You need to pass key type `--key-type gocryptfs` 
with each commands.

Once, you have imported keys, create a `operator-config.json` with 
following template:-

The below example shows how you can use the key names which will be taken from an alternate path using `gocryptfs` type.
```
{
  "watchtower_encrypted_keys": [
    "~/alternate/path/to/your/keys/.encrypted_keys/wt1"
  ],
  "operator_encrypted_key": "~/alternate/path/to/your/keys/.encrypted_keys/op1",
  "encrypted_key_type": "gocryptfs",
  "eth_rpc_url": "<Mainnet RPC URL>"
}
```
