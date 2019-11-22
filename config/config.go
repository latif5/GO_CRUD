package config

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnDB() (db *sql.DB, err error){
	dbDriver := "mysql"
	dbName := "golang_crud"
	dbUser := "root"
	dbPass := ""
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}