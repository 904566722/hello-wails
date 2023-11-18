package models

import "time"

type Operation struct {
	KeyType  string    `json:"keyType"`
	Action   string    `json:"action"`
	Message  string    `json:"message"`
	CreateAt time.Time `json:"createAt"`
}
