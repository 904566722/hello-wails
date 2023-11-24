package sqlite

import (
	"sync"

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
		select id, jsonFormat, etcdEndPoint from global_config
		where id = ?;
	`

	rows, err := db.Query(querySQL, id)
	if err != nil {
		log.Errorf("query global_config failed: [%v]", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		gc := &models.GlobalConfig{}
		err = rows.Scan(&gc.Id, &gc.JsonFormat, &gc.EtcdEndPoint)
		if err != nil {
			log.Errorf("scan global_config failed: [%v]", err)
			return nil, err
		}
		return gc, nil
	}

	return nil, ErrValueNotFound
}

func (g *GlobalConfigDb) Insert(gc *models.GlobalConfig) error {
	// 插入一条数据
	insertSQL := `
		INSERT INTO global_config (jsonFormat, etcdEndPoint)
		VALUES (?, ?);
	`

	_, err := db.Exec(insertSQL, gc.JsonFormat, gc.EtcdEndPoint)
	if err != nil {
		return err
	}
	return nil
}

func (g *GlobalConfigDb) Update(gc *models.GlobalConfig) error {
	// 更新一条数据
	updateSQL := `
		UPDATE global_config SET jsonFormat = ?, etcdEndPoint = ?
		WHERE id = ?;
	`

	_, err := db.Exec(updateSQL, gc.JsonFormat, gc.EtcdEndPoint, gc.Id)
	if err != nil {
		log.Errorf("update global_config failed: [%v]", err)
		return err
	}
	log.InfoWithFields(map[string]interface{}{
		"jsonFormat": gc.JsonFormat,
	}, "update global_config success")
	return nil
}
