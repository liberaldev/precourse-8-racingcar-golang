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
	Name  string
	Steps int
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
		c.cars = append(c.cars, Car{Name: names[i], Steps: 0})
	}
	return nil
}

func (c *Cars) MoveCarsByRandomNumber() {
	for i := range c.cars {
		if n, err := rand.Int(rand.Reader, big.NewInt(9)); err == nil && n.Int64() >= 4 {
			car := &c.cars[i]
			car.Steps += 1
		}
	}
}

func (c *Cars) CarsStepPrint() {
	for i := range c.cars {
		car := &c.cars[i]
		fmt.Println(car.Name+":", strings.Repeat("-", car.Steps))
	}
}

func (c *Cars) Sort() {
	sort.Slice(c.cars, func(i, j int) bool { return c.cars[i].Steps > c.cars[j].Steps })
}

func (c *Cars) GetCars() []Car {
	return c.cars
}

func (c *Cars) GetWinners() []string {
	carsData := c.GetCars()
	maxSteps := 0
	for _, car := range carsData {
		if car.Steps > maxSteps {
			maxSteps = car.Steps
		}
	}

	var winners []string
	for _, car := range carsData {
		if car.Steps == maxSteps {
			winners = append(winners, car.Name)
		}
	}
	return winners
}
