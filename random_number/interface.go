package random_number

import "math/big"

type RandomNumberGenerator interface {
	GetRandomNumber(maxNum *big.Int) (int64, error)
}
