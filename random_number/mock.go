package random_number

import (
	"errors"
	"math/big"
)

type MockRandomGenerator struct {
	numbers []int
	index   int
}

func InitMockRandomGenerator(numbers []int) *MockRandomGenerator {
	return &MockRandomGenerator{numbers: numbers, index: 0}
}

func (m *MockRandomGenerator) GetRandomNumber(maxNum *big.Int) (int64, error) {
	num := m.numbers[m.index]
	if int64(num) > maxNum.Int64() {
		return -1, errors.New("maxNum 인자보다 큰 수를 입력하셨습니다")
	}
	m.index++
	return int64(num), nil
}
