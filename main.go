package main

import (
	"fmt"
	"github.com/KumoTV/mds/api/v1/controllers"
	"github.com/KumoTV/mds/api/v1/models"
	"github.com/KumoTV/mds/config"
	"log"
)

func main() {
	dbConfig := config.ReadConfig().DbConfig
	// Initializing Db
	db, err := models.InitDb(
		"postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.Name,
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	controllers.StartHttpServer(db)

}
