package global

import (
	"changeme/pkg/etcd"
	"changeme/pkg/log"
	"changeme/pkg/models"
	"changeme/pkg/service"
	"changeme/pkg/sqlite"
)

var GlobalConfig models.GlobalConfig

func Init() error {
	if err := log.Init(); err != nil {
		return err
	}
	if err := etcd.Init(); err != nil {
		log.Log.Errorf("etcd init failed: %v", err)
		return err
	}
	if err := sqlite.Init(); err != nil {
		log.Log.Errorf("sqlite init failed: %v", err)
		return err
	}
	gcSvc := service.GetGlobalConfigService()
	if err := gcSvc.InitGlobalConfig(); err != nil {
		log.Log.Errorf("init global config failed: %v", err)
		return err
	}
	gc, err := gcSvc.GetConfig()
	if err != nil {
		log.Log.Errorf("get global config failed: %v", err)
		return err
	}
	GlobalConfig = *gc

	log.Log.Info("global init success")
	return nil
}
