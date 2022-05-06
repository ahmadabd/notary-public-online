package gorm

import (
	"fmt"
	"log"
	"notary-public-online/internal/configs/yaml"
	"notary-public-online/internal/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Gorm struct {
	Db *gorm.DB
}

func SetupDatabase(cfg *yaml.Config) (repository.DB, error) {
	db, err := gorm.Open(mysql.Open(databaseConfig(cfg)), &gorm.Config{})

	// db.AutoMigrate()

	if err != nil {
		log.Println("error while connecting to database: ", err)
		return nil, err
	}

	return &Gorm{Db: db}, nil
}

func databaseConfig(cfg *yaml.Config) string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name)
}
