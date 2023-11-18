package models

import "encoding/json"

type KeyVal struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (kv *KeyVal) String() string {
	b, _ := json.Marshal(kv)
	return string(b)
}
