package etcd

import (
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"

	"changeme/pkg/log"
)

func (e *EtcdClient) Del(key string) error {
	ctx, cancel := e.etcdOpCtx()
	defer cancel()
	if _, err := e.cli.Delete(ctx, key); err != nil {
		return err
	}
	return nil
}

// DelByPrefix delete all keys with the prefix
// return the number of keys deleted
func (e *EtcdClient) DelByPrefix(pfx string) (int64, error) {
	ctx, cancel := e.etcdOpCtx()
	defer cancel()

	resp, err := e.cli.Delete(ctx, pfx, clientv3.WithPrefix())
	if err != nil {
		return -1, err
	}

	log.InfoWithFields(map[string]interface{}{
		"prefix": pfx,
		"number": resp.Deleted,
	}, "delete by prefix key")
	return resp.Deleted, nil
}

func (e *EtcdClient) DelByKeyword(kw string) (int64, error) {
	kwKeys, err := e.ListKeyByKeyword(kw)
	if err != nil {
		log.Errorf("list key by keyword failed: [%v]", err)
		return -1, err
	}
	if len(kwKeys) == 0 {
		return 0, nil
	}
	ops := make([]clientv3.Op, 0, len(kwKeys))
	for _, key := range kwKeys {
		ops = append(ops, clientv3.OpDelete(key))
	}
	ctx, cancel := e.etcdBatchOpCtx()
	defer cancel()
	txn := e.cli.Txn(ctx)
	resp, err := txn.Then(ops...).Commit()
	if err != nil {
		log.Errorf("batch delete by keyword failed: [%v]", err)
		return -1, err
	}
	if !resp.Succeeded {
		return -1, fmt.Errorf("batch delete by keyword failed: [%v]", resp)
	}
	return int64(len(kwKeys)), nil
}
