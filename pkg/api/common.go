package api

import (
	"runtime/debug"

	"changeme/pkg/global"
	"changeme/pkg/log"
	"changeme/pkg/utils"
)

type BaseRequest struct {
	Data    string  `json:"data"`
	KeyType KeyType `json:"keyType"`
	Action  Action  `json:"action"`
	Options Options `json:"options"`
}

type Options struct {
	Json bool `json:"json"`
}

func (req BaseRequest) ParseOpType() (KeyType, Action) {
	return KeyType(req.KeyType), Action(req.Action)
}

type ListRequest struct {
	Limit int `json:"limit"`
}

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Total   int         `json:"total"`
}

const (
	CodeSuccess = 200
	CodeFail    = 500 + iota
	CodeActionNotAllowed
)

const (
	MsgSuccess = "success"
)

func RespSuccess(data interface{}) BaseResponse {
	// 是否格式化数据
	var (
		finalVal interface{}
	)

	switch val := data.(type) {
	case string:
		value, err := FormatValue(val)
		if err != nil {
			finalVal = err.Error()
		} else {
			finalVal = value
		}
	case []string:
		values, err := FormatValues(val)
		if err != nil {
			finalVal = err.Error()
		} else {
			finalVal = values
		}
	}

	return BaseResponse{
		Code:    CodeSuccess,
		Message: MsgSuccess,
		Data:    finalVal,
	}
}

func RespErr(code int, msg string) BaseResponse {
	return BaseResponse{
		Code:    code,
		Message: msg,
	}
}

type KeyType uint

const (
	DTypeWholeKey KeyType = iota // 完整的 etcd key
	DTypePrefix                  // etcd key 的前缀
	DTypeKeyword                 // etcd key 的关键字
)

func (d KeyType) String() string {
	switch d {
	case DTypeWholeKey:
		return "DTypeWholeKey"
	case DTypePrefix:
		return "DTypePrefix"
	case DTypeKeyword:
		return "DTypeKeyword"
	default:
		return "Unknown"
	}
}

func Str2KeyType(s string) KeyType {
	switch s {
	case "DTypeWholeKey":
		return DTypeWholeKey
	case "DTypePrefix":
		return DTypePrefix
	case "DTypeKeyword":
		return DTypeKeyword
	default:
		return DTypeKeyword
	}
}

func Str2Action(s string) Action {
	switch s {
	case "ActionGet":
		return ActionGet
	case "ActionPut":
		return ActionPut
	case "ActionDelete":
		return ActionDelete
	case "ActionListKey":
		return ActionListKey
	case "ActionListValue":
		return ActionListValue
	case "ActionList":
		return ActionList
	default:
		return ActionList
	}
}

type Action uint

const (
	ActionGet Action = iota
	ActionPut
	ActionDelete
	ActionListKey   // 获取 key 列表
	ActionListValue // 获取 value 列表
	ActionList      // 获取 key 和 value 列表
)

func (a Action) String() string {
	switch a {
	case ActionGet:
		return "ActionGet"
	case ActionPut:
		return "ActionPut"
	case ActionDelete:
		return "ActionDelete"
	case ActionList:
		return "ActionList"
	default:
		return "Unknown"
	}
}

// dTypeActionMap 定义 key 类型允许的操作
var dTypeActionMap = map[KeyType]map[Action]struct{}{
	// 允许通过完整的 key 进行：1.get 操作；2.put 操作；3.delete 操作；4.list 操作；5.list value 操作
	DTypeWholeKey: {
		ActionGet:       struct{}{}, // 允许通过完整的 key 进行 get 操作
		ActionPut:       struct{}{}, // 允许通过完整的 key 进行 put 操作
		ActionDelete:    struct{}{}, // 允许通过完整的 key 进行 delete 操作
		ActionList:      struct{}{}, // 允许通过完整的 key 进行 list 操作
		ActionListValue: struct{}{}, // 允许通过完整的 key 进行 list value 操作
	},
	// 允许通过前缀进行：1.删除所有key对应的值；2.查询所有key；3.查询所有value；4.查询所有key和value
	DTypePrefix: {
		ActionDelete:    struct{}{}, // 允许通过前缀进行 delete 操作 (需要询问是否删除全部的key)
		ActionListKey:   struct{}{}, // 允许通过前缀进行查询所有 key 操作
		ActionListValue: struct{}{}, // 允许通过前缀进行查询所有 value 操作
		ActionList:      struct{}{}, // 允许通过前缀进行查询所有 key 和 value 操作
	},
	// 允许通过关键字进行：1.删除所有key对应的值；2.查询所有key；3.查询所有value；4.查询所有key和value
	DTypeKeyword: {
		ActionDelete:    struct{}{}, // 允许通过关键字进行 delete 操作 (需要询问是否删除全部的key)
		ActionListKey:   struct{}{}, // 允许通过关键字进行查询所有 key 操作
		ActionListValue: struct{}{}, // 允许通过关键字进行查询所有 value 操作
		ActionList:      struct{}{}, // 允许通过关键字进行查询所有 key 和 value 操作
	},
}

func recoverFromPanic() {
	if err := recover(); err != nil {
		log.Log.Errorf("panic: [%v]", err)
		debug.PrintStack()
	}
}

func FormatValue(val string) (string, error) {
	if global.GlobalConfig.JsonFormat {
		jsonVal, err := utils.JQ(val)
		if err != nil {
			return "", err
		}
		return jsonVal, nil
	}
	return val, nil
}

func FormatValues(vals []string) ([]string, error) {
	if !global.GlobalConfig.JsonFormat {
		return vals, nil
	}
	var ret []string
	for _, val := range vals {
		jsonVal, err := FormatValue(val)
		if err != nil {
			return nil, err
		}
		ret = append(ret, jsonVal)
	}
	return ret, nil
}
