package database

import (
	"capstone-project/helper"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	DB *sql.DB
}

func GetDatabaseEnv() string {
	user := helper.GetEnv("DATABASE_USER")
	password := helper.GetEnv("DATABASE_PASSWORD")
	host := helper.GetEnv("DATABASE_HOST")
	port := helper.GetEnv("DATABASE_PORT")
	name := helper.GetEnv("DATABASE_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, name)
}

func NewDBConnection() (*Database, error) {
	db, err := sql.Open("mysql", GetDatabaseEnv())
	if err != nil {
		return nil, err
	}
	return &Database{DB: db}, nil
}

func (d *Database) Close() error {
	return d.DB.Close()
}

func (d *Database) Reset(tables []string) error {
	for _, table := range tables {
		_, err := d.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table))
		if err != nil {
			return err
		}
		_, err = d.DB.Exec(fmt.Sprintf("ALTER TABLE %s AUTO_INCREMENT = 1", table))
		if err != nil {
			return err
		}
	}
	return nil
}
