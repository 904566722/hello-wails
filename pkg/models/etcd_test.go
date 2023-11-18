package models

import (
	"fmt"
	"testing"
)

func TestKeyVal_String(t *testing.T) {
	kv := KeyVal{
		Key:   "hci/user/admin",
		Value: "{\"name\":\"admin\",\"password\":\"admin\",\"role\":\"admin\"}",
	}
	fmt.Println(kv.String())
}
