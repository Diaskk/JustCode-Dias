package tv

import . "awesomeProject/Lecture02/task3/devices/device"

type TV struct {
	DeviceProperties
}

func (t TV) TurnOn() string {
	return "TV is now turned on."
}

func (t TV) TurnOff() string {
	return "TV is now turned off."
}
