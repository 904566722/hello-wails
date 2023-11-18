package models

import (
	"strings"
	"time"
)

type OpResult uint

const (
	OpSuccess OpResult = iota
	OpFail
)

type Operation struct {
	Id       int       `json:"id"`
	KeyType  string    `json:"keyType"`
	Action   string    `json:"action"`
	Result   OpResult  `json:"result"`
	Message  string    `json:"message"`
	CreateAt time.Time `json:"createAt"`
}

func NewOperationSuccess(keyType string, action string, message ...string) *Operation {
	msg := "nil"
	if len(message) != 0 {
		msg = strings.Join(message, " ")
	}
	return &Operation{
		KeyType:  keyType,
		Action:   action,
		Result:   OpSuccess,
		Message:  msg,
		CreateAt: time.Now(),
	}
}

func NewOperationFailed(keyType string, action string, message ...string) *Operation {
	msg := "nil"
	if len(message) != 0 {
		msg = strings.Join(message, " ")
	}
	return &Operation{
		KeyType:  keyType,
		Action:   action,
		Result:   OpFail,
		Message:  msg,
		CreateAt: time.Now(),
	}
}
