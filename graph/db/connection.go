package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	// USERNAME := os.Getenv("USERNAME")
	// PASSWORD := os.Getenv("PASSWORD")
	// HOST := os.Getenv("HOST")
	// // PORT := os.Getenv("PORT")
	// DATABASE := os.Getenv("DATABASE")
	// dbconf := USERNAME + ":" + PASSWORD + "@tcp(" + HOST + ")/" + DATABASE + "?charset=utf8mb4" + "&parseTime=True"
	// dbconf := USERNAME + ":" + PASSWORD + "@tcp(" + HOST + ":" + PORT + ")/" + DATABASE + "?charset=utf8mb4" + "&parseTime=True"
	db, err := sql.Open(os.Getenv("DRIVER"), "root:root@tcp(mysql_container)/GOLANG_MYSQL?parseTime=true")
	// db, err := sql.Open(os.Getenv("DRIVER"), dbconf)
	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil, err
	}
	return db, err
}
