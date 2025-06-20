package stack

import (
	"xompose/internal/service"
)

type sStack struct{}

func init() {
	service.RegisterStack(New())
}

func New() *sStack {
	return &sStack{}
}
