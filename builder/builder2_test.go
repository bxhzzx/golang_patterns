package builder

import "testing"

func TestBuilder2(t *testing.T) {
	assembly := NewBuilder2().Paint(Red)

	familyCar := assembly.Wheels(SportWheels).TopSpeed(50 * MPH).Build()
	familyCar.Drive()

	sportsCar := assembly.Wheels(SteelWheels).TopSpeed(150 * MPH).Build()
	sportsCar.Drive()
}
