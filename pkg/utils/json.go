package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	jq = "jq"
)

func JQ(s string) (string, error) {
	return ExecCmdRetOut("sh", "-c", fmt.Sprintf("echo '%s' | %s", s, jq))
}

func Compact(s string) (string, error) {
	buffer := bytes.Buffer{}
	err := json.Compact(&buffer, []byte(s))
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

// IsJsonFormat 判断是否为 json 格式
func IsJsonFormat(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
