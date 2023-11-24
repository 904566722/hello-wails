package sqlite

import (
	"database/sql"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"

	"changeme/pkg/consts"
	"changeme/pkg/log"
)

var dbFile = filepath.Join(consts.AppFilePath, consts.DbName)

var db *sql.DB

func Init() error {
	var err error
	log.Infof("db file: %s", dbFile)
	db, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := initOperatorTable(db); err != nil {
		log.Errorf("init operator table failed: %v", err)
		return err
	}

	if err := initGlobalConfigTable(db); err != nil {
		log.Errorf("init global config table failed: %v", err)
		return err
	}

	if err := open(); err != nil {
		log.Errorf("open sqlite failed: %v", err)
		return err
	}

	return nil
}

func initOperatorTable(db *sql.DB) error {
	// 创建表
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS operators (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
			keyType INTEGER,	
			action INTEGER,
			key TEXT,
    		value TEXT,
			result INTEGER NOT NULL,
			message TEXT,
			desc TEXT,
			createAt TIMESTAMP
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

func initGlobalConfigTable(db *sql.DB) error {
	// 创建表
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS global_config (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    etcdEndPoint TEXT,
			jsonFormat INTEGER
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

func open() error {
	var err error
	db, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		return err
	}
	return nil
}

func close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
