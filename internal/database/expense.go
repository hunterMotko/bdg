package database

import (
	"fmt"
	"time"

	"github.com/hunterMotko/budgot/internal/data"
)

func (s *service) InsertExpense(rec *data.Record) map[string]string {
	query := "INSERT INTO expenses (date, amount, description, category_id) VALUES (?, ?, ?, ?)"
	_, err := s.db.Exec(query, rec.GetDateTime(), rec.GetAmount(), rec.Description, rec.GetCategory())
	if err != nil {
		return map[string]string{
			"message": err.Error(),
		}
	}
	return map[string]string{
		"message": "success",
	}
}

func (s *service) GetPlannedExpense() ([]data.Category, error) {
	rows, err := s.db.Query("SELECT * FROM categories WHERE id < 12")
	if err != nil {
		return nil, err
	}
	var categories []data.Category
	defer rows.Close()
	for rows.Next() {
		var rec data.Category
		err = rows.Scan(&rec.Id, &rec.Name, &rec.Planned)
		if err != nil {
			return nil, err
		}
		categories = append(categories, rec)
	}
	return categories, nil
}

func (s *service) GetExpenseRecords() (map[int]float64, error) {
	rows, err := s.db.Query(
		"select amount, category_id from expenses where STRFTIME('%m', date) = STRFTIME('%m', 'now')",
		fmt.Sprintf("%d", time.Now().Month()),
	)
	if err != nil {
		return nil, err
	}
	res := make(map[int]float64)
	defer rows.Close()
	for rows.Next() {
		var amount float64
		var categoryId int
		err = rows.Scan(&amount, &categoryId)
		if err != nil {
			return nil, err
		}

		res[categoryId] += amount
	}
	return res, nil
}
