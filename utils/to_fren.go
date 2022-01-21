package utils

import "math/big"

// ToFren number of FREN to Wei
func ToFren(fren uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(fren), big.NewInt(1e18))
}
