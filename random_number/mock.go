package random_number

type MockRandomGenerator struct {
	numbers []int
	index   int
}

func InitMockRandomGenerator(numbers []int) *MockRandomGenerator {
	return &MockRandomGenerator{numbers: numbers, index: 0}
}

func (m *MockRandomGenerator) GetRandomNumber() int64 {
	num := m.numbers[m.index]
	m.index++
	return int64(num)
}
