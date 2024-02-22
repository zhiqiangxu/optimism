// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
)

// SoulETHMetaData contains all meta data concerning the SoulETH contract.
var SoulETHMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addBurners\",\"inputs\":[{\"name\":\"burners_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addMinters\",\"inputs\":[{\"name\":\"minters_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchMint\",\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delBurners\",\"inputs\":[{\"name\":\"burners_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delMinters\",\"inputs\":[{\"name\":\"minters_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"name_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minters_\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"burners_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100de565b600054610100900460ff161561008a5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811610156100dc576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b611f13806100ed6000396000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c806371e2a657116100d8578063a457c2d71161008c578063dd62ed3e11610066578063dd62ed3e1461030c578063e04b818014610352578063f2fde38b1461036557600080fd5b8063a457c2d7146102d3578063a9059cbb146102e6578063b8de86b4146102f957600080fd5b80638da5cb5b116100bd5780638da5cb5b1461029057806395d89b41146102b8578063991fa7bd146102c057600080fd5b806371e2a6571461026a57806379cc67901461027d57600080fd5b8063395093511161012f5780636857310711610114578063685731071461021957806370a082311461022c578063715018a61461026257600080fd5b806339509351146101f15780633ab84dd91461020457600080fd5b806318160ddd1161016057806318160ddd146101bd57806323b872dd146101cf578063313ce567146101e257600080fd5b806306fdde031461017c578063095ea7b31461019a575b600080fd5b610184610378565b60405161019191906118f2565b60405180910390f35b6101ad6101a836600461198e565b61040a565b6040519015158152602001610191565b6035545b604051908152602001610191565b6101ad6101dd3660046119b8565b610474565b60405160128152602001610191565b6101ad6101ff36600461198e565b6104fe565b610217610212366004611a40565b610554565b005b610217610227366004611a82565b610622565b6101c161023a366004611aee565b73ffffffffffffffffffffffffffffffffffffffff1660009081526033602052604090205490565b610217610793565b610217610278366004611a40565b6107a7565b61021761028b36600461198e565b61086c565b60655460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610191565b610184610935565b6102176102ce366004611b52565b610944565b6101ad6102e136600461198e565b610d37565b6101ad6102f436600461198e565b610e13565b610217610307366004611a40565b610e77565b6101c161031a366004611c2c565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260346020908152604080832093909416825291909152205490565b610217610360366004611a40565b610f36565b610217610373366004611aee565b610ff2565b60606036805461038790611c5f565b80601f01602080910402602001604051908101604052809291908181526020018280546103b390611c5f565b80156104005780601f106103d557610100808354040283529160200191610400565b820191906000526020600020905b8154815290600101906020018083116103e357829003601f168201915b5050505050905090565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f617070726f76652069732064697361626c656420666f7220536f756c4554480060448201526000906064015b60405180910390fd5b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f7472616e7366657246726f6d2069732064697361626c656420666f7220536f7560448201527f6c45544800000000000000000000000000000000000000000000000000000000606482015260009060840161046b565b33600081815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061054a9082908690610545908790611ce1565b6110a9565b5060019392505050565b61055c61125c565b7f6bbe77d0d543accaa2114faf81aa28535044f3b83160e7815884a82039c7dc0060005b8281101561061c5760018260010160008686858181106105a2576105a2611cf9565b90506020020160208101906105b79190611aee565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790558061061481611d28565b915050610580565b50505050565b82811461068b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f696e76616c696420617267756d656e7473000000000000000000000000000000604482015260640161046b565b3360009081527f6bbe77d0d543accaa2114faf81aa28535044f3b83160e7815884a82039c7dc00602081905260409091205460ff16610726576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f6e6f742061206d696e7465720000000000000000000000000000000000000000604482015260640161046b565b60005b8481101561078b5761077986868381811061074657610746611cf9565b905060200201602081019061075b9190611aee565b85858481811061076d5761076d611cf9565b905060200201356112dd565b8061078381611d28565b915050610729565b505050505050565b61079b61125c565b6107a560006113fe565b565b6107af61125c565b7f6bbe77d0d543accaa2114faf81aa28535044f3b83160e7815884a82039c7dc0060005b8281101561061c5760018260008686858181106107f2576107f2611cf9565b90506020020160208101906108079190611aee565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790558061086481611d28565b9150506107d3565b3360009081527f6bbe77d0d543accaa2114faf81aa28535044f3b83160e7815884a82039c7dc0160205260409020547f6bbe77d0d543accaa2114faf81aa28535044f3b83160e7815884a82039c7dc009060ff16610926576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f6e6f742061206275726e65720000000000000000000000000000000000000000604482015260640161046b565b6109308383611475565b505050565b60606037805461038790611c5f565b600054610100900460ff16158080156109645750600054600160ff909116105b8061097e5750303b15801561097e575060005460ff166001145b610a0a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161046b565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610a6857600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff8616610ae5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f6f776e6572206973207a65726f00000000000000000000000000000000000000604482015260640161046b565b610b588a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8e018190048102820181019092528c815292508c91508b908190840183828082843760009201919091525061166292505050565b610b60611703565b610b6986610ff2565b7f6bbe77d0d543accaa2114faf81aa28535044f3b83160e7815884a82039c7dc0060005b85811015610c26576001826000898985818110610bac57610bac611cf9565b9050602002016020810190610bc19190611aee565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001691151591909117905580610c1e81611d28565b915050610b8d565b5060005b83811015610cc6576001826001016000878785818110610c4c57610c4c611cf9565b9050602002016020810190610c619190611aee565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001691151591909117905580610cbe81611d28565b915050610c2a565b50508015610d2b57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050505050565b33600081815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919083811015610dfb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f000000000000000000000000000000000000000000000000000000606482015260840161046b565b610e0882868684036110a9565b506001949350505050565b60006040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046b906020808252818101527f7472616e736665722069732064697361626c656420666f7220536f756c455448604082015260600190565b610e7f61125c565b7f6bbe77d0d543accaa2114faf81aa28535044f3b83160e7815884a82039c7dc0060005b8281101561061c57816001016000858584818110610ec357610ec3611cf9565b9050602002016020810190610ed89190611aee565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905580610f2e81611d28565b915050610ea3565b610f3e61125c565b7f6bbe77d0d543accaa2114faf81aa28535044f3b83160e7815884a82039c7dc0060005b8281101561061c57816000858584818110610f7f57610f7f611cf9565b9050602002016020810190610f949190611aee565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905580610fea81611d28565b915050610f62565b610ffa61125c565b73ffffffffffffffffffffffffffffffffffffffff811661109d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161046b565b6110a6816113fe565b50565b73ffffffffffffffffffffffffffffffffffffffff831661114b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f7265737300000000000000000000000000000000000000000000000000000000606482015260840161046b565b73ffffffffffffffffffffffffffffffffffffffff82166111ee576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f7373000000000000000000000000000000000000000000000000000000000000606482015260840161046b565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526034602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b60655473ffffffffffffffffffffffffffffffffffffffff1633146107a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161046b565b73ffffffffffffffffffffffffffffffffffffffff821661135a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640161046b565b806035600082825461136c9190611ce1565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600090815260336020526040812080548392906113a6908490611ce1565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35b5050565b6065805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b73ffffffffffffffffffffffffffffffffffffffff8216611518576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f7300000000000000000000000000000000000000000000000000000000000000606482015260840161046b565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260336020526040902054818110156115ce576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f6365000000000000000000000000000000000000000000000000000000000000606482015260840161046b565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260336020526040812083830390556035805484929061160a908490611d60565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a3505050565b600054610100900460ff166116f9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161046b565b6113fa82826117a2565b600054610100900460ff1661179a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161046b565b6107a5611852565b600054610100900460ff16611839576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161046b565b60366118458382611dec565b5060376109308282611dec565b600054610100900460ff166118e9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161046b565b6107a5336113fe565b600060208083528351808285015260005b8181101561191f57858101830151858201604001528201611903565b81811115611931576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461198957600080fd5b919050565b600080604083850312156119a157600080fd5b6119aa83611965565b946020939093013593505050565b6000806000606084860312156119cd57600080fd5b6119d684611965565b92506119e460208501611965565b9150604084013590509250925092565b60008083601f840112611a0657600080fd5b50813567ffffffffffffffff811115611a1e57600080fd5b6020830191508360208260051b8501011115611a3957600080fd5b9250929050565b60008060208385031215611a5357600080fd5b823567ffffffffffffffff811115611a6a57600080fd5b611a76858286016119f4565b90969095509350505050565b60008060008060408587031215611a9857600080fd5b843567ffffffffffffffff80821115611ab057600080fd5b611abc888389016119f4565b90965094506020870135915080821115611ad557600080fd5b50611ae2878288016119f4565b95989497509550505050565b600060208284031215611b0057600080fd5b611b0982611965565b9392505050565b60008083601f840112611b2257600080fd5b50813567ffffffffffffffff811115611b3a57600080fd5b602083019150836020828501011115611a3957600080fd5b600080600080600080600080600060a08a8c031215611b7057600080fd5b893567ffffffffffffffff80821115611b8857600080fd5b611b948d838e01611b10565b909b50995060208c0135915080821115611bad57600080fd5b611bb98d838e01611b10565b9099509750879150611bcd60408d01611965565b965060608c0135915080821115611be357600080fd5b611bef8d838e016119f4565b909650945060808c0135915080821115611c0857600080fd5b50611c158c828d016119f4565b915080935050809150509295985092959850929598565b60008060408385031215611c3f57600080fd5b611c4883611965565b9150611c5660208401611965565b90509250929050565b600181811c90821680611c7357607f821691505b602082108103611cac577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115611cf457611cf4611cb2565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611d5957611d59611cb2565b5060010190565b600082821015611d7257611d72611cb2565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f82111561093057600081815260208120601f850160051c81016020861015611dcd5750805b601f850160051c820191505b8181101561078b57828155600101611dd9565b815167ffffffffffffffff811115611e0657611e06611d77565b611e1a81611e148454611c5f565b84611da6565b602080601f831160018114611e6d5760008415611e375750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b17855561078b565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015611eba57888601518255948401946001909101908401611e9b565b5085821015611ef657878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b0190555056fea164736f6c634300080f000a",
}

// SoulETHABI is the input ABI used to generate the binding from.
// Deprecated: Use SoulETHMetaData.ABI instead.
var SoulETHABI = SoulETHMetaData.ABI

// SoulETHBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SoulETHMetaData.Bin instead.
var SoulETHBin = SoulETHMetaData.Bin

// DeploySoulETH deploys a new Ethereum contract, binding an instance of SoulETH to it.
func DeploySoulETH(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SoulETH, error) {
	parsed, err := SoulETHMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SoulETHBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SoulETH{SoulETHCaller: SoulETHCaller{contract: contract}, SoulETHTransactor: SoulETHTransactor{contract: contract}, SoulETHFilterer: SoulETHFilterer{contract: contract}}, nil
}

// SoulETH is an auto generated Go binding around an Ethereum contract.
type SoulETH struct {
	SoulETHCaller     // Read-only binding to the contract
	SoulETHTransactor // Write-only binding to the contract
	SoulETHFilterer   // Log filterer for contract events
}

// SoulETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type SoulETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SoulETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SoulETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SoulETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SoulETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SoulETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SoulETHSession struct {
	Contract     *SoulETH          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SoulETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SoulETHCallerSession struct {
	Contract *SoulETHCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SoulETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SoulETHTransactorSession struct {
	Contract     *SoulETHTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SoulETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type SoulETHRaw struct {
	Contract *SoulETH // Generic contract binding to access the raw methods on
}

// SoulETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SoulETHCallerRaw struct {
	Contract *SoulETHCaller // Generic read-only contract binding to access the raw methods on
}

// SoulETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SoulETHTransactorRaw struct {
	Contract *SoulETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSoulETH creates a new instance of SoulETH, bound to a specific deployed contract.
func NewSoulETH(address common.Address, backend bind.ContractBackend) (*SoulETH, error) {
	contract, err := bindSoulETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SoulETH{SoulETHCaller: SoulETHCaller{contract: contract}, SoulETHTransactor: SoulETHTransactor{contract: contract}, SoulETHFilterer: SoulETHFilterer{contract: contract}}, nil
}

// NewSoulETHCaller creates a new read-only instance of SoulETH, bound to a specific deployed contract.
func NewSoulETHCaller(address common.Address, caller bind.ContractCaller) (*SoulETHCaller, error) {
	contract, err := bindSoulETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SoulETHCaller{contract: contract}, nil
}

// NewSoulETHTransactor creates a new write-only instance of SoulETH, bound to a specific deployed contract.
func NewSoulETHTransactor(address common.Address, transactor bind.ContractTransactor) (*SoulETHTransactor, error) {
	contract, err := bindSoulETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SoulETHTransactor{contract: contract}, nil
}

// NewSoulETHFilterer creates a new log filterer instance of SoulETH, bound to a specific deployed contract.
func NewSoulETHFilterer(address common.Address, filterer bind.ContractFilterer) (*SoulETHFilterer, error) {
	contract, err := bindSoulETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SoulETHFilterer{contract: contract}, nil
}

// bindSoulETH binds a generic wrapper to an already deployed contract.
func bindSoulETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SoulETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SoulETH *SoulETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SoulETH.Contract.SoulETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SoulETH *SoulETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SoulETH.Contract.SoulETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SoulETH *SoulETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SoulETH.Contract.SoulETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SoulETH *SoulETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SoulETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SoulETH *SoulETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SoulETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SoulETH *SoulETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SoulETH.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SoulETH *SoulETHCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SoulETH.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SoulETH *SoulETHSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SoulETH.Contract.Allowance(&_SoulETH.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SoulETH *SoulETHCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SoulETH.Contract.Allowance(&_SoulETH.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SoulETH *SoulETHCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SoulETH.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SoulETH *SoulETHSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SoulETH.Contract.BalanceOf(&_SoulETH.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SoulETH *SoulETHCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SoulETH.Contract.BalanceOf(&_SoulETH.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SoulETH *SoulETHCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SoulETH.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SoulETH *SoulETHSession) Decimals() (uint8, error) {
	return _SoulETH.Contract.Decimals(&_SoulETH.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SoulETH *SoulETHCallerSession) Decimals() (uint8, error) {
	return _SoulETH.Contract.Decimals(&_SoulETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SoulETH *SoulETHCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SoulETH.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SoulETH *SoulETHSession) Name() (string, error) {
	return _SoulETH.Contract.Name(&_SoulETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SoulETH *SoulETHCallerSession) Name() (string, error) {
	return _SoulETH.Contract.Name(&_SoulETH.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SoulETH *SoulETHCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SoulETH.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SoulETH *SoulETHSession) Owner() (common.Address, error) {
	return _SoulETH.Contract.Owner(&_SoulETH.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SoulETH *SoulETHCallerSession) Owner() (common.Address, error) {
	return _SoulETH.Contract.Owner(&_SoulETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SoulETH *SoulETHCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SoulETH.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SoulETH *SoulETHSession) Symbol() (string, error) {
	return _SoulETH.Contract.Symbol(&_SoulETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SoulETH *SoulETHCallerSession) Symbol() (string, error) {
	return _SoulETH.Contract.Symbol(&_SoulETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SoulETH *SoulETHCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SoulETH.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SoulETH *SoulETHSession) TotalSupply() (*big.Int, error) {
	return _SoulETH.Contract.TotalSupply(&_SoulETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SoulETH *SoulETHCallerSession) TotalSupply() (*big.Int, error) {
	return _SoulETH.Contract.TotalSupply(&_SoulETH.CallOpts)
}

// AddBurners is a paid mutator transaction binding the contract method 0x3ab84dd9.
//
// Solidity: function addBurners(address[] burners_) returns()
func (_SoulETH *SoulETHTransactor) AddBurners(opts *bind.TransactOpts, burners_ []common.Address) (*types.Transaction, error) {
	return _SoulETH.contract.Transact(opts, "addBurners", burners_)
}

// AddBurners is a paid mutator transaction binding the contract method 0x3ab84dd9.
//
// Solidity: function addBurners(address[] burners_) returns()
func (_SoulETH *SoulETHSession) AddBurners(burners_ []common.Address) (*types.Transaction, error) {
	return _SoulETH.Contract.AddBurners(&_SoulETH.TransactOpts, burners_)
}

// AddBurners is a paid mutator transaction binding the contract method 0x3ab84dd9.
//
// Solidity: function addBurners(address[] burners_) returns()
func (_SoulETH *SoulETHTransactorSession) AddBurners(burners_ []common.Address) (*types.Transaction, error) {
	return _SoulETH.Contract.AddBurners(&_SoulETH.TransactOpts, burners_)
}

// AddMinters is a paid mutator transaction binding the contract method 0x71e2a657.
//
// Solidity: function addMinters(address[] minters_) returns()
func (_SoulETH *SoulETHTransactor) AddMinters(opts *bind.TransactOpts, minters_ []common.Address) (*types.Transaction, error) {
	return _SoulETH.contract.Transact(opts, "addMinters", minters_)
}

// AddMinters is a paid mutator transaction binding the contract method 0x71e2a657.
//
// Solidity: function addMinters(address[] minters_) returns()
func (_SoulETH *SoulETHSession) AddMinters(minters_ []common.Address) (*types.Transaction, error) {
	return _SoulETH.Contract.AddMinters(&_SoulETH.TransactOpts, minters_)
}

// AddMinters is a paid mutator transaction binding the contract method 0x71e2a657.
//
// Solidity: function addMinters(address[] minters_) returns()
func (_SoulETH *SoulETHTransactorSession) AddMinters(minters_ []common.Address) (*types.Transaction, error) {
	return _SoulETH.Contract.AddMinters(&_SoulETH.TransactOpts, minters_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_SoulETH *SoulETHTransactor) Approve(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _SoulETH.contract.Transact(opts, "approve", arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_SoulETH *SoulETHSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _SoulETH.Contract.Approve(&_SoulETH.TransactOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_SoulETH *SoulETHTransactorSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _SoulETH.Contract.Approve(&_SoulETH.TransactOpts, arg0, arg1)
}

// BatchMint is a paid mutator transaction binding the contract method 0x68573107.
//
// Solidity: function batchMint(address[] accounts, uint256[] values) returns()
func (_SoulETH *SoulETHTransactor) BatchMint(opts *bind.TransactOpts, accounts []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _SoulETH.contract.Transact(opts, "batchMint", accounts, values)
}

// BatchMint is a paid mutator transaction binding the contract method 0x68573107.
//
// Solidity: function batchMint(address[] accounts, uint256[] values) returns()
func (_SoulETH *SoulETHSession) BatchMint(accounts []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _SoulETH.Contract.BatchMint(&_SoulETH.TransactOpts, accounts, values)
}

// BatchMint is a paid mutator transaction binding the contract method 0x68573107.
//
// Solidity: function batchMint(address[] accounts, uint256[] values) returns()
func (_SoulETH *SoulETHTransactorSession) BatchMint(accounts []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _SoulETH.Contract.BatchMint(&_SoulETH.TransactOpts, accounts, values)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_SoulETH *SoulETHTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, value *big.Int) (*types.Transaction, error) {
	return _SoulETH.contract.Transact(opts, "burnFrom", account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_SoulETH *SoulETHSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _SoulETH.Contract.BurnFrom(&_SoulETH.TransactOpts, account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_SoulETH *SoulETHTransactorSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _SoulETH.Contract.BurnFrom(&_SoulETH.TransactOpts, account, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SoulETH *SoulETHTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SoulETH.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SoulETH *SoulETHSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SoulETH.Contract.DecreaseAllowance(&_SoulETH.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SoulETH *SoulETHTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SoulETH.Contract.DecreaseAllowance(&_SoulETH.TransactOpts, spender, subtractedValue)
}

// DelBurners is a paid mutator transaction binding the contract method 0xb8de86b4.
//
// Solidity: function delBurners(address[] burners_) returns()
func (_SoulETH *SoulETHTransactor) DelBurners(opts *bind.TransactOpts, burners_ []common.Address) (*types.Transaction, error) {
	return _SoulETH.contract.Transact(opts, "delBurners", burners_)
}

// DelBurners is a paid mutator transaction binding the contract method 0xb8de86b4.
//
// Solidity: function delBurners(address[] burners_) returns()
func (_SoulETH *SoulETHSession) DelBurners(burners_ []common.Address) (*types.Transaction, error) {
	return _SoulETH.Contract.DelBurners(&_SoulETH.TransactOpts, burners_)
}

// DelBurners is a paid mutator transaction binding the contract method 0xb8de86b4.
//
// Solidity: function delBurners(address[] burners_) returns()
func (_SoulETH *SoulETHTransactorSession) DelBurners(burners_ []common.Address) (*types.Transaction, error) {
	return _SoulETH.Contract.DelBurners(&_SoulETH.TransactOpts, burners_)
}

// DelMinters is a paid mutator transaction binding the contract method 0xe04b8180.
//
// Solidity: function delMinters(address[] minters_) returns()
func (_SoulETH *SoulETHTransactor) DelMinters(opts *bind.TransactOpts, minters_ []common.Address) (*types.Transaction, error) {
	return _SoulETH.contract.Transact(opts, "delMinters", minters_)
}

// DelMinters is a paid mutator transaction binding the contract method 0xe04b8180.
//
// Solidity: function delMinters(address[] minters_) returns()
func (_SoulETH *SoulETHSession) DelMinters(minters_ []common.Address) (*types.Transaction, error) {
	return _SoulETH.Contract.DelMinters(&_SoulETH.TransactOpts, minters_)
}

// DelMinters is a paid mutator transaction binding the contract method 0xe04b8180.
//
// Solidity: function delMinters(address[] minters_) returns()
func (_SoulETH *SoulETHTransactorSession) DelMinters(minters_ []common.Address) (*types.Transaction, error) {
	return _SoulETH.Contract.DelMinters(&_SoulETH.TransactOpts, minters_)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SoulETH *SoulETHTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SoulETH.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SoulETH *SoulETHSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SoulETH.Contract.IncreaseAllowance(&_SoulETH.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SoulETH *SoulETHTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SoulETH.Contract.IncreaseAllowance(&_SoulETH.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x991fa7bd.
//
// Solidity: function initialize(string name_, string symbol_, address owner_, address[] minters_, address[] burners_) returns()
func (_SoulETH *SoulETHTransactor) Initialize(opts *bind.TransactOpts, name_ string, symbol_ string, owner_ common.Address, minters_ []common.Address, burners_ []common.Address) (*types.Transaction, error) {
	return _SoulETH.contract.Transact(opts, "initialize", name_, symbol_, owner_, minters_, burners_)
}

// Initialize is a paid mutator transaction binding the contract method 0x991fa7bd.
//
// Solidity: function initialize(string name_, string symbol_, address owner_, address[] minters_, address[] burners_) returns()
func (_SoulETH *SoulETHSession) Initialize(name_ string, symbol_ string, owner_ common.Address, minters_ []common.Address, burners_ []common.Address) (*types.Transaction, error) {
	return _SoulETH.Contract.Initialize(&_SoulETH.TransactOpts, name_, symbol_, owner_, minters_, burners_)
}

// Initialize is a paid mutator transaction binding the contract method 0x991fa7bd.
//
// Solidity: function initialize(string name_, string symbol_, address owner_, address[] minters_, address[] burners_) returns()
func (_SoulETH *SoulETHTransactorSession) Initialize(name_ string, symbol_ string, owner_ common.Address, minters_ []common.Address, burners_ []common.Address) (*types.Transaction, error) {
	return _SoulETH.Contract.Initialize(&_SoulETH.TransactOpts, name_, symbol_, owner_, minters_, burners_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SoulETH *SoulETHTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SoulETH.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SoulETH *SoulETHSession) RenounceOwnership() (*types.Transaction, error) {
	return _SoulETH.Contract.RenounceOwnership(&_SoulETH.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SoulETH *SoulETHTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SoulETH.Contract.RenounceOwnership(&_SoulETH.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_SoulETH *SoulETHTransactor) Transfer(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _SoulETH.contract.Transact(opts, "transfer", arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_SoulETH *SoulETHSession) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _SoulETH.Contract.Transfer(&_SoulETH.TransactOpts, arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_SoulETH *SoulETHTransactorSession) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _SoulETH.Contract.Transfer(&_SoulETH.TransactOpts, arg0, arg1)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_SoulETH *SoulETHTransactor) TransferFrom(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _SoulETH.contract.Transact(opts, "transferFrom", arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_SoulETH *SoulETHSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _SoulETH.Contract.TransferFrom(&_SoulETH.TransactOpts, arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_SoulETH *SoulETHTransactorSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _SoulETH.Contract.TransferFrom(&_SoulETH.TransactOpts, arg0, arg1, arg2)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SoulETH *SoulETHTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SoulETH.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SoulETH *SoulETHSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SoulETH.Contract.TransferOwnership(&_SoulETH.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SoulETH *SoulETHTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SoulETH.Contract.TransferOwnership(&_SoulETH.TransactOpts, newOwner)
}

// SoulETHApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SoulETH contract.
type SoulETHApprovalIterator struct {
	Event *SoulETHApproval // Event containing the contract specifics and raw log

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
func (it *SoulETHApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulETHApproval)
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
		it.Event = new(SoulETHApproval)
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
func (it *SoulETHApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulETHApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulETHApproval represents a Approval event raised by the SoulETH contract.
type SoulETHApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SoulETH *SoulETHFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SoulETHApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SoulETH.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SoulETHApprovalIterator{contract: _SoulETH.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SoulETH *SoulETHFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SoulETHApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SoulETH.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulETHApproval)
				if err := _SoulETH.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SoulETH *SoulETHFilterer) ParseApproval(log types.Log) (*SoulETHApproval, error) {
	event := new(SoulETHApproval)
	if err := _SoulETH.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulETHInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SoulETH contract.
type SoulETHInitializedIterator struct {
	Event *SoulETHInitialized // Event containing the contract specifics and raw log

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
func (it *SoulETHInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulETHInitialized)
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
		it.Event = new(SoulETHInitialized)
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
func (it *SoulETHInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulETHInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulETHInitialized represents a Initialized event raised by the SoulETH contract.
type SoulETHInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SoulETH *SoulETHFilterer) FilterInitialized(opts *bind.FilterOpts) (*SoulETHInitializedIterator, error) {

	logs, sub, err := _SoulETH.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SoulETHInitializedIterator{contract: _SoulETH.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SoulETH *SoulETHFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SoulETHInitialized) (event.Subscription, error) {

	logs, sub, err := _SoulETH.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulETHInitialized)
				if err := _SoulETH.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SoulETH *SoulETHFilterer) ParseInitialized(log types.Log) (*SoulETHInitialized, error) {
	event := new(SoulETHInitialized)
	if err := _SoulETH.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulETHOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SoulETH contract.
type SoulETHOwnershipTransferredIterator struct {
	Event *SoulETHOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SoulETHOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulETHOwnershipTransferred)
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
		it.Event = new(SoulETHOwnershipTransferred)
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
func (it *SoulETHOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulETHOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulETHOwnershipTransferred represents a OwnershipTransferred event raised by the SoulETH contract.
type SoulETHOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SoulETH *SoulETHFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SoulETHOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SoulETH.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SoulETHOwnershipTransferredIterator{contract: _SoulETH.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SoulETH *SoulETHFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SoulETHOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SoulETH.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulETHOwnershipTransferred)
				if err := _SoulETH.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SoulETH *SoulETHFilterer) ParseOwnershipTransferred(log types.Log) (*SoulETHOwnershipTransferred, error) {
	event := new(SoulETHOwnershipTransferred)
	if err := _SoulETH.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SoulETHTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SoulETH contract.
type SoulETHTransferIterator struct {
	Event *SoulETHTransfer // Event containing the contract specifics and raw log

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
func (it *SoulETHTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SoulETHTransfer)
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
		it.Event = new(SoulETHTransfer)
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
func (it *SoulETHTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SoulETHTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SoulETHTransfer represents a Transfer event raised by the SoulETH contract.
type SoulETHTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SoulETH *SoulETHFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SoulETHTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SoulETH.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SoulETHTransferIterator{contract: _SoulETH.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SoulETH *SoulETHFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SoulETHTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SoulETH.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SoulETHTransfer)
				if err := _SoulETH.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SoulETH *SoulETHFilterer) ParseTransfer(log types.Log) (*SoulETHTransfer, error) {
	event := new(SoulETHTransfer)
	if err := _SoulETH.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
