package lib

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var DB *gorm.DB

//
// ConnectDatabase to database
//
func ConnectDatabase() {
	connectionString := fmt.Sprintf(`host=%v port=%v dbname=%v user=%v password=%v sslmode=disable`,
		Config.Production.DBHost,
		Config.Production.DBPort,
		Config.Production.DBName,
		Config.Production.DBUser,
		Config.Production.DBPassword)

	if Config.Environment == "development" {
		connectionString = fmt.Sprintf(`host=%v port=%v dbname=%v user=%v password=%v sslmode=disable`,
			Config.Development.DBHost,
			Config.Development.DBPort,
			Config.Development.DBName,
			Config.Development.DBUser,
			Config.Development.DBPassword)

	} else if Config.Environment == "staging" {
		connectionString = fmt.Sprintf(`host=%v port=%v dbname=%v user=%v password=%v sslmode=disable`,
			Config.Staging.DBHost,
			Config.Staging.DBPort,
			Config.Staging.DBName,
			Config.Staging.DBUser,
			Config.Staging.DBPassword)
	}

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Println(connectionString)
		panic(err)
	}

	db.LogMode(Config.Environment != "production")
	DB = db
}
