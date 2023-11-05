package smartphone

import . "awesomeProject/Lecture02/task3/devices/device"

type Smartphone struct {
	DeviceProperties
}

func (s Smartphone) TurnOn() string {
	return "Smartphone is now turned on."
}

func (s Smartphone) TurnOff() string {
	return "Smartphone is now turned off."
}
