package etcd

import (
	"fmt"
	"strings"

	clientv3 "go.etcd.io/etcd/client/v3"

	"changeme/pkg/log"
	"changeme/pkg/models"
	"changeme/pkg/utils"
)

func (e *EtcdClient) ListKeyByPrefix(prefix string) ([]string, error) {
	ctx, cancel := e.etcdOpCtx()
	defer cancel()
	resp, err := e.cli.Get(ctx, prefix, clientv3.WithPrefix(), clientv3.WithKeysOnly())
	if err != nil {
		return nil, err
	}
	var list []string
	for _, kv := range resp.Kvs {
		list = append(list, string(kv.Key))
	}
	return list, nil
}

func (e *EtcdClient) ListKeyByKeyword(kw string) ([]string, error) {
	ctx, cancel := e.etcdOpCtx()
	defer cancel()
	resp, err := e.cli.Get(ctx, "", clientv3.WithPrefix(), clientv3.WithKeysOnly())
	if err != nil {
		log.Errorf("list all key failed: [%v]", err)
		return nil, err
	}
	var list []string
	for _, kv := range resp.Kvs {
		if strings.Contains(string(kv.Key), kw) {
			list = append(list, string(kv.Key))
		}
	}
	return list, nil
}

func (e *EtcdClient) listAllKey() ([]string, error) {
	return nil, nil
}

func (e *EtcdClient) listKeyByKwBySh(kw string) (string, error) {
	return utils.ExecCmdRetOut("sh", "-c", fmt.Sprintf("etcdctl --prefix --keys-only get '' | grep -i %s", kw))
}

func (e *EtcdClient) ListValueByPrefix(prefix string) ([]string, error) {
	ctx, cancel := e.etcdOpCtx()
	defer cancel()
	resp, err := e.cli.Get(ctx, prefix, clientv3.WithPrefix(), clientv3.WithKeysOnly())
	if err != nil {
		return nil, err
	}
	var list []string
	for _, kv := range resp.Kvs {
		list = append(list, string(kv.Value))
	}
	return list, nil
}

func (e *EtcdClient) ListValueByKeyword(kw string) ([]string, error) {
	kwKeys, err := e.ListKeyByKeyword(kw)
	if err != nil {
		log.Errorf("list key by keyword failed: [%v]", err)
		return nil, err
	}
	var res []string
	for _, key := range kwKeys {
		value, err := e.Get(key)
		if err != nil {
			continue
		}
		res = append(res, value.Value)
	}
	return res, nil
}

func (e *EtcdClient) ListKeyValueByPrefix(pfx string) ([]models.KeyVal, error) {
	ctx, cancel := e.etcdOpCtx()
	defer cancel()
	resp, err := e.cli.Get(ctx, pfx, clientv3.WithPrefix(), clientv3.WithKeysOnly())
	if err != nil {
		return nil, err
	}
	var list []models.KeyVal
	for _, kv := range resp.Kvs {
		list = append(list, models.KeyVal{
			Key:   string(kv.Key),
			Value: string(kv.Value),
		})
	}
	return list, nil
}

func splitByLine(out string) []string {
	return strings.Split(strings.TrimSpace(out), "\n")
}
