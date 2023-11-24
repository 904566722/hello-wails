package etcd

import (
	"context"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	endpoint            = "localhost:2379"
	dialTimeout         = 5 * time.Second
	etcdOpTimeout       = 10 * time.Second
	etcdBatchOpTimeout  = 1 * time.Minute
	etcdTestConnTimeout = 4 * time.Second
)

var (
	_etcdClient     *EtcdClient
	_etcdClientOnce sync.Once
	_etcdCli        *clientv3.Client
	_etcdCliOnce    sync.Once
)

func Init() error {
	_, err := NewEtcdClient()
	if err != nil {
		return err
	}
	return nil
}

type EtcdClient struct {
	cli *clientv3.Client
}

func NewEtcdClient() (*EtcdClient, error) {
	_etcdClientOnce.Do(func() {
		client, err := GetClient()
		if err != nil {
			panic(err)
		}
		_etcdClient = &EtcdClient{
			cli: client,
		}
	})
	return _etcdClient, nil
}

func MustGetEtcdClient() *EtcdClient {
	return _etcdClient
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

func (e *EtcdClient) etcdOpCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.TODO(), etcdOpTimeout)
}

func (e *EtcdClient) etcdBatchOpCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.TODO(), etcdBatchOpTimeout)
}

func (e *EtcdClient) testConnCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.TODO(), etcdTestConnTimeout)
}

// TestConnecting 测试连接
func (e *EtcdClient) TestConnecting(endpoint string) error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{endpoint},
		DialTimeout: etcdTestConnTimeout,
	})
	if err != nil {
		return err
	}
	defer cli.Close()
	ctx, cancel := e.testConnCtx()
	defer cancel()
	_, err = cli.Get(ctx, "test")
	if err != nil {
		return err
	}
	return nil
}

func (e *EtcdClient) ChangeEndpoint(endpoint string) {
	e.cli.SetEndpoints(endpoint)
}

func (e *EtcdClient) EndPoints() []string {
	return e.cli.Endpoints()
}
