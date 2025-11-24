package random_number

import (
	"errors"
	"math/big"
)

type MockRandomGenerator struct {
	numbers []int64
	index   int
}

func InitMockRandomGenerator(numbers []int64) *MockRandomGenerator {
	return &MockRandomGenerator{numbers: numbers, index: 0}
}

func (m *MockRandomGenerator) GetRandomNumber(maxNum *big.Int) (int64, error) {
	num := m.numbers[m.index]
	if num > maxNum.Int64() {
		return -1, errors.New("maxNum 인자보다 큰 수가 numbers 필드에 들어가 있습니다")
	}
	m.index++
	return num, nil
}
