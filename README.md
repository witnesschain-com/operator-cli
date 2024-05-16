# WitnessChain Operator CLI

## Description
watchtower-operator is a command-line interface (CLI) tool for interacting with some functionalities provided by the WatchTower(EigenLayer AVS) contracts . You can refer to the [How to use the config files](#how-to-use-the_config-files) section to understand how to use the config files.

## Installation
You can get the watchtower-operator cli prebuilt, or build from source

### Pre-requisite

gocryptfs should be installed on the machine you will use the CLI tool on. The installation and running of gocryptfs is easier and smoother on Unbuntu/Debian systems. So, to use the CLI, those would be the ideal systems

1. **Prebuilt**
   
    You can run the following command in your terminal and follow instructions provided by the script to use the cli
    ```
    curl -sSfL witnesschain.com/install-operator-cli | bash
    ```

1. **Building from source**
   - clone the repository 
    ```
    git clone https://github.com/witnesschain-com/operator-cli.git
    ```

   - Build the binaries
    ```
    cd operator-cli/watchtower-operator-cli
    ./build
    ```


## How to use
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

## How to use the config files
The structure and details in the config file might differ based on the functionality you are using. Config files related to operator have been explained below. Enter the following details in the config file. Changing the key names in the json file will lead to misbehavior

### Operator config file

#### for command - registerOperatorToAVS, deRegisterOperatorFromAVS
Default file: config/l1-operator-config.json.template (reference file)

| Field | Description |
|----------|----------|
|watchtower_private_key | Private key of the watchtower |
|operator_private_key | Private key of the operator(on which the actions will be performed) |
|operator_registry_address | Address of the OperatorRegistry contract(used to register and deregister operators) |
|witnesshub_address | Address of the WitnessHub contract(used to register operator to the AVS) |
|avs_directory_address | Address of the AvsDirectory contract(used while registering operator to the AVS) |
|eth_rpc_url | The RPC URL where you want to perform the transactions |
|chain_id | Chain ID of the respective chain |
|gas_limit | The gas limit you want to set while sending the transactions |
|tx_receipt_timeout| Timeout in seconds for waiting of tx receipts |
|expiry| Expiry in days after which the operator signature becomes invalid |
|use_encrypted_keys| Indicates if you want to store and use the private keys in encrypted format |

#### for commands - registerWatchtower, deRegisterWatchtower
Default file: config/l2-operator-config.json.template (reference file)

| Field | Description |
|----------|----------|
|watchtower_private_key | Private key of the watchtower |
|operator_private_key | Private key of the operator(on which the actions will be performed) |
|operator_registry_address | Address of the OperatorRegistry contract(used to register and deregister operators) |
|eth_rpc_url | The RPC URL where you want to perform the transactions |
|chain_id | Chain ID of the respective chain |
|gas_limit | The gas limit you want to set while sending the transactions |
|tx_receipt_timeout| Timeout in seconds for waiting of tx receipts |
|expiry| Expiry in days after which the operator signature becomes invalid |
|use_encrypted_keys| Indicates if you want to store and use the private keys in encrypted format |

> **Note:** The same l1-operator-config.json file can be used for the registerWatchtower, deRegisterWatchtower too. But vice-versa is not possible

## How to use the encrypted keys

To store and use the keys in an encrypted format, use the `use_encrypted_keys` field in the config. If it is set to `true`, the CLI tool will use the keys stored in the encrypted format in the location, `~/.witnesschain/cli/encrypted_keys`. The CLI uses gocrytpfs to achieve this encryption. Read more about it [here](https://github.com/rfjakob/gocryptfs)

### Sub-commands
The `keys` command has 4 sub-commands. Let's look at them one by one -
1. `init` - This command is used to initiate a new gocryptfs file system. This will ask you to input a password which will be required whenever we need to access these file. It also performs password validations for stromger passwords. To bypass this validation in the test environment, there is a flag `--insecure(-i)`. Once the command is successfully run, all other actions will need this password.
```
$ watchtower-operator keys init

Creating directory:  .encrypted_keys
Creating directory:  .decrypted_keys
Enter password to init: **********
```
After this command, two directories .encrypted_keys and .decrypted_keys are created. The names indicate their functions. Once this is done, we don't need to do it again, unless the .encrypted_keys or .decrypted_keys are tampered with.

2. `create` - This command will create a new key. This command will need a flag --key-name(-k). This will be the name of the key which will be referred in the future by the CLI. This name should be mentioned in the config file to extract the corresponding private key. When you run this command, it will ask for password to mount and then ask you to enter the private key
```
$ watchtower-operator keys create -k wt1

Enter password to mount: **********
Enter private key: ****************************************************************
Created key: wt1
```
3. `list` - This command will list all the keys currently present (only by file name and created date)
```
$ watchtower-operator keys list

Enter password to mount: **********
   -------------------------------------------------------
   Name                           Created
   -------------------------------------------------------
   wt1                            30-04-2024 11:36:16
   -------------------------------------------------------
```
4. `delete` - This command will delete a key. This command will need a flag --key-name(-k). The name given will be searched in the `decrypted_keys` and deleted
```
$ watchtower-operator keys delete -k wt1

Enter password to mount: **********
Deleted key: wt1

$ watchtower-operator keys list

Enter password to mount: **********
   -------------------------------------------------------
   Name                           Created
   -------------------------------------------------------
   -------------------------------------------------------
```

Going forward, the CLI will ask for the mount password to decrpyt and use these keys. This is how the config will look like when using encrypted keys

```
{
  "watchtower_private_keys": [
    "wt1"
  ],
  "operator_private_key": "op1",
  "operator_registry_address": "0xEf1a89841fd189ba28e780A977ca70eb1A5e985D",
  "witnesshub_address": "0xD25c2c5802198CB8541987b73A8db4c9BCaE5cC7",
  "avs_directory_address": "0x135dda560e946695d6f155dacafc6f1f25c1f5af",
  "eth_rpc_url": "<Mainnet RPC URL>",
  "chain_id": 1,
  "gas_limit": 5000000,
  "tx_receipt_timeout": 600,
  "expiry_in_days": 1,
  "use_encrypted_keys": true
}
```
