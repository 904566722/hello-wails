package sqlite

import (
	"sync"

	"changeme/pkg/log"
	"changeme/pkg/models"
)

const (
	TableOperators = "operators"
)

type OperatorDb struct {
	table string
}

var (
	operatorDb     *OperatorDb
	operatorDbOnce sync.Once
)

func GetOperatorDb() *OperatorDb {
	operatorDbOnce.Do(func() {
		operatorDb = &OperatorDb{
			TableOperators,
		}
	})
	return operatorDb
}

func (o *OperatorDb) Insert(op *models.Operation) error {
	// 插入一条数据
	insertSQL := `
		INSERT INTO operators (keyType, action, result, message, createAt)
		VALUES (?, ?, ?, ?, ?);
	`

	_, err := db.Exec(insertSQL, op.KeyType, op.Action, op.Result, op.Message, op.CreateAt)
	if err != nil {
		log.Log.Errorf("insert operator failed: [%v]", err)
		return err
	}
	log.Log.Debugf("insert operator success: [%v]", op)
	return nil
}

func (o *OperatorDb) List(limit int) ([]*models.Operation, error) {
	// 根据创建时间查询最新的 limit 条数据
	querySQL := `
		SELECT id, keyType, action, result, message, createAt FROM operators
		ORDER BY createAt DESC
		LIMIT ?;
	`

	rows, err := db.Query(querySQL, limit)
	if err != nil {
		log.Log.Errorf("query operator failed: [%v]", err)
		return nil, err
	}

	var ops []*models.Operation
	for rows.Next() {
		op := &models.Operation{}
		if err := rows.Scan(&op.Id, &op.KeyType, &op.Action, &op.Result, &op.Message, &op.CreateAt); err != nil {
			log.Log.Errorf("parse operator failed: [%v]", err)
			return nil, err
		}
		ops = append(ops, op)
	}
	if err := rows.Err(); err != nil {
		log.Log.Errorf("error happened when range rows: [%v]", err)
		return nil, err
	}
	return ops, nil
}
