**# WitnessChain Operator CLI

## Description
watchtower-operator is a command-line interface (CLI) tool for interacting with some functionalities provided by the WatchTower(EigenLayer AVS) contracts . You can refer to the 'How to use the config files' section to understand how to use the config files.


## Installation
You can get the watchtower-operator cli prebuilt, or build from source

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
|registerWatchtower | Used to register watch tower |
|deRegisterWatchtower | Used to deregister watch tower |
|registerOperatorToAVS | Used to notify EigenLayer that an operator is registered to the AVS |
|deRegisterOperatorToAVS | Used to notify EigenLayer that an operator is de-registered from the AVS |

## How to use the config files
The structure and details in the config file might differ based on the functionality you are using. Config files related to both, the operator and the admin functionality have been explained below. Enter the following details in the config file. Changing the key names in the json file will lead to misbehavior

### Operator config file

#### for command - registerOperatorToAVS
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
|tx_receipt_timeout | Transaction receipt timeout in seconds |
|expiry_in_days | Signature expiry in days |


> **Note:** The same l1-operator-config.json file can be used for the registerWatchtower, deRegisterWatchtower too. But vice-versa is not possible
