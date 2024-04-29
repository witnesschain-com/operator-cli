# WitnessChain Operator CLI

## Description
watchtower-operator is a command-line interface (CLI) tool for interacting with some functionalities provided by the WatchTower(EigenLayer AVS) contracts . You can refer to the [How to use the config files](#how-to-use-the_config-files) section to understand how to use the config files.

## How to use
To build the CLI, use the following command - 
```
./build
```

If the module is already built you can directly use the exectable -
```
watchtower-operator command [command options] [arguments...]
```

## Commands available
| Command | Description |
|----------|----------|
|keys | Used to store private keys in an encrypted format |
|registerWatchtower | Used to register watch tower |
|deRegisterWatchtower | Used to deregister watch tower |
|registerOperatorToAVS | Used to notify EigenLayer that an operator is registered to the AVS |
|deRegisterOperatorToAVS | Used to notify EigenLayer that an operator is de-registered from the AVS |

## How to use the config files
The structure and details in the config file might differ based on the functionality you are using. Config files related to both, the operator and the admin functionality have been explained below. Enter the following details in the config file. Changing the key names in the yaml file will lead to misbehavior

### Operator config file

#### for command - registerOperatorToAVS
Default file: config/l1-operator-config.yaml

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
Default file: config/l2-operator-config.yaml

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

> **Note:** The same l1-operator-config.yaml file can be used for the registerWatchtower, deRegisterWatchtower too. But vice-versa is not possible

## How to use the encrypted keys

To store and use the keys in an encrypted format, use the `use_encrypted_keys` field in the config. If it is set to `true`, the CLI tool will use the keys stored in the encrypted format in the location, `~/.witnesschain/cli/encrypted_keys`. The CLI uses gocrytpfs to achieve this encryption. Read more about it [here](https://github.com/rfjakob/gocryptfs)

### Sub-commands
The `keys` command has 4 sub-commands. Let's look at them one by one -
1. `init` - This command is used to initiate a new gocryptfs file system. This will ask you to input a password which will be required whenever we need to access these file. It also performs password validations for stromger passwords. To bypass this validation in the test environment, there is a flag `--insecure(-i)`. Once the command is successfully run, all other actions will need this password.
2. `create` - This command will create a new key. This command will need a flag --key-name(-k). This will be the name of the key which will be referred in the future by the CLI. This name should be mentioned in the config file to extract the corresponding private key. When you run this command, it will ask for password to mount and then ask you to enter the private key
3. `delete` - This command will delete a key. This command will need a flag --key-name(-k). The name given will be searched in the `decrypted_keys` and deleted
4. `list` - This command will list all the keys currently present (only by file name and created date)