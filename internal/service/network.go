// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"xompose/internal/model"
)

type (
	INetwork interface {
		NetworkList(ctx context.Context) ([]*model.NetworkItem, error)
	}
)

var (
	localNetwork INetwork
)

func Network() INetwork {
	if localNetwork == nil {
		panic("implement not found for interface INetwork, forgot register?")
	}
	return localNetwork
}

func RegisterNetwork(i INetwork) {
	localNetwork = i
}
