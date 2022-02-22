package models

import (
	"fmt"
	"strconv"

	"github.com/TarikTz/user-menagment-system/helpers"
)

type User struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Username  string `json:"username" db:"username"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	Status    int    `json:"status" db:"status"`
}

type FilterUser struct {
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
	Order  string `json:"sort"`
	Status string    `json:"status"`
}

type Permission struct {
	ID          int    `json:"id" db:"id"`
	UserID      int    `json:"user_id" db:"user_id"`
	Code        int    `json:"code" db:"code"`
	Description string `json:"description" db:"description"`
}

var UserTable = `CREATE TABLE 
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

// ListUsers ...
func ListUsers(filters FilterUser) ([]User, int, error) {
	var users []User

	db := helpers.Access.GetSQLDB()

	filter := ""
	order := " ORDER BY id DESC"

	if filters.Status != "-1" {
		filter = " WHERE status = " + filters.Status
	}

	if filters.Order != "" {
		order = " ORDER BY " + filters.Order
	}

	query := `SELECT 
					id,first_name,last_name,username,email,status
				FROM users`

	if filter != "" {
		query += filter
	}

	query += order

	query += " LIMIT ? OFFSET ?"

	fmt.Println(query)


	err := db.Select(&users, query, filters.Limit, (filters.Limit*filters.Page)-filters.Limit)

	if err != nil {
		return nil, 0, err
	}

	var count int
	
	err = db.Get(&count, "SELECT COUNT(id) FROM users" + filter)

	return users, count, err
}

// CreateUser ...
func (u *User) CreateUser() error {
	db := helpers.Access.GetSQLDB()

	query := `INSERT INTO users
					(first_name,last_name,username,password,email,status)
				VALUES
					(:first_name,:last_name,:username,:password,:email,:status)`

	_, err := db.NamedExec(query, u)

	return err
}

// Delete ...
func (u *User) Delete(id string) error {
	db := helpers.Access.GetSQLDB()

	query := `DELETE FROM users WHERE id = ?`

	_, err := db.Exec(query, id)

	return err
}

// User ...
func (u *User) User(id string) error {
	db := helpers.Access.GetSQLDB()

	query := `SELECT 
					id,first_name,last_name,username,email,status
				FROM users WHERE id = ?`

	err := db.QueryRowx(query, id).StructScan(u)

	return err

}

// Update ...
func (u *User) Update(id string) error {
	db := helpers.Access.GetSQLDB()
	var err error

	u.ID, err = strconv.Atoi(id)

	if err != nil {
		fmt.Println(err)
		return err
	}

	query := `UPDATE users
				SET first_name=:first_name,last_name=:last_name,username=:username,
					password=:password,email=:email,status=:status
				WHERE id = :id`

	_, err = db.NamedExec(query, u)

	return err
}
