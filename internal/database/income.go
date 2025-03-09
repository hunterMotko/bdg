package database

import (
	"fmt"
	"time"

	"github.com/hunterMotko/bdg/internal/data"
)

func (s *service) InsertIncome(rec *data.Record) map[string]string {
	query := "INSERT INTO income (date, amount, description, category_id) VALUES (?, ?, ?, ?)"
	_, err := s.db.Exec(query, rec.GetDateTime(), rec.GetAmount(), rec.Description, rec.GetCategory())
	if err != nil {
		return map[string]string{
			"message": "failed exec income record",
		}
	}
	return map[string]string{
		"message": "success",
	}
}

func (s *service) GetPlannedIncome() ([]data.Category, error) {
	rows, err := s.db.Query("SELECT * FROM categories WHERE id > 11")
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

func (s *service) GetIncomeRecords() (map[int]float64, error) {
	rows, err := s.db.Query(
		"select amount, category_id from income where STRFTIME('%m', date) = STRFTIME('%m', 'now')",
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
