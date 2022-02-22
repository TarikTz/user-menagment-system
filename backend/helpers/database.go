package helpers

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Access *App

var (
	dbName = "ums_db"

	// SQL
	sqlDriver   = "mysql"
	sqlUser     = "root"
	sqlPassword = "root"
	sqlProtocol = "tcp"       
	sqlAddress  = "localhost"

	sqlPort = "8889" // "mysql"
)

type App struct {
	DB *sqlx.DB
}

func DBAccess() *App {
	a := new(App)

	dbName = "ums_db"
	sqlUser = "root"
	sqlPassword = "root"
	sqlAddress = "localhost"

	var err error

	a.DB, err = sqlx.Connect(sqlDriver, a.getDSN())

	if err != nil {
		fmt.Println("Error connecting to Portal DB", err)
	} else {
		fmt.Println("Connected to Portal DB")
	}

	Access = a

	return a
}

// getDSN ->
func (a App) getDSN() string {
	return sqlUser + ":" + sqlPassword + "@" + sqlProtocol + "(" + sqlAddress + ":" + sqlPort + ")" + "/" + dbName // + "?tls=false"
}

// GetSQLDB ->
func (a App) GetSQLDB() *sqlx.DB {
	return a.DB
}
