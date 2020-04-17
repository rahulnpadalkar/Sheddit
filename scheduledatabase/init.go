package scheduledatabase

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var dbSingelton DBService

// InitilaizeService : Depending upon env initialize either Postgres or Bolt
func InitilaizeService() {
	postgresurl := os.Getenv("postgres_url")
	if postgresurl != "" {
		dbSingelton = &PostgresClient{}
		dbSingelton.InitializeDB()
	} else {
		dbSingelton = &BoltService{}
		dbSingelton.InitializeDB()
	}
}

// GetInstance : Get DBService instance
func GetInstance() DBService {
	return dbSingelton
}
