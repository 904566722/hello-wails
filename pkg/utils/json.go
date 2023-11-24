package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"changeme/pkg/log"
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

// UnCompact 解除紧凑的 json 格式
func UnCompact(s string) (string, error) {
	if !IsCompactJson(s) {
		return s, nil
	}

	var (
		js  map[string]interface{}
		err error
	)
	if err = json.Unmarshal([]byte(s), &js); err != nil {
		return "", err
	}

	indent, err := json.MarshalIndent(js, "", " ")
	if err != nil {
		log.Errorf("marshal indent failed: [%v]", err)
		return "", err
	}
	return string(indent), nil
}

// IsJsonFormat 判断是否为 json 格式
func IsJsonFormat(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

// IsCompactJson 判断是否为紧凑的 json 格式
func IsCompactJson(s string) bool {
	return !strings.Contains(s, "\n")
}

type DiffStatistic struct {
	Diff map[string]Change      `json:"diff"` // key: field, value: change
	Del  map[string]interface{} `json:"del"`  // key: field, value: old value
	Add  map[string]interface{} `json:"add"`  // key: field, value: new value
}

type Change struct {
	Old interface{} `json:"old"`
	New interface{} `json:"new"`
}

// JsonDiff 比较两个 json 字符串的差异
func JsonDiff(oldJs, newJs string) (*DiffStatistic, error) {
	var (
		diffStatistic = &DiffStatistic{
			Diff: make(map[string]Change),
			Del:  make(map[string]interface{}),
			Add:  make(map[string]interface{}),
		}
		oldObj map[string]interface{}
		newObj map[string]interface{}
	)

	if err := json.Unmarshal([]byte(oldJs), &oldObj); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(newJs), &newObj); err != nil {
		return nil, err
	}

	for fieldName, oldVal := range oldObj {
		newVal, exist := newObj[fieldName]
		if !exist {
			diffStatistic.Del[fieldName] = oldVal
			continue
		}

		if !reflect.DeepEqual(oldVal, newVal) {
			diffStatistic.Diff[fieldName] = Change{
				Old: oldVal,
				New: newVal,
			}
		}
		delete(newObj, fieldName)
	}

	for fieldName, val := range newObj {
		diffStatistic.Add[fieldName] = val
	}

	return diffStatistic, nil
}
