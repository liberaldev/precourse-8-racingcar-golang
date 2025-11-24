package random_number

import (
	"crypto/rand"
	"math/big"
)

type DefaultRandomGenerator struct{}

func (d *DefaultRandomGenerator) GetRandomNumber(maxNum *big.Int) (int64, error) {
	n, err := rand.Int(rand.Reader, maxNum)
	return n.Int64(), err
}
