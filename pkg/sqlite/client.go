package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"changeme/pkg/log"
)

var db *sql.DB

func Init() error {
	var err error
	db, err = sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return err
	}
	defer db.Close()

	if err := initOperatorTable(db); err != nil {
		log.Log.Errorf("init operator table failed: %v", err)
		return err
	}

	if err := initGlobalConfigTable(db); err != nil {
		log.Log.Errorf("init global config table failed: %v", err)
		return err
	}

	if err := open(); err != nil {
		log.Log.Errorf("open sqlite failed: %v", err)
		return err
	}

	return nil
}

func initOperatorTable(db *sql.DB) error {
	// 创建表
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS operators (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
			keyType TEXT,	
			action TEXT,
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
	db, err = sql.Open("sqlite3", "./foo.db")
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
