package db

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err := sql.Open("sqlite3", "budget.db")

	if err != nil {
		panic("could not connect to database.")
	}

	createTable()
}

func createTable() {
	createBudgetTable := `
	CREATE TABLE IF NOT EXISTS budget(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		amount REAL NOT NULL,
		year INTEGER NOT NULL,
		month INTEGER NOT NULL
	)
	`

	_, err := DB.Exec(createBudgetTable)
	if err != nil {
		fmt.Println(err)
		panic("could not create budget table.")
	}
}

type BudgetRecord struct {
	ID     int
	year   string
	month  string
	Amount float64
}

func GetAllBudgetRecords(year int, month int) ([]BudgetRecord, error) {
	query := `SELECT amount, description FROM budget WHERE year = ? AND month = ?`
}

func Save() error {

}