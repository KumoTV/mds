package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Db struct {
	*sql.DB
}

func InitDb(protocol string, dataSourceUrl string) (*Db, error) {
	db, err := sql.Open(protocol, dataSourceUrl)
	if err != nil {
		return nil, err
	}
	/*
		if err = db.Ping(); err != nil {
			return nil, err
		}
	*/
	return &Db{db}, nil
}
