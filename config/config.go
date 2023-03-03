package config

import "database/sql"

func Connect() *sql.DB {
	dbDriver := "mysql"
	dbUser := "jwtuser"
	dbPass := "jwtpass"
	dbName := "jwtdb"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(10.225.5.31:3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
