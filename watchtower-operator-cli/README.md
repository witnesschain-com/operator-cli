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
|registerWatchtower | Used to register watch tower |
|deRegisterWatchtower | Used to deregister watch tower |
|registerOperatorToAVS | Used to notify EigenLayer that an operator is registered to the AVS |

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

> **Note:** The same l1-operator-config.yaml file can be used for the registerWatchtower, deRegisterWatchtower too. But vice-versa is not possible
