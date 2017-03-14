package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Db struct {
	*sql.DB
}

func InitDb(protocol string, dataSourceUrl string) (*Db, error) {
	/*
		dataSource := fmt.Scanf("%s//%s:%s@%s/%s",
			config.Protocol,
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.DbName
		)*/
	fmt.Println(dataSourceUrl)
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
