package main

import (
	"awesomeProject/Lecture02/task3/devices/device"
	"awesomeProject/Lecture02/task3/devices/laptop"
	"awesomeProject/Lecture02/task3/devices/smartphone"
	"awesomeProject/Lecture02/task3/devices/tv"
	"fmt"
)

func PerformOperations(e device.Electronic) {
	fmt.Println(e.TurnOn())
	fmt.Println(e.TurnOff())
}

func main() {
	myLaptop := laptop.Laptop{device.DeviceProperties{Brand: "Dell", Model: "XPS"}}
	mySmartphone := smartphone.Smartphone{device.DeviceProperties{Brand: "Apple", Model: "iPhone"}}
	myTV := tv.TV{device.DeviceProperties{Brand: "Samsung", Model: "QLED"}}

	PerformOperations(myLaptop)
	PerformOperations(mySmartphone)
	PerformOperations(myTV)
}
