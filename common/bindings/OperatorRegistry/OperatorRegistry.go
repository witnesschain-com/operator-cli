// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OperatorRegistry

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

// OperatorRegistryMetaData contains all meta data concerning the OperatorRegistry contract.
var OperatorRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"OperatorSuspended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"operator\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"OperatorsWhiteListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"watchtower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"WatchtowerDeRegisteredFromOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"watchtower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"WatchtowerRegisteredToOperator\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"operatorsList\",\"type\":\"address[]\"}],\"name\":\"addToOperatorWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"calculateWatchtowerRegistrationMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkIsDelegatedOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"watchtowerAddress\",\"type\":\"address\"}],\"name\":\"deRegister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationManagerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableCheckIsDelegatedOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableCheckIsDelegatedOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllActiveOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"watchtower\",\"type\":\"address\"}],\"name\":\"getOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegationManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_slasherAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isActiveOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"watchtower\",\"type\":\"address\"}],\"name\":\"isValidWatchtower\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorDetails\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"watchtower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signedMessage\",\"type\":\"bytes\"}],\"name\":\"registerWatchtowerAsOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegationManagerAddress\",\"type\":\"address\"}],\"name\":\"setDelegationManagerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_slasherAddress\",\"type\":\"address\"}],\"name\":\"setSlasherAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slasherAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"name\":\"suspend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6040608081526004908136101561001557600080fd5b600091823560e01c806311d2c708146114f857806317fa0ce81461145e578063286781c7146111eb5780633367cca514610f135780633659cfe614610f1f5780633af32abf14610f1357806345adeda014610ee3578063485cc95514610d485780634f1ef28614610a2857806352d1902d146109665780635865c60c1461092a5780635b114af61461077c5780635c468db3146106b75780635c975abb14610693578063715018a6146106335780637831ad20146105ea578063827a1bf5146105545780638da5cb5b1461052b5780639053c5b3146105055780639a521382146104d6578063addd9cf51461048d578063b15e668914610464578063c5e480db14610419578063c8525c3e14610201578063d53c61bf146101d45763f2fde38b14610140575b600080fd5b346101d05760203660031901126101d05761015961190e565b90610162611a37565b6001600160a01b0382161561017e575061017b90611a8f565b51f35b608490602084519162461bcd60e51b8352820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152fd5b8280fd5b5050346101fd57816003193601126101fd5760ff5490516001600160a01b039091168152602090f35b5080fd5b50346101d0576020806003193601126104155781359067ffffffffffffffff928383116104115736602384011215610411578281013593841161041157602490818401938236918760051b01011161040d5761025b611df9565b610263611a37565b865b8581106102e8575050508351928085850186865252606084019290865b8181106102bc575050509180917fdd7d462090c5f3ddd88b4c509269e35d5e148454dc3a9bd24812f78efc8f306f9343908301520390a151f35b90919384359060018060a01b0382168092036102e45790815283019383019190600101610282565b8880fd5b6102fb6102f6828888611d70565b611d96565b8751633af32abf60e01b81526001600160a01b039182168482015285818681305afa908115610403578a916103d6575b501561033b575b50600101610265565b6103496102f6838989611d70565b16885260fd84528688206001908160ff1982541617905561036e6102f6838989611d70565b9060fc805491680100000000000000008310156103c45782019055600192916103be9161039a90611dc2565b90919082549060031b9160018060a01b039283811b93849216901b16911916179055565b90610332565b634e487b7160e01b8c5260418652868cfd5b6103f69150863d88116103fc575b6103ee8183611972565b810190611daa565b3861032b565b503d6103e4565b89513d8c823e3d90fd5b8680fd5b8580fd5b8380fd5b828434610461576020366003190112610461576001600160a01b039060ff9083908361044361190e565b16815260fb60205220548351928116835260a01c1615156020820152f35b80fd5b5050346101fd57816003193601126101fd5760fe5490516001600160a01b039091168152602090f35b5050346101fd5760203660031901126101fd576104a861190e565b6104b0611df9565b6104b8611a37565b60018060a01b03166001600160601b0360a01b60ff54161760ff5551f35b5050346101fd5760203660031901126101fd576020906104fc6104f761190e565b61201e565b90519015158152f35b5050346101fd57816003193601126101fd5760209060ff805460a01c1690519015158152f35b5050346101fd57816003193601126101fd5760335490516001600160a01b039091168152602090f35b50346101d057826003193601126101d05761056d611a37565b60ff549060ff8260a01c161561058b575060ff60a01b191660ff5551f35b608490602084519162461bcd60e51b8352820152603360248201527f5769746e6573734875623a20454c2064656c65676174696f6e20636865636b206044820152721a5cc8185b1c9958591e48191a5cd8589b1959606a1b6064820152fd5b5050346101fd5760203660031901126101fd5761060561190e565b61060d611df9565b610615611a37565b60018060a01b03166001600160601b0360a01b60fe54161760fe5551f35b5050346101fd57816003193601126101fd5761064d611a37565b603380546001600160a01b0319811690915590519082906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08284a3f35b5050346101fd57816003193601126101fd5760209060ff6065541690519015158152f35b82843461046157806003193601126104615790805191829060fc54918285526020809501809360fc84527f371f36870d18f32a11fea0f144b021c8b407bb50f8e0267c711123f454b963c090845b81811061075f575050508161071b910382611972565b83519485948186019282875251809352850193925b82811061073f57505050500390f35b83516001600160a01b031685528695509381019392810192600101610730565b82546001600160a01b031684529288019260019283019201610705565b50346101d057602090816003193601126104155761079861190e565b916107a1611df9565b6001600160a01b0383811680875261010080845286882054909491921633036108aa5733875260fd835260ff868820541615610835575085525281832080546001600160a01b031916905581513381526001600160a01b0390911660208201524360408201527fecfc290ff8c3aac71e14aee07653f81e5aa316be4d2315ba2ec1bff9dc50cd799080606081015b0390a151f35b60a4908387519162461bcd60e51b8352820152604460248201527f5769746e6573734875623a204465726567697374726174696f6e2063616e206260448201527f6520646f6e65206f6e6c79206f6e2077686974656c6973746564206f70657261606482015263746f727360e01b6084820152fd5b855162461bcd60e51b8152908101839052604e60248201527f5769746e6573734875623a204465726567697374726174696f6e2073686f756c60448201527f6420626520646f6e65206f6e206f70657261746f72277320726567697374657260648201526d6564207761746368746f7765727360901b608482015260a490fd5b5050346101fd5760203660031901126101fd576020916001600160a01b039082908261095461190e565b16815261010085522054169051908152f35b508234610461578060031936011261046157507f000000000000000000000000f824d48eeb6759ec9a6ac7e5bdda68f11764e7d36001600160a01b031630036109c057602082516000805160206120b88339815191528152f35b6020608492519162461bcd60e51b8352820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152fd5b5090806003193601126101d057610a3d61190e565b60243567ffffffffffffffff8111610d4457610a5c90369085016119b0565b6001600160a01b037f000000000000000000000000f824d48eeb6759ec9a6ac7e5bdda68f11764e7d381169190610a9530841415611ada565b610ab26000805160206120b8833981519152938285541614611b3b565b610aba611a37565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610af25750505061017b919250611b9c565b84516352d1902d60e01b8152602095949293918316919086818981865afa899181610d11575b50610b7757855162461bcd60e51b8152808901889052602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608490fd5b96919293949603610cbc57610b8b83611b9c565b8551917fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8884a2835115801590610cb4575b610bca575b505050505051f35b823b15610c655750508482819285610c529695519201905af43d15610c5d573d90610bf482611994565b91610c0186519384611972565b82523d868484013e5b7f416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c855193610c3785611956565b60278552840152660819985a5b195960ca1b85840152611c2c565b503880808080610bc2565b606090610c0a565b62461bcd60e51b82528101849052602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608490fd5b506001610bbd565b855162461bcd60e51b8152908101859052602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608490fd5b9091508781813d8311610d3d575b610d298183611972565b81010312610d3957519038610b18565b8980fd5b503d610d1f565b8480fd5b50346101d057816003193601126101d057610d6161190e565b6001600160a01b039190602435838116908190036104115785549160ff94858460081c161594858096610ed7575b8015610ec1575b15610e67575090610e1a93929160ff1993866001868316178b55610e56575b506001600160601b0360a01b91168160fe54161760fe55855416178455610dea84875460081c16610de581611cb3565b611cb3565b610df333611a8f565b8554938460081c1690610e0582611cb3565b610e0e82611cb3565b60655416606555611cb3565b610e22575051f35b61ff00191682557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986020825160018152a151f35b61ffff191661010117895538610db5565b608490602089519162461bcd60e51b8352820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152fd5b50303b158015610d965750600187861614610d96565b50600187861610610d8f565b5050346101fd57806003193601126101fd57602090610f0c610f0361190e565b60243590611d13565b9051908152f35b5050505061013b6119f7565b5090346101d05760208060031936011261041557610f3b61190e565b906001600160a01b037f000000000000000000000000f824d48eeb6759ec9a6ac7e5bdda68f11764e7d38116610f7330821415611ada565b610f906000805160206120b8833981519152918383541614611b3b565b610f98611a37565b8451908382019282841067ffffffffffffffff8511176111d8578387528883527f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610ff257505050505061017b919250611b9c565b85949695169085516352d1902d60e01b815287818a81865afa8a91816111a5575b5061107257865162461bcd60e51b8152808a01899052602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608490fd5b97919293949597036111505761108785611b9c565b8651917fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8984a2825115801590611149575b6110c7575b50505050505051f35b843b156110fa5750506110ee93928692839251915af43d15610c5d573d90610bf482611994565b503880808080806110be565b62461bcd60e51b82528101859052602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608490fd5b50876110b9565b865162461bcd60e51b8152908101869052602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608490fd5b9091508881813d83116111d1575b6111bd8183611972565b810103126111cd57519038611013565b8a80fd5b503d6111b3565b634e487b7160e01b895260418852602489fd5b5090346101d0576020806003193601126104155761120761190e565b90611210611df9565b611218611a37565b6001600160a01b0382811680875260fd83528487205490919060ff16156113de57611241611a37565b865b60fc8054808310156113a8579083918561125c85611dc2565b949054600395861b1c161461127657505050600101611243565b91939495969798909260018310611395576112aa9061039a8761129d600019809701611dc2565b905490891b1c1691611dc2565b82549081156113825750927f98ff055e829cd4bfba35f84d3b43877fd53084d2deebbb1cda57404f44df4def97959285611363969361082f999601926112ef84611dc2565b81939154921b1b19169055555b81895260fd835287892060ff19815416905560fb835260fb888a2093828a519561132587611924565b541685528085018b8152938b525287892092518354925160ff60a01b90151560a01b166001600160a81b0319909316911660ff60a01b191617179055565b83516001600160a01b0390911681524360208201529081906040820190565b634e487b7160e01b8b526031905260248afd5b634e487b7160e01b8b526011825260248bfd5b5050507f98ff055e829cd4bfba35f84d3b43877fd53084d2deebbb1cda57404f44df4def949550916113639161082f94936112fc565b845162461bcd60e51b8152808701849052604e60248201527f5769746e6573734875623a2043616e6e6f742073757370656e64206966206f7060448201527f657261746f7220697320616c72656164792073757370656e646564206f72206e60648201526d1bdd081dda1a5d195b1a5cdd195960921b608482015260a490fd5b50346101d057826003193601126101d057611477611a37565b60ff549060ff8260a01c1661149a575060ff60a01b1916600160a01b1760ff5551f35b608490602084519162461bcd60e51b8352820152603260248201527f5769746e6573734875623a20454c2064656c65676174696f6e20636865636b206044820152711a5cc8185b1c9958591e48195b98589b195960721b6064820152fd5b50346101d05760603660031901126101d05761151261190e565b90602490813591604490813567ffffffffffffffff811161190a5761153a90369085016119b0565b92611543611df9565b33885260209260fd845260ff888a205416156118a55760ff805460a01c166117e8575b60018060a01b039283881695868b5261010097888752858b8d20541661178d57871561172d574281106116db576115b36115ab89936115a6899433611d13565b611f57565b919091611e3d565b160361167c57505050927f3d521ffddd8cbc2d603c1e7fe9af4b70adbbc43675a7281c5624d8053e526f0694926116579261082f95338a5260fb82528851936115fb85611924565b338086526001848701908152948c52908352898b2080546001600160a01b031916821790558a5260fb90915287892092518354925160ff60a01b90151560a01b166001600160a81b0319909316911660ff60a01b191617179055565b83513381526001600160a01b0390911660208201524360408201529081906060820190565b885162461bcd60e51b81529283018590526035908301527f5769746e6573734875623a20526567697374726174696f6e207369676e657220908201527434b9903737ba103a3432903bb0ba31b43a37bbb2b960591b6064820152608490fd5b8a5162461bcd60e51b81528086018890526028818501527f5769746e6573734875623a207761746368746f776572207369676e6174757265818601526708195e1c1a5c995960c21b6064820152608490fd5b8a5162461bcd60e51b81528086018890526036818501527f5769746e6573734875623a205761746368746f776572206164647265737320638186015275616e6e6f74206265207468652030206164647265737360501b6064820152608490fd5b8a5162461bcd60e51b81528086018890526031818501527f5769746e6573734875623a205761746368746f7765722061646472657373206181860152701b1c9958591e481c9959da5cdd195c9959607a1b6064820152608490fd5b60fe5488516336b87bd760e11b815233848201529085908290869082906001600160a01b03165afa908115610403578a91611888575b5061156657603f7f5769746e6573734875623a20596f75206e65656420746f20626520612064656c92936084958a519562461bcd60e51b87528601528401528201527f656761746564206f70657261746f72207769746820456967656e4c61796572006064820152fd5b61189f9150853d87116103fc576103ee8183611972565b3861181e565b603e7f5769746e6573734875623a204f70657261746f72206973206e6f74207768697492936084958a519562461bcd60e51b87528601528401528201527f656c69737465642077697468205769746e65737320436861696e2041565300006064820152fd5b8780fd5b600435906001600160a01b038216820361013b57565b6040810190811067ffffffffffffffff82111761194057604052565b634e487b7160e01b600052604160045260246000fd5b6060810190811067ffffffffffffffff82111761194057604052565b90601f8019910116810190811067ffffffffffffffff82111761194057604052565b67ffffffffffffffff811161194057601f01601f191660200190565b81601f8201121561013b578035906119c782611994565b926119d56040519485611972565b8284526020838301011161013b57816000926020809301838601378301015290565b503461013b57602036600319011261013b576001600160a01b03611a1961190e565b1660005260fd602052602060ff604060002054166040519015158152f35b6033546001600160a01b03163303611a4b57565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b6033549060018060a01b0380911691826001600160601b0360a01b821617603355167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e06000604051a3565b15611ae157565b60405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b6064820152608490fd5b15611b4257565b60405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b6163746976652070726f787960a01b6064820152608490fd5b803b15611bd1576000805160206120b883398151915280546001600160a01b0319166001600160a01b03909216919091179055565b60405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608490fd5b90919015611c38575090565b815115611c485750805190602001fd5b6040519062461bcd60e51b82528160208060048301528251928360248401526000915b848310611c9a575050918060449311611c8d575b601f01601f19168101030190fd5b6000838284010152611c7f565b8183018101518684016044015285935091820191611c6b565b15611cba57565b60405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608490fd5b604080516001600160a01b039290921660208301908152828201939093528152611d3e606082611972565b5190207f19457468657265756d205369676e6564204d6573736167653a0a333200000000600052601c52603c60002090565b9190811015611d805760051b0190565b634e487b7160e01b600052603260045260246000fd5b356001600160a01b038116810361013b5790565b9081602091031261013b5751801515810361013b5790565b60fc54811015611d805760fc6000527f371f36870d18f32a11fea0f144b021c8b407bb50f8e0267c711123f454b963c00190600090565b60ff60655416611e0557565b60405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606490fd5b6005811015611f415780611e4e5750565b60018103611e9b5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606490fd5b60028103611ee85760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606490fd5b600314611ef157565b60405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608490fd5b634e487b7160e01b600052602160045260246000fd5b906041815114600014611f8557611f81916020820151906060604084015193015160001a90611f8f565b9091565b5050600090600290565b9291907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083116120125791608094939160ff602094604051948552168484015260408301526060820152600093849182805260015afa156120055781516001600160a01b03811615611fff579190565b50600190565b50604051903d90823e3d90fd5b50505050600090600390565b6001600160a01b03908116801561205f576000526101006020526040600020541680156120595760005260fd60205260ff6040600020541690565b50600090565b60405162461bcd60e51b815260206004820152602a60248201527f5769746e6573734875623a205761746368746f776572206164647265737320636044820152690616e6e6f7420626520360b41b6064820152608490fdfe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca26469706673582212205bdf1a237006b07caee1b524628f7e114c6b78b5522fa5cf1bd5a4c7ba8779b064736f6c634300080f0033",
}

// OperatorRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use OperatorRegistryMetaData.ABI instead.
var OperatorRegistryABI = OperatorRegistryMetaData.ABI

// OperatorRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OperatorRegistryMetaData.Bin instead.
var OperatorRegistryBin = OperatorRegistryMetaData.Bin

// DeployOperatorRegistry deploys a new Ethereum contract, binding an instance of OperatorRegistry to it.
func DeployOperatorRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OperatorRegistry, error) {
	parsed, err := OperatorRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OperatorRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OperatorRegistry{OperatorRegistryCaller: OperatorRegistryCaller{contract: contract}, OperatorRegistryTransactor: OperatorRegistryTransactor{contract: contract}, OperatorRegistryFilterer: OperatorRegistryFilterer{contract: contract}}, nil
}

// OperatorRegistry is an auto generated Go binding around an Ethereum contract.
type OperatorRegistry struct {
	OperatorRegistryCaller     // Read-only binding to the contract
	OperatorRegistryTransactor // Write-only binding to the contract
	OperatorRegistryFilterer   // Log filterer for contract events
}

// OperatorRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OperatorRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OperatorRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OperatorRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OperatorRegistrySession struct {
	Contract     *OperatorRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OperatorRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OperatorRegistryCallerSession struct {
	Contract *OperatorRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// OperatorRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OperatorRegistryTransactorSession struct {
	Contract     *OperatorRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// OperatorRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OperatorRegistryRaw struct {
	Contract *OperatorRegistry // Generic contract binding to access the raw methods on
}

// OperatorRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OperatorRegistryCallerRaw struct {
	Contract *OperatorRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// OperatorRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OperatorRegistryTransactorRaw struct {
	Contract *OperatorRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOperatorRegistry creates a new instance of OperatorRegistry, bound to a specific deployed contract.
func NewOperatorRegistry(address common.Address, backend bind.ContractBackend) (*OperatorRegistry, error) {
	contract, err := bindOperatorRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OperatorRegistry{OperatorRegistryCaller: OperatorRegistryCaller{contract: contract}, OperatorRegistryTransactor: OperatorRegistryTransactor{contract: contract}, OperatorRegistryFilterer: OperatorRegistryFilterer{contract: contract}}, nil
}

// NewOperatorRegistryCaller creates a new read-only instance of OperatorRegistry, bound to a specific deployed contract.
func NewOperatorRegistryCaller(address common.Address, caller bind.ContractCaller) (*OperatorRegistryCaller, error) {
	contract, err := bindOperatorRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryCaller{contract: contract}, nil
}

// NewOperatorRegistryTransactor creates a new write-only instance of OperatorRegistry, bound to a specific deployed contract.
func NewOperatorRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*OperatorRegistryTransactor, error) {
	contract, err := bindOperatorRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryTransactor{contract: contract}, nil
}

// NewOperatorRegistryFilterer creates a new log filterer instance of OperatorRegistry, bound to a specific deployed contract.
func NewOperatorRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*OperatorRegistryFilterer, error) {
	contract, err := bindOperatorRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryFilterer{contract: contract}, nil
}

// bindOperatorRegistry binds a generic wrapper to an already deployed contract.
func bindOperatorRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OperatorRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperatorRegistry *OperatorRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperatorRegistry.Contract.OperatorRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperatorRegistry *OperatorRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.OperatorRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperatorRegistry *OperatorRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.OperatorRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperatorRegistry *OperatorRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperatorRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperatorRegistry *OperatorRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperatorRegistry *OperatorRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.contract.Transact(opts, method, params...)
}

// CalculateWatchtowerRegistrationMessageHash is a free data retrieval call binding the contract method 0x45adeda0.
//
// Solidity: function calculateWatchtowerRegistrationMessageHash(address operator, uint256 expiry) pure returns(bytes32)
func (_OperatorRegistry *OperatorRegistryCaller) CalculateWatchtowerRegistrationMessageHash(opts *bind.CallOpts, operator common.Address, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "calculateWatchtowerRegistrationMessageHash", operator, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateWatchtowerRegistrationMessageHash is a free data retrieval call binding the contract method 0x45adeda0.
//
// Solidity: function calculateWatchtowerRegistrationMessageHash(address operator, uint256 expiry) pure returns(bytes32)
func (_OperatorRegistry *OperatorRegistrySession) CalculateWatchtowerRegistrationMessageHash(operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _OperatorRegistry.Contract.CalculateWatchtowerRegistrationMessageHash(&_OperatorRegistry.CallOpts, operator, expiry)
}

// CalculateWatchtowerRegistrationMessageHash is a free data retrieval call binding the contract method 0x45adeda0.
//
// Solidity: function calculateWatchtowerRegistrationMessageHash(address operator, uint256 expiry) pure returns(bytes32)
func (_OperatorRegistry *OperatorRegistryCallerSession) CalculateWatchtowerRegistrationMessageHash(operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _OperatorRegistry.Contract.CalculateWatchtowerRegistrationMessageHash(&_OperatorRegistry.CallOpts, operator, expiry)
}

// CheckIsDelegatedOperator is a free data retrieval call binding the contract method 0x9053c5b3.
//
// Solidity: function checkIsDelegatedOperator() view returns(bool)
func (_OperatorRegistry *OperatorRegistryCaller) CheckIsDelegatedOperator(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "checkIsDelegatedOperator")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckIsDelegatedOperator is a free data retrieval call binding the contract method 0x9053c5b3.
//
// Solidity: function checkIsDelegatedOperator() view returns(bool)
func (_OperatorRegistry *OperatorRegistrySession) CheckIsDelegatedOperator() (bool, error) {
	return _OperatorRegistry.Contract.CheckIsDelegatedOperator(&_OperatorRegistry.CallOpts)
}

// CheckIsDelegatedOperator is a free data retrieval call binding the contract method 0x9053c5b3.
//
// Solidity: function checkIsDelegatedOperator() view returns(bool)
func (_OperatorRegistry *OperatorRegistryCallerSession) CheckIsDelegatedOperator() (bool, error) {
	return _OperatorRegistry.Contract.CheckIsDelegatedOperator(&_OperatorRegistry.CallOpts)
}

// DelegationManagerAddress is a free data retrieval call binding the contract method 0xb15e6689.
//
// Solidity: function delegationManagerAddress() view returns(address)
func (_OperatorRegistry *OperatorRegistryCaller) DelegationManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "delegationManagerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManagerAddress is a free data retrieval call binding the contract method 0xb15e6689.
//
// Solidity: function delegationManagerAddress() view returns(address)
func (_OperatorRegistry *OperatorRegistrySession) DelegationManagerAddress() (common.Address, error) {
	return _OperatorRegistry.Contract.DelegationManagerAddress(&_OperatorRegistry.CallOpts)
}

// DelegationManagerAddress is a free data retrieval call binding the contract method 0xb15e6689.
//
// Solidity: function delegationManagerAddress() view returns(address)
func (_OperatorRegistry *OperatorRegistryCallerSession) DelegationManagerAddress() (common.Address, error) {
	return _OperatorRegistry.Contract.DelegationManagerAddress(&_OperatorRegistry.CallOpts)
}

// GetAllActiveOperators is a free data retrieval call binding the contract method 0x5c468db3.
//
// Solidity: function getAllActiveOperators() view returns(address[])
func (_OperatorRegistry *OperatorRegistryCaller) GetAllActiveOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "getAllActiveOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllActiveOperators is a free data retrieval call binding the contract method 0x5c468db3.
//
// Solidity: function getAllActiveOperators() view returns(address[])
func (_OperatorRegistry *OperatorRegistrySession) GetAllActiveOperators() ([]common.Address, error) {
	return _OperatorRegistry.Contract.GetAllActiveOperators(&_OperatorRegistry.CallOpts)
}

// GetAllActiveOperators is a free data retrieval call binding the contract method 0x5c468db3.
//
// Solidity: function getAllActiveOperators() view returns(address[])
func (_OperatorRegistry *OperatorRegistryCallerSession) GetAllActiveOperators() ([]common.Address, error) {
	return _OperatorRegistry.Contract.GetAllActiveOperators(&_OperatorRegistry.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address watchtower) view returns(address operator)
func (_OperatorRegistry *OperatorRegistryCaller) GetOperator(opts *bind.CallOpts, watchtower common.Address) (common.Address, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "getOperator", watchtower)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address watchtower) view returns(address operator)
func (_OperatorRegistry *OperatorRegistrySession) GetOperator(watchtower common.Address) (common.Address, error) {
	return _OperatorRegistry.Contract.GetOperator(&_OperatorRegistry.CallOpts, watchtower)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address watchtower) view returns(address operator)
func (_OperatorRegistry *OperatorRegistryCallerSession) GetOperator(watchtower common.Address) (common.Address, error) {
	return _OperatorRegistry.Contract.GetOperator(&_OperatorRegistry.CallOpts, watchtower)
}

// IsActiveOperator is a free data retrieval call binding the contract method 0x3367cca5.
//
// Solidity: function isActiveOperator(address operator) view returns(bool)
func (_OperatorRegistry *OperatorRegistryCaller) IsActiveOperator(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "isActiveOperator", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveOperator is a free data retrieval call binding the contract method 0x3367cca5.
//
// Solidity: function isActiveOperator(address operator) view returns(bool)
func (_OperatorRegistry *OperatorRegistrySession) IsActiveOperator(operator common.Address) (bool, error) {
	return _OperatorRegistry.Contract.IsActiveOperator(&_OperatorRegistry.CallOpts, operator)
}

// IsActiveOperator is a free data retrieval call binding the contract method 0x3367cca5.
//
// Solidity: function isActiveOperator(address operator) view returns(bool)
func (_OperatorRegistry *OperatorRegistryCallerSession) IsActiveOperator(operator common.Address) (bool, error) {
	return _OperatorRegistry.Contract.IsActiveOperator(&_OperatorRegistry.CallOpts, operator)
}

// IsValidWatchtower is a free data retrieval call binding the contract method 0x9a521382.
//
// Solidity: function isValidWatchtower(address watchtower) view returns(bool)
func (_OperatorRegistry *OperatorRegistryCaller) IsValidWatchtower(opts *bind.CallOpts, watchtower common.Address) (bool, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "isValidWatchtower", watchtower)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidWatchtower is a free data retrieval call binding the contract method 0x9a521382.
//
// Solidity: function isValidWatchtower(address watchtower) view returns(bool)
func (_OperatorRegistry *OperatorRegistrySession) IsValidWatchtower(watchtower common.Address) (bool, error) {
	return _OperatorRegistry.Contract.IsValidWatchtower(&_OperatorRegistry.CallOpts, watchtower)
}

// IsValidWatchtower is a free data retrieval call binding the contract method 0x9a521382.
//
// Solidity: function isValidWatchtower(address watchtower) view returns(bool)
func (_OperatorRegistry *OperatorRegistryCallerSession) IsValidWatchtower(watchtower common.Address) (bool, error) {
	return _OperatorRegistry.Contract.IsValidWatchtower(&_OperatorRegistry.CallOpts, watchtower)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address operator) view returns(bool)
func (_OperatorRegistry *OperatorRegistryCaller) IsWhitelisted(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "isWhitelisted", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address operator) view returns(bool)
func (_OperatorRegistry *OperatorRegistrySession) IsWhitelisted(operator common.Address) (bool, error) {
	return _OperatorRegistry.Contract.IsWhitelisted(&_OperatorRegistry.CallOpts, operator)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address operator) view returns(bool)
func (_OperatorRegistry *OperatorRegistryCallerSession) IsWhitelisted(operator common.Address) (bool, error) {
	return _OperatorRegistry.Contract.IsWhitelisted(&_OperatorRegistry.CallOpts, operator)
}

// OperatorDetails is a free data retrieval call binding the contract method 0xc5e480db.
//
// Solidity: function operatorDetails(address ) view returns(address operatorAddress, bool isActive)
func (_OperatorRegistry *OperatorRegistryCaller) OperatorDetails(opts *bind.CallOpts, arg0 common.Address) (struct {
	OperatorAddress common.Address
	IsActive        bool
}, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "operatorDetails", arg0)

	outstruct := new(struct {
		OperatorAddress common.Address
		IsActive        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OperatorAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IsActive = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// OperatorDetails is a free data retrieval call binding the contract method 0xc5e480db.
//
// Solidity: function operatorDetails(address ) view returns(address operatorAddress, bool isActive)
func (_OperatorRegistry *OperatorRegistrySession) OperatorDetails(arg0 common.Address) (struct {
	OperatorAddress common.Address
	IsActive        bool
}, error) {
	return _OperatorRegistry.Contract.OperatorDetails(&_OperatorRegistry.CallOpts, arg0)
}

// OperatorDetails is a free data retrieval call binding the contract method 0xc5e480db.
//
// Solidity: function operatorDetails(address ) view returns(address operatorAddress, bool isActive)
func (_OperatorRegistry *OperatorRegistryCallerSession) OperatorDetails(arg0 common.Address) (struct {
	OperatorAddress common.Address
	IsActive        bool
}, error) {
	return _OperatorRegistry.Contract.OperatorDetails(&_OperatorRegistry.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OperatorRegistry *OperatorRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OperatorRegistry *OperatorRegistrySession) Owner() (common.Address, error) {
	return _OperatorRegistry.Contract.Owner(&_OperatorRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OperatorRegistry *OperatorRegistryCallerSession) Owner() (common.Address, error) {
	return _OperatorRegistry.Contract.Owner(&_OperatorRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OperatorRegistry *OperatorRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OperatorRegistry *OperatorRegistrySession) Paused() (bool, error) {
	return _OperatorRegistry.Contract.Paused(&_OperatorRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OperatorRegistry *OperatorRegistryCallerSession) Paused() (bool, error) {
	return _OperatorRegistry.Contract.Paused(&_OperatorRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_OperatorRegistry *OperatorRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_OperatorRegistry *OperatorRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _OperatorRegistry.Contract.ProxiableUUID(&_OperatorRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_OperatorRegistry *OperatorRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _OperatorRegistry.Contract.ProxiableUUID(&_OperatorRegistry.CallOpts)
}

// SlasherAddress is a free data retrieval call binding the contract method 0xd53c61bf.
//
// Solidity: function slasherAddress() view returns(address)
func (_OperatorRegistry *OperatorRegistryCaller) SlasherAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OperatorRegistry.contract.Call(opts, &out, "slasherAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SlasherAddress is a free data retrieval call binding the contract method 0xd53c61bf.
//
// Solidity: function slasherAddress() view returns(address)
func (_OperatorRegistry *OperatorRegistrySession) SlasherAddress() (common.Address, error) {
	return _OperatorRegistry.Contract.SlasherAddress(&_OperatorRegistry.CallOpts)
}

// SlasherAddress is a free data retrieval call binding the contract method 0xd53c61bf.
//
// Solidity: function slasherAddress() view returns(address)
func (_OperatorRegistry *OperatorRegistryCallerSession) SlasherAddress() (common.Address, error) {
	return _OperatorRegistry.Contract.SlasherAddress(&_OperatorRegistry.CallOpts)
}

// AddToOperatorWhitelist is a paid mutator transaction binding the contract method 0xc8525c3e.
//
// Solidity: function addToOperatorWhitelist(address[] operatorsList) returns()
func (_OperatorRegistry *OperatorRegistryTransactor) AddToOperatorWhitelist(opts *bind.TransactOpts, operatorsList []common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "addToOperatorWhitelist", operatorsList)
}

// AddToOperatorWhitelist is a paid mutator transaction binding the contract method 0xc8525c3e.
//
// Solidity: function addToOperatorWhitelist(address[] operatorsList) returns()
func (_OperatorRegistry *OperatorRegistrySession) AddToOperatorWhitelist(operatorsList []common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.AddToOperatorWhitelist(&_OperatorRegistry.TransactOpts, operatorsList)
}

// AddToOperatorWhitelist is a paid mutator transaction binding the contract method 0xc8525c3e.
//
// Solidity: function addToOperatorWhitelist(address[] operatorsList) returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) AddToOperatorWhitelist(operatorsList []common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.AddToOperatorWhitelist(&_OperatorRegistry.TransactOpts, operatorsList)
}

// DeRegister is a paid mutator transaction binding the contract method 0x5b114af6.
//
// Solidity: function deRegister(address watchtowerAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactor) DeRegister(opts *bind.TransactOpts, watchtowerAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "deRegister", watchtowerAddress)
}

// DeRegister is a paid mutator transaction binding the contract method 0x5b114af6.
//
// Solidity: function deRegister(address watchtowerAddress) returns()
func (_OperatorRegistry *OperatorRegistrySession) DeRegister(watchtowerAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.DeRegister(&_OperatorRegistry.TransactOpts, watchtowerAddress)
}

// DeRegister is a paid mutator transaction binding the contract method 0x5b114af6.
//
// Solidity: function deRegister(address watchtowerAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) DeRegister(watchtowerAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.DeRegister(&_OperatorRegistry.TransactOpts, watchtowerAddress)
}

// DisableCheckIsDelegatedOperator is a paid mutator transaction binding the contract method 0x827a1bf5.
//
// Solidity: function disableCheckIsDelegatedOperator() returns()
func (_OperatorRegistry *OperatorRegistryTransactor) DisableCheckIsDelegatedOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "disableCheckIsDelegatedOperator")
}

// DisableCheckIsDelegatedOperator is a paid mutator transaction binding the contract method 0x827a1bf5.
//
// Solidity: function disableCheckIsDelegatedOperator() returns()
func (_OperatorRegistry *OperatorRegistrySession) DisableCheckIsDelegatedOperator() (*types.Transaction, error) {
	return _OperatorRegistry.Contract.DisableCheckIsDelegatedOperator(&_OperatorRegistry.TransactOpts)
}

// DisableCheckIsDelegatedOperator is a paid mutator transaction binding the contract method 0x827a1bf5.
//
// Solidity: function disableCheckIsDelegatedOperator() returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) DisableCheckIsDelegatedOperator() (*types.Transaction, error) {
	return _OperatorRegistry.Contract.DisableCheckIsDelegatedOperator(&_OperatorRegistry.TransactOpts)
}

// EnableCheckIsDelegatedOperator is a paid mutator transaction binding the contract method 0x17fa0ce8.
//
// Solidity: function enableCheckIsDelegatedOperator() returns()
func (_OperatorRegistry *OperatorRegistryTransactor) EnableCheckIsDelegatedOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "enableCheckIsDelegatedOperator")
}

// EnableCheckIsDelegatedOperator is a paid mutator transaction binding the contract method 0x17fa0ce8.
//
// Solidity: function enableCheckIsDelegatedOperator() returns()
func (_OperatorRegistry *OperatorRegistrySession) EnableCheckIsDelegatedOperator() (*types.Transaction, error) {
	return _OperatorRegistry.Contract.EnableCheckIsDelegatedOperator(&_OperatorRegistry.TransactOpts)
}

// EnableCheckIsDelegatedOperator is a paid mutator transaction binding the contract method 0x17fa0ce8.
//
// Solidity: function enableCheckIsDelegatedOperator() returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) EnableCheckIsDelegatedOperator() (*types.Transaction, error) {
	return _OperatorRegistry.Contract.EnableCheckIsDelegatedOperator(&_OperatorRegistry.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _delegationManagerAddress, address _slasherAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactor) Initialize(opts *bind.TransactOpts, _delegationManagerAddress common.Address, _slasherAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "initialize", _delegationManagerAddress, _slasherAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _delegationManagerAddress, address _slasherAddress) returns()
func (_OperatorRegistry *OperatorRegistrySession) Initialize(_delegationManagerAddress common.Address, _slasherAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.Initialize(&_OperatorRegistry.TransactOpts, _delegationManagerAddress, _slasherAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _delegationManagerAddress, address _slasherAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) Initialize(_delegationManagerAddress common.Address, _slasherAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.Initialize(&_OperatorRegistry.TransactOpts, _delegationManagerAddress, _slasherAddress)
}

// RegisterWatchtowerAsOperator is a paid mutator transaction binding the contract method 0x11d2c708.
//
// Solidity: function registerWatchtowerAsOperator(address watchtower, uint256 expiry, bytes signedMessage) returns()
func (_OperatorRegistry *OperatorRegistryTransactor) RegisterWatchtowerAsOperator(opts *bind.TransactOpts, watchtower common.Address, expiry *big.Int, signedMessage []byte) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "registerWatchtowerAsOperator", watchtower, expiry, signedMessage)
}

// RegisterWatchtowerAsOperator is a paid mutator transaction binding the contract method 0x11d2c708.
//
// Solidity: function registerWatchtowerAsOperator(address watchtower, uint256 expiry, bytes signedMessage) returns()
func (_OperatorRegistry *OperatorRegistrySession) RegisterWatchtowerAsOperator(watchtower common.Address, expiry *big.Int, signedMessage []byte) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.RegisterWatchtowerAsOperator(&_OperatorRegistry.TransactOpts, watchtower, expiry, signedMessage)
}

// RegisterWatchtowerAsOperator is a paid mutator transaction binding the contract method 0x11d2c708.
//
// Solidity: function registerWatchtowerAsOperator(address watchtower, uint256 expiry, bytes signedMessage) returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) RegisterWatchtowerAsOperator(watchtower common.Address, expiry *big.Int, signedMessage []byte) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.RegisterWatchtowerAsOperator(&_OperatorRegistry.TransactOpts, watchtower, expiry, signedMessage)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OperatorRegistry *OperatorRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OperatorRegistry *OperatorRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _OperatorRegistry.Contract.RenounceOwnership(&_OperatorRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OperatorRegistry.Contract.RenounceOwnership(&_OperatorRegistry.TransactOpts)
}

// SetDelegationManagerAddress is a paid mutator transaction binding the contract method 0x7831ad20.
//
// Solidity: function setDelegationManagerAddress(address _delegationManagerAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactor) SetDelegationManagerAddress(opts *bind.TransactOpts, _delegationManagerAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "setDelegationManagerAddress", _delegationManagerAddress)
}

// SetDelegationManagerAddress is a paid mutator transaction binding the contract method 0x7831ad20.
//
// Solidity: function setDelegationManagerAddress(address _delegationManagerAddress) returns()
func (_OperatorRegistry *OperatorRegistrySession) SetDelegationManagerAddress(_delegationManagerAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.SetDelegationManagerAddress(&_OperatorRegistry.TransactOpts, _delegationManagerAddress)
}

// SetDelegationManagerAddress is a paid mutator transaction binding the contract method 0x7831ad20.
//
// Solidity: function setDelegationManagerAddress(address _delegationManagerAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) SetDelegationManagerAddress(_delegationManagerAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.SetDelegationManagerAddress(&_OperatorRegistry.TransactOpts, _delegationManagerAddress)
}

// SetSlasherAddress is a paid mutator transaction binding the contract method 0xaddd9cf5.
//
// Solidity: function setSlasherAddress(address _slasherAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactor) SetSlasherAddress(opts *bind.TransactOpts, _slasherAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "setSlasherAddress", _slasherAddress)
}

// SetSlasherAddress is a paid mutator transaction binding the contract method 0xaddd9cf5.
//
// Solidity: function setSlasherAddress(address _slasherAddress) returns()
func (_OperatorRegistry *OperatorRegistrySession) SetSlasherAddress(_slasherAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.SetSlasherAddress(&_OperatorRegistry.TransactOpts, _slasherAddress)
}

// SetSlasherAddress is a paid mutator transaction binding the contract method 0xaddd9cf5.
//
// Solidity: function setSlasherAddress(address _slasherAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) SetSlasherAddress(_slasherAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.SetSlasherAddress(&_OperatorRegistry.TransactOpts, _slasherAddress)
}

// Suspend is a paid mutator transaction binding the contract method 0x286781c7.
//
// Solidity: function suspend(address operatorAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactor) Suspend(opts *bind.TransactOpts, operatorAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "suspend", operatorAddress)
}

// Suspend is a paid mutator transaction binding the contract method 0x286781c7.
//
// Solidity: function suspend(address operatorAddress) returns()
func (_OperatorRegistry *OperatorRegistrySession) Suspend(operatorAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.Suspend(&_OperatorRegistry.TransactOpts, operatorAddress)
}

// Suspend is a paid mutator transaction binding the contract method 0x286781c7.
//
// Solidity: function suspend(address operatorAddress) returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) Suspend(operatorAddress common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.Suspend(&_OperatorRegistry.TransactOpts, operatorAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OperatorRegistry *OperatorRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OperatorRegistry *OperatorRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.TransferOwnership(&_OperatorRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.TransferOwnership(&_OperatorRegistry.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_OperatorRegistry *OperatorRegistryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_OperatorRegistry *OperatorRegistrySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.UpgradeTo(&_OperatorRegistry.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.UpgradeTo(&_OperatorRegistry.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_OperatorRegistry *OperatorRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_OperatorRegistry *OperatorRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.UpgradeToAndCall(&_OperatorRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.UpgradeToAndCall(&_OperatorRegistry.TransactOpts, newImplementation, data)
}

// OperatorRegistryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the OperatorRegistry contract.
type OperatorRegistryAdminChangedIterator struct {
	Event *OperatorRegistryAdminChanged // Event containing the contract specifics and raw log

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
func (it *OperatorRegistryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRegistryAdminChanged)
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
		it.Event = new(OperatorRegistryAdminChanged)
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
func (it *OperatorRegistryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRegistryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRegistryAdminChanged represents a AdminChanged event raised by the OperatorRegistry contract.
type OperatorRegistryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_OperatorRegistry *OperatorRegistryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*OperatorRegistryAdminChangedIterator, error) {

	logs, sub, err := _OperatorRegistry.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryAdminChangedIterator{contract: _OperatorRegistry.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_OperatorRegistry *OperatorRegistryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *OperatorRegistryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _OperatorRegistry.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRegistryAdminChanged)
				if err := _OperatorRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_OperatorRegistry *OperatorRegistryFilterer) ParseAdminChanged(log types.Log) (*OperatorRegistryAdminChanged, error) {
	event := new(OperatorRegistryAdminChanged)
	if err := _OperatorRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRegistryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the OperatorRegistry contract.
type OperatorRegistryBeaconUpgradedIterator struct {
	Event *OperatorRegistryBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *OperatorRegistryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRegistryBeaconUpgraded)
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
		it.Event = new(OperatorRegistryBeaconUpgraded)
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
func (it *OperatorRegistryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRegistryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRegistryBeaconUpgraded represents a BeaconUpgraded event raised by the OperatorRegistry contract.
type OperatorRegistryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_OperatorRegistry *OperatorRegistryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*OperatorRegistryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _OperatorRegistry.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryBeaconUpgradedIterator{contract: _OperatorRegistry.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_OperatorRegistry *OperatorRegistryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *OperatorRegistryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _OperatorRegistry.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRegistryBeaconUpgraded)
				if err := _OperatorRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_OperatorRegistry *OperatorRegistryFilterer) ParseBeaconUpgraded(log types.Log) (*OperatorRegistryBeaconUpgraded, error) {
	event := new(OperatorRegistryBeaconUpgraded)
	if err := _OperatorRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OperatorRegistry contract.
type OperatorRegistryInitializedIterator struct {
	Event *OperatorRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *OperatorRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRegistryInitialized)
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
		it.Event = new(OperatorRegistryInitialized)
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
func (it *OperatorRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRegistryInitialized represents a Initialized event raised by the OperatorRegistry contract.
type OperatorRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OperatorRegistry *OperatorRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*OperatorRegistryInitializedIterator, error) {

	logs, sub, err := _OperatorRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryInitializedIterator{contract: _OperatorRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OperatorRegistry *OperatorRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OperatorRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _OperatorRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRegistryInitialized)
				if err := _OperatorRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_OperatorRegistry *OperatorRegistryFilterer) ParseInitialized(log types.Log) (*OperatorRegistryInitialized, error) {
	event := new(OperatorRegistryInitialized)
	if err := _OperatorRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRegistryOperatorSuspendedIterator is returned from FilterOperatorSuspended and is used to iterate over the raw logs and unpacked data for OperatorSuspended events raised by the OperatorRegistry contract.
type OperatorRegistryOperatorSuspendedIterator struct {
	Event *OperatorRegistryOperatorSuspended // Event containing the contract specifics and raw log

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
func (it *OperatorRegistryOperatorSuspendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRegistryOperatorSuspended)
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
		it.Event = new(OperatorRegistryOperatorSuspended)
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
func (it *OperatorRegistryOperatorSuspendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRegistryOperatorSuspendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRegistryOperatorSuspended represents a OperatorSuspended event raised by the OperatorRegistry contract.
type OperatorRegistryOperatorSuspended struct {
	Operator    common.Address
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSuspended is a free log retrieval operation binding the contract event 0x98ff055e829cd4bfba35f84d3b43877fd53084d2deebbb1cda57404f44df4def.
//
// Solidity: event OperatorSuspended(address operator, uint256 blockNumber)
func (_OperatorRegistry *OperatorRegistryFilterer) FilterOperatorSuspended(opts *bind.FilterOpts) (*OperatorRegistryOperatorSuspendedIterator, error) {

	logs, sub, err := _OperatorRegistry.contract.FilterLogs(opts, "OperatorSuspended")
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryOperatorSuspendedIterator{contract: _OperatorRegistry.contract, event: "OperatorSuspended", logs: logs, sub: sub}, nil
}

// WatchOperatorSuspended is a free log subscription operation binding the contract event 0x98ff055e829cd4bfba35f84d3b43877fd53084d2deebbb1cda57404f44df4def.
//
// Solidity: event OperatorSuspended(address operator, uint256 blockNumber)
func (_OperatorRegistry *OperatorRegistryFilterer) WatchOperatorSuspended(opts *bind.WatchOpts, sink chan<- *OperatorRegistryOperatorSuspended) (event.Subscription, error) {

	logs, sub, err := _OperatorRegistry.contract.WatchLogs(opts, "OperatorSuspended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRegistryOperatorSuspended)
				if err := _OperatorRegistry.contract.UnpackLog(event, "OperatorSuspended", log); err != nil {
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

// ParseOperatorSuspended is a log parse operation binding the contract event 0x98ff055e829cd4bfba35f84d3b43877fd53084d2deebbb1cda57404f44df4def.
//
// Solidity: event OperatorSuspended(address operator, uint256 blockNumber)
func (_OperatorRegistry *OperatorRegistryFilterer) ParseOperatorSuspended(log types.Log) (*OperatorRegistryOperatorSuspended, error) {
	event := new(OperatorRegistryOperatorSuspended)
	if err := _OperatorRegistry.contract.UnpackLog(event, "OperatorSuspended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRegistryOperatorsWhiteListedIterator is returned from FilterOperatorsWhiteListed and is used to iterate over the raw logs and unpacked data for OperatorsWhiteListed events raised by the OperatorRegistry contract.
type OperatorRegistryOperatorsWhiteListedIterator struct {
	Event *OperatorRegistryOperatorsWhiteListed // Event containing the contract specifics and raw log

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
func (it *OperatorRegistryOperatorsWhiteListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRegistryOperatorsWhiteListed)
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
		it.Event = new(OperatorRegistryOperatorsWhiteListed)
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
func (it *OperatorRegistryOperatorsWhiteListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRegistryOperatorsWhiteListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRegistryOperatorsWhiteListed represents a OperatorsWhiteListed event raised by the OperatorRegistry contract.
type OperatorRegistryOperatorsWhiteListed struct {
	Operator    []common.Address
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorsWhiteListed is a free log retrieval operation binding the contract event 0xdd7d462090c5f3ddd88b4c509269e35d5e148454dc3a9bd24812f78efc8f306f.
//
// Solidity: event OperatorsWhiteListed(address[] operator, uint256 blockNumber)
func (_OperatorRegistry *OperatorRegistryFilterer) FilterOperatorsWhiteListed(opts *bind.FilterOpts) (*OperatorRegistryOperatorsWhiteListedIterator, error) {

	logs, sub, err := _OperatorRegistry.contract.FilterLogs(opts, "OperatorsWhiteListed")
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryOperatorsWhiteListedIterator{contract: _OperatorRegistry.contract, event: "OperatorsWhiteListed", logs: logs, sub: sub}, nil
}

// WatchOperatorsWhiteListed is a free log subscription operation binding the contract event 0xdd7d462090c5f3ddd88b4c509269e35d5e148454dc3a9bd24812f78efc8f306f.
//
// Solidity: event OperatorsWhiteListed(address[] operator, uint256 blockNumber)
func (_OperatorRegistry *OperatorRegistryFilterer) WatchOperatorsWhiteListed(opts *bind.WatchOpts, sink chan<- *OperatorRegistryOperatorsWhiteListed) (event.Subscription, error) {

	logs, sub, err := _OperatorRegistry.contract.WatchLogs(opts, "OperatorsWhiteListed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRegistryOperatorsWhiteListed)
				if err := _OperatorRegistry.contract.UnpackLog(event, "OperatorsWhiteListed", log); err != nil {
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

// ParseOperatorsWhiteListed is a log parse operation binding the contract event 0xdd7d462090c5f3ddd88b4c509269e35d5e148454dc3a9bd24812f78efc8f306f.
//
// Solidity: event OperatorsWhiteListed(address[] operator, uint256 blockNumber)
func (_OperatorRegistry *OperatorRegistryFilterer) ParseOperatorsWhiteListed(log types.Log) (*OperatorRegistryOperatorsWhiteListed, error) {
	event := new(OperatorRegistryOperatorsWhiteListed)
	if err := _OperatorRegistry.contract.UnpackLog(event, "OperatorsWhiteListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OperatorRegistry contract.
type OperatorRegistryOwnershipTransferredIterator struct {
	Event *OperatorRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OperatorRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRegistryOwnershipTransferred)
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
		it.Event = new(OperatorRegistryOwnershipTransferred)
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
func (it *OperatorRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the OperatorRegistry contract.
type OperatorRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OperatorRegistry *OperatorRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OperatorRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OperatorRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryOwnershipTransferredIterator{contract: _OperatorRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OperatorRegistry *OperatorRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OperatorRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OperatorRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRegistryOwnershipTransferred)
				if err := _OperatorRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OperatorRegistry *OperatorRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*OperatorRegistryOwnershipTransferred, error) {
	event := new(OperatorRegistryOwnershipTransferred)
	if err := _OperatorRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the OperatorRegistry contract.
type OperatorRegistryPausedIterator struct {
	Event *OperatorRegistryPaused // Event containing the contract specifics and raw log

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
func (it *OperatorRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRegistryPaused)
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
		it.Event = new(OperatorRegistryPaused)
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
func (it *OperatorRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRegistryPaused represents a Paused event raised by the OperatorRegistry contract.
type OperatorRegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OperatorRegistry *OperatorRegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*OperatorRegistryPausedIterator, error) {

	logs, sub, err := _OperatorRegistry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryPausedIterator{contract: _OperatorRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OperatorRegistry *OperatorRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OperatorRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _OperatorRegistry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRegistryPaused)
				if err := _OperatorRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_OperatorRegistry *OperatorRegistryFilterer) ParsePaused(log types.Log) (*OperatorRegistryPaused, error) {
	event := new(OperatorRegistryPaused)
	if err := _OperatorRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the OperatorRegistry contract.
type OperatorRegistryUnpausedIterator struct {
	Event *OperatorRegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *OperatorRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRegistryUnpaused)
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
		it.Event = new(OperatorRegistryUnpaused)
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
func (it *OperatorRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRegistryUnpaused represents a Unpaused event raised by the OperatorRegistry contract.
type OperatorRegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OperatorRegistry *OperatorRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*OperatorRegistryUnpausedIterator, error) {

	logs, sub, err := _OperatorRegistry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryUnpausedIterator{contract: _OperatorRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OperatorRegistry *OperatorRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OperatorRegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _OperatorRegistry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRegistryUnpaused)
				if err := _OperatorRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_OperatorRegistry *OperatorRegistryFilterer) ParseUnpaused(log types.Log) (*OperatorRegistryUnpaused, error) {
	event := new(OperatorRegistryUnpaused)
	if err := _OperatorRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the OperatorRegistry contract.
type OperatorRegistryUpgradedIterator struct {
	Event *OperatorRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *OperatorRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRegistryUpgraded)
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
		it.Event = new(OperatorRegistryUpgraded)
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
func (it *OperatorRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRegistryUpgraded represents a Upgraded event raised by the OperatorRegistry contract.
type OperatorRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_OperatorRegistry *OperatorRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*OperatorRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _OperatorRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryUpgradedIterator{contract: _OperatorRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_OperatorRegistry *OperatorRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *OperatorRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _OperatorRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRegistryUpgraded)
				if err := _OperatorRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_OperatorRegistry *OperatorRegistryFilterer) ParseUpgraded(log types.Log) (*OperatorRegistryUpgraded, error) {
	event := new(OperatorRegistryUpgraded)
	if err := _OperatorRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRegistryWatchtowerDeRegisteredFromOperatorIterator is returned from FilterWatchtowerDeRegisteredFromOperator and is used to iterate over the raw logs and unpacked data for WatchtowerDeRegisteredFromOperator events raised by the OperatorRegistry contract.
type OperatorRegistryWatchtowerDeRegisteredFromOperatorIterator struct {
	Event *OperatorRegistryWatchtowerDeRegisteredFromOperator // Event containing the contract specifics and raw log

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
func (it *OperatorRegistryWatchtowerDeRegisteredFromOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRegistryWatchtowerDeRegisteredFromOperator)
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
		it.Event = new(OperatorRegistryWatchtowerDeRegisteredFromOperator)
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
func (it *OperatorRegistryWatchtowerDeRegisteredFromOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRegistryWatchtowerDeRegisteredFromOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRegistryWatchtowerDeRegisteredFromOperator represents a WatchtowerDeRegisteredFromOperator event raised by the OperatorRegistry contract.
type OperatorRegistryWatchtowerDeRegisteredFromOperator struct {
	Operator    common.Address
	Watchtower  common.Address
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWatchtowerDeRegisteredFromOperator is a free log retrieval operation binding the contract event 0xecfc290ff8c3aac71e14aee07653f81e5aa316be4d2315ba2ec1bff9dc50cd79.
//
// Solidity: event WatchtowerDeRegisteredFromOperator(address operator, address watchtower, uint256 blockNumber)
func (_OperatorRegistry *OperatorRegistryFilterer) FilterWatchtowerDeRegisteredFromOperator(opts *bind.FilterOpts) (*OperatorRegistryWatchtowerDeRegisteredFromOperatorIterator, error) {

	logs, sub, err := _OperatorRegistry.contract.FilterLogs(opts, "WatchtowerDeRegisteredFromOperator")
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryWatchtowerDeRegisteredFromOperatorIterator{contract: _OperatorRegistry.contract, event: "WatchtowerDeRegisteredFromOperator", logs: logs, sub: sub}, nil
}

// WatchWatchtowerDeRegisteredFromOperator is a free log subscription operation binding the contract event 0xecfc290ff8c3aac71e14aee07653f81e5aa316be4d2315ba2ec1bff9dc50cd79.
//
// Solidity: event WatchtowerDeRegisteredFromOperator(address operator, address watchtower, uint256 blockNumber)
func (_OperatorRegistry *OperatorRegistryFilterer) WatchWatchtowerDeRegisteredFromOperator(opts *bind.WatchOpts, sink chan<- *OperatorRegistryWatchtowerDeRegisteredFromOperator) (event.Subscription, error) {

	logs, sub, err := _OperatorRegistry.contract.WatchLogs(opts, "WatchtowerDeRegisteredFromOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRegistryWatchtowerDeRegisteredFromOperator)
				if err := _OperatorRegistry.contract.UnpackLog(event, "WatchtowerDeRegisteredFromOperator", log); err != nil {
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

// ParseWatchtowerDeRegisteredFromOperator is a log parse operation binding the contract event 0xecfc290ff8c3aac71e14aee07653f81e5aa316be4d2315ba2ec1bff9dc50cd79.
//
// Solidity: event WatchtowerDeRegisteredFromOperator(address operator, address watchtower, uint256 blockNumber)
func (_OperatorRegistry *OperatorRegistryFilterer) ParseWatchtowerDeRegisteredFromOperator(log types.Log) (*OperatorRegistryWatchtowerDeRegisteredFromOperator, error) {
	event := new(OperatorRegistryWatchtowerDeRegisteredFromOperator)
	if err := _OperatorRegistry.contract.UnpackLog(event, "WatchtowerDeRegisteredFromOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRegistryWatchtowerRegisteredToOperatorIterator is returned from FilterWatchtowerRegisteredToOperator and is used to iterate over the raw logs and unpacked data for WatchtowerRegisteredToOperator events raised by the OperatorRegistry contract.
type OperatorRegistryWatchtowerRegisteredToOperatorIterator struct {
	Event *OperatorRegistryWatchtowerRegisteredToOperator // Event containing the contract specifics and raw log

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
func (it *OperatorRegistryWatchtowerRegisteredToOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRegistryWatchtowerRegisteredToOperator)
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
		it.Event = new(OperatorRegistryWatchtowerRegisteredToOperator)
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
func (it *OperatorRegistryWatchtowerRegisteredToOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRegistryWatchtowerRegisteredToOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRegistryWatchtowerRegisteredToOperator represents a WatchtowerRegisteredToOperator event raised by the OperatorRegistry contract.
type OperatorRegistryWatchtowerRegisteredToOperator struct {
	Operator    common.Address
	Watchtower  common.Address
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWatchtowerRegisteredToOperator is a free log retrieval operation binding the contract event 0x3d521ffddd8cbc2d603c1e7fe9af4b70adbbc43675a7281c5624d8053e526f06.
//
// Solidity: event WatchtowerRegisteredToOperator(address operator, address watchtower, uint256 blockNumber)
func (_OperatorRegistry *OperatorRegistryFilterer) FilterWatchtowerRegisteredToOperator(opts *bind.FilterOpts) (*OperatorRegistryWatchtowerRegisteredToOperatorIterator, error) {

	logs, sub, err := _OperatorRegistry.contract.FilterLogs(opts, "WatchtowerRegisteredToOperator")
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryWatchtowerRegisteredToOperatorIterator{contract: _OperatorRegistry.contract, event: "WatchtowerRegisteredToOperator", logs: logs, sub: sub}, nil
}

// WatchWatchtowerRegisteredToOperator is a free log subscription operation binding the contract event 0x3d521ffddd8cbc2d603c1e7fe9af4b70adbbc43675a7281c5624d8053e526f06.
//
// Solidity: event WatchtowerRegisteredToOperator(address operator, address watchtower, uint256 blockNumber)
func (_OperatorRegistry *OperatorRegistryFilterer) WatchWatchtowerRegisteredToOperator(opts *bind.WatchOpts, sink chan<- *OperatorRegistryWatchtowerRegisteredToOperator) (event.Subscription, error) {

	logs, sub, err := _OperatorRegistry.contract.WatchLogs(opts, "WatchtowerRegisteredToOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRegistryWatchtowerRegisteredToOperator)
				if err := _OperatorRegistry.contract.UnpackLog(event, "WatchtowerRegisteredToOperator", log); err != nil {
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

// ParseWatchtowerRegisteredToOperator is a log parse operation binding the contract event 0x3d521ffddd8cbc2d603c1e7fe9af4b70adbbc43675a7281c5624d8053e526f06.
//
// Solidity: event WatchtowerRegisteredToOperator(address operator, address watchtower, uint256 blockNumber)
func (_OperatorRegistry *OperatorRegistryFilterer) ParseWatchtowerRegisteredToOperator(log types.Log) (*OperatorRegistryWatchtowerRegisteredToOperator, error) {
	event := new(OperatorRegistryWatchtowerRegisteredToOperator)
	if err := _OperatorRegistry.contract.UnpackLog(event, "WatchtowerRegisteredToOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
