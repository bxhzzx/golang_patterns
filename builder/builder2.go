package builder

import "fmt"

type Speed float64

const (
	MPH Speed = 1
)

type Color string

const (
	Blue  Color = "blue"
	Green       = "green"
	Red         = "red"
)

type Wheels string

const (
	SportWheels Wheels = "sport"
	SteelWheels        = "steel"
)

type Builder2 interface {
	Paint(Color) Builder2 // 返回Builder2，使得可以支持链式调用
	TopSpeed(Speed) Builder2
	Wheels(Wheels) Builder2
	Build() AbsCar
}

type AbsCar interface {
	Drive() error
	Stop() error
}

// implements
func NewBuilder2() Builder2 {
	return &CarBuilder2{}
}

type Car2 struct {
	MyColor  Color
	MyWheels Wheels
	MySpeed  Speed
}

func (c Car2) Drive() error {
	fmt.Printf("Drive: my color is %s, my wheels is %s, my speed is %f\n", string(c.MyColor), string(c.MyWheels), float64(c.MySpeed))
	return nil
}

func (c Car2) Stop() error {
	fmt.Println("Stop")
	return nil
}

type CarBuilder2 struct {
	Car Car2
}

func (c *CarBuilder2) Paint(color Color) Builder2 {
	c.Car.MyColor = color
	return c
}

func (c *CarBuilder2) TopSpeed(speed Speed) Builder2 {
	c.Car.MySpeed = speed
	return c
}

func (c *CarBuilder2) Wheels(wheels Wheels) Builder2 {
	c.Car.MyWheels = wheels
	return c
}

func (c CarBuilder2) Build() AbsCar {
	return c.Car
}
