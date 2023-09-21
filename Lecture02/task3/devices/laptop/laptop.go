package laptop

import . "awesomeProject/Lecture02/task3/devices/device"

type Laptop struct {
	DeviceProperties
}

func (l Laptop) TurnOn() string {
	return "Laptop is now turned on."
}

func (l Laptop) TurnOff() string {

	return "Laptop is now turned off."
}
