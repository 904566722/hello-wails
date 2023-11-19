package etcd

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"

	"changeme/pkg/log"
	"changeme/pkg/utils"
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

	log.Log.WithFields(logrus.Fields{
		"prefix": pfx,
		"number": resp.Deleted,
	}).Info("delete by prefix key")
	return resp.Deleted, nil
}

func (e *EtcdClient) DelByKeyword(kw string) (int64, error) {
	out, err := utils.ExecCmdRetOut("sh", "-c", fmt.Sprintf("etcdctl --prefix --keys-only get '' | grep -i %s | wc -l", kw))
	if err != nil {
		return -1, err
	}
	n, err := strconv.Atoi(out)
	if err != nil {
		return -1, err
	}

	_, err = utils.ExecCmdRetOut("sh", "-c", fmt.Sprintf("etcdctl --prefix --keys-only get '' | grep -i %s | xargs -I {} etcdctl del {}", kw))
	if err != nil {
		return -1, err
	}
	log.Log.WithFields(logrus.Fields{
		"keyword": kw,
		"number":  n,
	}).Info("delete by keyword")
	return int64(n), nil
}
