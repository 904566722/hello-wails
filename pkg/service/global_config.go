package service

import (
	"errors"
	"sync"

	"changeme/pkg/log"
	"changeme/pkg/models"
	"changeme/pkg/sqlite"
)

const (
	gcId = 1
)

type GlobalConfigService struct {
	gcDb *sqlite.GlobalConfigDb
}

var (
	globalConfigSvc     *GlobalConfigService
	globalConfigSvcOnce sync.Once
)

func GetGlobalConfigService() *GlobalConfigService {
	globalConfigSvcOnce.Do(func() {
		globalConfigSvc = &GlobalConfigService{
			gcDb: sqlite.GetGlobalConfigDb(),
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
				JsonFormat: false,
			}
			if err := g.gcDb.Insert(gc); err != nil {
				log.Log.Errorf("insert global_config failed: [%v]", err)
				return err
			}
			log.Log.Info("init global_config success")
		} else {
			log.Log.Errorf("select global_config failed: [%v]", err)
			return err
		}
	}
	// 已经存在
	log.Log.Debugf("global_config already exist")
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
	if err := g.gcDb.Update(gc); err != nil {
		return err
	}
	return nil
}
