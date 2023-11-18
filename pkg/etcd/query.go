package etcd

import (
	"fmt"
	"strings"

	clientv3 "go.etcd.io/etcd/client/v3"

	"changeme/pkg/utils"
)

func (e *EtcdClient) ListKeyByPrefix(prefix string) (string, error) {
	ctx, cancel := e.etcdOpCtx()
	defer cancel()
	resp, err := e.cli.Get(ctx, prefix, clientv3.WithPrefix(), clientv3.WithKeysOnly())
	if err != nil {
		return "", err
	}
	var list []string
	for _, kv := range resp.Kvs {
		list = append(list, string(kv.Key))
	}
	return listRet(list), nil
}

func (e *EtcdClient) ListKeyByKeyword(kw string) (string, error) {
	out, err := utils.ExecCmdRetOut("sh", "-c", fmt.Sprintf("etcdctl --prefix --keys-only get '' | grep -i %s", kw))
	if err != nil {
		return "", err
	}
	return out, nil
}

func (e *EtcdClient) ListValueByKeyword(kw string) (string, error) {
	out, err := utils.ExecCmdRetOut("sh", "-c", fmt.Sprintf("etcdctl --prefix --keys-only get '' | grep -i %s | xargs -I {} etcdctl --prefix --print-value-only get {} | jq", kw))
	if err != nil {
		return "", err
	}
	return out, nil
}

func listRet(list []string) string {
	return strings.Join(list, "\n")
}
