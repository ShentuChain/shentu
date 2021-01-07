package keeper

import (
	"github.com/hyperledger/burrow/execution/errors"

	"github.com/certikfoundation/shentu/x/cvm/types"
)

var (
	InsufficientGasErrorCode   = types.BurrowErrorCodeStart + errors.Codes.InsufficientGas.ErrorCode().Number
	CodeOutOfBoundsErrorCode   = types.BurrowErrorCodeStart + errors.Codes.CodeOutOfBounds.ErrorCode().Number
	ExecutionRevertedErrorCode = types.BurrowErrorCodeStart + errors.Codes.ExecutionReverted.ErrorCode().Number
	/*
		pragma solidity >=0.4.22 <0.6.0;
		contract Hello55 {
			function sayHi() public pure returns (uint) {
				return 55;
			}
		}
	*/
	Hello55BytecodeString     = "6080604052348015600f57600080fd5b5060ac8061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80630c49c36c14602d575b600080fd5b60336049565b6040518082815260200191505060405180910390f35b6000603790509056fea26469706673582212207e355097621bbba6da748b17d355f936e7b5ad809077d16433f46dbfda2cae0364736f6c637826302e362e342d646576656c6f702e323032302e332e352b636f6d6d69742e33326361316135650057"
	Hello55AbiJsonString      = `[{"inputs":[],"name":"sayHi","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"pure","type":"function"}]`
	Hello55MetadataJsonString = `{"compiler":{"version":"0.6.4-develop.2020.3.5+commit.32ca1a5e"},"language":"Solidity","output":{"abi":[{"inputs":[],"name":"sayHi","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"pure","type":"function"}],"devdoc":{"methods":{}},"userdoc":{"methods":{}}},"settings":{"compilationTarget":{"tests/hello55.sol":"Hello55"},"evmVersion":"istanbul","libraries":{},"metadata":{"bytecodeHash":"ipfs"},"optimizer":{"enabled":false,"runs":200},"remappings":[]},"sources":{"tests/hello55.sol":{"keccak256":"0x743a9d971915a1fe43a196d542a13ab0f4f36f9e3c27579eb3e64b78a0469182","urls":["bzz-raw://55a4dce17dbdde893afbfde0cbe9dd4ff7b9a281e1e3e4a47a1ea7a8733aa1de","dweb:/ipfs/QmTR4LxBT2F38RzuCogTjpk8iEXRpz2RRNjeLB8Bwc6SQs"]}},"version":1}`

	/*
		pragma solidity >=0.4.22 <0.6.0;
		contract BasicTests {
			uint myFavoriteNumber = 34;
			function addTwoNumbers(uint a, uint b) public pure returns (uint) {
				return a + b;
			}
			function failureFunction() public pure {
				revert("Go away!!");
			}
			function setMyFavoriteNumber(uint newFavNum) public {
				myFavoriteNumber = newFavNum;
			}
		}
	*/
	BasicTestsBytecodeString = "6080604052602260005534801561001557600080fd5b50610184806100256000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630b30d76414610046578063a0eb379f14610074578063e2276f1c1461007e575b600080fd5b6100726004803603602081101561005c57600080fd5b81019080803590602001909291905050506100ca565b005b61007c6100d4565b005b6100b46004803603604081101561009457600080fd5b810190808035906020019092919080359060200190929190505050610142565b6040518082815260200191505060405180910390f35b8060008190555050565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260098152602001807f476f20617761792121000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b600081830190509291505056fea265627a7a7231582029e87152c00d34140b78a06d51e5b41bdd4eab369148d1b9540394dcc93f1d5e64736f6c634300050b0032"
	BasicTestsAbiJsonString  = `
	[
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "uint256",
				"name": "newFavNum",
				"type": "uint256"
			}
		],
		"name": "setMyFavoriteNumber",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "failureFunction",
		"outputs": [],
		"payable": false,
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"internalType": "uint256",
				"name": "a",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "b",
				"type": "uint256"
			}
		],
		"name": "addTwoNumbers",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "pure",
		"type": "function"
	}
]`
	/*
		pragma solidity >=0.4.22 <0.6.0;
		contract GoofyContract {
			uint public goofyNumber;
			function setGoofyNumber(uint a) public {
				goofyNumber = a;
			}
		}
		contract GasTests {
			GoofyContract gc;
			function addTwoNumbers(uint a, uint b) public pure returns (uint) {
				return a + b;
			}
			function hashMe(bytes memory b) public pure returns (bytes32) {
				return keccak256(b);
			}
			function deployAnotherContract() public {
				gc = new GoofyContract();
			}
			function setGoofyNumber(uint a) public returns (uint) {
				gc.setGoofyNumber(a);
				return gc.goofyNumber();
			}
		}
	*/
	GasTestsBytecodeString = "608060405234801561001057600080fd5b5061049a806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80632e60e461146100515780635422426314610120578063b71b449014610162578063e2276f1c1461016c575b600080fd5b61010a6004803603602081101561006757600080fd5b810190808035906020019064010000000081111561008457600080fd5b82018360208201111561009657600080fd5b803590602001918460018302840111640100000000831117156100b857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506101b8565b6040518082815260200191505060405180910390f35b61014c6004803603602081101561013657600080fd5b81019080803590602001909291905050506101c9565b6040518082815260200191505060405180910390f35b61016a610300565b005b6101a26004803603604081101561018257600080fd5b81019080803590602001909291908035906020019092919050505061036a565b6040518082815260200191505060405180910390f35b600081805190602001209050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166354224263836040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b15801561023f57600080fd5b505af1158015610253573d6000803e3d6000fd5b505050506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630a2346c66040518163ffffffff1660e01b815260040160206040518083038186803b1580156102be57600080fd5b505afa1580156102d2573d6000803e3d6000fd5b505050506040513d60208110156102e857600080fd5b81019080805190602001909291905050509050919050565b60405161030c90610377565b604051809103906000f080158015610328573d6000803e3d6000fd5b506000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000818301905092915050565b60e2806103848339019056fe608060405234801561001057600080fd5b5060c38061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80630a2346c614603757806354224263146053575b600080fd5b603d607e565b6040518082815260200191505060405180910390f35b607c60048036036020811015606757600080fd5b81019080803590602001909291905050506084565b005b60005481565b806000819055505056fea265627a7a72315820bd31260ca27a654607a2d4b452d03506665f2e48c15e43ae924552f645d26de864736f6c634300050c0032a265627a7a723158209366e5562db198cc3463c52ac5b103071ee3b30d645e47b876037d90f7c4b6d564736f6c634300050c0032"
	GasTestsAbiJsonString  = `
	[
		{
			"constant": true,
			"inputs": [
				{
					"internalType": "bytes",
					"name": "b",
					"type": "bytes"
				}
			],
			"name": "hashMe",
			"outputs": [
				{
					"internalType": "bytes32",
					"name": "",
					"type": "bytes32"
				}
			],
			"payable": false,
			"stateMutability": "pure",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [
				{
					"internalType": "uint256",
					"name": "a",
					"type": "uint256"
				}
			],
			"name": "setGoofyNumber",
			"outputs": [
				{
					"internalType": "uint256",
					"name": "",
					"type": "uint256"
				}
			],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": false,
			"inputs": [],
			"name": "deployAnotherContract",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"constant": true,
			"inputs": [
				{
					"internalType": "uint256",
					"name": "a",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "b",
					"type": "uint256"
				}
			],
			"name": "addTwoNumbers",
			"outputs": [
				{
					"internalType": "uint256",
					"name": "",
					"type": "uint256"
				}
			],
			"payable": false,
			"stateMutability": "pure",
			"type": "function"
		}
	]
	`
	//derived using remix.ethereum.org
	AddTwoNumbersGasCost         uint64 = 13055
	HashMeGasCost                uint64 = 13099
	DeployAnotherContractGasCost uint64 = 113549 + 92277
	SetGoofyNumberGasCost        uint64 = 47895 + 26367

	/*
		pragma solidity >=0.4.22 <0.6.0;
		contract GasRefund {
			uint stupidNumber;
			constructor() public {
				stupidNumber=10000000;
			}
			function iWillRevert() public {
				uint a = 4 + 5;
				revert("thats enough work for now");
			}
			function iWillFail() public {
				uint a = 4 - 4;
				uint b = 6 / a;
			}
			function deleteFromStorage() public {
				stupidNumber = 0;
			}
			function die() public {
				selfdestruct(address(0x0));
			}
		}
	*/
	//GasRefundBytecodeString = /* version with self destruct address as longer string */ "608060405234801561001057600080fd5b50629896806000819055506101708061002a6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806335f469941461005157806360a5f1871461005b57806383bfecfa146100655780639e2ff0741461006f575b600080fd5b610059610079565b005b6100636100a6565b005b61006d6100af565b005b6100776100c7565b005b73ab35ee8df2f8dd950cc1cfd38fef86857374e97173ffffffffffffffffffffffffffffffffffffffff16ff5b60008081905550565b60008090506000816006816100c057fe5b0490505050565b6000600990506040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f746861747320656e6f75676820776f726b20666f72206e6f770000000000000081525060200191505060405180910390fdfea265627a7a72315820e6b239ca4c62450752d8bb2877fe37bc15f1cc95afb6f9387df4199c97f6b6c864736f6c634300050c0032"
	GasRefundBytecodeString = /* version with self destruct address as zero */ "608060405234801561001057600080fd5b506298968060008190555061015d8061002a6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806335f469941461005157806360a5f1871461005b57806383bfecfa146100655780639e2ff0741461006f575b600080fd5b610059610079565b005b610063610093565b005b61006d61009c565b005b6100776100b4565b005b600073ffffffffffffffffffffffffffffffffffffffff16ff5b60008081905550565b60008090506000816006816100ad57fe5b0490505050565b6000600990506040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f746861747320656e6f75676820776f726b20666f72206e6f770000000000000081525060200191505060405180910390fdfea265627a7a7231582086b971c2ec8ef8b8fb46a876f47f527b6e879ecd64bb16d3d864febac60b29c364736f6c634300050c0032"
	GasRefundAbiJsonString  = `
		[
			{
				"inputs": [],
				"payable": false,
				"stateMutability": "nonpayable",
				"type": "constructor"
			},
			{
				"constant": false,
				"inputs": [],
				"name": "deleteFromStorage",
				"outputs": [],
				"payable": false,
				"stateMutability": "nonpayable",
				"type": "function"
			},
			{
				"constant": false,
				"inputs": [],
				"name": "die",
				"outputs": [],
				"payable": false,
				"stateMutability": "nonpayable",
				"type": "function"
			},
			{
				"constant": false,
				"inputs": [],
				"name": "iWillFail",
				"outputs": [],
				"payable": false,
				"stateMutability": "nonpayable",
				"type": "function"
			},
			{
				"constant": false,
				"inputs": [],
				"name": "iWillRevert",
				"outputs": [],
				"payable": false,
				"stateMutability": "nonpayable",
				"type": "function"
			}
		]
		`

	CtkTransferTestBytecodeString = "608060405234801561001057600080fd5b50610145806100206000396000f3fe6080604052600436106100295760003560e01c80634897e0631461002e578063eb53b14e14610072575b600080fd5b6100706004803603602081101561004457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061009d565b005b34801561007e57600080fd5b506100876100f1565b6040518082815260200191505060405180910390f35b8073ffffffffffffffffffffffffffffffffffffffff166108fc600234816100c157fe5b049081150290604051600060405180830381858888f193505050501580156100ed573d6000803e3d6000fd5b5050565b60003073ffffffffffffffffffffffffffffffffffffffff163190509056fea265627a7a72315820995e9941f88fede53e75c0509196b74084b2cb6778b668a6cddee82c542e075564736f6c634300050c0032"
	CtkTransferTestAbiJsonString  = `
			[
				{
					"constant": false,
					"inputs": [
						{
							"internalType": "address payable",
							"name": "friend",
							"type": "address"
						}
					],
					"name": "sendToAFriend",
					"outputs": [],
					"payable": true,
					"stateMutability": "payable",
					"type": "function"
				},
				{
					"constant": true,
					"inputs": [],
					"name": "whatsMyBalance",
					"outputs": [
						{
							"internalType": "uint256",
							"name": "",
							"type": "uint256"
						}
					],
					"payable": false,
					"stateMutability": "view",
					"type": "function"
				}
			]
			`

	TestCheckBytecodeString = "60806040526040518060600160405280602d8152602001610900602d91396000908051906020019061003292919061014d565b506040518060600160405280602d8152602001610987602d91396001908051906020019061006192919061014d565b506040518060600160405280602d81526020016109b4602d91396002908051906020019061009092919061014d565b506040518060600160405280602d815260200161092d602d9139600390805190602001906100bf92919061014d565b506040518060600160405280602d815260200161095a602d9139600490805190602001906100ee92919061014d565b506040518060400160405280601381526020017f64756d6d79736f75726365636f646568617368000000000000000000000000008152506005908051906020019061013a92919061014d565b5034801561014757600080fd5b506101f2565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061018e57805160ff19168380011785556101bc565b828001600101855582156101bc579182015b828111156101bb5782518255916020019190600101906101a0565b5b5090506101c991906101cd565b5090565b6101ef91905b808211156101eb5760008160009055506001016101d3565b5090565b90565b6106ff806102016000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630bfe74e81461005c5780634a6abcda146100df5780638504647c146101625780639ca131af146101e5578063de78f9fc14610268575b600080fd5b6100646102eb565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100a4578082015181840152602081019050610089565b50505050905090810190601f1680156100d15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6100e76103c9565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561012757808201518184015260208101905061010c565b50505050905090810190601f1680156101545780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61016a610480565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101aa57808201518184015260208101905061018f565b50505050905090810190601f1680156101d75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101ed610537565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561022d578082015181840152602081019050610212565b50505050905090810190601f16801561025a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102706105ee565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102b0578082015181840152602081019050610295565b50505050905090810190601f1680156102dd5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60608060048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103845780601f1061035957610100808354040283529160200191610384565b820191906000526020600020905b81548152906001019060200180831161036757829003601f168201915b50505050509050600181516001828260208601600961c350fa6001838360208701600a61c350fa60018214156103c35760018114156103c257600180f35b5b60016000f35b60608060018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104625780601f1061043757610100808354040283529160200191610462565b820191906000526020600020905b81548152906001019060200180831161044557829003601f168201915b50505050509050600181516001828260208601600961c350fa600183f35b60608060008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105195780601f106104ee57610100808354040283529160200191610519565b820191906000526020600020905b8154815290600101906020018083116104fc57829003601f168201915b50505050509050600181516001828260208601600961c350fa600183f35b60608060028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105d05780601f106105a5576101008083540402835291602001916105d0565b820191906000526020600020905b8154815290600101906020018083116105b357829003601f168201915b50505050509050600181516001828260208601600a61c350fa600183f35b60608060058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106875780601f1061065c57610100808354040283529160200191610687565b820191906000526020600020905b81548152906001019060200180831161066a57829003601f168201915b50505050509050600181516001828260208601600b61c350fa600183f3fea264697066735822122079782cdcd176d9d80a7b3dc3550f6232424a9c779c74a91ec463a87b4934587c64736f6c637826302e362e342d646576656c6f702e323032302e332e352b636f6d6d69742e3332636131613565005763657274696b316a3974686e7733363764373274787a726d73753277736b35717237756764383668707438716763657274696b3175647a7032337477663461367066343765723236756a67637a7a683374637374376335617a6563657274696b316e7a677664346b33347a7a6636766b3371657675683578783878736867367579306c3872643563657274696b31713768356532677770796b71323765746e64356b356663766332617a707565756a716376617063657274696b316d39356b6676616a7735646d6e75396536683336357471636e617a6d39617564646e676b6c76"
	TestCheckAbiJsonString  = `[{"inputs":[],"name":"callCheck","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"callCheckNotCertified","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"compilationCheck","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"proofAndAuditingCheck","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"proofCheck","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"nonpayable","type":"function"}]`

	TestCertifyValidatorString        = "60806040526040518060800160405280605381526020016102bf6053913960009080519060200190610032929190610045565b5034801561003f57600080fd5b506100ea565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061008657805160ff19168380011785556100b4565b828001600101855582156100b4579182015b828111156100b3578251825591602001919060010190610098565b5b5090506100c191906100c5565b5090565b6100e791905b808211156100e35760008160009055506001016100cb565b5090565b90565b6101c6806100f96000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80633c1bf57b14610030575b600080fd5b6100386100b3565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561007857808201518184015260208101905061005d565b50505050905090810190601f1680156100a55780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60608060008054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561014c5780601f106101215761010080835404028352916020019161014c565b820191906000526020600020905b81548152906001019060200180831161012f57829003601f168201915b505050505090506001815160018282602086016000600c61c350f1600183f3fea2646970667358221220fd5a2e955887ad15e2ad81a4830c4742eb2884f797e0e5bf73c915354ad2e69a64736f6c637826302e362e342d646576656c6f702e323032302e332e352b636f6d6d69742e3332636131613565005763657274696b76616c636f6e73707562317a636a647565707133327636356565676b327976677a6479613564716e6c6e6330363375376d7433646836367a3278797639726464676d367439347334706a656174"
	TestCertifyValidatorAbiJsonString = `[{"inputs":[],"name":"certifyValidator","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"nonpayable","type":"function"}]`
)
