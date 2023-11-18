package api

import (
	"fmt"

	"changeme/pkg/service"
)

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

// return: [是否允许操作，不允许操作的原因]
func checkActionAllow(kType KeyType, action Action) (bool, string) {
	allowAction, exist := dTypeActionMap[kType]
	if !exist {
		return false, fmt.Sprintf("key type %s not exist", kType)
	}
	_, exist = allowAction[action]
	if !exist {
		return false, fmt.Sprintf("key type %s does not allow operation %s", kType, action)
	}
	return true, ""
}

func DoAction(req BaseRequest) (BaseResponse, error) {
	allow, reason := checkActionAllow(req.KeyType, req.Action)
	if !allow {
		return ErrResp(CodeActionNotAllowed, reason), nil
	}

	switch req.KeyType {
	case DTypeWholeKey:
		switch req.Action {
		case ActionGet:
			val, err := service.GetWholeKeyService(req.Data).Get()
			if err != nil {
				return ErrResp(CodeFail, err.Error()), nil
			}
			return SuccessResp(val), nil
		default:
			return ErrResp(CodeFail, fmt.Sprintf("key type %s not support action %s", req.KeyType, req.Action)), nil
		}
	default:
		return ErrResp(CodeFail, fmt.Sprintf("key type %s not support action %s", req.KeyType, req.Action)), nil
	}
}
