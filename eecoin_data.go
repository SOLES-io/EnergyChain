package eecoin

import (
	"bytes"

	"github.com/tendermint/basecoin/types"
	wire "github.com/tendermint/go-wire"
)

type EECoinState struct {
	Bankers [][]byte
}

func (s *EECoinState) AddBanker(addr []byte) {
	if !s.IsBanker(addr) {
		s.Bankers = append(s.Bankers, addr)
	}
}

func (s *EECoinState) RemoveBanker(addr []byte) {
	b := s.Bankers
	for i := range b {
		if bytes.Equal(addr, b[i]) {
			s.Bankers = append(b[:i], b[i+1:]...)
			return
		}
	}
}

func (s *EECoinState) IsBanker(addr []byte) bool {
	for _, b := range s.Bankers {
		if bytes.Equal(b, addr) {
			return true
		}
	}
	return false
}

type EECoinTx struct {
	Receivers []Receiver
}

type Receiver struct {
	Addr   []byte
	Amount types.Coins
}

func (tx EECoinTx) Serialize() []byte {
	return wire.BinaryBytes(tx)
}
