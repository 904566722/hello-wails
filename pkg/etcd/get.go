package etcd

import (
	"errors"

	"changeme/pkg/models"
	"changeme/pkg/utils"
)

var (
	ErrValueNotFound = errors.New("value not found")
)

func (e *EtcdClient) Get(key string) (*models.KeyVal, error) {
	ctx, cancel := e.etcdOpCtx()
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
	return nil, ErrValueNotFound
}

func (e *EtcdClient) Get2String(key string) (string, error) {
	kv, err := e.Get(key)
	if err != nil || kv == nil {
		return "", err
	}
	if kv.Value == "" {
		return "", nil
	}
	jsonExpand, err := utils.JQ(kv.Value)
	if err != nil {
		return "", err
	}
	return jsonExpand, err
}
