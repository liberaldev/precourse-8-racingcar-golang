package random_number

import (
	"crypto/rand"
	"math/big"
)

type DefaultRandomGenerator struct{}

func (d *DefaultRandomGenerator) GetRandomNumber() int64 {
	if n, err := rand.Int(rand.Reader, big.NewInt(10)); err == nil {
		return n.Int64()
	}
	return 0
}
