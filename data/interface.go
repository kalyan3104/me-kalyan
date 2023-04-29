package data

import (
	"math/big"
)

// HeaderHandler defines getters and setters for header data holder
type HeaderHandler interface {
	GetShardID() uint32
	GetNonce() uint64
	GetEpoch() uint32
	GetRound() uint64
	GetRootHash() []byte
	GetValidatorStatsRootHash() []byte
	GetPrevHash() []byte
	GetPrevRandSeed() []byte
	GetRandSeed() []byte
	GetPubKeysBitmap() []byte
	GetSignature() []byte
	GetLeaderSignature() []byte
	GetChainID() []byte
	GetSoftwareVersion() []byte
	GetTimeStamp() uint64
	GetTxCount() uint32
	GetReceiptsHash() []byte
	GetAccumulatedFees() *big.Int
	GetDeveloperFees() *big.Int
	GetEpochStartMetaHash() []byte
	GetReserved() []byte

	SetAccumulatedFees(value *big.Int)
	SetDeveloperFees(value *big.Int)
	SetShardID(shId uint32)
	SetNonce(n uint64)
	SetEpoch(e uint32)
	SetRound(r uint64)
	SetTimeStamp(ts uint64)
	SetRootHash(rHash []byte)
	SetValidatorStatsRootHash(rHash []byte)
	SetPrevHash(pvHash []byte)
	SetPrevRandSeed(pvRandSeed []byte)
	SetRandSeed(randSeed []byte)
	SetPubKeysBitmap(pkbm []byte)
	SetSignature(sg []byte)
	SetLeaderSignature(sg []byte)
	SetChainID(chainID []byte)
	SetSoftwareVersion(version []byte)
	SetTxCount(txCount uint32)

	IsStartOfEpochBlock() bool
	GetMiniBlockHeadersWithDst(destId uint32) map[string]uint32
	GetOrderedCrossMiniblocksWithDst(destId uint32) []*MiniBlockInfo
	GetMiniBlockHeadersHashes() [][]byte

	IsInterfaceNil() bool
	Clone() HeaderHandler
}

// BodyHandler interface for a block body
type BodyHandler interface {
	Clone() BodyHandler
	// IntegrityAndValidity checks the integrity and validity of the block
	IntegrityAndValidity() error
	// IsInterfaceNil returns true if there is no value under the interface
	IsInterfaceNil() bool
}

// ChainHandler is the interface defining the functionality a blockchain should implement
type ChainHandler interface {
	GetGenesisHeader() HeaderHandler
	SetGenesisHeader(gb HeaderHandler) error
	GetGenesisHeaderHash() []byte
	SetGenesisHeaderHash(hash []byte)
	GetCurrentBlockHeader() HeaderHandler
	SetCurrentBlockHeader(bh HeaderHandler) error
	GetCurrentBlockHeaderHash() []byte
	SetCurrentBlockHeaderHash(hash []byte)
	IsInterfaceNil() bool
	CreateNewHeader() HeaderHandler
}

// TransactionHandler defines the type of executable transaction
type TransactionHandler interface {
	IsInterfaceNil() bool

	GetValue() *big.Int
	GetNonce() uint64
	GetData() []byte
	GetRcvAddr() []byte
	GetRcvUserName() []byte
	GetSndAddr() []byte
	GetGasLimit() uint64
	GetGasPrice() uint64

	SetValue(*big.Int)
	SetData([]byte)
	SetRcvAddr([]byte)
	SetSndAddr([]byte)
	Size() int

	CheckIntegrity() error
}

// LogHandler defines the type for a log resulted from executing a transaction or smart contract call
type LogHandler interface {
	// GetAddress returns the address of the sc that was originally called by the user
	GetAddress() []byte
	// GetLogEvents returns the events from a transaction log entry
	GetLogEvents() []EventHandler

	IsInterfaceNil() bool
}

// EventHandler defines the type for an event resulted from a smart contract call contained in a log
type EventHandler interface {
	// GetAddress returns the address of the contract that generated this event
	//  - in sc calling another sc situation this will differ from the
	//    LogHandler's GetAddress, whereas in the single sc situation
	//    they will be the same
	GetAddress() []byte
	// GetIdentifier returns identifier of the event, that together with the ABI can
	//   be used to understand the type of the event by other applications
	GetIdentifier() []byte
	// GetTopics returns the data that can be indexed so that it would be searchable
	//  by other applications
	GetTopics() [][]byte
	// GetData returns the rest of the event data, which will not be indexed, so storing
	//  information here should be cheaper
	GetData() []byte

	IsInterfaceNil() bool
}

// ValidatorInfoHandler is used to store multiple validatorInfo properties
type ValidatorInfoHandler interface {
	GetPublicKey() []byte
	GetShardId() uint32
	GetList() string
	GetIndex() uint32
	GetTempRating() uint32
	GetRating() uint32
	String() string
	IsInterfaceNil() bool
}

// ShardValidatorInfoHandler is used to store multiple validatorInfo properties required in shards
type ShardValidatorInfoHandler interface {
	GetPublicKey() []byte
	GetTempRating() uint32
	String() string
	IsInterfaceNil() bool
}

// GoRoutineThrottler can monitor the number of the currently running go routines
type GoRoutineThrottler interface {
	CanProcess() bool
	StartProcessing()
	EndProcessing()
	IsInterfaceNil() bool
}

// MiniBlockInfo holds information about a cross miniblock referenced in a received block
type MiniBlockInfo struct {
	Hash          []byte
	SenderShardID uint32
	Round         uint64
}

// SyncStatisticsHandler defines the methods for a component able to store the sync statistics for a trie
type SyncStatisticsHandler interface {
	Reset()
	AddNumReceived(value int)
	AddNumLarge(value int)
	SetNumMissing(rootHash []byte, value int)
	NumReceived() int
	NumLarge() int
	NumMissing() int
	IsInterfaceNil() bool
}

// TransactionWithFeeHandler represents a transaction structure that has economics variables defined
type TransactionWithFeeHandler interface {
	GetGasLimit() uint64
	GetGasPrice() uint64
	GetData() []byte
	GetRcvAddr() []byte
	GetValue() *big.Int
}

// UserAccountHandler models a user account
type UserAccountHandler interface {
	RetrieveValueFromDataTrieTracker(key []byte) ([]byte, error)
	GetBalance() *big.Int
	GetNonce() uint64
	AddressBytes() []byte
	IsInterfaceNil() bool
}
