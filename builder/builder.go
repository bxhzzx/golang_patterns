package builder

type Builder interface {
	SetWheels() Builder // 返回Builder，使得可以支持链式调用
	SetSeats() Builder
	GetVehicle() Vehicle
}

type Vehicle struct {
	Wheels int
	Seats  int
}

type ManufacturingDirector struct {
	b Builder
}

func (m *ManufacturingDirector) SetBuilder(b Builder) {
	m.b = b
}

func (m *ManufacturingDirector) Construct() {
	m.b.SetWheels().SetSeats()
}

type CarBuilder struct {
	v Vehicle
}

func (c *CarBuilder) SetWheels() Builder {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() Builder {
	c.v.Seats = 4
	return c
}

func (c *CarBuilder) GetVehicle() Vehicle {
	return c.v
}

type BikeBuilder struct {
	v Vehicle
}

func (c *BikeBuilder) SetWheels() Builder {
	c.v.Wheels = 2
	return c
}

func (c *BikeBuilder) SetSeats() Builder {
	c.v.Seats = 2
	return c
}

func (c *BikeBuilder) GetVehicle() Vehicle {
	return c.v
}
