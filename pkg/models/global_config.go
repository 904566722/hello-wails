package models

// GlobalConfig 保存应用的全局配置
type GlobalConfig struct {
	Id         int  `json:"id"`
	JsonFormat bool `json:"jsonFormat"` // value 是否以 json 格式返回
}
