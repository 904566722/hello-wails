package service

import (
	"sync"

	"changeme/pkg/etcd"
	"changeme/pkg/log"
)

var (
	prefixKeySvc  *PrefixKeyService
	prefixKeyOnce sync.Once
)

type PrefixKeyService struct {
	etcdCli *etcd.EtcdClient
}

func GetPrefixKeyService() *PrefixKeyService {
	prefixKeyOnce.Do(func() {
		prefixKeySvc = &PrefixKeyService{
			etcdCli: etcd.MustGetEtcdClient(),
		}
	})
	return prefixKeySvc
}

func (p *PrefixKeyService) Del(pfx string) (int64, error) {
	n, err := p.etcdCli.DelByPrefix(pfx)
	if err != nil {
		return 0, err
	}
	log.Log.Infof("delete by prefix key: [%s], del number: [%d]", pfx, n)
	return n, nil
}

func (p *PrefixKeyService) ListKey(pfx string) ([]string, error) {
	keys, err := p.etcdCli.ListKeyByPrefix(pfx)
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func (p *PrefixKeyService) ListValue(pfx string) ([]string, error) {
	values, err := p.etcdCli.ListValueByKeyword(pfx)
	if err != nil {
		return nil, err
	}
	return values, nil
}
