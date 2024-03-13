package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	user     = "postgres"
	port     = 5432
	password = "Bagus062912"
	dbname   = "digitalent"

)
var (
	db       *sql.DB
	err      error
)

func HandlerConnection() {
	// appConfig := config.GetAppConfig()
	// psqlInfo := fmt.Sprintf("host")
}

func InitialaizeDatabase(){
	HandlerConnection()
	// handlerCreateRequiredDatabase()
}

func GetDatabaseInstance() *sql.DB {
	if db == nil {
		log.Panic("Database tidak boleh nill")
		
	}
	return  db
}