// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package WitnessHub

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// IWitnessHubStrategyParam is an auto generated low-level Go binding around an user-defined struct.
type IWitnessHubStrategyParam struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// TypesBountyRewards is an auto generated low-level Go binding around an user-defined struct.
type TypesBountyRewards struct {
	InclusionProofBounties *big.Int
	DiligenceProofBounties *big.Int
}

// WitnessHubMetaData contains all meta data concerning the WitnessHub contract.
var WitnessHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIAVSDirectory\",\"name\":\"__avsDirectory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAggregator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAggregator\",\"type\":\"address\"}],\"name\":\"AggregatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"InvalidOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldL2ChainMapping\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newL2ChainMapping\",\"type\":\"address\"}],\"name\":\"L2ChainMappingUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2BlockNumberEnd\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"rewardHash\",\"type\":\"bytes32\"}],\"name\":\"NewRewardsUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRegistry\",\"type\":\"address\"}],\"name\":\"RegistryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"indexed\":false,\"internalType\":\"structIWitnessHub.StrategyParam[]\",\"name\":\"params\",\"type\":\"tuple[]\"}],\"name\":\"SetStrategyParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_proofCommitments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2BlockNumberBegin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2BlockNumberEnd\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"rewardHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"submissionBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsDirectory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"deregisterOperatorFromAVS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"}],\"name\":\"getNextBlockByChainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperatorRestakedStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"}],\"name\":\"getOperatorRewardsByChainID\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"inclusionProofBounties\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"diligenceProofBounties\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.BountyRewards\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRestakeableStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOperatorRegistry\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"contractIL2ChainMapping\",\"name\":\"_l2ChainMapping\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_agg\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ChainMapping\",\"outputs\":[{\"internalType\":\"contractIL2ChainMapping\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastUpdateBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"operatorSignature\",\"type\":\"tuple\"}],\"name\":\"registerOperatorToAVS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIOperatorRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"setAggregatorAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIL2ChainMapping\",\"name\":\"_l2chainmapping\",\"type\":\"address\"}],\"name\":\"setL2ChainMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOperatorRegistry\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structIWitnessHub.StrategyParam[]\",\"name\":\"params\",\"type\":\"tuple[]\"}],\"name\":\"setStrategyParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_metadataURI\",\"type\":\"string\"}],\"name\":\"updateAVSMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumBegin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumEnd\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_operatorsList\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"inclusionProofBounties\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"diligenceProofBounties\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.BountyRewards[]\",\"name\":\"_proofRewards\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"_rewardHash\",\"type\":\"bytes32\"}],\"name\":\"updateReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608080604052600436101561001357600080fd5b600090813560e01c90816322062c1414611cfa57508063245a7bfc14611cd157806333cfb7b714611cab5780633659cfe6146119ed578063395373b0146113ec5780633f4ba83a1461138a57806347b32448146112985780634f1ef28614610f7457806352d1902d14610eb05780635472534d14610e865780635c975abb14610e635780636b3aa72e14610e1e578063715018a614610dbe5780637b10399914610d95578063820b6f3914610d6c5780638456cb5914610d075780638d134f8d14610c805780638da5cb5b14610c5757806393a949fb14610b7b5780639926ee7d146109e8578063a364f4da14610872578063a91ee0dc1461077b578063a98fb35514610698578063ac359bd314610671578063ae30f16d146103d6578063c0c53b8b14610216578063e481af9d146101ea5763f2fde38b1461015557600080fd5b346101e75760203660031901126101e75761016e611df4565b610176611f6e565b6001600160a01b038116156101935761018e90611fc6565b604051f35b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b80fd5b50346101e757806003193601126101e7576102126102066124a7565b60405191829182611e0a565b0390f35b50346101e75760603660031901126101e757610230611df4565b6001600160a01b0360243581811692908390036103d157604435908282168092036103d15784549360ff8560081c1615938480956103c4575b80156103ad575b156103515760ff1986811660011788559585610340575b506001600160601b0360a01b92168260fb54161760fb558160fc54161760fc5560fd54161760fd556102c860ff845460081c166102c3816121f8565b6121f8565b6102d133611fc6565b6102fd83549260ff8460081c16906102e8826121f8565b6102f1826121f8565b606554166065556121f8565b610308575b50604051f35b61ff00191681557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160018152a138610302565b61ffff191661010117875538610287565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b50303b1580156102705750600160ff871614610270565b50600160ff871610610269565b600080fd5b50346101e75760208060031936011261066d576004356001600160401b03811161066957610408903690600401611e4e565b610410612258565b6033546001600160a01b039190821633036106155760fe8054868255806105d4575b5085949291945b8581106104cf57505060405193808386018487525260408501939286925b82841061048a57877f78f9b85d45475399beaeb2ca0a5bb0eb3cdc5e60dd4be21f77e1257e641f33fb88880389a1604051f35b909192939485358281168091036104cb57815282860135906001600160601b0382168092036104cb5760408160019386839401520196019401929190610457565b8880fd5b826104e76104e28389899a97989a612449565b612435565b161561058f576104f8818487612449565b90825491600160401b83101561057b576105186001938481018655612459565b919091610567578661052982612435565b1690896001600160601b0360a01b91838386541617855501356001600160601b03811681036105635760a01b161790550194929194610439565b8b80fd5b634e487b7160e01b8a5260048a905260248afd5b634e487b7160e01b89526041600452602489fd5b60405162461bcd60e51b815260048101879052601f60248201527f5769746e6573734875623a20206e6f204e756c6c2073747261746567696573006044820152606490fd5b8187527f54075df80ec1ae6ac9100e1fd0ebf3246c17f5c933137af392011f4c5f61513a908101905b81811061060a5750610432565b8781556001016105fd565b60405162461bcd60e51b815260048101859052602660248201527f5769746e6573734875623a204f776e65722073686f756c64206265207468652060448201526539b2b73232b960d11b6064820152608490fd5b8280fd5b5080fd5b50346101e75760203660031901126101e75760206106906004356123d1565b604051908152f35b50346101e75760203660031901126101e757806004356001600160401b0381116107785736602382011215610778576106db903690602481600401359101611f19565b6106e3611f6e565b7f000000000000000000000000055733000064333caddbc92763c58bf0192ffebf6001600160a01b0316803b156107745760405163a98fb35560e01b8152602060048201529183918391829084908290610741906024830190612163565b03925af18015610769576107555750604051f35b61075e90611eaf565b6101e7578038610302565b6040513d84823e3d90fd5b5050fd5b50fd5b50346101e75760203660031901126101e757610795611df4565b61079d611f6e565b60fb546001600160a01b0391821691811690828214610814576001600160a01b03198116831760fb55604080516001600160a01b03938416815291851690931790911660208201527f482b97c53e48ffa324a976e2738053e9aff6eee04d8aac63b10e19411d869b8291819081015b0390a1604051f35b60405162461bcd60e51b815260206004820152603060248201527f5769746e6573734875623a20526567697374727920616c72656164792073657460448201526f20746f2074686973206164647265737360801b6064820152608490fd5b50346101e75760203660031901126101e75761088c611df4565b610894612258565b6001600160a01b0390811633811480156109db575b1561097b578160246020859460fb541660405192838092633367cca560e01b82528760048301525afa8015610970576108e9918591610942575b506122fd565b7f000000000000000000000000055733000064333caddbc92763c58bf0192ffebf1690813b156107745782916024839260405194859384926351b27a6d60e11b845260048401525af18015610769576107555750604051f35b610963915060203d8111610969575b61095b8183611edd565b8101906122e5565b386108e3565b503d610951565b6040513d86823e3d90fd5b60405162461bcd60e51b815260206004820152603260248201527f5769746e6573734875623a204f70657261746f72206f72206f776e657220736860448201527137bab632103132903a34329039b2b73232b960711b6064820152608490fd5b50816033541633146108a9565b50346101e75760031960403682011261066d57610a03611df4565b602435916001600160401b0390818411610b54576060908436030112610b6357604051906060820182811082821117610b6757604052849291906004850135908111610b6357610a599060043691870101611f50565b81526020810190602485013582526044604082019501358552610a7a612258565b60fb54604051633367cca560e01b81526001600160a01b039485166004820181905294916020908290602490829086165afa8015610b5857610ac291879161094257506122fd565b7f000000000000000000000000055733000064333caddbc92763c58bf0192ffebf1691823b15610b5457610b2493859160405197889687958694639926ee7d60e01b8652600486015260406024860152516060604486015260a4850190612163565b9151606484015251608483015203925af1801561076957610b455750604051f35b610b4e90611eaf565b38610302565b8480fd5b6040513d88823e3d90fd5b8380fd5b634e487b7160e01b86526041600452602486fd5b50346101e75760403660031901126101e757610b95611df4565b60405190602435610ba583611e7e565b838352602092830184905260fc54604051633d69d91d60e01b8152600481018390526001600160a01b03929185908290602490829087165afa908115610b585791610bff8794926040989489979591610c3a575b50612363565b835260ff855260018484200191168252835220908251610c1e81611e7e565b8160018454948584520154910190815283519283525190820152f35b610c519150883d8a116109695761095b8183611edd565b38610bf9565b50346101e757806003193601126101e7576033546040516001600160a01b039091168152602090f35b50346101e75760203660031901126101e757610c9a611df4565b50610ca3611f6e565b60405162461bcd60e51b815260206004820152603660248201527f5769746e6573734875623a204c32436861696e4d617070696e6720616c72656160448201527564792073657420746f2074686973206164647265737360501b6064820152608490fd5b50346101e757806003193601126101e757610d20612258565b610d28611f6e565b610d30612258565b600160ff1960655416176065557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1604051f35b50346101e757806003193601126101e75760fc546040516001600160a01b039091168152602090f35b50346101e757806003193601126101e75760fb546040516001600160a01b039091168152602090f35b50346101e757806003193601126101e757610dd7611f6e565b603380546001600160a01b031981169091556040519082906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08284a3f35b50346101e757806003193601126101e7576040517f000000000000000000000000055733000064333caddbc92763c58bf0192ffebf6001600160a01b03168152602090f35b50346101e757806003193601126101e757602060ff606554166040519015158152f35b50346101e75760203660031901126101e7576040602091600435815260ff83522054604051908152f35b50346101e757806003193601126101e7577f000000000000000000000000e8a3c4e7701028a9fab27e468a1be5a8161ec9686001600160a01b03163003610f095760206040516000805160206125348339815191528152f35b60405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608490fd5b5060403660031901126101e757610f89611df4565b906024356001600160401b03811161066d57610fa9903690600401611f50565b916001600160a01b037f000000000000000000000000e8a3c4e7701028a9fab27e468a1be5a8161ec968811690610fe230831415612011565b610fff600080516020612534833981519152928284541614612072565b611007611f6e565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff161561103e57505061018e9192506120d3565b6040516352d1902d60e01b8152602094939291831691908581600481865afa859181611265575b506110c65760405162461bcd60e51b815260048101879052602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608490fd5b9493940361120e576110d7826120d3565b604051907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8583a2845115801590611206575b611119575b5050509050604051f35b813b156111b557509180846111a394848397519201905af43d156111ad573d9061114282611efe565b916111506040519384611edd565b82523d858484013e5b7f416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c6040519361118785611ec2565b60278552840152660819985a5b195960ca1b60408401526121b4565b508038808061110f565b606090611159565b62461bcd60e51b815260048101839052602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608490fd5b50600161110a565b60405162461bcd60e51b815260048101849052602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608490fd5b9091508681813d8311611291575b61127d8183611edd565b8101031261128d57519038611065565b8580fd5b503d611273565b50346101e75760203660031901126101e7576112b2611df4565b6112ba611f6e565b60fd54906001600160a01b0390811690821681811461132a576001600160a01b0319909216811760fd55604080516001600160a01b0393841681529290911660208301527f89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f291908190810161080c565b60405162461bcd60e51b815260206004820152603260248201527f5769746e6573734875623a2041676772656761746f7220616c72656164792073604482015271657420746f2074686973206164647265737360701b6064820152608490fd5b50346101e757806003193601126101e7576113a361229c565b6113ab611f6e565b6113b361229c565b60ff19606554166065557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1604051f35b50346101e75760c03660031901126101e7576064356001600160401b03811161066d573660238201121561066d576001600160401b0381600401351161066d57366024826004013560051b8301011161066d576084356001600160401b0381116106695761145e903690600401611e4e565b90611467612258565b60fd546001600160a01b031633036119985760fc54604051633d69d91d60e01b8152600480359082015290602090829060249082906001600160a01b03165afa801561198d576114bd9186916119745750612363565b604435602435101561190957818360040135036118a7576114df6004356123d1565b6024350361185657835b82811061169c5784600435815260ff60205260443560408220556040519060a43591826044356004357f36b8234ba3e9f8897b5448ed687509b776a4ad94a88ffd28d6aef8bee37ee6108585a460a081018181106001600160401b03821117611688576040526004358152602081019260243584526040820190604435825260608301908152608083019143835261010095865496600160401b88101561167457600188018082558810156116605786976005917f45e010b9ae401e2eb71529478da8bd513a9bdc2d095a111e324f5b95c09ed87f979852029551867f45e010b9ae401e2eb71529478da8bd513a9bdc2d095a111e324f5b95c09ed87b0155517f45e010b9ae401e2eb71529478da8bd513a9bdc2d095a111e324f5b95c09ed87c860155517f45e010b9ae401e2eb71529478da8bd513a9bdc2d095a111e324f5b95c09ed87d850155517f45e010b9ae401e2eb71529478da8bd513a9bdc2d095a111e324f5b95c09ed87e84015551910155604051f35b634e487b7160e01b87526032600452602487fd5b634e487b7160e01b87526041600452602487fd5b634e487b7160e01b83526041600452602483fd5b60fb546001600160a01b031660206116bf6104e284600489013560248a0161240f565b604051633367cca560e01b81526001600160a01b03909116600482015291829060249082905afa908115610b58578691611837575b50156117a857806117086001928585612449565b3560043587526117a060ff918260205284908160408b2001868060a01b0361173c6104e2888d60248160040135910161240f565b168b5260205261175160408b209182546123af565b90556020611760858989612449565b0135926004358a526020528060408a2001858060a01b0361178d6104e2878c60248160040135910161240f565b168a5260205260408920019182546123af565b9055016114e9565b60206117e16104e27f3eaa03e38745f9ef282a3d3ae730e428a187a9075ccc2a3b7a1a318fe83d25be938760248160040135910161240f565b6040516001600160a01b039091168152a160405162461bcd60e51b815260206004820152601d60248201527f5769746e6573734875623a20496e616374697665206f70657261746f720000006044820152606490fd5b611850915060203d6020116109695761095b8183611edd565b386116f4565b60405162461bcd60e51b8152602060048201526024808201527f5769746e6573734875623a20496e636f7272656374205f626c6f636b4e756d4260448201526332b3b4b760e11b6064820152608490fd5b60405162461bcd60e51b815260206004820152603460248201527f5769746e6573734875623a20756e657175616c206f70657261746f727320616e6044820152730c840e4caeec2e4c840d8d2e6e840d8cadccee8d60631b6064820152608490fd5b60405162461bcd60e51b815260206004820152603b60248201527f5769746e6573734875623a205f626c6f636b4e756d426567696e2073686f756c60448201527f64206265206c657373207468616e205f626c6f636b4e756d456e6400000000006064820152608490fd5b610c51915060203d6020116109695761095b8183611edd565b6040513d87823e3d90fd5b60405162461bcd60e51b815260206004820152602760248201527f5769746e6573734875623a20596f7520617265206e6f74207468652061676772604482015266656761746f722160c81b6064820152608490fd5b50346101e75760208060031936011261066d57611a08611df4565b906001600160a01b037f000000000000000000000000e8a3c4e7701028a9fab27e468a1be5a8161ec9688116611a4030821415612011565b611a5d600080516020612534833981519152918383541614612072565b611a65611f6e565b60405190838201928284106001600160401b03851117611674578360405286835260ff7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435416600014611ac057505050505061018e906120d3565b8596949516906040516352d1902d60e01b81528681600481865afa869181611c78575b50611b445760405162461bcd60e51b815260048101889052602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608490fd5b95949503611c2157611b55866120d3565b604051907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8683a2815115801590611c1a575b611b98575b505050509050604051f35b853b15611bc957509280948192611bbe9551915af43d156111ad573d9061114282611efe565b508038808080611b8d565b62461bcd60e51b815260048101849052602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608490fd5b5084611b88565b60405162461bcd60e51b815260048101859052602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608490fd5b9091508781813d8311611ca4575b611c908183611edd565b81010312611ca057519038611ae3565b8680fd5b503d611c86565b50346101e75760203660031901126101e757611cc5611df4565b506102126102066124a7565b50346101e757806003193601126101e75760fd546040516001600160a01b039091168152602090f35b90503461066d57602036600319011261066d576004356101008054821015610b63579060059160a0945202807f45e010b9ae401e2eb71529478da8bd513a9bdc2d095a111e324f5b95c09ed87b0154907f45e010b9ae401e2eb71529478da8bd513a9bdc2d095a111e324f5b95c09ed87c8101547f45e010b9ae401e2eb71529478da8bd513a9bdc2d095a111e324f5b95c09ed87d820154907f45e010b9ae401e2eb71529478da8bd513a9bdc2d095a111e324f5b95c09ed87f7f45e010b9ae401e2eb71529478da8bd513a9bdc2d095a111e324f5b95c09ed87e8401549301549385526020850152604084015260608301526080820152f35b600435906001600160a01b03821682036103d157565b6020908160408183019282815285518094520193019160005b828110611e31575050505090565b83516001600160a01b031685529381019392810192600101611e23565b9181601f840112156103d1578235916001600160401b0383116103d1576020808501948460061b0101116103d157565b604081019081106001600160401b03821117611e9957604052565b634e487b7160e01b600052604160045260246000fd5b6001600160401b038111611e9957604052565b606081019081106001600160401b03821117611e9957604052565b90601f801991011681019081106001600160401b03821117611e9957604052565b6001600160401b038111611e9957601f01601f191660200190565b929192611f2582611efe565b91611f336040519384611edd565b8294818452818301116103d1578281602093846000960137010152565b9080601f830112156103d157816020611f6b93359101611f19565b90565b6033546001600160a01b03163303611f8257565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b6033549060018060a01b0380911691826001600160601b0360a01b821617603355167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e06000604051a3565b1561201857565b60405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b6064820152608490fd5b1561207957565b60405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b6163746976652070726f787960a01b6064820152608490fd5b803b156121085760008051602061253483398151915280546001600160a01b0319166001600160a01b03909216919091179055565b60405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608490fd5b91908251928382526000905b84821061219c57509280602093941161218f575b601f01601f1916010190565b6000838284010152612183565b9060209081808285010151908286010152019061216f565b909190156121c0575090565b8151156121d05750805190602001fd5b60405162461bcd60e51b8152602060048201529081906121f4906024830190612163565b0390fd5b156121ff57565b60405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608490fd5b60ff6065541661226457565b60405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606490fd5b60ff60655416156122a957565b60405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606490fd5b908160209103126103d1575180151581036103d15790565b1561230457565b60405162461bcd60e51b815260206004820152603160248201527f5769746e6573734875623a204f70657261746f722073686f756c64206265206160448201527031ba34bb32902ba19037b832b930ba37b960791b6064820152608490fd5b1561236a57565b60405162461bcd60e51b815260206004820152601c60248201527f5769746e6573734875623a20496e76616c696420436861696e204944000000006044820152606490fd5b811981116123bb570190565b634e487b7160e01b600052601160045260246000fd5b600090815260ff602052604081205415611f6b5760408120549060011982116123fb575060010190565b634e487b7160e01b81526011600452602490fd5b919081101561241f5760051b0190565b634e487b7160e01b600052603260045260246000fd5b356001600160a01b03811681036103d15790565b919081101561241f5760061b0190565b60fe5481101561241f5760fe6000527f54075df80ec1ae6ac9100e1fd0ebf3246c17f5c933137af392011f4c5f61513a0190600090565b6001600160401b038111611e995760051b60200190565b60fe546124b381612490565b906124c16040519283611edd565b808252601f196124d082612490565b01906020913683850137600091825b8281106124ed575050505090565b6124f681612459565b505485516001600160a01b039091169082101561251f57600582901b86018301526001016124df565b634e487b7160e01b85526032600452602485fdfe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca2646970667358221220d49561065e2fd54e3f57147a64d32b6743935b252d53d41f4a09e863b1594c4364736f6c634300080f0033",
}

// WitnessHubABI is the input ABI used to generate the binding from.
// Deprecated: Use WitnessHubMetaData.ABI instead.
var WitnessHubABI = WitnessHubMetaData.ABI

// WitnessHubBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WitnessHubMetaData.Bin instead.
var WitnessHubBin = WitnessHubMetaData.Bin

// DeployWitnessHub deploys a new Ethereum contract, binding an instance of WitnessHub to it.
func DeployWitnessHub(auth *bind.TransactOpts, backend bind.ContractBackend, __avsDirectory common.Address) (common.Address, *types.Transaction, *WitnessHub, error) {
	parsed, err := WitnessHubMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WitnessHubBin), backend, __avsDirectory)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WitnessHub{WitnessHubCaller: WitnessHubCaller{contract: contract}, WitnessHubTransactor: WitnessHubTransactor{contract: contract}, WitnessHubFilterer: WitnessHubFilterer{contract: contract}}, nil
}

// WitnessHub is an auto generated Go binding around an Ethereum contract.
type WitnessHub struct {
	WitnessHubCaller     // Read-only binding to the contract
	WitnessHubTransactor // Write-only binding to the contract
	WitnessHubFilterer   // Log filterer for contract events
}

// WitnessHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type WitnessHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WitnessHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WitnessHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WitnessHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WitnessHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WitnessHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WitnessHubSession struct {
	Contract     *WitnessHub       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WitnessHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WitnessHubCallerSession struct {
	Contract *WitnessHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// WitnessHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WitnessHubTransactorSession struct {
	Contract     *WitnessHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// WitnessHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type WitnessHubRaw struct {
	Contract *WitnessHub // Generic contract binding to access the raw methods on
}

// WitnessHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WitnessHubCallerRaw struct {
	Contract *WitnessHubCaller // Generic read-only contract binding to access the raw methods on
}

// WitnessHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WitnessHubTransactorRaw struct {
	Contract *WitnessHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWitnessHub creates a new instance of WitnessHub, bound to a specific deployed contract.
func NewWitnessHub(address common.Address, backend bind.ContractBackend) (*WitnessHub, error) {
	contract, err := bindWitnessHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WitnessHub{WitnessHubCaller: WitnessHubCaller{contract: contract}, WitnessHubTransactor: WitnessHubTransactor{contract: contract}, WitnessHubFilterer: WitnessHubFilterer{contract: contract}}, nil
}

// NewWitnessHubCaller creates a new read-only instance of WitnessHub, bound to a specific deployed contract.
func NewWitnessHubCaller(address common.Address, caller bind.ContractCaller) (*WitnessHubCaller, error) {
	contract, err := bindWitnessHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WitnessHubCaller{contract: contract}, nil
}

// NewWitnessHubTransactor creates a new write-only instance of WitnessHub, bound to a specific deployed contract.
func NewWitnessHubTransactor(address common.Address, transactor bind.ContractTransactor) (*WitnessHubTransactor, error) {
	contract, err := bindWitnessHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WitnessHubTransactor{contract: contract}, nil
}

// NewWitnessHubFilterer creates a new log filterer instance of WitnessHub, bound to a specific deployed contract.
func NewWitnessHubFilterer(address common.Address, filterer bind.ContractFilterer) (*WitnessHubFilterer, error) {
	contract, err := bindWitnessHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WitnessHubFilterer{contract: contract}, nil
}

// bindWitnessHub binds a generic wrapper to an already deployed contract.
func bindWitnessHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WitnessHubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WitnessHub *WitnessHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WitnessHub.Contract.WitnessHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WitnessHub *WitnessHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessHub.Contract.WitnessHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WitnessHub *WitnessHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WitnessHub.Contract.WitnessHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WitnessHub *WitnessHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WitnessHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WitnessHub *WitnessHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WitnessHub *WitnessHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WitnessHub.Contract.contract.Transact(opts, method, params...)
}

// ProofCommitments is a free data retrieval call binding the contract method 0x22062c14.
//
// Solidity: function _proofCommitments(uint256 ) view returns(uint256 chainID, uint256 l2BlockNumberBegin, uint256 l2BlockNumberEnd, bytes32 rewardHash, uint256 submissionBlock)
func (_WitnessHub *WitnessHubCaller) ProofCommitments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ChainID            *big.Int
	L2BlockNumberBegin *big.Int
	L2BlockNumberEnd   *big.Int
	RewardHash         [32]byte
	SubmissionBlock    *big.Int
}, error) {
	var out []interface{}
	err := _WitnessHub.contract.Call(opts, &out, "_proofCommitments", arg0)

	outstruct := new(struct {
		ChainID            *big.Int
		L2BlockNumberBegin *big.Int
		L2BlockNumberEnd   *big.Int
		RewardHash         [32]byte
		SubmissionBlock    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.L2BlockNumberBegin = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.L2BlockNumberEnd = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RewardHash = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.SubmissionBlock = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProofCommitments is a free data retrieval call binding the contract method 0x22062c14.
//
// Solidity: function _proofCommitments(uint256 ) view returns(uint256 chainID, uint256 l2BlockNumberBegin, uint256 l2BlockNumberEnd, bytes32 rewardHash, uint256 submissionBlock)
func (_WitnessHub *WitnessHubSession) ProofCommitments(arg0 *big.Int) (struct {
	ChainID            *big.Int
	L2BlockNumberBegin *big.Int
	L2BlockNumberEnd   *big.Int
	RewardHash         [32]byte
	SubmissionBlock    *big.Int
}, error) {
	return _WitnessHub.Contract.ProofCommitments(&_WitnessHub.CallOpts, arg0)
}

// ProofCommitments is a free data retrieval call binding the contract method 0x22062c14.
//
// Solidity: function _proofCommitments(uint256 ) view returns(uint256 chainID, uint256 l2BlockNumberBegin, uint256 l2BlockNumberEnd, bytes32 rewardHash, uint256 submissionBlock)
func (_WitnessHub *WitnessHubCallerSession) ProofCommitments(arg0 *big.Int) (struct {
	ChainID            *big.Int
	L2BlockNumberBegin *big.Int
	L2BlockNumberEnd   *big.Int
	RewardHash         [32]byte
	SubmissionBlock    *big.Int
}, error) {
	return _WitnessHub.Contract.ProofCommitments(&_WitnessHub.CallOpts, arg0)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_WitnessHub *WitnessHubCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WitnessHub.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_WitnessHub *WitnessHubSession) Aggregator() (common.Address, error) {
	return _WitnessHub.Contract.Aggregator(&_WitnessHub.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_WitnessHub *WitnessHubCallerSession) Aggregator() (common.Address, error) {
	return _WitnessHub.Contract.Aggregator(&_WitnessHub.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_WitnessHub *WitnessHubCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WitnessHub.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_WitnessHub *WitnessHubSession) AvsDirectory() (common.Address, error) {
	return _WitnessHub.Contract.AvsDirectory(&_WitnessHub.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_WitnessHub *WitnessHubCallerSession) AvsDirectory() (common.Address, error) {
	return _WitnessHub.Contract.AvsDirectory(&_WitnessHub.CallOpts)
}

// GetNextBlockByChainID is a free data retrieval call binding the contract method 0xac359bd3.
//
// Solidity: function getNextBlockByChainID(uint256 _chainID) view returns(uint256)
func (_WitnessHub *WitnessHubCaller) GetNextBlockByChainID(opts *bind.CallOpts, _chainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WitnessHub.contract.Call(opts, &out, "getNextBlockByChainID", _chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextBlockByChainID is a free data retrieval call binding the contract method 0xac359bd3.
//
// Solidity: function getNextBlockByChainID(uint256 _chainID) view returns(uint256)
func (_WitnessHub *WitnessHubSession) GetNextBlockByChainID(_chainID *big.Int) (*big.Int, error) {
	return _WitnessHub.Contract.GetNextBlockByChainID(&_WitnessHub.CallOpts, _chainID)
}

// GetNextBlockByChainID is a free data retrieval call binding the contract method 0xac359bd3.
//
// Solidity: function getNextBlockByChainID(uint256 _chainID) view returns(uint256)
func (_WitnessHub *WitnessHubCallerSession) GetNextBlockByChainID(_chainID *big.Int) (*big.Int, error) {
	return _WitnessHub.Contract.GetNextBlockByChainID(&_WitnessHub.CallOpts, _chainID)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_WitnessHub *WitnessHubCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _WitnessHub.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_WitnessHub *WitnessHubSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _WitnessHub.Contract.GetOperatorRestakedStrategies(&_WitnessHub.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_WitnessHub *WitnessHubCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _WitnessHub.Contract.GetOperatorRestakedStrategies(&_WitnessHub.CallOpts, operator)
}

// GetOperatorRewardsByChainID is a free data retrieval call binding the contract method 0x93a949fb.
//
// Solidity: function getOperatorRewardsByChainID(address _operator, uint256 _chainID) view returns((uint256,uint256))
func (_WitnessHub *WitnessHubCaller) GetOperatorRewardsByChainID(opts *bind.CallOpts, _operator common.Address, _chainID *big.Int) (TypesBountyRewards, error) {
	var out []interface{}
	err := _WitnessHub.contract.Call(opts, &out, "getOperatorRewardsByChainID", _operator, _chainID)

	if err != nil {
		return *new(TypesBountyRewards), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesBountyRewards)).(*TypesBountyRewards)

	return out0, err

}

// GetOperatorRewardsByChainID is a free data retrieval call binding the contract method 0x93a949fb.
//
// Solidity: function getOperatorRewardsByChainID(address _operator, uint256 _chainID) view returns((uint256,uint256))
func (_WitnessHub *WitnessHubSession) GetOperatorRewardsByChainID(_operator common.Address, _chainID *big.Int) (TypesBountyRewards, error) {
	return _WitnessHub.Contract.GetOperatorRewardsByChainID(&_WitnessHub.CallOpts, _operator, _chainID)
}

// GetOperatorRewardsByChainID is a free data retrieval call binding the contract method 0x93a949fb.
//
// Solidity: function getOperatorRewardsByChainID(address _operator, uint256 _chainID) view returns((uint256,uint256))
func (_WitnessHub *WitnessHubCallerSession) GetOperatorRewardsByChainID(_operator common.Address, _chainID *big.Int) (TypesBountyRewards, error) {
	return _WitnessHub.Contract.GetOperatorRewardsByChainID(&_WitnessHub.CallOpts, _operator, _chainID)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_WitnessHub *WitnessHubCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _WitnessHub.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_WitnessHub *WitnessHubSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _WitnessHub.Contract.GetRestakeableStrategies(&_WitnessHub.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_WitnessHub *WitnessHubCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _WitnessHub.Contract.GetRestakeableStrategies(&_WitnessHub.CallOpts)
}

// L2ChainMapping is a free data retrieval call binding the contract method 0x820b6f39.
//
// Solidity: function l2ChainMapping() view returns(address)
func (_WitnessHub *WitnessHubCaller) L2ChainMapping(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WitnessHub.contract.Call(opts, &out, "l2ChainMapping")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2ChainMapping is a free data retrieval call binding the contract method 0x820b6f39.
//
// Solidity: function l2ChainMapping() view returns(address)
func (_WitnessHub *WitnessHubSession) L2ChainMapping() (common.Address, error) {
	return _WitnessHub.Contract.L2ChainMapping(&_WitnessHub.CallOpts)
}

// L2ChainMapping is a free data retrieval call binding the contract method 0x820b6f39.
//
// Solidity: function l2ChainMapping() view returns(address)
func (_WitnessHub *WitnessHubCallerSession) L2ChainMapping() (common.Address, error) {
	return _WitnessHub.Contract.L2ChainMapping(&_WitnessHub.CallOpts)
}

// OperatorRewards is a free data retrieval call binding the contract method 0x5472534d.
//
// Solidity: function operatorRewards(uint256 ) view returns(uint256 lastUpdateBlock)
func (_WitnessHub *WitnessHubCaller) OperatorRewards(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WitnessHub.contract.Call(opts, &out, "operatorRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorRewards is a free data retrieval call binding the contract method 0x5472534d.
//
// Solidity: function operatorRewards(uint256 ) view returns(uint256 lastUpdateBlock)
func (_WitnessHub *WitnessHubSession) OperatorRewards(arg0 *big.Int) (*big.Int, error) {
	return _WitnessHub.Contract.OperatorRewards(&_WitnessHub.CallOpts, arg0)
}

// OperatorRewards is a free data retrieval call binding the contract method 0x5472534d.
//
// Solidity: function operatorRewards(uint256 ) view returns(uint256 lastUpdateBlock)
func (_WitnessHub *WitnessHubCallerSession) OperatorRewards(arg0 *big.Int) (*big.Int, error) {
	return _WitnessHub.Contract.OperatorRewards(&_WitnessHub.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WitnessHub *WitnessHubCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WitnessHub.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WitnessHub *WitnessHubSession) Owner() (common.Address, error) {
	return _WitnessHub.Contract.Owner(&_WitnessHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WitnessHub *WitnessHubCallerSession) Owner() (common.Address, error) {
	return _WitnessHub.Contract.Owner(&_WitnessHub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WitnessHub *WitnessHubCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WitnessHub.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WitnessHub *WitnessHubSession) Paused() (bool, error) {
	return _WitnessHub.Contract.Paused(&_WitnessHub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WitnessHub *WitnessHubCallerSession) Paused() (bool, error) {
	return _WitnessHub.Contract.Paused(&_WitnessHub.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WitnessHub *WitnessHubCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WitnessHub.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WitnessHub *WitnessHubSession) ProxiableUUID() ([32]byte, error) {
	return _WitnessHub.Contract.ProxiableUUID(&_WitnessHub.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WitnessHub *WitnessHubCallerSession) ProxiableUUID() ([32]byte, error) {
	return _WitnessHub.Contract.ProxiableUUID(&_WitnessHub.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_WitnessHub *WitnessHubCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WitnessHub.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_WitnessHub *WitnessHubSession) Registry() (common.Address, error) {
	return _WitnessHub.Contract.Registry(&_WitnessHub.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_WitnessHub *WitnessHubCallerSession) Registry() (common.Address, error) {
	return _WitnessHub.Contract.Registry(&_WitnessHub.CallOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_WitnessHub *WitnessHubTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _WitnessHub.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_WitnessHub *WitnessHubSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _WitnessHub.Contract.DeregisterOperatorFromAVS(&_WitnessHub.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_WitnessHub *WitnessHubTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _WitnessHub.Contract.DeregisterOperatorFromAVS(&_WitnessHub.TransactOpts, operator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _registry, address _l2ChainMapping, address _agg) returns()
func (_WitnessHub *WitnessHubTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _l2ChainMapping common.Address, _agg common.Address) (*types.Transaction, error) {
	return _WitnessHub.contract.Transact(opts, "initialize", _registry, _l2ChainMapping, _agg)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _registry, address _l2ChainMapping, address _agg) returns()
func (_WitnessHub *WitnessHubSession) Initialize(_registry common.Address, _l2ChainMapping common.Address, _agg common.Address) (*types.Transaction, error) {
	return _WitnessHub.Contract.Initialize(&_WitnessHub.TransactOpts, _registry, _l2ChainMapping, _agg)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _registry, address _l2ChainMapping, address _agg) returns()
func (_WitnessHub *WitnessHubTransactorSession) Initialize(_registry common.Address, _l2ChainMapping common.Address, _agg common.Address) (*types.Transaction, error) {
	return _WitnessHub.Contract.Initialize(&_WitnessHub.TransactOpts, _registry, _l2ChainMapping, _agg)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WitnessHub *WitnessHubTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessHub.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WitnessHub *WitnessHubSession) Pause() (*types.Transaction, error) {
	return _WitnessHub.Contract.Pause(&_WitnessHub.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WitnessHub *WitnessHubTransactorSession) Pause() (*types.Transaction, error) {
	return _WitnessHub.Contract.Pause(&_WitnessHub.TransactOpts)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_WitnessHub *WitnessHubTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _WitnessHub.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_WitnessHub *WitnessHubSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _WitnessHub.Contract.RegisterOperatorToAVS(&_WitnessHub.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_WitnessHub *WitnessHubTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _WitnessHub.Contract.RegisterOperatorToAVS(&_WitnessHub.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WitnessHub *WitnessHubTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessHub.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WitnessHub *WitnessHubSession) RenounceOwnership() (*types.Transaction, error) {
	return _WitnessHub.Contract.RenounceOwnership(&_WitnessHub.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WitnessHub *WitnessHubTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WitnessHub.Contract.RenounceOwnership(&_WitnessHub.TransactOpts)
}

// SetAggregatorAddress is a paid mutator transaction binding the contract method 0x47b32448.
//
// Solidity: function setAggregatorAddress(address _aggregator) returns()
func (_WitnessHub *WitnessHubTransactor) SetAggregatorAddress(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _WitnessHub.contract.Transact(opts, "setAggregatorAddress", _aggregator)
}

// SetAggregatorAddress is a paid mutator transaction binding the contract method 0x47b32448.
//
// Solidity: function setAggregatorAddress(address _aggregator) returns()
func (_WitnessHub *WitnessHubSession) SetAggregatorAddress(_aggregator common.Address) (*types.Transaction, error) {
	return _WitnessHub.Contract.SetAggregatorAddress(&_WitnessHub.TransactOpts, _aggregator)
}

// SetAggregatorAddress is a paid mutator transaction binding the contract method 0x47b32448.
//
// Solidity: function setAggregatorAddress(address _aggregator) returns()
func (_WitnessHub *WitnessHubTransactorSession) SetAggregatorAddress(_aggregator common.Address) (*types.Transaction, error) {
	return _WitnessHub.Contract.SetAggregatorAddress(&_WitnessHub.TransactOpts, _aggregator)
}

// SetL2ChainMapping is a paid mutator transaction binding the contract method 0x8d134f8d.
//
// Solidity: function setL2ChainMapping(address _l2chainmapping) returns()
func (_WitnessHub *WitnessHubTransactor) SetL2ChainMapping(opts *bind.TransactOpts, _l2chainmapping common.Address) (*types.Transaction, error) {
	return _WitnessHub.contract.Transact(opts, "setL2ChainMapping", _l2chainmapping)
}

// SetL2ChainMapping is a paid mutator transaction binding the contract method 0x8d134f8d.
//
// Solidity: function setL2ChainMapping(address _l2chainmapping) returns()
func (_WitnessHub *WitnessHubSession) SetL2ChainMapping(_l2chainmapping common.Address) (*types.Transaction, error) {
	return _WitnessHub.Contract.SetL2ChainMapping(&_WitnessHub.TransactOpts, _l2chainmapping)
}

// SetL2ChainMapping is a paid mutator transaction binding the contract method 0x8d134f8d.
//
// Solidity: function setL2ChainMapping(address _l2chainmapping) returns()
func (_WitnessHub *WitnessHubTransactorSession) SetL2ChainMapping(_l2chainmapping common.Address) (*types.Transaction, error) {
	return _WitnessHub.Contract.SetL2ChainMapping(&_WitnessHub.TransactOpts, _l2chainmapping)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_WitnessHub *WitnessHubTransactor) SetRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _WitnessHub.contract.Transact(opts, "setRegistry", _registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_WitnessHub *WitnessHubSession) SetRegistry(_registry common.Address) (*types.Transaction, error) {
	return _WitnessHub.Contract.SetRegistry(&_WitnessHub.TransactOpts, _registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_WitnessHub *WitnessHubTransactorSession) SetRegistry(_registry common.Address) (*types.Transaction, error) {
	return _WitnessHub.Contract.SetRegistry(&_WitnessHub.TransactOpts, _registry)
}

// SetStrategyParams is a paid mutator transaction binding the contract method 0xae30f16d.
//
// Solidity: function setStrategyParams((address,uint96)[] params) returns()
func (_WitnessHub *WitnessHubTransactor) SetStrategyParams(opts *bind.TransactOpts, params []IWitnessHubStrategyParam) (*types.Transaction, error) {
	return _WitnessHub.contract.Transact(opts, "setStrategyParams", params)
}

// SetStrategyParams is a paid mutator transaction binding the contract method 0xae30f16d.
//
// Solidity: function setStrategyParams((address,uint96)[] params) returns()
func (_WitnessHub *WitnessHubSession) SetStrategyParams(params []IWitnessHubStrategyParam) (*types.Transaction, error) {
	return _WitnessHub.Contract.SetStrategyParams(&_WitnessHub.TransactOpts, params)
}

// SetStrategyParams is a paid mutator transaction binding the contract method 0xae30f16d.
//
// Solidity: function setStrategyParams((address,uint96)[] params) returns()
func (_WitnessHub *WitnessHubTransactorSession) SetStrategyParams(params []IWitnessHubStrategyParam) (*types.Transaction, error) {
	return _WitnessHub.Contract.SetStrategyParams(&_WitnessHub.TransactOpts, params)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WitnessHub *WitnessHubTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WitnessHub.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WitnessHub *WitnessHubSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WitnessHub.Contract.TransferOwnership(&_WitnessHub.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WitnessHub *WitnessHubTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WitnessHub.Contract.TransferOwnership(&_WitnessHub.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WitnessHub *WitnessHubTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessHub.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WitnessHub *WitnessHubSession) Unpause() (*types.Transaction, error) {
	return _WitnessHub.Contract.Unpause(&_WitnessHub.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WitnessHub *WitnessHubTransactorSession) Unpause() (*types.Transaction, error) {
	return _WitnessHub.Contract.Unpause(&_WitnessHub.TransactOpts)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_WitnessHub *WitnessHubTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _WitnessHub.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_WitnessHub *WitnessHubSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _WitnessHub.Contract.UpdateAVSMetadataURI(&_WitnessHub.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_WitnessHub *WitnessHubTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _WitnessHub.Contract.UpdateAVSMetadataURI(&_WitnessHub.TransactOpts, _metadataURI)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x395373b0.
//
// Solidity: function updateReward(uint256 _chainID, uint256 _blockNumBegin, uint256 _blockNumEnd, address[] _operatorsList, (uint256,uint256)[] _proofRewards, bytes32 _rewardHash) returns()
func (_WitnessHub *WitnessHubTransactor) UpdateReward(opts *bind.TransactOpts, _chainID *big.Int, _blockNumBegin *big.Int, _blockNumEnd *big.Int, _operatorsList []common.Address, _proofRewards []TypesBountyRewards, _rewardHash [32]byte) (*types.Transaction, error) {
	return _WitnessHub.contract.Transact(opts, "updateReward", _chainID, _blockNumBegin, _blockNumEnd, _operatorsList, _proofRewards, _rewardHash)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x395373b0.
//
// Solidity: function updateReward(uint256 _chainID, uint256 _blockNumBegin, uint256 _blockNumEnd, address[] _operatorsList, (uint256,uint256)[] _proofRewards, bytes32 _rewardHash) returns()
func (_WitnessHub *WitnessHubSession) UpdateReward(_chainID *big.Int, _blockNumBegin *big.Int, _blockNumEnd *big.Int, _operatorsList []common.Address, _proofRewards []TypesBountyRewards, _rewardHash [32]byte) (*types.Transaction, error) {
	return _WitnessHub.Contract.UpdateReward(&_WitnessHub.TransactOpts, _chainID, _blockNumBegin, _blockNumEnd, _operatorsList, _proofRewards, _rewardHash)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x395373b0.
//
// Solidity: function updateReward(uint256 _chainID, uint256 _blockNumBegin, uint256 _blockNumEnd, address[] _operatorsList, (uint256,uint256)[] _proofRewards, bytes32 _rewardHash) returns()
func (_WitnessHub *WitnessHubTransactorSession) UpdateReward(_chainID *big.Int, _blockNumBegin *big.Int, _blockNumEnd *big.Int, _operatorsList []common.Address, _proofRewards []TypesBountyRewards, _rewardHash [32]byte) (*types.Transaction, error) {
	return _WitnessHub.Contract.UpdateReward(&_WitnessHub.TransactOpts, _chainID, _blockNumBegin, _blockNumEnd, _operatorsList, _proofRewards, _rewardHash)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_WitnessHub *WitnessHubTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _WitnessHub.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_WitnessHub *WitnessHubSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _WitnessHub.Contract.UpgradeTo(&_WitnessHub.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_WitnessHub *WitnessHubTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _WitnessHub.Contract.UpgradeTo(&_WitnessHub.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WitnessHub *WitnessHubTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WitnessHub.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WitnessHub *WitnessHubSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WitnessHub.Contract.UpgradeToAndCall(&_WitnessHub.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WitnessHub *WitnessHubTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WitnessHub.Contract.UpgradeToAndCall(&_WitnessHub.TransactOpts, newImplementation, data)
}

// WitnessHubAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the WitnessHub contract.
type WitnessHubAdminChangedIterator struct {
	Event *WitnessHubAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WitnessHubAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessHubAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WitnessHubAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WitnessHubAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessHubAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessHubAdminChanged represents a AdminChanged event raised by the WitnessHub contract.
type WitnessHubAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_WitnessHub *WitnessHubFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*WitnessHubAdminChangedIterator, error) {

	logs, sub, err := _WitnessHub.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &WitnessHubAdminChangedIterator{contract: _WitnessHub.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_WitnessHub *WitnessHubFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *WitnessHubAdminChanged) (event.Subscription, error) {

	logs, sub, err := _WitnessHub.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessHubAdminChanged)
				if err := _WitnessHub.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_WitnessHub *WitnessHubFilterer) ParseAdminChanged(log types.Log) (*WitnessHubAdminChanged, error) {
	event := new(WitnessHubAdminChanged)
	if err := _WitnessHub.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessHubAggregatorUpdatedIterator is returned from FilterAggregatorUpdated and is used to iterate over the raw logs and unpacked data for AggregatorUpdated events raised by the WitnessHub contract.
type WitnessHubAggregatorUpdatedIterator struct {
	Event *WitnessHubAggregatorUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WitnessHubAggregatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessHubAggregatorUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WitnessHubAggregatorUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WitnessHubAggregatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessHubAggregatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessHubAggregatorUpdated represents a AggregatorUpdated event raised by the WitnessHub contract.
type WitnessHubAggregatorUpdated struct {
	OldAggregator common.Address
	NewAggregator common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAggregatorUpdated is a free log retrieval operation binding the contract event 0x89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f2.
//
// Solidity: event AggregatorUpdated(address oldAggregator, address newAggregator)
func (_WitnessHub *WitnessHubFilterer) FilterAggregatorUpdated(opts *bind.FilterOpts) (*WitnessHubAggregatorUpdatedIterator, error) {

	logs, sub, err := _WitnessHub.contract.FilterLogs(opts, "AggregatorUpdated")
	if err != nil {
		return nil, err
	}
	return &WitnessHubAggregatorUpdatedIterator{contract: _WitnessHub.contract, event: "AggregatorUpdated", logs: logs, sub: sub}, nil
}

// WatchAggregatorUpdated is a free log subscription operation binding the contract event 0x89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f2.
//
// Solidity: event AggregatorUpdated(address oldAggregator, address newAggregator)
func (_WitnessHub *WitnessHubFilterer) WatchAggregatorUpdated(opts *bind.WatchOpts, sink chan<- *WitnessHubAggregatorUpdated) (event.Subscription, error) {

	logs, sub, err := _WitnessHub.contract.WatchLogs(opts, "AggregatorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessHubAggregatorUpdated)
				if err := _WitnessHub.contract.UnpackLog(event, "AggregatorUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAggregatorUpdated is a log parse operation binding the contract event 0x89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f2.
//
// Solidity: event AggregatorUpdated(address oldAggregator, address newAggregator)
func (_WitnessHub *WitnessHubFilterer) ParseAggregatorUpdated(log types.Log) (*WitnessHubAggregatorUpdated, error) {
	event := new(WitnessHubAggregatorUpdated)
	if err := _WitnessHub.contract.UnpackLog(event, "AggregatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessHubBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the WitnessHub contract.
type WitnessHubBeaconUpgradedIterator struct {
	Event *WitnessHubBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WitnessHubBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessHubBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WitnessHubBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WitnessHubBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessHubBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessHubBeaconUpgraded represents a BeaconUpgraded event raised by the WitnessHub contract.
type WitnessHubBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_WitnessHub *WitnessHubFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*WitnessHubBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _WitnessHub.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &WitnessHubBeaconUpgradedIterator{contract: _WitnessHub.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_WitnessHub *WitnessHubFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *WitnessHubBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _WitnessHub.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessHubBeaconUpgraded)
				if err := _WitnessHub.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_WitnessHub *WitnessHubFilterer) ParseBeaconUpgraded(log types.Log) (*WitnessHubBeaconUpgraded, error) {
	event := new(WitnessHubBeaconUpgraded)
	if err := _WitnessHub.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessHubInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WitnessHub contract.
type WitnessHubInitializedIterator struct {
	Event *WitnessHubInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WitnessHubInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessHubInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WitnessHubInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WitnessHubInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessHubInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessHubInitialized represents a Initialized event raised by the WitnessHub contract.
type WitnessHubInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WitnessHub *WitnessHubFilterer) FilterInitialized(opts *bind.FilterOpts) (*WitnessHubInitializedIterator, error) {

	logs, sub, err := _WitnessHub.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WitnessHubInitializedIterator{contract: _WitnessHub.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WitnessHub *WitnessHubFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WitnessHubInitialized) (event.Subscription, error) {

	logs, sub, err := _WitnessHub.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessHubInitialized)
				if err := _WitnessHub.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WitnessHub *WitnessHubFilterer) ParseInitialized(log types.Log) (*WitnessHubInitialized, error) {
	event := new(WitnessHubInitialized)
	if err := _WitnessHub.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessHubInvalidOperatorIterator is returned from FilterInvalidOperator and is used to iterate over the raw logs and unpacked data for InvalidOperator events raised by the WitnessHub contract.
type WitnessHubInvalidOperatorIterator struct {
	Event *WitnessHubInvalidOperator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WitnessHubInvalidOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessHubInvalidOperator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WitnessHubInvalidOperator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WitnessHubInvalidOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessHubInvalidOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessHubInvalidOperator represents a InvalidOperator event raised by the WitnessHub contract.
type WitnessHubInvalidOperator struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInvalidOperator is a free log retrieval operation binding the contract event 0x3eaa03e38745f9ef282a3d3ae730e428a187a9075ccc2a3b7a1a318fe83d25be.
//
// Solidity: event InvalidOperator(address operator)
func (_WitnessHub *WitnessHubFilterer) FilterInvalidOperator(opts *bind.FilterOpts) (*WitnessHubInvalidOperatorIterator, error) {

	logs, sub, err := _WitnessHub.contract.FilterLogs(opts, "InvalidOperator")
	if err != nil {
		return nil, err
	}
	return &WitnessHubInvalidOperatorIterator{contract: _WitnessHub.contract, event: "InvalidOperator", logs: logs, sub: sub}, nil
}

// WatchInvalidOperator is a free log subscription operation binding the contract event 0x3eaa03e38745f9ef282a3d3ae730e428a187a9075ccc2a3b7a1a318fe83d25be.
//
// Solidity: event InvalidOperator(address operator)
func (_WitnessHub *WitnessHubFilterer) WatchInvalidOperator(opts *bind.WatchOpts, sink chan<- *WitnessHubInvalidOperator) (event.Subscription, error) {

	logs, sub, err := _WitnessHub.contract.WatchLogs(opts, "InvalidOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessHubInvalidOperator)
				if err := _WitnessHub.contract.UnpackLog(event, "InvalidOperator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInvalidOperator is a log parse operation binding the contract event 0x3eaa03e38745f9ef282a3d3ae730e428a187a9075ccc2a3b7a1a318fe83d25be.
//
// Solidity: event InvalidOperator(address operator)
func (_WitnessHub *WitnessHubFilterer) ParseInvalidOperator(log types.Log) (*WitnessHubInvalidOperator, error) {
	event := new(WitnessHubInvalidOperator)
	if err := _WitnessHub.contract.UnpackLog(event, "InvalidOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessHubL2ChainMappingUpdatedIterator is returned from FilterL2ChainMappingUpdated and is used to iterate over the raw logs and unpacked data for L2ChainMappingUpdated events raised by the WitnessHub contract.
type WitnessHubL2ChainMappingUpdatedIterator struct {
	Event *WitnessHubL2ChainMappingUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WitnessHubL2ChainMappingUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessHubL2ChainMappingUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WitnessHubL2ChainMappingUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WitnessHubL2ChainMappingUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessHubL2ChainMappingUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessHubL2ChainMappingUpdated represents a L2ChainMappingUpdated event raised by the WitnessHub contract.
type WitnessHubL2ChainMappingUpdated struct {
	OldL2ChainMapping common.Address
	NewL2ChainMapping common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterL2ChainMappingUpdated is a free log retrieval operation binding the contract event 0xafcfe8816011d6bd68dd7f59cbd0865f721cc58085f72778b69065c75a19b5ca.
//
// Solidity: event L2ChainMappingUpdated(address oldL2ChainMapping, address newL2ChainMapping)
func (_WitnessHub *WitnessHubFilterer) FilterL2ChainMappingUpdated(opts *bind.FilterOpts) (*WitnessHubL2ChainMappingUpdatedIterator, error) {

	logs, sub, err := _WitnessHub.contract.FilterLogs(opts, "L2ChainMappingUpdated")
	if err != nil {
		return nil, err
	}
	return &WitnessHubL2ChainMappingUpdatedIterator{contract: _WitnessHub.contract, event: "L2ChainMappingUpdated", logs: logs, sub: sub}, nil
}

// WatchL2ChainMappingUpdated is a free log subscription operation binding the contract event 0xafcfe8816011d6bd68dd7f59cbd0865f721cc58085f72778b69065c75a19b5ca.
//
// Solidity: event L2ChainMappingUpdated(address oldL2ChainMapping, address newL2ChainMapping)
func (_WitnessHub *WitnessHubFilterer) WatchL2ChainMappingUpdated(opts *bind.WatchOpts, sink chan<- *WitnessHubL2ChainMappingUpdated) (event.Subscription, error) {

	logs, sub, err := _WitnessHub.contract.WatchLogs(opts, "L2ChainMappingUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessHubL2ChainMappingUpdated)
				if err := _WitnessHub.contract.UnpackLog(event, "L2ChainMappingUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseL2ChainMappingUpdated is a log parse operation binding the contract event 0xafcfe8816011d6bd68dd7f59cbd0865f721cc58085f72778b69065c75a19b5ca.
//
// Solidity: event L2ChainMappingUpdated(address oldL2ChainMapping, address newL2ChainMapping)
func (_WitnessHub *WitnessHubFilterer) ParseL2ChainMappingUpdated(log types.Log) (*WitnessHubL2ChainMappingUpdated, error) {
	event := new(WitnessHubL2ChainMappingUpdated)
	if err := _WitnessHub.contract.UnpackLog(event, "L2ChainMappingUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessHubNewRewardsUpdateIterator is returned from FilterNewRewardsUpdate and is used to iterate over the raw logs and unpacked data for NewRewardsUpdate events raised by the WitnessHub contract.
type WitnessHubNewRewardsUpdateIterator struct {
	Event *WitnessHubNewRewardsUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WitnessHubNewRewardsUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessHubNewRewardsUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WitnessHubNewRewardsUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WitnessHubNewRewardsUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessHubNewRewardsUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessHubNewRewardsUpdate represents a NewRewardsUpdate event raised by the WitnessHub contract.
type WitnessHubNewRewardsUpdate struct {
	ChainId          *big.Int
	L2BlockNumberEnd *big.Int
	RewardHash       [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewRewardsUpdate is a free log retrieval operation binding the contract event 0x36b8234ba3e9f8897b5448ed687509b776a4ad94a88ffd28d6aef8bee37ee610.
//
// Solidity: event NewRewardsUpdate(uint256 indexed chainId, uint256 indexed l2BlockNumberEnd, bytes32 indexed rewardHash)
func (_WitnessHub *WitnessHubFilterer) FilterNewRewardsUpdate(opts *bind.FilterOpts, chainId []*big.Int, l2BlockNumberEnd []*big.Int, rewardHash [][32]byte) (*WitnessHubNewRewardsUpdateIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var l2BlockNumberEndRule []interface{}
	for _, l2BlockNumberEndItem := range l2BlockNumberEnd {
		l2BlockNumberEndRule = append(l2BlockNumberEndRule, l2BlockNumberEndItem)
	}
	var rewardHashRule []interface{}
	for _, rewardHashItem := range rewardHash {
		rewardHashRule = append(rewardHashRule, rewardHashItem)
	}

	logs, sub, err := _WitnessHub.contract.FilterLogs(opts, "NewRewardsUpdate", chainIdRule, l2BlockNumberEndRule, rewardHashRule)
	if err != nil {
		return nil, err
	}
	return &WitnessHubNewRewardsUpdateIterator{contract: _WitnessHub.contract, event: "NewRewardsUpdate", logs: logs, sub: sub}, nil
}

// WatchNewRewardsUpdate is a free log subscription operation binding the contract event 0x36b8234ba3e9f8897b5448ed687509b776a4ad94a88ffd28d6aef8bee37ee610.
//
// Solidity: event NewRewardsUpdate(uint256 indexed chainId, uint256 indexed l2BlockNumberEnd, bytes32 indexed rewardHash)
func (_WitnessHub *WitnessHubFilterer) WatchNewRewardsUpdate(opts *bind.WatchOpts, sink chan<- *WitnessHubNewRewardsUpdate, chainId []*big.Int, l2BlockNumberEnd []*big.Int, rewardHash [][32]byte) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var l2BlockNumberEndRule []interface{}
	for _, l2BlockNumberEndItem := range l2BlockNumberEnd {
		l2BlockNumberEndRule = append(l2BlockNumberEndRule, l2BlockNumberEndItem)
	}
	var rewardHashRule []interface{}
	for _, rewardHashItem := range rewardHash {
		rewardHashRule = append(rewardHashRule, rewardHashItem)
	}

	logs, sub, err := _WitnessHub.contract.WatchLogs(opts, "NewRewardsUpdate", chainIdRule, l2BlockNumberEndRule, rewardHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessHubNewRewardsUpdate)
				if err := _WitnessHub.contract.UnpackLog(event, "NewRewardsUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewRewardsUpdate is a log parse operation binding the contract event 0x36b8234ba3e9f8897b5448ed687509b776a4ad94a88ffd28d6aef8bee37ee610.
//
// Solidity: event NewRewardsUpdate(uint256 indexed chainId, uint256 indexed l2BlockNumberEnd, bytes32 indexed rewardHash)
func (_WitnessHub *WitnessHubFilterer) ParseNewRewardsUpdate(log types.Log) (*WitnessHubNewRewardsUpdate, error) {
	event := new(WitnessHubNewRewardsUpdate)
	if err := _WitnessHub.contract.UnpackLog(event, "NewRewardsUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessHubOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WitnessHub contract.
type WitnessHubOwnershipTransferredIterator struct {
	Event *WitnessHubOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WitnessHubOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessHubOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WitnessHubOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WitnessHubOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessHubOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessHubOwnershipTransferred represents a OwnershipTransferred event raised by the WitnessHub contract.
type WitnessHubOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WitnessHub *WitnessHubFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WitnessHubOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WitnessHub.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WitnessHubOwnershipTransferredIterator{contract: _WitnessHub.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WitnessHub *WitnessHubFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WitnessHubOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WitnessHub.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessHubOwnershipTransferred)
				if err := _WitnessHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WitnessHub *WitnessHubFilterer) ParseOwnershipTransferred(log types.Log) (*WitnessHubOwnershipTransferred, error) {
	event := new(WitnessHubOwnershipTransferred)
	if err := _WitnessHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessHubPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the WitnessHub contract.
type WitnessHubPausedIterator struct {
	Event *WitnessHubPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WitnessHubPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessHubPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WitnessHubPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WitnessHubPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessHubPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessHubPaused represents a Paused event raised by the WitnessHub contract.
type WitnessHubPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WitnessHub *WitnessHubFilterer) FilterPaused(opts *bind.FilterOpts) (*WitnessHubPausedIterator, error) {

	logs, sub, err := _WitnessHub.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &WitnessHubPausedIterator{contract: _WitnessHub.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WitnessHub *WitnessHubFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *WitnessHubPaused) (event.Subscription, error) {

	logs, sub, err := _WitnessHub.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessHubPaused)
				if err := _WitnessHub.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WitnessHub *WitnessHubFilterer) ParsePaused(log types.Log) (*WitnessHubPaused, error) {
	event := new(WitnessHubPaused)
	if err := _WitnessHub.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessHubRegistryUpdatedIterator is returned from FilterRegistryUpdated and is used to iterate over the raw logs and unpacked data for RegistryUpdated events raised by the WitnessHub contract.
type WitnessHubRegistryUpdatedIterator struct {
	Event *WitnessHubRegistryUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WitnessHubRegistryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessHubRegistryUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WitnessHubRegistryUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WitnessHubRegistryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessHubRegistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessHubRegistryUpdated represents a RegistryUpdated event raised by the WitnessHub contract.
type WitnessHubRegistryUpdated struct {
	OldRegistry common.Address
	NewRegistry common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegistryUpdated is a free log retrieval operation binding the contract event 0x482b97c53e48ffa324a976e2738053e9aff6eee04d8aac63b10e19411d869b82.
//
// Solidity: event RegistryUpdated(address oldRegistry, address newRegistry)
func (_WitnessHub *WitnessHubFilterer) FilterRegistryUpdated(opts *bind.FilterOpts) (*WitnessHubRegistryUpdatedIterator, error) {

	logs, sub, err := _WitnessHub.contract.FilterLogs(opts, "RegistryUpdated")
	if err != nil {
		return nil, err
	}
	return &WitnessHubRegistryUpdatedIterator{contract: _WitnessHub.contract, event: "RegistryUpdated", logs: logs, sub: sub}, nil
}

// WatchRegistryUpdated is a free log subscription operation binding the contract event 0x482b97c53e48ffa324a976e2738053e9aff6eee04d8aac63b10e19411d869b82.
//
// Solidity: event RegistryUpdated(address oldRegistry, address newRegistry)
func (_WitnessHub *WitnessHubFilterer) WatchRegistryUpdated(opts *bind.WatchOpts, sink chan<- *WitnessHubRegistryUpdated) (event.Subscription, error) {

	logs, sub, err := _WitnessHub.contract.WatchLogs(opts, "RegistryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessHubRegistryUpdated)
				if err := _WitnessHub.contract.UnpackLog(event, "RegistryUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegistryUpdated is a log parse operation binding the contract event 0x482b97c53e48ffa324a976e2738053e9aff6eee04d8aac63b10e19411d869b82.
//
// Solidity: event RegistryUpdated(address oldRegistry, address newRegistry)
func (_WitnessHub *WitnessHubFilterer) ParseRegistryUpdated(log types.Log) (*WitnessHubRegistryUpdated, error) {
	event := new(WitnessHubRegistryUpdated)
	if err := _WitnessHub.contract.UnpackLog(event, "RegistryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessHubSetStrategyParamsIterator is returned from FilterSetStrategyParams and is used to iterate over the raw logs and unpacked data for SetStrategyParams events raised by the WitnessHub contract.
type WitnessHubSetStrategyParamsIterator struct {
	Event *WitnessHubSetStrategyParams // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WitnessHubSetStrategyParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessHubSetStrategyParams)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WitnessHubSetStrategyParams)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WitnessHubSetStrategyParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessHubSetStrategyParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessHubSetStrategyParams represents a SetStrategyParams event raised by the WitnessHub contract.
type WitnessHubSetStrategyParams struct {
	Params []IWitnessHubStrategyParam
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetStrategyParams is a free log retrieval operation binding the contract event 0x78f9b85d45475399beaeb2ca0a5bb0eb3cdc5e60dd4be21f77e1257e641f33fb.
//
// Solidity: event SetStrategyParams((address,uint96)[] params)
func (_WitnessHub *WitnessHubFilterer) FilterSetStrategyParams(opts *bind.FilterOpts) (*WitnessHubSetStrategyParamsIterator, error) {

	logs, sub, err := _WitnessHub.contract.FilterLogs(opts, "SetStrategyParams")
	if err != nil {
		return nil, err
	}
	return &WitnessHubSetStrategyParamsIterator{contract: _WitnessHub.contract, event: "SetStrategyParams", logs: logs, sub: sub}, nil
}

// WatchSetStrategyParams is a free log subscription operation binding the contract event 0x78f9b85d45475399beaeb2ca0a5bb0eb3cdc5e60dd4be21f77e1257e641f33fb.
//
// Solidity: event SetStrategyParams((address,uint96)[] params)
func (_WitnessHub *WitnessHubFilterer) WatchSetStrategyParams(opts *bind.WatchOpts, sink chan<- *WitnessHubSetStrategyParams) (event.Subscription, error) {

	logs, sub, err := _WitnessHub.contract.WatchLogs(opts, "SetStrategyParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessHubSetStrategyParams)
				if err := _WitnessHub.contract.UnpackLog(event, "SetStrategyParams", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetStrategyParams is a log parse operation binding the contract event 0x78f9b85d45475399beaeb2ca0a5bb0eb3cdc5e60dd4be21f77e1257e641f33fb.
//
// Solidity: event SetStrategyParams((address,uint96)[] params)
func (_WitnessHub *WitnessHubFilterer) ParseSetStrategyParams(log types.Log) (*WitnessHubSetStrategyParams, error) {
	event := new(WitnessHubSetStrategyParams)
	if err := _WitnessHub.contract.UnpackLog(event, "SetStrategyParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessHubUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the WitnessHub contract.
type WitnessHubUnpausedIterator struct {
	Event *WitnessHubUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WitnessHubUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessHubUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WitnessHubUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WitnessHubUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessHubUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessHubUnpaused represents a Unpaused event raised by the WitnessHub contract.
type WitnessHubUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WitnessHub *WitnessHubFilterer) FilterUnpaused(opts *bind.FilterOpts) (*WitnessHubUnpausedIterator, error) {

	logs, sub, err := _WitnessHub.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &WitnessHubUnpausedIterator{contract: _WitnessHub.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WitnessHub *WitnessHubFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *WitnessHubUnpaused) (event.Subscription, error) {

	logs, sub, err := _WitnessHub.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessHubUnpaused)
				if err := _WitnessHub.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WitnessHub *WitnessHubFilterer) ParseUnpaused(log types.Log) (*WitnessHubUnpaused, error) {
	event := new(WitnessHubUnpaused)
	if err := _WitnessHub.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessHubUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the WitnessHub contract.
type WitnessHubUpgradedIterator struct {
	Event *WitnessHubUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WitnessHubUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessHubUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WitnessHubUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WitnessHubUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessHubUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessHubUpgraded represents a Upgraded event raised by the WitnessHub contract.
type WitnessHubUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WitnessHub *WitnessHubFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*WitnessHubUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _WitnessHub.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &WitnessHubUpgradedIterator{contract: _WitnessHub.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WitnessHub *WitnessHubFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *WitnessHubUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _WitnessHub.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessHubUpgraded)
				if err := _WitnessHub.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WitnessHub *WitnessHubFilterer) ParseUpgraded(log types.Log) (*WitnessHubUpgraded, error) {
	event := new(WitnessHubUpgraded)
	if err := _WitnessHub.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
