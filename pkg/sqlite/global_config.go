package sqlite

import (
	"sync"

	"github.com/sirupsen/logrus"

	"changeme/pkg/log"
	"changeme/pkg/models"
)

type GlobalConfigDb struct {
}

var (
	globalConfigDb     *GlobalConfigDb
	globalConfigDbOnce sync.Once
)

func GetGlobalConfigDb() *GlobalConfigDb {
	globalConfigDbOnce.Do(func() {
		globalConfigDb = &GlobalConfigDb{}
	})
	return globalConfigDb
}

func (g *GlobalConfigDb) Select(id int) (*models.GlobalConfig, error) {
	// 根据 id 查询一条数据
	querySQL := `
		select id, jsonFormat from global_config
		where id = ?;
	`

	rows, err := db.Query(querySQL, id)
	if err != nil {
		log.Log.Errorf("query global_config failed: [%v]", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		gc := &models.GlobalConfig{}
		err = rows.Scan(&gc.Id, &gc.JsonFormat)
		if err != nil {
			log.Log.Errorf("scan global_config failed: [%v]", err)
			return nil, err
		}
		return gc, nil
	}

	return nil, ErrValueNotFound
}

func (g *GlobalConfigDb) Insert(gc *models.GlobalConfig) error {
	// 插入一条数据
	insertSQL := `
		INSERT INTO global_config (jsonFormat)
		VALUES (?);
	`

	_, err := db.Exec(insertSQL, gc.JsonFormat)
	if err != nil {
		return err
	}
	return nil
}

func (g *GlobalConfigDb) Update(gc *models.GlobalConfig) error {
	// 更新一条数据
	updateSQL := `
		UPDATE global_config SET jsonFormat = ?
		WHERE id = ?;
	`

	_, err := db.Exec(updateSQL, gc.JsonFormat, gc.Id)
	if err != nil {
		log.Log.Errorf("update global_config failed: [%v]", err)
		return err
	}
	log.Log.WithFields(logrus.Fields{
		"jsonFormat": gc.JsonFormat,
	}).Info("update global_config success")
	return nil
}
