package system

import (
	"xompose/internal/service"
)

type sSystem struct{}

func init() {
	service.RegisterSystem(New())
}

func New() *sSystem {
	return &sSystem{}
}
