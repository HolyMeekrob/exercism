package elon

import "fmt"

func (car *Car) IsOutOfJuice() bool {
	return car.batteryDrain > car.battery
}

func (car *Car) Drive() {
	if car.IsOutOfJuice() {
		return
	}

	car.distance = car.distance + car.speed
	car.battery = car.battery - car.batteryDrain
}

func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

func (car *Car) CanFinish(trackDistance int) bool {
	stretches := car.battery / car.batteryDrain
	carRange := car.speed * stretches
	return trackDistance <= carRange
}
