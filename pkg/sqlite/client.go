package sqlite

import (
	"database/sql"
	"log"
)

func Init() error {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return err
	}
	defer db.Close()

	if err := initOperatorTable(db); err != nil {
		return err
	}

	return nil
}

func initOperatorTable(db *sql.DB) error {
	// 创建表
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS action_records (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
			keyType TEXT,	
			action TEXT,
			message TEXT,
			createAt TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
