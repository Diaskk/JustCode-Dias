package device

type Electronic interface {
	TurnOn() string
	TurnOff() string
}

type DeviceProperties struct {
	Brand string
	Model string
}
