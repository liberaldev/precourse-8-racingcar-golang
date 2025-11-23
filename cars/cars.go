package cars

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
