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
	key     string
	etcdCli *etcd.EtcdClient
}

func GetWholeKeyService(key string) *WholeKeyService {
	wholeKeyOnce.Do(func() {
		wholeKeySvc = &WholeKeyService{
			key:     key,
			etcdCli: etcd.MustGetEtcdClient(),
		}
	})
	return wholeKeySvc
}

func (w *WholeKeyService) Get() (string, error) {
	kv, err := w.etcdCli.Get(w.key)
	if err != nil || kv == nil {
		return "", fmt.Errorf("get key %s failed: %v", w.key, err)
	}
	return kv.Value, nil
}
