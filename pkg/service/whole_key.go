package service

import (
	"fmt"
	"sync"

	"changeme/pkg/etcd"
)

var (
	wholeKeySvc  *WholeKeyService
	wholeKeyOnce sync.Once
)

type WholeKeyService struct {
	etcdCli *etcd.EtcdClient
}

func GetWholeKeyService() *WholeKeyService {
	wholeKeyOnce.Do(func() {
		wholeKeySvc = &WholeKeyService{
			etcdCli: etcd.MustGetEtcdClient(),
		}
	})
	return wholeKeySvc
}

func (w *WholeKeyService) Get(key string) (string, error) {
	kv, err := w.etcdCli.Get(key)
	if err != nil || kv == nil {
		return "", fmt.Errorf("get key %s failed: %v", key, err)
	}
	return kv.Value, nil
}
