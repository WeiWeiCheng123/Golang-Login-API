package model

import (
	"database/sql"
	"time"
)

func Init(mysql_db *sql.DB) {
	mydb = mysql_db
}

var (
	mydb *sql.DB
	Id   int
	user string
	pass string
)

type User struct {
	Username string `xorm:"pk" json:"username"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return "user"
}

func Connect_mysql(user string, password string, port string, dbname string) *sql.DB {
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+port+")/"+dbname)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	return db
}

func CheckExist(username string, password string) (string, string, error) {
	stmt, err := mydb.Prepare("SELECT * FROM users WHERE username = $1")
	if err != nil {
		return "", "", err
	}

	err = stmt.QueryRow(username).Scan(&Id, &user, &pass)
	defer stmt.Close()
	if err != nil {
		return "", "", err
	}

	return user, pass, nil
}

func CheckAccept(username string) (string, error) {
	stmt, err := mydb.Prepare("SELECT * FROM users WHERE username = $1")
	if err != nil {
		return "", err
	}

	err = stmt.QueryRow(username).Scan(&Id, &user, &pass)
	defer stmt.Close()
	if err != nil {
		return "", err
	}

	return user, nil
}

func Register(username string, password string) error {
	stmt, err := mydb.Prepare("INSERT INTO users(username,password) values($1,$2)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(username, password)
	defer stmt.Close()
	if err != nil {
		return err
	}
	return nil
}
