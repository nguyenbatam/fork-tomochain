// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"github.com/fns/fns"
	"github.com/fns/fns/accounts/abi"
	"github.com/fns/fns/accounts/abi/bind"
	"github.com/fns/fns/common"
	"github.com/fns/fns/core/types"
	"github.com/fns/fns/event"
	"math/big"
	"strings"
)

// FNSValidatorABI is the input ABI used to generate the binding from.
const FNSValidatorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"propose\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"unvote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCandidates\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getWithdrawCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getVoters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWithdrawBlockNumbers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"candidateCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voterWithdrawDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"resign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxValidatorNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"candidateWithdrawDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"isCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minCandidateCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minVoterCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_candidates\",\"type\":\"address[]\"},{\"name\":\"_caps\",\"type\":\"uint256[]\"},{\"name\":\"_firstOwner\",\"type\":\"address\"},{\"name\":\"_minCandidateCap\",\"type\":\"uint256\"},{\"name\":\"_minVoterCap\",\"type\":\"uint256\"},{\"name\":\"_maxValidatorNumber\",\"type\":\"uint256\"},{\"name\":\"_candidateWithdrawDelay\",\"type\":\"uint256\"},{\"name\":\"_voterWithdrawDelay\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Unvote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Propose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"Resign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]"

// FNSValidatorBin is the compiled bytecode used for deploying new contracts.
const FNSValidatorBin = `0x6080604052600060045534801561001557600080fd5b506040516111d73803806111d783398101604090815281516020830151918301516060840151608085015160a086015160c087015160e088015160058590556006849055600783905560088290556009819055958801805160045597969096019593949293919290919060005b885181101561024d576003898281518110151561009b57fe5b602090810291909101810151825460018082018555600094855293839020018054600160a060020a031916600160a060020a0392831617905560408051606081018252918b168252918101929092528951908201908a90849081106100fc57fe5b90602001906020020151815250600160008b8481518110151561011b57fe5b602090810291909101810151600160a060020a03908116835282820193909352604091820160009081208551815493870151600160a060020a031990941695169490941760a060020a60ff02191674010000000000000000000000000000000000000000921515929092029190911783559201516001909101558951600291908b90849081106101a757fe5b602090810291909101810151600160a060020a03908116835282820193909352604090910160009081208054600180820183559183529282209092018054600160a060020a031916938b16939093179092556005548b519092908c908590811061020d57fe5b6020908102909101810151600160a060020a0390811683528282019390935260409182016000908120938c16815260029093019052902055600101610082565b505050505050505050610f72806102656000396000f3006080604052600436106101115763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301267951811461011657806302aa9be21461012c57806306a49fce1461015057806315febd68146101b55780632d15cc04146101df5780632f9c4bba14610200578063302b6872146102155780633477ee2e1461023c578063441a3e701461027057806358e7525f1461028b5780636dd7d8ea146102ac578063a9a981a3146102c0578063a9ff959e146102d5578063ae6e43f5146102ea578063b642facd1461030b578063d09f1ab41461032c578063d161c76714610341578063d51b9e9314610356578063d55b7dff1461038b578063f8ac9dd5146103a0575b600080fd5b61012a600160a060020a03600435166103b5565b005b34801561013857600080fd5b5061012a600160a060020a03600435166024356105d5565b34801561015c57600080fd5b506101656107cb565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101a1578181015183820152602001610189565b505050509050019250505060405180910390f35b3480156101c157600080fd5b506101cd60043561082d565b60408051918252519081900360200190f35b3480156101eb57600080fd5b50610165600160a060020a0360043516610848565b34801561020c57600080fd5b506101656108be565b34801561022157600080fd5b506101cd600160a060020a036004358116906024351661091f565b34801561024857600080fd5b5061025460043561094e565b60408051600160a060020a039092168252519081900360200190f35b34801561027c57600080fd5b5061012a600435602435610976565b34801561029757600080fd5b506101cd600160a060020a0360043516610aa6565b61012a600160a060020a0360043516610ac5565b3480156102cc57600080fd5b506101cd610c50565b3480156102e157600080fd5b506101cd610c56565b3480156102f657600080fd5b5061012a600160a060020a0360043516610c5c565b34801561031757600080fd5b50610254600160a060020a0360043516610ec3565b34801561033857600080fd5b506101cd610ee1565b34801561034d57600080fd5b506101cd610ee7565b34801561036257600080fd5b50610377600160a060020a0360043516610eed565b604080519115158252519081900360200190f35b34801561039757600080fd5b506101cd610f12565b3480156103ac57600080fd5b506101cd610f18565b6005546000903410156103c757600080fd5b600160a060020a038216600090815260016020526040902054829060a060020a900460ff16156103f657600080fd5b600160a060020a03831660009081526001602081905260409091200154610423903463ffffffff610f1e16565b6003805460018181019092557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b018054600160a060020a0380881673ffffffffffffffffffffffffffffffffffffffff199283168117909355604080516060810182523380825260208281018881528385018a81526000988952898352858920945185549251151560a060020a0274ff0000000000000000000000000000000000000000199190981692909816919091179690961694909417825593519581019590955591835260029093019092522054909250610507903463ffffffff610f1e16565b600160a060020a0384166000908152600160208181526040808420338552600201909152909120919091556004546105449163ffffffff610f1e16565b600455600160a060020a0383166000818152600260209081526040808320805460018101825590845292829020909201805473ffffffffffffffffffffffffffffffffffffffff1916339081179091558251908152908101929092523482820152517f7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c19181900360600190a1505050565b600160a060020a03821660009081526001602090815260408083203384526002019091528120548390839081111561060c57600080fd5b600160a060020a038281166000908152600160205260409020541633141561067357600554600160a060020a0383166000908152600160209081526040808320338452600201909152902054610668908363ffffffff610f3416565b101561067357600080fd5b600160a060020a038516600090815260016020819052604090912001546106a0908563ffffffff610f3416565b600160a060020a0386166000908152600160208181526040808420928301949094553383526002909101905220546106de908563ffffffff610f3416565b600160a060020a0386166000908152600160209081526040808320338452600201909152902055600954610718904363ffffffff610f1e16565b33600090815260208181526040808320848452909152902054909350610744908563ffffffff610f1e16565b33600081815260208181526040808320888452808352818420959095558282526001948501805495860181558352918190209093018690558051918252600160a060020a0388169282019290925280820186905290517faa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa29181900360600190a15050505050565b6060600380548060200260200160405190810160405280929190818152602001828054801561082357602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610805575b5050505050905090565b33600090815260208181526040808320938352929052205490565b600160a060020a0381166000908152600260209081526040918290208054835181840281018401909452808452606093928301828280156108b257602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610894575b50505050509050919050565b336000908152602081815260409182902060010180548351818402810184019094528084526060939283018282801561082357602002820191906000526020600020905b815481526020019060010190808311610902575050505050905090565b600160a060020a0391821660009081526001602090815260408083209390941682526002909201909152205490565b600380548290811061095c57fe5b600091825260209091200154600160a060020a0316905081565b6000828282821161098657600080fd5b4382111561099357600080fd5b33600090815260208181526040808320858452909152812054116109b657600080fd5b3360009081526020819052604090206001018054839190839081106109d757fe5b90600052602060002001541415156109ee57600080fd5b3360008181526020818152604080832089845280835290832080549084905593835291905260010180549194509085908110610a2657fe5b60009182526020822001819055604051339185156108fc02918691818181858888f19350505050158015610a5e573d6000803e3d6000fd5b50604080513381526020810187905280820185905290517ff279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b5689181900360600190a15050505050565b600160a060020a03166000908152600160208190526040909120015490565b600654341015610ad457600080fd5b600160a060020a038116600090815260016020526040902054819060a060020a900460ff161515610b0457600080fd5b600160a060020a03821660009081526001602081905260409091200154610b31903463ffffffff610f1e16565b600160a060020a0383166000908152600160208181526040808420928301949094553383526002909101905220541515610bab57600160a060020a038216600090815260026020908152604082208054600181018255908352912001805473ffffffffffffffffffffffffffffffffffffffff1916331790555b600160a060020a0382166000908152600160209081526040808320338452600201909152902054610be2903463ffffffff610f1e16565b600160a060020a0383166000818152600160209081526040808320338085526002909101835292819020949094558351918252810191909152348183015290517f66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc9181900360600190a15050565b60045481565b60095481565b600160a060020a038082166000908152600160205260408120549091829182918591163314610c8a57600080fd5b600160a060020a038516600090815260016020526040902054859060a060020a900460ff161515610cba57600080fd5b600160a060020a0386166000908152600160208190526040909120805474ff000000000000000000000000000000000000000019169055600454610d039163ffffffff610f3416565b600455600094505b600354851015610d8d5785600160a060020a0316600386815481101515610d2e57fe5b600091825260209091200154600160a060020a03161415610d82576003805486908110610d5757fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff19169055610d8d565b600190940193610d0b565b600160a060020a038616600081815260016020818152604080842033855260028101835290842054949093528190520154909450610dd1908563ffffffff610f3416565b600160a060020a0387166000908152600160208181526040808420928301949094553383526002909101905290812055600854610e14904363ffffffff610f1e16565b33600090815260208181526040808320848452909152902054909350610e40908563ffffffff610f1e16565b33600081815260208181526040808320888452808352818420959095558282526001948501805495860181558352918190209093018690558051918252600160a060020a0389169282019290925281517f4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d3929181900390910190a1505050505050565b600160a060020a039081166000908152600160205260409020541690565b60075481565b60085481565b600160a060020a031660009081526001602052604090205460a060020a900460ff1690565b60055481565b60065481565b600082820183811015610f2d57fe5b9392505050565b600082821115610f4057fe5b509003905600a165627a7a72305820589982ee93090a327109f5b0224efff3a58e920f21dab66c74c4523ad6895d2b0029`

// DeployFNSValidator deploys a new Ethereum contract, binding an instance of FNSValidator to it.
func DeployFNSValidator(auth *bind.TransactOpts, backend bind.ContractBackend, _candidates []common.Address, _caps []*big.Int, _firstOwner common.Address, _minCandidateCap *big.Int, _minVoterCap *big.Int, _maxValidatorNumber *big.Int, _candidateWithdrawDelay *big.Int, _voterWithdrawDelay *big.Int) (common.Address, *types.Transaction, *FNSValidator, error) {
	parsed, err := abi.JSON(strings.NewReader(FNSValidatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FNSValidatorBin), backend, _candidates, _caps, _firstOwner, _minCandidateCap, _minVoterCap, _maxValidatorNumber, _candidateWithdrawDelay, _voterWithdrawDelay)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FNSValidator{FNSValidatorCaller: FNSValidatorCaller{contract: contract}, FNSValidatorTransactor: FNSValidatorTransactor{contract: contract}, FNSValidatorFilterer: FNSValidatorFilterer{contract: contract}}, nil
}

// FNSValidator is an auto generated Go binding around an Ethereum contract.
type FNSValidator struct {
	FNSValidatorCaller     // Read-only binding to the contract
	FNSValidatorTransactor // Write-only binding to the contract
	FNSValidatorFilterer   // Log filterer for contract events
}

// FNSValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type FNSValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FNSValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FNSValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FNSValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FNSValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FNSValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FNSValidatorSession struct {
	Contract     *FNSValidator     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FNSValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FNSValidatorCallerSession struct {
	Contract *FNSValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// FNSValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FNSValidatorTransactorSession struct {
	Contract     *FNSValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FNSValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type FNSValidatorRaw struct {
	Contract *FNSValidator // Generic contract binding to access the raw methods on
}

// FNSValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FNSValidatorCallerRaw struct {
	Contract *FNSValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// FNSValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FNSValidatorTransactorRaw struct {
	Contract *FNSValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFNSValidator creates a new instance of FNSValidator, bound to a specific deployed contract.
func NewFNSValidator(address common.Address, backend bind.ContractBackend) (*FNSValidator, error) {
	contract, err := bindFNSValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FNSValidator{FNSValidatorCaller: FNSValidatorCaller{contract: contract}, FNSValidatorTransactor: FNSValidatorTransactor{contract: contract}, FNSValidatorFilterer: FNSValidatorFilterer{contract: contract}}, nil
}

// NewFNSValidatorCaller creates a new read-only instance of FNSValidator, bound to a specific deployed contract.
func NewFNSValidatorCaller(address common.Address, caller bind.ContractCaller) (*FNSValidatorCaller, error) {
	contract, err := bindFNSValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FNSValidatorCaller{contract: contract}, nil
}

// NewFNSValidatorTransactor creates a new write-only instance of FNSValidator, bound to a specific deployed contract.
func NewFNSValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*FNSValidatorTransactor, error) {
	contract, err := bindFNSValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FNSValidatorTransactor{contract: contract}, nil
}

// NewFNSValidatorFilterer creates a new log filterer instance of FNSValidator, bound to a specific deployed contract.
func NewFNSValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*FNSValidatorFilterer, error) {
	contract, err := bindFNSValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FNSValidatorFilterer{contract: contract}, nil
}

// bindFNSValidator binds a generic wrapper to an already deployed contract.
func bindFNSValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FNSValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FNSValidator *FNSValidatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FNSValidator.Contract.FNSValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FNSValidator *FNSValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FNSValidator.Contract.FNSValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FNSValidator *FNSValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FNSValidator.Contract.FNSValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FNSValidator *FNSValidatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FNSValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FNSValidator *FNSValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FNSValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FNSValidator *FNSValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FNSValidator.Contract.contract.Transact(opts, method, params...)
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() constant returns(uint256)
func (_FNSValidator *FNSValidatorCaller) CandidateCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FNSValidator.contract.Call(opts, out, "candidateCount")
	return *ret0, err
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() constant returns(uint256)
func (_FNSValidator *FNSValidatorSession) CandidateCount() (*big.Int, error) {
	return _FNSValidator.Contract.CandidateCount(&_FNSValidator.CallOpts)
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() constant returns(uint256)
func (_FNSValidator *FNSValidatorCallerSession) CandidateCount() (*big.Int, error) {
	return _FNSValidator.Contract.CandidateCount(&_FNSValidator.CallOpts)
}

// CandidateWithdrawDelay is a free data retrieval call binding the contract method 0xd161c767.
//
// Solidity: function candidateWithdrawDelay() constant returns(uint256)
func (_FNSValidator *FNSValidatorCaller) CandidateWithdrawDelay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FNSValidator.contract.Call(opts, out, "candidateWithdrawDelay")
	return *ret0, err
}

// CandidateWithdrawDelay is a free data retrieval call binding the contract method 0xd161c767.
//
// Solidity: function candidateWithdrawDelay() constant returns(uint256)
func (_FNSValidator *FNSValidatorSession) CandidateWithdrawDelay() (*big.Int, error) {
	return _FNSValidator.Contract.CandidateWithdrawDelay(&_FNSValidator.CallOpts)
}

// CandidateWithdrawDelay is a free data retrieval call binding the contract method 0xd161c767.
//
// Solidity: function candidateWithdrawDelay() constant returns(uint256)
func (_FNSValidator *FNSValidatorCallerSession) CandidateWithdrawDelay() (*big.Int, error) {
	return _FNSValidator.Contract.CandidateWithdrawDelay(&_FNSValidator.CallOpts)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates( uint256) constant returns(address)
func (_FNSValidator *FNSValidatorCaller) Candidates(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FNSValidator.contract.Call(opts, out, "candidates", arg0)
	return *ret0, err
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates( uint256) constant returns(address)
func (_FNSValidator *FNSValidatorSession) Candidates(arg0 *big.Int) (common.Address, error) {
	return _FNSValidator.Contract.Candidates(&_FNSValidator.CallOpts, arg0)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates( uint256) constant returns(address)
func (_FNSValidator *FNSValidatorCallerSession) Candidates(arg0 *big.Int) (common.Address, error) {
	return _FNSValidator.Contract.Candidates(&_FNSValidator.CallOpts, arg0)
}

// GetCandidateCap is a free data retrieval call binding the contract method 0x58e7525f.
//
// Solidity: function getCandidateCap(_candidate address) constant returns(uint256)
func (_FNSValidator *FNSValidatorCaller) GetCandidateCap(opts *bind.CallOpts, _candidate common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FNSValidator.contract.Call(opts, out, "getCandidateCap", _candidate)
	return *ret0, err
}

// GetCandidateCap is a free data retrieval call binding the contract method 0x58e7525f.
//
// Solidity: function getCandidateCap(_candidate address) constant returns(uint256)
func (_FNSValidator *FNSValidatorSession) GetCandidateCap(_candidate common.Address) (*big.Int, error) {
	return _FNSValidator.Contract.GetCandidateCap(&_FNSValidator.CallOpts, _candidate)
}

// GetCandidateCap is a free data retrieval call binding the contract method 0x58e7525f.
//
// Solidity: function getCandidateCap(_candidate address) constant returns(uint256)
func (_FNSValidator *FNSValidatorCallerSession) GetCandidateCap(_candidate common.Address) (*big.Int, error) {
	return _FNSValidator.Contract.GetCandidateCap(&_FNSValidator.CallOpts, _candidate)
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xb642facd.
//
// Solidity: function getCandidateOwner(_candidate address) constant returns(address)
func (_FNSValidator *FNSValidatorCaller) GetCandidateOwner(opts *bind.CallOpts, _candidate common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FNSValidator.contract.Call(opts, out, "getCandidateOwner", _candidate)
	return *ret0, err
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xb642facd.
//
// Solidity: function getCandidateOwner(_candidate address) constant returns(address)
func (_FNSValidator *FNSValidatorSession) GetCandidateOwner(_candidate common.Address) (common.Address, error) {
	return _FNSValidator.Contract.GetCandidateOwner(&_FNSValidator.CallOpts, _candidate)
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xb642facd.
//
// Solidity: function getCandidateOwner(_candidate address) constant returns(address)
func (_FNSValidator *FNSValidatorCallerSession) GetCandidateOwner(_candidate common.Address) (common.Address, error) {
	return _FNSValidator.Contract.GetCandidateOwner(&_FNSValidator.CallOpts, _candidate)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(address[])
func (_FNSValidator *FNSValidatorCaller) GetCandidates(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _FNSValidator.contract.Call(opts, out, "getCandidates")
	return *ret0, err
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(address[])
func (_FNSValidator *FNSValidatorSession) GetCandidates() ([]common.Address, error) {
	return _FNSValidator.Contract.GetCandidates(&_FNSValidator.CallOpts)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(address[])
func (_FNSValidator *FNSValidatorCallerSession) GetCandidates() ([]common.Address, error) {
	return _FNSValidator.Contract.GetCandidates(&_FNSValidator.CallOpts)
}

// GetVoterCap is a free data retrieval call binding the contract method 0x302b6872.
//
// Solidity: function getVoterCap(_candidate address, _voter address) constant returns(uint256)
func (_FNSValidator *FNSValidatorCaller) GetVoterCap(opts *bind.CallOpts, _candidate common.Address, _voter common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FNSValidator.contract.Call(opts, out, "getVoterCap", _candidate, _voter)
	return *ret0, err
}

// GetVoterCap is a free data retrieval call binding the contract method 0x302b6872.
//
// Solidity: function getVoterCap(_candidate address, _voter address) constant returns(uint256)
func (_FNSValidator *FNSValidatorSession) GetVoterCap(_candidate common.Address, _voter common.Address) (*big.Int, error) {
	return _FNSValidator.Contract.GetVoterCap(&_FNSValidator.CallOpts, _candidate, _voter)
}

// GetVoterCap is a free data retrieval call binding the contract method 0x302b6872.
//
// Solidity: function getVoterCap(_candidate address, _voter address) constant returns(uint256)
func (_FNSValidator *FNSValidatorCallerSession) GetVoterCap(_candidate common.Address, _voter common.Address) (*big.Int, error) {
	return _FNSValidator.Contract.GetVoterCap(&_FNSValidator.CallOpts, _candidate, _voter)
}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(_candidate address) constant returns(address[])
func (_FNSValidator *FNSValidatorCaller) GetVoters(opts *bind.CallOpts, _candidate common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _FNSValidator.contract.Call(opts, out, "getVoters", _candidate)
	return *ret0, err
}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(_candidate address) constant returns(address[])
func (_FNSValidator *FNSValidatorSession) GetVoters(_candidate common.Address) ([]common.Address, error) {
	return _FNSValidator.Contract.GetVoters(&_FNSValidator.CallOpts, _candidate)
}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(_candidate address) constant returns(address[])
func (_FNSValidator *FNSValidatorCallerSession) GetVoters(_candidate common.Address) ([]common.Address, error) {
	return _FNSValidator.Contract.GetVoters(&_FNSValidator.CallOpts, _candidate)
}

// GetWithdrawBlockNumbers is a free data retrieval call binding the contract method 0x2f9c4bba.
//
// Solidity: function getWithdrawBlockNumbers() constant returns(uint256[])
func (_FNSValidator *FNSValidatorCaller) GetWithdrawBlockNumbers(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _FNSValidator.contract.Call(opts, out, "getWithdrawBlockNumbers")
	return *ret0, err
}

// GetWithdrawBlockNumbers is a free data retrieval call binding the contract method 0x2f9c4bba.
//
// Solidity: function getWithdrawBlockNumbers() constant returns(uint256[])
func (_FNSValidator *FNSValidatorSession) GetWithdrawBlockNumbers() ([]*big.Int, error) {
	return _FNSValidator.Contract.GetWithdrawBlockNumbers(&_FNSValidator.CallOpts)
}

// GetWithdrawBlockNumbers is a free data retrieval call binding the contract method 0x2f9c4bba.
//
// Solidity: function getWithdrawBlockNumbers() constant returns(uint256[])
func (_FNSValidator *FNSValidatorCallerSession) GetWithdrawBlockNumbers() ([]*big.Int, error) {
	return _FNSValidator.Contract.GetWithdrawBlockNumbers(&_FNSValidator.CallOpts)
}

// GetWithdrawCap is a free data retrieval call binding the contract method 0x15febd68.
//
// Solidity: function getWithdrawCap(_blockNumber uint256) constant returns(uint256)
func (_FNSValidator *FNSValidatorCaller) GetWithdrawCap(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FNSValidator.contract.Call(opts, out, "getWithdrawCap", _blockNumber)
	return *ret0, err
}

// GetWithdrawCap is a free data retrieval call binding the contract method 0x15febd68.
//
// Solidity: function getWithdrawCap(_blockNumber uint256) constant returns(uint256)
func (_FNSValidator *FNSValidatorSession) GetWithdrawCap(_blockNumber *big.Int) (*big.Int, error) {
	return _FNSValidator.Contract.GetWithdrawCap(&_FNSValidator.CallOpts, _blockNumber)
}

// GetWithdrawCap is a free data retrieval call binding the contract method 0x15febd68.
//
// Solidity: function getWithdrawCap(_blockNumber uint256) constant returns(uint256)
func (_FNSValidator *FNSValidatorCallerSession) GetWithdrawCap(_blockNumber *big.Int) (*big.Int, error) {
	return _FNSValidator.Contract.GetWithdrawCap(&_FNSValidator.CallOpts, _blockNumber)
}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(_candidate address) constant returns(bool)
func (_FNSValidator *FNSValidatorCaller) IsCandidate(opts *bind.CallOpts, _candidate common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FNSValidator.contract.Call(opts, out, "isCandidate", _candidate)
	return *ret0, err
}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(_candidate address) constant returns(bool)
func (_FNSValidator *FNSValidatorSession) IsCandidate(_candidate common.Address) (bool, error) {
	return _FNSValidator.Contract.IsCandidate(&_FNSValidator.CallOpts, _candidate)
}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(_candidate address) constant returns(bool)
func (_FNSValidator *FNSValidatorCallerSession) IsCandidate(_candidate common.Address) (bool, error) {
	return _FNSValidator.Contract.IsCandidate(&_FNSValidator.CallOpts, _candidate)
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() constant returns(uint256)
func (_FNSValidator *FNSValidatorCaller) MaxValidatorNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FNSValidator.contract.Call(opts, out, "maxValidatorNumber")
	return *ret0, err
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() constant returns(uint256)
func (_FNSValidator *FNSValidatorSession) MaxValidatorNumber() (*big.Int, error) {
	return _FNSValidator.Contract.MaxValidatorNumber(&_FNSValidator.CallOpts)
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() constant returns(uint256)
func (_FNSValidator *FNSValidatorCallerSession) MaxValidatorNumber() (*big.Int, error) {
	return _FNSValidator.Contract.MaxValidatorNumber(&_FNSValidator.CallOpts)
}

// MinCandidateCap is a free data retrieval call binding the contract method 0xd55b7dff.
//
// Solidity: function minCandidateCap() constant returns(uint256)
func (_FNSValidator *FNSValidatorCaller) MinCandidateCap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FNSValidator.contract.Call(opts, out, "minCandidateCap")
	return *ret0, err
}

// MinCandidateCap is a free data retrieval call binding the contract method 0xd55b7dff.
//
// Solidity: function minCandidateCap() constant returns(uint256)
func (_FNSValidator *FNSValidatorSession) MinCandidateCap() (*big.Int, error) {
	return _FNSValidator.Contract.MinCandidateCap(&_FNSValidator.CallOpts)
}

// MinCandidateCap is a free data retrieval call binding the contract method 0xd55b7dff.
//
// Solidity: function minCandidateCap() constant returns(uint256)
func (_FNSValidator *FNSValidatorCallerSession) MinCandidateCap() (*big.Int, error) {
	return _FNSValidator.Contract.MinCandidateCap(&_FNSValidator.CallOpts)
}

// MinVoterCap is a free data retrieval call binding the contract method 0xf8ac9dd5.
//
// Solidity: function minVoterCap() constant returns(uint256)
func (_FNSValidator *FNSValidatorCaller) MinVoterCap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FNSValidator.contract.Call(opts, out, "minVoterCap")
	return *ret0, err
}

// MinVoterCap is a free data retrieval call binding the contract method 0xf8ac9dd5.
//
// Solidity: function minVoterCap() constant returns(uint256)
func (_FNSValidator *FNSValidatorSession) MinVoterCap() (*big.Int, error) {
	return _FNSValidator.Contract.MinVoterCap(&_FNSValidator.CallOpts)
}

// MinVoterCap is a free data retrieval call binding the contract method 0xf8ac9dd5.
//
// Solidity: function minVoterCap() constant returns(uint256)
func (_FNSValidator *FNSValidatorCallerSession) MinVoterCap() (*big.Int, error) {
	return _FNSValidator.Contract.MinVoterCap(&_FNSValidator.CallOpts)
}

// VoterWithdrawDelay is a free data retrieval call binding the contract method 0xa9ff959e.
//
// Solidity: function voterWithdrawDelay() constant returns(uint256)
func (_FNSValidator *FNSValidatorCaller) VoterWithdrawDelay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FNSValidator.contract.Call(opts, out, "voterWithdrawDelay")
	return *ret0, err
}

// VoterWithdrawDelay is a free data retrieval call binding the contract method 0xa9ff959e.
//
// Solidity: function voterWithdrawDelay() constant returns(uint256)
func (_FNSValidator *FNSValidatorSession) VoterWithdrawDelay() (*big.Int, error) {
	return _FNSValidator.Contract.VoterWithdrawDelay(&_FNSValidator.CallOpts)
}

// VoterWithdrawDelay is a free data retrieval call binding the contract method 0xa9ff959e.
//
// Solidity: function voterWithdrawDelay() constant returns(uint256)
func (_FNSValidator *FNSValidatorCallerSession) VoterWithdrawDelay() (*big.Int, error) {
	return _FNSValidator.Contract.VoterWithdrawDelay(&_FNSValidator.CallOpts)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(_candidate address) returns()
func (_FNSValidator *FNSValidatorTransactor) Propose(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
	return _FNSValidator.contract.Transact(opts, "propose", _candidate)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(_candidate address) returns()
func (_FNSValidator *FNSValidatorSession) Propose(_candidate common.Address) (*types.Transaction, error) {
	return _FNSValidator.Contract.Propose(&_FNSValidator.TransactOpts, _candidate)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(_candidate address) returns()
func (_FNSValidator *FNSValidatorTransactorSession) Propose(_candidate common.Address) (*types.Transaction, error) {
	return _FNSValidator.Contract.Propose(&_FNSValidator.TransactOpts, _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(_candidate address) returns()
func (_FNSValidator *FNSValidatorTransactor) Resign(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
	return _FNSValidator.contract.Transact(opts, "resign", _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(_candidate address) returns()
func (_FNSValidator *FNSValidatorSession) Resign(_candidate common.Address) (*types.Transaction, error) {
	return _FNSValidator.Contract.Resign(&_FNSValidator.TransactOpts, _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(_candidate address) returns()
func (_FNSValidator *FNSValidatorTransactorSession) Resign(_candidate common.Address) (*types.Transaction, error) {
	return _FNSValidator.Contract.Resign(&_FNSValidator.TransactOpts, _candidate)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(_candidate address, _cap uint256) returns()
func (_FNSValidator *FNSValidatorTransactor) Unvote(opts *bind.TransactOpts, _candidate common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _FNSValidator.contract.Transact(opts, "unvote", _candidate, _cap)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(_candidate address, _cap uint256) returns()
func (_FNSValidator *FNSValidatorSession) Unvote(_candidate common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _FNSValidator.Contract.Unvote(&_FNSValidator.TransactOpts, _candidate, _cap)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(_candidate address, _cap uint256) returns()
func (_FNSValidator *FNSValidatorTransactorSession) Unvote(_candidate common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _FNSValidator.Contract.Unvote(&_FNSValidator.TransactOpts, _candidate, _cap)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(_candidate address) returns()
func (_FNSValidator *FNSValidatorTransactor) Vote(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
	return _FNSValidator.contract.Transact(opts, "vote", _candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(_candidate address) returns()
func (_FNSValidator *FNSValidatorSession) Vote(_candidate common.Address) (*types.Transaction, error) {
	return _FNSValidator.Contract.Vote(&_FNSValidator.TransactOpts, _candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(_candidate address) returns()
func (_FNSValidator *FNSValidatorTransactorSession) Vote(_candidate common.Address) (*types.Transaction, error) {
	return _FNSValidator.Contract.Vote(&_FNSValidator.TransactOpts, _candidate)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(_blockNumber uint256, _index uint256) returns()
func (_FNSValidator *FNSValidatorTransactor) Withdraw(opts *bind.TransactOpts, _blockNumber *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _FNSValidator.contract.Transact(opts, "withdraw", _blockNumber, _index)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(_blockNumber uint256, _index uint256) returns()
func (_FNSValidator *FNSValidatorSession) Withdraw(_blockNumber *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _FNSValidator.Contract.Withdraw(&_FNSValidator.TransactOpts, _blockNumber, _index)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(_blockNumber uint256, _index uint256) returns()
func (_FNSValidator *FNSValidatorTransactorSession) Withdraw(_blockNumber *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _FNSValidator.Contract.Withdraw(&_FNSValidator.TransactOpts, _blockNumber, _index)
}

// FNSValidatorProposeIterator is returned from FilterPropose and is used to iterate over the raw logs and unpacked data for Propose events raised by the FNSValidator contract.
type FNSValidatorProposeIterator struct {
	Event *FNSValidatorPropose // Event containing the contract specifics and raw log

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
func (it *FNSValidatorProposeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FNSValidatorPropose)
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
		it.Event = new(FNSValidatorPropose)
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
func (it *FNSValidatorProposeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FNSValidatorProposeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FNSValidatorPropose represents a Propose event raised by the FNSValidator contract.
type FNSValidatorPropose struct {
	Owner     common.Address
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPropose is a free log retrieval operation binding the contract event 0x7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c1.
//
// Solidity: event Propose(_owner address, _candidate address, _cap uint256)
func (_FNSValidator *FNSValidatorFilterer) FilterPropose(opts *bind.FilterOpts) (*FNSValidatorProposeIterator, error) {

	logs, sub, err := _FNSValidator.contract.FilterLogs(opts, "Propose")
	if err != nil {
		return nil, err
	}
	return &FNSValidatorProposeIterator{contract: _FNSValidator.contract, event: "Propose", logs: logs, sub: sub}, nil
}

// WatchPropose is a free log subscription operation binding the contract event 0x7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c1.
//
// Solidity: event Propose(_owner address, _candidate address, _cap uint256)
func (_FNSValidator *FNSValidatorFilterer) WatchPropose(opts *bind.WatchOpts, sink chan<- *FNSValidatorPropose) (event.Subscription, error) {

	logs, sub, err := _FNSValidator.contract.WatchLogs(opts, "Propose")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FNSValidatorPropose)
				if err := _FNSValidator.contract.UnpackLog(event, "Propose", log); err != nil {
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

// FNSValidatorResignIterator is returned from FilterResign and is used to iterate over the raw logs and unpacked data for Resign events raised by the FNSValidator contract.
type FNSValidatorResignIterator struct {
	Event *FNSValidatorResign // Event containing the contract specifics and raw log

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
func (it *FNSValidatorResignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FNSValidatorResign)
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
		it.Event = new(FNSValidatorResign)
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
func (it *FNSValidatorResignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FNSValidatorResignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FNSValidatorResign represents a Resign event raised by the FNSValidator contract.
type FNSValidatorResign struct {
	Owner     common.Address
	Candidate common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterResign is a free log retrieval operation binding the contract event 0x4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d3.
//
// Solidity: event Resign(_owner address, _candidate address)
func (_FNSValidator *FNSValidatorFilterer) FilterResign(opts *bind.FilterOpts) (*FNSValidatorResignIterator, error) {

	logs, sub, err := _FNSValidator.contract.FilterLogs(opts, "Resign")
	if err != nil {
		return nil, err
	}
	return &FNSValidatorResignIterator{contract: _FNSValidator.contract, event: "Resign", logs: logs, sub: sub}, nil
}

// WatchResign is a free log subscription operation binding the contract event 0x4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d3.
//
// Solidity: event Resign(_owner address, _candidate address)
func (_FNSValidator *FNSValidatorFilterer) WatchResign(opts *bind.WatchOpts, sink chan<- *FNSValidatorResign) (event.Subscription, error) {

	logs, sub, err := _FNSValidator.contract.WatchLogs(opts, "Resign")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FNSValidatorResign)
				if err := _FNSValidator.contract.UnpackLog(event, "Resign", log); err != nil {
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

// FNSValidatorUnvoteIterator is returned from FilterUnvote and is used to iterate over the raw logs and unpacked data for Unvote events raised by the FNSValidator contract.
type FNSValidatorUnvoteIterator struct {
	Event *FNSValidatorUnvote // Event containing the contract specifics and raw log

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
func (it *FNSValidatorUnvoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FNSValidatorUnvote)
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
		it.Event = new(FNSValidatorUnvote)
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
func (it *FNSValidatorUnvoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FNSValidatorUnvoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FNSValidatorUnvote represents a Unvote event raised by the FNSValidator contract.
type FNSValidatorUnvote struct {
	Voter     common.Address
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnvote is a free log retrieval operation binding the contract event 0xaa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2.
//
// Solidity: event Unvote(_voter address, _candidate address, _cap uint256)
func (_FNSValidator *FNSValidatorFilterer) FilterUnvote(opts *bind.FilterOpts) (*FNSValidatorUnvoteIterator, error) {

	logs, sub, err := _FNSValidator.contract.FilterLogs(opts, "Unvote")
	if err != nil {
		return nil, err
	}
	return &FNSValidatorUnvoteIterator{contract: _FNSValidator.contract, event: "Unvote", logs: logs, sub: sub}, nil
}

// WatchUnvote is a free log subscription operation binding the contract event 0xaa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2.
//
// Solidity: event Unvote(_voter address, _candidate address, _cap uint256)
func (_FNSValidator *FNSValidatorFilterer) WatchUnvote(opts *bind.WatchOpts, sink chan<- *FNSValidatorUnvote) (event.Subscription, error) {

	logs, sub, err := _FNSValidator.contract.WatchLogs(opts, "Unvote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FNSValidatorUnvote)
				if err := _FNSValidator.contract.UnpackLog(event, "Unvote", log); err != nil {
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

// FNSValidatorVoteIterator is returned from FilterVote and is used to iterate over the raw logs and unpacked data for Vote events raised by the FNSValidator contract.
type FNSValidatorVoteIterator struct {
	Event *FNSValidatorVote // Event containing the contract specifics and raw log

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
func (it *FNSValidatorVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FNSValidatorVote)
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
		it.Event = new(FNSValidatorVote)
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
func (it *FNSValidatorVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FNSValidatorVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FNSValidatorVote represents a Vote event raised by the FNSValidator contract.
type FNSValidatorVote struct {
	Voter     common.Address
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0x66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc.
//
// Solidity: event Vote(_voter address, _candidate address, _cap uint256)
func (_FNSValidator *FNSValidatorFilterer) FilterVote(opts *bind.FilterOpts) (*FNSValidatorVoteIterator, error) {

	logs, sub, err := _FNSValidator.contract.FilterLogs(opts, "Vote")
	if err != nil {
		return nil, err
	}
	return &FNSValidatorVoteIterator{contract: _FNSValidator.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0x66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc.
//
// Solidity: event Vote(_voter address, _candidate address, _cap uint256)
func (_FNSValidator *FNSValidatorFilterer) WatchVote(opts *bind.WatchOpts, sink chan<- *FNSValidatorVote) (event.Subscription, error) {

	logs, sub, err := _FNSValidator.contract.WatchLogs(opts, "Vote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FNSValidatorVote)
				if err := _FNSValidator.contract.UnpackLog(event, "Vote", log); err != nil {
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

// FNSValidatorWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the FNSValidator contract.
type FNSValidatorWithdrawIterator struct {
	Event *FNSValidatorWithdraw // Event containing the contract specifics and raw log

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
func (it *FNSValidatorWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FNSValidatorWithdraw)
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
		it.Event = new(FNSValidatorWithdraw)
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
func (it *FNSValidatorWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FNSValidatorWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FNSValidatorWithdraw represents a Withdraw event raised by the FNSValidator contract.
type FNSValidatorWithdraw struct {
	Owner       common.Address
	BlockNumber *big.Int
	Cap         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(_owner address, _blockNumber uint256, _cap uint256)
func (_FNSValidator *FNSValidatorFilterer) FilterWithdraw(opts *bind.FilterOpts) (*FNSValidatorWithdrawIterator, error) {

	logs, sub, err := _FNSValidator.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &FNSValidatorWithdrawIterator{contract: _FNSValidator.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(_owner address, _blockNumber uint256, _cap uint256)
func (_FNSValidator *FNSValidatorFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *FNSValidatorWithdraw) (event.Subscription, error) {

	logs, sub, err := _FNSValidator.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FNSValidatorWithdraw)
				if err := _FNSValidator.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058204e9ff712f75026048f9d46bb29317eda21df26c9adbab8fa62492c441878d48a0029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}
