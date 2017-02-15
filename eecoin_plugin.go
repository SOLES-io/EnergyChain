package eecoin

import (
	"encoding/hex"
	"fmt"

	abci "github.com/tendermint/abci/types"
	"github.com/tendermint/basecoin/state"
	"github.com/tendermint/basecoin/types"
	wire "github.com/tendermint/go-wire"
)

const (
	AddBanker    = "add"
	RemoveBanker = "remove"
)

// EECoinPlugin is a plugin, storing all state prefixed with it's unique name
type EECoinPlugin struct {
	name string
}

func New(name string) EECoinPlugin {
	return EECoinPlugin{name: name}
}

func (eep EECoinPlugin) Name() string {
	return eep.name
}

// Set initial minters
func (eep EECoinPlugin) SetOption(store types.KVStore, key string, value string) (log string) {
	// value is always a hex-encoded address
	addr, err := hex.DecodeString(value)
	if err != nil {
		return fmt.Sprintf("Invalid address: %s: %v", addr, err)
	}

	switch key {
	case AddBanker:
		s := eep.loadState(store)
		s.AddBanker(addr)
		eep.saveState(store, s)
		eep.saveState(store, s)
		return fmt.Sprintf("Added: %s", addr)
	case RemoveBanker:                           // eecoin may not need this, SOLES is the only minter
		s := eep.loadState(store)
		s.RemoveBanker(addr)
		eep.saveState(store, s)
		return fmt.Sprintf("Removed: %s", addr)
	default:
		return fmt.Sprintf("Unknown key: %s", key)
	}
}

// This allows
func (eep EECoinPlugin) RunTx(store types.KVStore, ctx types.CallContext, txBytes []byte) (res abci.Result) {
	// parse transaction
	var tx EECoinTx
	err := wire.ReadBinaryBytes(txBytes, &tx)
	if err != nil {
		return abci.ErrEncodingError
	}

	// make sure it was signed by a banker
	s := eep.loadState(store)
	if !s.IsBanker(ctx.CallerAddress) {
		return abci.ErrUnauthorized
	}

	// now, send all this money!
	for _, receiver := range tx.Receivers {
		// load or create account
		acct := state.GetAccount(store, receiver.Addr)
		if acct == nil {
			acct = &types.Account{
				PubKey:   nil,
				Sequence: 0,
			}
		}

		// add the money
		acct.Balance = acct.Balance.Plus(receiver.Amount)

		// and save the new balance
		state.SetAccount(store, receiver.Addr, acct)
	}

	return abci.Result{}
}

// placeholders empty to fulfill interface
func (eep EECoinPlugin) InitChain(store types.KVStore, vals []*abci.Validator) {}
func (eep EECoinPlugin) BeginBlock(store types.KVStore, height uint64)         {}
func (eep EECoinPlugin) EndBlock(store types.KVStore, height uint64) []*abci.Validator {
	return nil
}

/*** implementation ***/

func (eep EECoinPlugin) stateKey() []byte {
	key := fmt.Sprintf("*%s*", eep.name)
	return []byte(key)
}

func (eep EECoinPlugin) loadState(store types.KVStore) *EECoinState {
	var s EECoinState
	data := store.Get(eep.stateKey())
	// here return an uninitialized state
	if len(data) == 0 {
		return &s
	}

	err := wire.ReadBinaryBytes(data, &s)
	// this should never happen, but we should also never panic....
	if err != nil {
		panic(err)
	}
	return &s
}

func (eep EECoinPlugin) saveState(store types.KVStore, state *EECoinState) {
	value := wire.BinaryBytes(*state)
	store.Set(eep.stateKey(), value)
}
