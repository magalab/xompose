package user

import (
	"xompose/internal/service"
)

type sNetwork struct{}

func init() {
	service.RegisterNetwork(New())
}

func New() *sNetwork {
	return &sNetwork{}
}
