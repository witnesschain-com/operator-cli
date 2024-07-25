# WitnessChain Operator CLI

## Description
watchtower-operator is a command-line interface (CLI) tool for interacting with some functionalities provided by the WatchTower(EigenLayer AVS) contracts . You can refer to the [How to use the config files](#how-to-use-the_config-files) section to understand how to use the config files.

## Installation
You can get the watchtower-operator cli prebuilt, or build from source


### Prebuilt
   
You can run the following command in your terminal and follow instructions provided by the script to use the cli
```
curl -sSfL witnesschain.com/install-operator-cli | bash
```

Installation instructions for building from source is available in 
[docs/install.md](docs/install.md).


## watchtower-operator cli usage
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

## Key management

You need access to operator and watchtower private keys to register your 
operator and watchtowers to WitnessChain. The following sections guide 
how to setup operator and watchtower keys.

###  Initialise keystore

```
$ watchtower-operator keys init
Creating directory:  .w3secretkeys
Init keystore done
```

### Import existing keys
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
kripashanker@fedora:~$
```

## Setup config file

Now create a new file, `operator-config.json`, and fill in the operator 
private keys and watchtower private keys. You must also change the 
`eth_rpc_url` to the L1 Ethereum node that you trust.

Create a new configuration file with the following template for mainnet:

```
{
  "watchtower_encrypted_keys": [
    "watchtower1.ecdsa.json",
    "watchtower2.ecdsa.json"
  ],
  "operator_encrypted_key": "operator.ecdsa.json",
  "eth_rpc_url": "wss://ethereum-rpc.publicnode.com"
}
```

In case you are running on testnet, replace "eth_rpc_url" with holesky.

Here, `operator_encrypted_key` corresponds to the operator private key 
that you imported earlier. Similarly, `watchtower_encrypted_keys` 
represent the imported keys of your watchotwers.

You can read more about other customization in 
[docs/config.md](docs/config.md)

### Register oeprator to AVS

### deRegister oeprator from AVS
### Register watchtower
### deRegister watchtower

