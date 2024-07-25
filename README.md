# WitnessChain Operator CLI

## Description
watchtower-operator is a command-line interface (CLI) tool for interacting with some functionalities provided by the WatchTower(EigenLayer AVS) contracts . You can refer to the [How to use the config files](#how-to-use-the_config-files) section to understand how to use the config files.

## 1. Installation
You can get the watchtower-operator cli prebuilt, or build from source


### Prebuilt
   
You can run the following command in your terminal and follow instructions provided by the script to use the cli
```
curl -sSfL witnesschain.com/install-operator-cli | bash
```

Installation instructions for building from source is available in 
[docs/install.md](docs/install.md).


##  watchtower-operator cli usage
Once you have the watchtower-operator installed, you can directly use the exectable -

```
watchtower-operator command [command options] [arguments...]
```
Note: In case you haven't exported the path for watchtower-operator executable, you can start the cli by `./watchtower-operator`

## Commands available
| Command | Description |
|----------|----------|
|keys | Used to store private keys in an encrypted format |
|registerWatchtower | Used to register watch tower |
|deRegisterWatchtower | Used to deregister watch tower |
|registerOperatorToAVS | Used to notify EigenLayer that an operator is registered to the AVS |
|deRegisterOperatorFromAVS | Used to notify EigenLayer that an operator is de-registered from the AVS |

## 2. Key management

You need access to operator and watchtower private keys to register your 
operator and watchtowers to WitnessChain. The following sections guide 
how to setup operator and watchtower keys.

###  Initialise web3 secret storage keystore

```
$ watchtower-operator keys init
Creating directory:  .w3secretkeys
Init keystore done
```

### Import existing operator and watchtower keys
```
$ watchtower-operator keys import --key-name operator
Enter password to import: **********
Enter private key: ******************************************************************
Imported key: operator

$ watchtower-operator keys import --key-name watchtower1
Enter password to import: **********
Enter private key: ******************************************************************
Imported key: watchtower1
```


### list imported or created keys in the keystore

```
$ watchtower-operator keys list
   -----------------------------------------------------------------------------------------------
   Name                                                                   Created
   -----------------------------------------------------------------------------------------------
   operator.ecdsa.key.json                                                25-07-2024 14:57:20
   watchtower1.ecdsa.key.json                                             25-07-2024 14:57:36
   -----------------------------------------------------------------------------------------------
```

These keys are stored in web3 secret storage format recommended by 
[ethereum 
foundation](https://ethereum.org/en/developers/docs/data-structures-and-encoding/web3-secret-storage/). 
`watchtower-operator` cli also support [gocryptfs](docs/gocryptfs.md) 
and [plaintext](docs/plaintext.md) format.

## 3. Setup config file

Now create a new file, `operator-config.json`, and fill in the operator 
private keys and watchtower private keys. You must also change the 
`eth_rpc_url` to the L1 Ethereum node that you trust.

Create a new configuration file with the following template for testnet:

```
{
  "watchtower_encrypted_keys": [
    "watchtower1.ecdsa.key.json",
    "watchtower2.ecdsa.key.json"
  ],
  "operator_encrypted_key": "operator.ecdsa.key.json",
  "eth_rpc_url": "https://ethereum-holesky-rpc.publicnode.com",
  "proof_submission_rpc_url": "https://blue-orangutan-rpc.eu-north-2.gateway.fm/"
}
```

In case you are running on mainnet, replace `eth_rpc_url` and 
`proof_submission_rpc_url`.

Template config for mainnet

```
{
  "watchtower_encrypted_keys": [
    "watchtower1.ecdsa.json",
    "watchtower2.ecdsa.json"
  ],
  "operator_encrypted_key": "operator.ecdsa.json",
  "eth_rpc_url": "wss://ethereum-rpc.publicnode.com",
  "proof_submission_rpc_url": "https://rpc.witnesschain.com"
}
```

Here, `operator_encrypted_key` corresponds to the operator private key 
that you imported earlier. Similarly, `watchtower_encrypted_keys` 
represent the imported keys of your watchotwers.

You can read more about other customization in 
[docs/config.md](docs/config.md)

### 4. Register oeprator to AVS
```
$ watchtower-operator registerOperatorToAVS --config-file operator-config.json
Using config file path : operator-config.json
Using the key path : .w3secretkeys
Enter password to export web3 secret storage keys: **********
Connection successful :  17000
github.com/witnes .. │ Jul 25 16:09:23 2024 │ ➤ keystore: raw://0x621593B9Ae270C418e9190714e7786Ba69398834
Tx sent: https://holesky.etherscan.io/tx/0x36ead44cfaa8b9d3e0b25f03399a0b0517b59e77e407b3574b5dc09dc7479b4a
Transaction executed successfully, logs are ...
[0xc010af8000]
```


### 5. Register watchtower
```
$ watchtower-operator registerWatchtower --config-file operator-config.json
Using config file path : operator-config.json
Using the key path : .w3secretkeys
Enter password to export web3 secret storage keys: **********
Connection successful :  17000
github.com/witnes .. │ Jul 25 16:52:40 2024 │ ➤ keystore: raw://0x621593B9Ae270C418e9190714e7786Ba69398834
watchtowerAddress: 0x621593B9Ae270C418e9190714e7786Ba69398834
github.com/witnes .. │ Jul 25 16:52:41 2024 │ ➤ keystore: raw://0x621593B9Ae270C418e9190714e7786Ba69398834
Tx sent: https://holesky.etherscan.io/tx/0x4f5d9ac9f8b425cbd2d32ac32625e6441e00c7692a57d7d884b842ff92be8901
Transaction executed successfully, logs are ...
[0xc010966000]
```
Congratulations! Your watchtower is successfully registered. Now you can 
proceed to install watchtower-client and submit proofs on chain.
