package cars

import (
	"crypto/rand"
	"math/big"
)

type Car struct {
	name  string
	steps int
}

type Cars struct {
	cars []Car
}

func (c *Cars) Init(names []string) {
	for i := 0; i < len(names); i++ {
		c.cars = append(c.cars, Car{name: names[i], steps: 0})
	}
}

func (c *Cars) MoveCarsByRandomNumber() {
	for _, car := range c.cars {
		if n, err := rand.Int(rand.Reader, big.NewInt(9)); err == nil && n.Int64() >= 4 {
			car.steps += 1
		}
	}
}
