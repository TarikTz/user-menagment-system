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
	sqlProtocol = "tcp"       // "unix"
	sqlAddress  = "localhost" // "/tmp/mysql.sock"

	// Common port for MySQL services
	sqlPort = "8889" // "mysql"
)

type App struct {
	DB *sqlx.DB
}

func DBAccess(deploy bool, runas string) *App {
	a := new(App)

	// a.devMode = (runas == RunModeDev) || (runas == RunModeTest) || (runas == RunModeLocal)

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

	a.DB.SetMaxIdleConns(10)
	a.DB.SetMaxOpenConns(5)

	Access = a

	return a
}

func (a App) getDSN() string {
	return sqlUser + ":" + sqlPassword + "@" + sqlProtocol + "(" + sqlAddress + ":" + sqlPort + ")" + "/" + dbName // + "?tls=false"
}

// GetSQLDB ->
func (a App) GetSQLDB() *sqlx.DB {
	return a.DB
}
