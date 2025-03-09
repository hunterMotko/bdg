package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/hunterMotko/bdg/internal/data"
	_ "github.com/mattn/go-sqlite3"
)

type Service interface {
	Health() map[string]string
	Init() map[string]string
	AddPlanned(map[int]float64) map[string]string
	InsertIncome(*data.Record) map[string]string
	InsertExpense(*data.Record) map[string]string
  GetSums() (*data.Sums, error)
	GetPlannedIncome() ([]data.Category, error)
	GetPlannedExpense() ([]data.Category, error)
  GetIncomeRecords() (map[int]float64, error)
  GetExpenseRecords() (map[int]float64, error)
}

type service struct {
	db *sql.DB
}

var (
	dbInstance *service
)

func New(dburl string) Service {
	if dbInstance != nil {
		return dbInstance
	}
	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		log.Fatal(err)
	}
	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err := s.db.PingContext(ctx)
	if err != nil {
		log.Fatalf("db down: %v", err)
	}
	return map[string]string{
		"message": "It's healthy",
	}
}

func (s *service) GetSums() (*data.Sums, error) {
  var sums data.Sums
	err := s.db.QueryRow(
    `SELECT 
      (SELECT balance FROM start) as start,
      (SELECT COALESCE(SUM(planned), 0.0) FROM categories WHERE id < 12) as planned_expense,
      (SELECT COALESCE(SUM(planned), 0.0) FROM categories WHERE id > 11) as planned_income,
      (SELECT COALESCE(SUM(amount), 0.0) FROM expenses WHERE strftime('%m', date) = STRFTIME('%m', 'now')) as total_expense,
      (SELECT COALESCE(SUM(amount), 0.0) FROM income WHERE strftime('%m', date) = STRFTIME('%m', 'now')) as total_income
    `,
  ).Scan(&sums.Start, &sums.PlannedExpense, &sums.PlannedIncome, &sums.TotalExpense, &sums.TotalIncome)
	if err != nil {
		return nil, err
	}
	return &sums, nil
}

