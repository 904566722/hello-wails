package global

import (
	"os"

	"changeme/pkg/consts"
	"changeme/pkg/etcd"
	"changeme/pkg/log"
	"changeme/pkg/service"
	"changeme/pkg/sqlite"
)

func Init() error {
	// 如果目录不存在，递归创建该目录
	_, err := os.Stat(consts.AppFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(consts.AppFilePath, 0755)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	if err := log.Init(); err != nil {
		return err
	}
	if err := etcd.Init(); err != nil {
		log.Errorf("etcd init failed: %v", err)
		return err
	}
	if err := sqlite.Init(); err != nil {
		log.Errorf("sqlite init failed: %v", err)
		return err
	}
	gcSvc := service.GetGlobalConfigService()
	if err := gcSvc.InitGlobalConfig(); err != nil {
		log.Errorf("init global config failed: %v", err)
		return err
	}
	if err := gcSvc.Sync(); err != nil {
		log.Errorf("sync global config failed: %v", err)
		return err
	}

	log.Info("global init success")
	return nil
}
