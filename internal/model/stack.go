package model

import "xompose/internal/model/entity"

type StackModel struct {
	entity.Stack
}

type StackItem struct {
	StackName   string      `json:"stackName"`
	StackStatus StackStatus `json:"stackStatus"`
}

type StackAddReq struct {
	StackName   string `json:"stackName"`
	YamlContent string `json:"yamlContent"`
	Envs        []map[string]any
}

type StackGetReq struct {
	StackName string `v:"required#请输入 stack 名称"`
}

type StackListReq struct {
	// PageReq
}

type StackStatus string

const (
	StackRunning StackStatus = "running"
	StackCreated StackStatus = "created"
)

type StackListItem struct {
	*StackItem
	IsManaged bool `json:"isManaged"`
}

func (m *StackModel) ToItem() *StackItem {
	return &StackItem{
		StackName:   m.StackName,
		StackStatus: StackStatus(m.StackStatus),
	}
}

type StackDeleteReq struct {
	StackName string `json:"stackName"`
}
