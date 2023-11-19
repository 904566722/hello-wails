package models

type GlobalConfigReq struct {
	JsonFormat bool `json:"jsonFormat"`
}

type BaseRequest struct {
	Key     string  `json:"key"`
	Value   string  `json:"value"`
	KeyType KeyType `json:"keyType"`
	Action  Action  `json:"action"`
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
	CodeUnSupportFormatType
)

const (
	MsgSuccess = "success"
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

func (d KeyType) Chinese() string {
	switch d {
	case DTypeWholeKey:
		return "完整的 key"
	case DTypePrefix:
		return "前缀"
	case DTypeKeyword:
		return "关键字"
	default:
		return "未知"
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
	case ActionListKey:
		return "ActionListKey"
	case ActionListValue:
		return "ActionListValue"
	default:
		return "Unknown"
	}
}

func (a Action) Chinese() string {
	switch a {
	case ActionGet:
		return "获取"
	case ActionPut:
		return "添加｜修改"
	case ActionDelete:
		return "删除"
	case ActionListKey:
		return "获取 key 列表"
	case ActionListValue:
		return "获取 value 列表"
	case ActionList:
		return "获取 key 和 value 列表"
	default:
		return "未知"
	}
}
