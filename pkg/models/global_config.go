package models

// GlobalConfig 保存应用的全局配置
type GlobalConfig struct {
	Id           int    `json:"id"`
	EtcdEndPoint string `json:"etcdEndPoint"` // etcd 的地址
	JsonFormat   bool   `json:"jsonFormat"`   // value 是否以 json 格式返回
}
