package app

import (
	"math/big"

	"github.com/frenchie-foundation/go-lachesis/inter/idx"
)

// SetSfcConstants stores SfcConstants
func (s *Store) SetSfcConstants(epoch idx.Epoch, constants SfcConstants) {
	s.set(s.table.SfcConstants, epoch.Bytes(), constants)
}

// HasSfcConstants returns true if SFC constants are stored
func (s *Store) HasSfcConstants(epoch idx.Epoch) bool {
	ok, err := s.table.SfcConstants.Has(epoch.Bytes())
	if err != nil {
		s.Log.Crit("Failed to check key-value")
	}
	return ok
}

// GetSfcConstants returns stored SfcConstants
func (s *Store) GetSfcConstants(epoch idx.Epoch) SfcConstants {
	w, _ := s.get(s.table.SfcConstants, epoch.Bytes(), &SfcConstants{}).(*SfcConstants)

	if w == nil {
		w = &SfcConstants{
			ShortGasPowerAllocPerSec: 0,
			LongGasPowerAllocPerSec:  0,
			BaseRewardPerSec:         new(big.Int),
			OfflinePenaltyThreshold: BlocksMissed{
				Num:    0,
				Period: 0,
			},
			MinGasPrice: new(big.Int),
		}
	}

	return *w
}
