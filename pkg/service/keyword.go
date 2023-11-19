package service

import (
	"sync"

	"changeme/pkg/etcd"
)

var (
	keywordSvc  *KeywordService
	keywordOnce sync.Once
)

type KeywordService struct {
	etcdCli *etcd.EtcdClient
}

func GetKeywordService() *KeywordService {
	keywordOnce.Do(func() {
		keywordSvc = &KeywordService{
			etcdCli: etcd.MustGetEtcdClient(),
		}
	})
	return keywordSvc
}

func (k *KeywordService) Del(kw string) (int64, error) {
	return k.etcdCli.DelByKeyword(kw)
}

func (k *KeywordService) ListKey(kw string) ([]string, error) {
	keys, err := k.etcdCli.ListKeyByKeyword(kw)
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func (k *KeywordService) ListValue(kw string) ([]string, error) {
	values, err := k.etcdCli.ListValueByKeyword(kw)
	if err != nil {
		return nil, err
	}
	return values, nil
}

func (k *KeywordService) ListKeyValue(kw string) ([]string, error) {
	// todo
	return nil, nil
}
