package model

type Driver string

const (
	DriverBridge Driver = "bridge"
	DriverHost   Driver = "host"
	DriverNone   Driver = "none"
)

type NetworkItem struct {
	Name   string `json:"name"`
	Scope  string `json:"scope"`
	Driver Driver `json:"driver"`
}
