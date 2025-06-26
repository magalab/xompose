package model

import "xompose/internal/model/entity"

type StackModel struct {
	entity.Stack
}

type StackItem struct {
	Id          int         `json:"id"`
	StackName   string      `json:"stackName"`
	YamlPath    string      `json:"yamlPath"`
	YamlContent string      `json:"yamlContent"`
	EnvContent  string      `json:"envContent"`
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
		Id:          m.Id,
		StackName:   m.StackName,
		StackStatus: StackStatus(m.StackStatus),
		YamlPath:    m.YamlContent,
		YamlContent: m.YamlContent,
		EnvContent:  m.EnvContent,
	}
}

type StackDeleteReq struct {
	StackName string `json:"stackName"`
}

type StackStatusItem struct {
	Image   string      `json:"image"`
	Service string      `json:"service"`
	State   StackStatus `json:"state"`
	Ports   []*Port     `json:"ports"`
}

type Port struct {
	PublishPort int    `json:"publishPort"`
	TargetPort  int    `json:"targetPort"`
	Protocol    string `json:"protocol"`
}
