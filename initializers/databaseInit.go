package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

)

var DB *gorm.DB
func ConnectToDB() {
	var err error
	// connection_string := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
	connection_string := LoadEnvVariables("DB_CONN_STRING")
	DB,err = gorm.Open(postgres.Open(connection_string), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to db")
	}
	fmt.Println("Established a successful connection with database")
	
}