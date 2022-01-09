package models

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)
var _ error = godotenv.Load(".env")

var db_user string = os.Getenv("DB_USER")
var db_pass string = os.Getenv("DB_PASSWORD")
var db_name string = os.Getenv("DB_NAME")

var Dsn string = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",db_user,db_pass,db_name)


func Create_db(){
	
	var Root_sql string = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/",db_user,db_pass)
	db, err := sql.Open("mysql", Root_sql)
	if err != nil {
		panic(err)
	}
	_,err = db.Exec("CREATE DATABASE IF NOT EXISTS "+db_name)

	if err != nil{
		panic(err)
	}

	defer db.Close()
}