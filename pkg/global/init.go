package global

import (
	"changeme/pkg/etcd"
	"changeme/pkg/log"
	"changeme/pkg/sqlite"
)

func Init() error {
	if err := log.Init(); err != nil {
		return err
	}
	if err := etcd.Init(); err != nil {
		return err
	}
	if err := sqlite.Init(); err != nil {
		return err
	}

	log.Log.Info("global init success")
	return nil
}
