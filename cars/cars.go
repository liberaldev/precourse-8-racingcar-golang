package cars

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"strings"
)

type Car struct {
	name  string
	steps int
}

type Cars struct {
	cars []Car
}

func validateCarName(name string) error {
	if len(strings.TrimSpace(name)) == 0 {
		return errors.New("이름이 비어있습니다")
	}
	if len(name) >= 5 {
		return errors.New("5자 이하로 이름을 입력하세요")
	}
	return nil
}

func (c *Cars) Init(names []string) error {
	for i := 0; i < len(names); i++ {
		if err := validateCarName(names[i]); err != nil {
			return err
		}
		c.cars = append(c.cars, Car{name: names[i], steps: 0})
	}
	return nil
}

func (c *Cars) MoveCarsByRandomNumber() {
	for i := range c.cars {
		if n, err := rand.Int(rand.Reader, big.NewInt(9)); err == nil && n.Int64() >= 4 {
			car := &c.cars[i]
			car.steps += 1
		}
	}
}

func (c *Cars) CarsStepPrint() {
	for i := range c.cars {
		car := &c.cars[i]
		fmt.Println(car.name+":", strings.Repeat("-", car.steps))
	}
}

func (c *Cars) Sort() {
	sort.Slice(c.cars, func(i, j int) bool { return c.cars[i].steps > c.cars[j].steps })
}
