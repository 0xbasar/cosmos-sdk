package types

import (
	"context"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	tmtypes "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ProtoApp is what we would need to implement in the new BaseApp.
type ProtoApp interface {
	ChainID() string
	Name() string
	Version() string

	InitialHeight() int64
	LastBlockHeight() int64
	AppVersion() (uint64, error)

	// TODO: Replace this with Marko's TX validation
	ValidateTX([]byte) (gInfo sdk.GasInfo, result *sdk.Result, anteEvents []abci.Event, err error) // TODO: I'm just replicating what runTx replies here

	// New interface
	InitChain(RequestInitChain) (ResponseInitChain, error)
	Query(abci.RequestQuery) (*abci.ResponseQuery, error) // TODO: replace queries req/res with something more generic
	DeliverBlock(Header, [][]byte) ([]interface{}, error)
	Commit() error

	// COMET BFT specific stuff below (tbd where to put them)
	Validators() []abci.ValidatorUpdate
	ConsensusParams() *tmtypes.ConsensusParams
	AppHash() []byte
	GetBlockRetentionHeight(commitHeight int64) int64
}

type ProposalHandler interface {
	PrepareProposal(context.Context, *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error)
	ProcessProposal(context.Context, *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error)
}

type VoteExtensionsHandler interface {
	ExtendVote(context.Context, *abci.RequestExtendVote) (*abci.ResponseExtendVote, error)
	VerifyVoteExtension(context.Context, *abci.RequestVerifyVoteExtension) (*abci.ResponseVerifyVoteExtension, error)
}

type RequestInitChain struct {
	StateBytes []byte
}

type ResponseInitChain struct{}

// Header defines a generic header interface. // TODO: move to core?
type Header interface {
	GetHeight() uint64  // GetHeight returns the height of the block
	GetHash() []byte    // GetHash returns the hash of the block header
	GetTime() time.Time // GetTime returns the time of the block
	GetChainID() string // GetChainID returns the chain ID of the chain
	GetAppHash() []byte // GetAppHash used in the current block header
}

// CometBFTHeader
type CometBFTHeader struct {
	Height  int64     // Height returns the height of the block
	Hash    []byte    // Hash returns the hash of the block header
	Time    time.Time // Time returns the time of the block
	ChainID string    // ChainId returns the chain ID of the block
	AppHash []byte    // AppHash used in the current block header

	NextValidatorsHash []byte
	ProposerAddress    []byte
	LastCommit         abci.CommitInfo
	Misbehavior        []abci.Misbehavior
}

func (h CometBFTHeader) GetHeight() uint64 {
	return uint64(h.Height)
}

func (h CometBFTHeader) GetHash() []byte {
	return h.Hash
}

func (h CometBFTHeader) GetTime() time.Time {
	return h.Time
}

func (h CometBFTHeader) GetChainID() string {
	return h.ChainID
}

func (h CometBFTHeader) GetAppHash() []byte {
	return h.AppHash
}
