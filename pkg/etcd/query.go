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
	out, err := utils.ExecCmdRetOut("sh", "-c", fmt.Sprintf("etcdctl --prefix --keys-only get '' | grep -i %s", kw))
	if err != nil {
		log.Log.Errorf("list key by keyword failed: [%v]", err)
		return nil, err
	}
	return splitByLine(out), nil
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
	out, err := utils.ExecCmdRetOut("sh", "-c", fmt.Sprintf("etcdctl --prefix --keys-only get '' | grep -i %s | xargs -I {} etcdctl --prefix --print-value-only get {} | jq", kw))
	if err != nil {
		return nil, err
	}
	return splitByLine(out), nil
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
