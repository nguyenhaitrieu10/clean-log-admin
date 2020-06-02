package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func clear_operation_log() {
	database, err := sql.Open("sqlite3", "/data/admin.db")
	if err != nil {
		fmt.Println(1)
		fmt.Println(err)
		return
	}

	RETENTION_DAYS := os.Getenv("RETENTION_DAYS")
	if RETENTION_DAYS == "" {
		RETENTION_DAYS = "14"
	}

	query := fmt.Sprintf("DELETE FROM goadmin_operation_log WHERE method='GET' AND created_at <= date('now','-%s day')", RETENTION_DAYS)
	statement, err := database.Prepare(query)

	if err != nil {
		fmt.Println(2)
		fmt.Println(err)
		return
	}

	result, err := statement.Exec()
	if err != nil {
		fmt.Println(3)
		fmt.Println(err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(4)
		fmt.Println(err)
		return
	}

	fmt.Printf("RowsAffected: %v\n", rowsAffected)
}

func main() {
	clear_operation_log()
}