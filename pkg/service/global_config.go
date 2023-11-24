package service

import (
	"errors"
	"strings"
	"sync"

	"changeme/pkg/log"
	"changeme/pkg/models"
	"changeme/pkg/sqlite"
)

const (
	gcId = 1
)

var GlobalConfig models.GlobalConfig

type GlobalConfigService struct {
	gcDb       *sqlite.GlobalConfigDb
	etcdCliSvc *EtcdClientService
}

var (
	globalConfigSvc     *GlobalConfigService
	globalConfigSvcOnce sync.Once
)

func GetGlobalConfigService() *GlobalConfigService {
	globalConfigSvcOnce.Do(func() {
		globalConfigSvc = &GlobalConfigService{
			gcDb:       sqlite.GetGlobalConfigDb(),
			etcdCliSvc: GetEtcdClientService(),
		}
	})
	return globalConfigSvc
}

func (g *GlobalConfigService) InitGlobalConfig() error {
	_, err := g.gcDb.Select(gcId)
	if err != nil {
		if errors.Is(err, sqlite.ErrValueNotFound) {
			// 如果不存在，则初始化一条数据
			// todo
			gc := &models.GlobalConfig{
				JsonFormat:   false,
				EtcdEndPoint: "localhost:2379",
			}
			if err := g.gcDb.Insert(gc); err != nil {
				log.Errorf("insert global_config failed: [%v]", err)
				return err
			}
			log.Info("init global_config success")
			return nil
		} else {
			log.Errorf("select global_config failed: [%v]", err)
			return err
		}
	}
	// 已经存在
	log.Debugf("global_config already exist")
	return nil
}

func (g *GlobalConfigService) GetConfig() (*models.GlobalConfig, error) {
	gc, err := g.gcDb.Select(gcId)
	if err != nil {
		return nil, err
	}
	return gc, nil
}

func (g *GlobalConfigService) Update(gc *models.GlobalConfig) error {
	gc.Id = gcId
	if err := g.gcDb.Update(gc); err != nil {
		return err
	}
	return nil
}

func (g *GlobalConfigService) Sync() error {
	gc, err := g.gcDb.Select(gcId)
	if err != nil {
		return err
	}
	GlobalConfig = *gc

	oldEps := strings.Join(g.etcdCliSvc.GetEndPoints(), ";")
	if !strings.Contains(oldEps, GlobalConfig.EtcdEndPoint) {
		if err := g.etcdCliSvc.ChangeEp(gc.EtcdEndPoint); err != nil {
			return err
		} else {
			log.Infof("sync global config success, etcd endpoint change from [%s] to [%s]", oldEps, GlobalConfig)
		}
	}

	return nil
}
