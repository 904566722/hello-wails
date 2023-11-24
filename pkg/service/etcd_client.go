package service

import (
	"sync"

	"changeme/pkg/etcd"
)

var (
	etcdCliSvc     *EtcdClientService
	etcdCliSvcOnce sync.Once
)

type EtcdClientService struct {
	etcdCli *etcd.EtcdClient
}

func GetEtcdClientService() *EtcdClientService {
	etcdCliSvcOnce.Do(func() {
		etcdCliSvc = &EtcdClientService{
			etcdCli: etcd.MustGetEtcdClient(),
		}
	})
	return etcdCliSvc
}

func (e *EtcdClientService) TestConnecting(endpoint string) error {
	return e.etcdCli.TestConnecting(endpoint)
}

func (e *EtcdClientService) ChangeEp(endpoint string) error {
	e.etcdCli.ChangeEndpoint(endpoint)
	return nil
}

func (e *EtcdClientService) GetEndPoints() []string {
	return e.etcdCli.EndPoints()
}
