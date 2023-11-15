package etcd

import (
	"changeme/pkg/models"
	"changeme/pkg/utils"
)

func (e *EtcdClient) Get(key string) (*models.KeyVal, error) {
	ctx, cancel := etcdOpCtx()
	defer cancel()
	resp, err := e.cli.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	for _, kv := range resp.Kvs {
		return &models.KeyVal{
			Key:   string(kv.Key),
			Value: string(kv.Value),
		}, nil
	}
	return nil, nil
}

func (e *EtcdClient) Get2String(key string) (string, error) {
	kv, err := e.Get(key)
	if err != nil {
		return "", err
	}
	jsonExpand, err := utils.JQ(kv.Value)
	if err != nil {
		return "", err
	}
	return jsonExpand, err
}
