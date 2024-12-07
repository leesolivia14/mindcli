package models

import (
	"github.com/leesolivia14/mindcli/db"
)

type BudgetRecord struct {
	ID     int
	year   string
	month  string
	Amount float64
}

func GetAllBudgetRecords(year int, month int) ([]BudgetRecord, error) {
	query := `SELECT amount, description FROM budget WHERE year = ? AND month = ?`
	rows, err := db.DB.Query(query)
}

func Save() error {

}
