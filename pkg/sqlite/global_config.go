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
		SELECT id, jsonFormat FROM global_config
		WHERE id = ?;
	`

	rows, err := db.Query(querySQL, id)
	if err != nil {
		log.Log.Errorf("query global_config failed: [%v]", err)
		return nil, err
	}
	defer rows.Close()

	// 如果未找到该条数据，返回 nil
	if !rows.Next() {
		return nil, ErrValueNotFound
	}

	// 解析该条数据
	gc := &models.GlobalConfig{}
	for rows.Next() {
		err = rows.Scan(&gc.Id, &gc.JsonFormat)
		if err != nil {
			log.Log.Errorf("scan global_config failed: [%v]", err)
			return nil, err
		}
	}

	return gc, nil
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
	return nil
}
