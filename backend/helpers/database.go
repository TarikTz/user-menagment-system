package helpers

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Access *App

var (
	dbName = "ums_db"
	sqlDriver   = "mysql"
	sqlUser     = "user"
	sqlPassword = "user"
	sqlProtocol = "tcp"       
	sqlAddress  = "localhost"
	sqlPort = "3306"
)

var UserTable = `CREATE TABLE IF NOT EXISTS
				ums_db.users 
				( 
					id INT(11) NOT NULL AUTO_INCREMENT , 
					first_name VARCHAR(50) NOT NULL , 
					last_name VARCHAR(50) NOT NULL , 
					username VARCHAR(50) NOT NULL , 
					password CHAR(60) NOT NULL , 
					email VARCHAR(100) NOT NULL , 
					status INT NOT NULL DEFAULT '0' , 
					PRIMARY KEY (id)
				) 
				ENGINE = InnoDB;`

type App struct {
	DB *sqlx.DB
}

func DBAccess() *App {
	a := new(App)

	dbName = "ums_db"
	sqlUser = "root"
	sqlPassword = "root"
	sqlAddress = "mysqlDB"

	var err error

	a.DB, err = sqlx.Connect(sqlDriver, a.getMySQLAddress())

	if err != nil {
		fmt.Println("Error connecting to database", err)
		fmt.Println(err)
	} else {
		fmt.Println("Connected to database")
	}

	err = createUsersTable(a.DB)

	if err != nil {
        fmt.Printf("Create product table failed with error %s", err)
    }

	Access = a

	return a
}

// getMySQLAddress ->
func (a App) getMySQLAddress() string {
	return sqlUser + ":" + sqlPassword + "@" + sqlProtocol + "(" + sqlAddress + ":" + sqlPort + ")" + "/" + dbName // + "?tls=false"
}

// GetSQLDB ->
func (a App) GetSQLDB() *sqlx.DB {
	return a.DB
}

func createUsersTable(db *sqlx.DB) error {  
    _, err := db.Exec(UserTable)

	if err != nil {
		fmt.Println("Error creating users table", err)
	}

	return err
}
