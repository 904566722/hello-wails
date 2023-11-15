package etcd

import (
	"context"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	endpoint      = "localhost:2379"
	dialTimeout   = 5 * time.Second
	etcdOpTimeout = 10 * time.Second
)

var (
	_etcdCli     *clientv3.Client
	_etcdCliOnce sync.Once
)

type EtcdClient struct {
	cli *clientv3.Client
}

func NewEtcdClient() (*EtcdClient, error) {
	client, err := GetClient()
	if err != nil {
		return nil, err
	}
	return &EtcdClient{
		cli: client,
	}, nil
}

func GetClient() (*clientv3.Client, error) {
	if _etcdCli == nil {
		return clientv3.New(clientv3.Config{
			Endpoints:   []string{endpoint},
			DialTimeout: dialTimeout,
		})
	}
	return _etcdCli, nil
}

func MustGetClient() *clientv3.Client {
	return _etcdCli
}

func InitEtcd() error {
	_etcdCliOnce.Do(func() {
		client, err := clientv3.New(clientv3.Config{
			Endpoints:   []string{endpoint},
			DialTimeout: dialTimeout,
		})
		if err != nil {
			return
		}
		_etcdCli = client
	})
	return nil
}

func etcdOpCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.TODO(), etcdOpTimeout)
}
