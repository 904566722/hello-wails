package models

import (
	"fmt"
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
	KeyType  KeyType   `json:"keyType"`
	Action   Action    `json:"action"`
	Key      string    `json:"key"`
	Value    string    `json:"value"`
	Result   OpResult  `json:"result"`
	Message  string    `json:"message"`
	Desc     string    `json:"desc"`
	CreateAt time.Time `json:"createAt"`
}

func (o *Operation) fillDesc() {
	res := "失败"
	if o.Result == OpSuccess {
		res = "成功"
	}
	o.Desc = fmt.Sprintf("根据「%s」(%s)-> %s%s", o.KeyType.Chinese(), o.Key, o.Action.Chinese(), res)
}

func (o *Operation) String() string {
	return fmt.Sprintf("Operation{Id: %d, KeyType: %s, Action: %s, Key: %s, Value: %s, Result: %d, Message: %s, Desc: %s, CreateAt: %s}",
		o.Id, o.KeyType, o.Action, o.Key, o.Value, o.Result, o.Message, o.Desc, o.CreateAt)
}

func NewOperationSuccess(keyType KeyType, action Action, key, val string, message ...string) *Operation {
	msg := ""
	if len(message) != 0 {
		msg = strings.Join(message, " ")
	}

	op := &Operation{
		KeyType:  keyType,
		Action:   action,
		Key:      key,
		Value:    val,
		Result:   OpSuccess,
		Message:  msg,
		CreateAt: time.Now(),
	}
	op.fillDesc()
	return op
}

func NewOperationFailed(keyType KeyType, action Action, key, val string, message ...string) *Operation {
	msg := ""
	if len(message) != 0 {
		msg = strings.Join(message, " ")
	}

	op := &Operation{
		KeyType:  keyType,
		Action:   action,
		Key:      key,
		Value:    val,
		Result:   OpFail,
		Message:  msg,
		CreateAt: time.Now(),
	}
	op.fillDesc()
	return op
}
