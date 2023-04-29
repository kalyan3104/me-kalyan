package core

// NodeType represents the node's role in the network
type NodeType string

// NodeTypeObserver signals that a node is running as observer node
const NodeTypeObserver NodeType = "observer"

// NodeTypeValidator signals that a node is running as validator node
const NodeTypeValidator NodeType = "validator"

// pkPrefixSize specifies the max numbers of chars to be displayed from one publc key
const pkPrefixSize = 12

// FileModeUserReadWrite represents the permission for a file which allows the user for reading and writing
const FileModeUserReadWrite = 0600

// FileModeReadWrite represents the permission for a file which allows reading and writing for user and group and read
// for others
const FileModeReadWrite = 0664

// MetachainShardId will be used to identify a shard ID as metachain
const MetachainShardId = uint32(0xFFFFFFFF)

// AllShardId will be used to identify that a message is for all shards
const AllShardId = uint32(0xFFFFFFF0)

// MegabyteSize represents the size in bytes of a megabyte
const MegabyteSize = 1024 * 1024

// BuiltInFunctionClaimDeveloperRewards is the key for the claim developer rewards built-in function
const BuiltInFunctionClaimDeveloperRewards = "ClaimDeveloperRewards"

// BuiltInFunctionChangeOwnerAddress is the key for the change owner built in function built-in function
const BuiltInFunctionChangeOwnerAddress = "ChangeOwnerAddress"

// BuiltInFunctionSetUserName is the key for the set user name built-in function
const BuiltInFunctionSetUserName = "SetUserName"

// BuiltInFunctionSaveKeyValue is the key for the save key value built-in function
const BuiltInFunctionSaveKeyValue = "SaveKeyValue"

// BuiltInFunctionDCTTransfer is the key for the Dharitri core token (Dct) transfer built-in function
const BuiltInFunctionDCTTransfer = "DCTTransfer"

// BuiltInFunctionDCTBurn is the key for the Dharitri core token (Dct) burn built-in function
const BuiltInFunctionDCTBurn = "DCTBurn"

// BuiltInFunctionDCTFreeze is the key for the Dharitri core token (Dct) freeze built-in function
const BuiltInFunctionDCTFreeze = "DCTFreeze"

// BuiltInFunctionDCTUnFreeze is the key for the Dharitri core token (Dct) unfreeze built-in function
const BuiltInFunctionDCTUnFreeze = "DCTUnFreeze"

// BuiltInFunctionDCTWipe is the key for the Dharitri core token (Dct) wipe built-in function
const BuiltInFunctionDCTWipe = "DCTWipe"

// BuiltInFunctionDCTPause is the key for the Dharitri core token (Dct) pause built-in function
const BuiltInFunctionDCTPause = "DCTPause"

// BuiltInFunctionDCTUnPause is the key for the Dharitri core token (Dct) unpause built-in function
const BuiltInFunctionDCTUnPause = "DCTUnPause"

// BuiltInFunctionSetDCTRole is the key for the Dharitri core token (Dct) set built-in function
const BuiltInFunctionSetDCTRole = "DCTSetRole"

// BuiltInFunctionUnSetDCTRole is the key for the Dharitri core token (Dct) unset built-in function
const BuiltInFunctionUnSetDCTRole = "DCTUnSetRole"

// BuiltInFunctionDCTSetLimitedTransfer is the key for the Dharitri core token (Dct) built-in function which sets the property
// for the token to be transferable only through accounts with transfer roles
const BuiltInFunctionDCTSetLimitedTransfer = "DCTSetLimitedTransfer"

// BuiltInFunctionDCTUnSetLimitedTransfer is the key for the Dharitri core token (Dct) built-in function which unsets the property
// for the token to be transferable only through accounts with transfer roles
const BuiltInFunctionDCTUnSetLimitedTransfer = "DCTUnSetLimitedTransfer"

// BuiltInFunctionDCTLocalMint is the key for the Dharitri core token (Dct) local mint built-in function
const BuiltInFunctionDCTLocalMint = "DCTLocalMint"

// BuiltInFunctionDCTLocalBurn is the key for the Dharitri core token (Dct) local burn built-in function
const BuiltInFunctionDCTLocalBurn = "DCTLocalBurn"

// BuiltInFunctionDCTNFTTransfer is the key for the Dharitri core token (Dct) NFT transfer built-in function
const BuiltInFunctionDCTNFTTransfer = "DCTNFTTransfer"

// BuiltInFunctionDCTNFTCreate is the key for the Dharitri core token (Dct) NFT create built-in function
const BuiltInFunctionDCTNFTCreate = "DCTNFTCreate"

// BuiltInFunctionDCTNFTAddQuantity is the key for the Dharitri core token (Dct) NFT add quantity built-in function
const BuiltInFunctionDCTNFTAddQuantity = "DCTNFTAddQuantity"

// BuiltInFunctionDCTNFTCreateRoleTransfer is the key for the Dharitri core token (Dct) create role transfer function
const BuiltInFunctionDCTNFTCreateRoleTransfer = "DCTNFTCreateRoleTransfer"

// BuiltInFunctionDCTNFTBurn is the key for the Dharitri core token (Dct) NFT burn built-in function
const BuiltInFunctionDCTNFTBurn = "DCTNFTBurn"

// BuiltInFunctionDCTNFTAddURI is the key for the Dharitri core token (Dct) NFT add URI built-in function
const BuiltInFunctionDCTNFTAddURI = "DCTNFTAddURI"

// BuiltInFunctionDCTNFTUpdateAttributes is the key for the Dharitri core token (Dct) NFT update attributes built-in function
const BuiltInFunctionDCTNFTUpdateAttributes = "DCTNFTUpdateAttributes"

// BuiltInFunctionMultiDCTNFTTransfer is the key for the Dharitri core token (Dct) multi transfer built-in function
const BuiltInFunctionMultiDCTNFTTransfer = "MultiDCTNFTTransfer"

// DCTRoleLocalMint is the constant string for the local role of mint for DCT tokens
const DCTRoleLocalMint = "DCTRoleLocalMint"

// DCTRoleLocalBurn is the constant string for the local role of burn for DCT tokens
const DCTRoleLocalBurn = "DCTRoleLocalBurn"

// DCTRoleNFTCreate is the constant string for the local role of create for DCT NFT tokens
const DCTRoleNFTCreate = "DCTRoleNFTCreate"

// DCTRoleNFTCreateMultiShard is the constant string for the local role of create for DCT NFT tokens multishard
const DCTRoleNFTCreateMultiShard = "DCTRoleNFTCreateMultiShard"

// DCTRoleNFTAddQuantity is the constant string for the local role of adding quantity for existing DCT NFT tokens
const DCTRoleNFTAddQuantity = "DCTRoleNFTAddQuantity"

// DCTRoleNFTBurn is the constant string for the local role of burn for DCT NFT tokens
const DCTRoleNFTBurn = "DCTRoleNFTBurn"

// DCTRoleNFTAddURI is the constant string for the local role of adding a URI for DCT NFT tokens
const DCTRoleNFTAddURI = "DCTRoleNFTAddURI"

// DCTRoleNFTUpdateAttributes is the constant string for the local role of updating attributes for DCT NFT tokens
const DCTRoleNFTUpdateAttributes = "DCTRoleNFTUpdateAttributes"

// DCTRoleTransfer is the constant string for the local role to transfer DCT, only for special tokens
const DCTRoleTransfer = "DCTTransferRole"

// DCTType defines the possible types in case of DCT tokens
type DCTType uint32

const (
	// Fungible defines the token type for DCT fungible tokens
	Fungible DCTType = iota
	// NonFungible defines the token type for DCT non fungible tokens
	NonFungible
)

// FungibleDCT defines the string for the token type of fungible DCT
const FungibleDCT = "FungibleDCT"

// NonFungibleDCT defines the string for the token type of non fungible DCT
const NonFungibleDCT = "NonFungibleDCT"

// SemiFungibleDCT defines the string for the token type of semi fungible DCT
const SemiFungibleDCT = "SemiFungibleDCT"

// MaxRoyalty defines 100% as uint32
const MaxRoyalty = uint32(10000)

// RelayedTransaction is the key for the dharitri meta/gassless/relayed transaction standard
const RelayedTransaction = "relayedTx"

// RelayedTransactionV2 is the key for the optimized dharitri meta/gassless/relayed transaction standard
const RelayedTransactionV2 = "relayedTxV2"

// SCDeployInitFunctionName is the key for the function which is called at smart contract deploy time
const SCDeployInitFunctionName = "_init"

// DharitriProtectedKeyPrefix is the key prefix which is protected from writing in the trie - only for special builtin functions
const DharitriProtectedKeyPrefix = "DHARITRI"

// DelegationSystemSCKey is the key under which there is data in case of system delegation smart contracts
const DelegationSystemSCKey = "delegation"

// DCTKeyIdentifier is the key prefix for dct tokens
const DCTKeyIdentifier = "dct"

// DCTRoleIdentifier is the key prefix for dct role identifier
const DCTRoleIdentifier = "role"

// DCTNFTLatestNonceIdentifier is the key prefix for dct latest nonce identifier
const DCTNFTLatestNonceIdentifier = "nonce"

// MaxNumShards represents the maximum number of shards possible in the system
const MaxNumShards = 256

// MinMetaTxExtraGasCost is the constant defined for minimum gas value to be sent in meta transaction
const MinMetaTxExtraGasCost = uint64(1_000_000)

// MaxLeafSize represents maximum amount of data which can be saved under one leaf
const MaxLeafSize = uint64(1 << 26) //64MB

// MaxBufferSizeToSendTrieNodes represents max buffer size to send in bytes used when resolving trie nodes
// Every trie node that has a greater size than this constant is considered a large trie node and should be split in
// smaller chunks
const MaxBufferSizeToSendTrieNodes = 1 << 18 //256KB

// MaxUserNameLength represents the maximum number of bytes a UserName can have
const MaxUserNameLength = 32

// MinLenArgumentsDCTTransfer defines the min length of arguments for the DCT transfer
const MinLenArgumentsDCTTransfer = 2

// MinLenArgumentsDCTNFTTransfer defines the minimum length for dct nft transfer
const MinLenArgumentsDCTNFTTransfer = 4

// MaxLenForDCTIssueMint defines the maximum length in bytes for the issued/minted balance
const MaxLenForDCTIssueMint = 100

// BaseOperationCostString represents the field name for base operation costs
const BaseOperationCostString = "BaseOperationCost"

// BuiltInCostString represents the field name for built in operation costs
const BuiltInCostString = "BuiltInCost"

// DCTSCAddress is the hard-coded address for dct issuing smart contract
var DCTSCAddress = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 255, 255}

// SCDeployIdentifier is the identifier for a smart contract deploy
const SCDeployIdentifier = "SCDeploy"

// SCUpgradeIdentifier is the identifier for a smart contract upgrade
const SCUpgradeIdentifier = "SCUpgrade"
