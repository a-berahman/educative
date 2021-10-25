package models

import (
	"database/sql"
)

//Configuration presents LoadConfig result model
type Configuration struct {
	PostgresConnection *sql.DB
}
