package database

import (
	"github.com/hunterMotko/bdg/internal/data"
	"github.com/hunterMotko/bdg/internal/utils"
)

func (s *service) Init() map[string]string {
	init := `
    CREATE TABLE start (
      id INTEGER PRIMARY KEY AUTOINCREMENT, 
      balance REAL,
      date TIMESTAMP NOT NULL
    );
    CREATE TABLE categories (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      name TEXT NOT NULL,
      planned REAL DEFAULT 0.0
    );
    CREATE TABLE income (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      date TIMESTAMP NOT NULL,
      amount REAL DEFAULT 0.0,
      description TEXT,
      category_id INT,
      FOREIGN KEY(category_id) REFERENCES categories(id)
    );
    CREATE TABLE expenses (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      date TIMESTAMP NOT NULL,
      amount REAL DEFAULT 0.0,
      description TEXT,
      category_id INT,
      FOREIGN KEY(category_id) REFERENCES categories(id)
    );
  `
	_, err := s.db.Exec(init)
	if err != nil {
		return map[string]string{
		"message": err.Error(),
	}
 
	}
	return map[string]string{
		"message": "success",
	}
}

func (s *service) AddPlanned(mp map[int]float64) map[string]string {
	_, err := s.db.Exec("INSERT INTO start (balance, date) VALUES (?, ?)", mp[0], utils.FormatTimeNow())
	if err != nil {
		return map[string]string{
			"message": err.Error(),
		}
	}
	tx, err := s.db.Begin()
	if err != nil {
		return map[string]string{
			"message": err.Error(),
		}
	}
	stmt, err := tx.Prepare("INSERT INTO categories(name, planned) VALUES(?, ?)")
	if err != nil {
		return map[string]string{
			"message": err.Error(),
		}
	}
	defer stmt.Close()
  for i := 1; i < 17; i++ {
    key := data.CategoryString(i)
		_, err := stmt.Exec(key, mp[i])
		if err != nil {
			return map[string]string{
				"message": err.Error(),
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		return map[string]string{
			"message": err.Error(),
		}
	}
	return map[string]string{
		"message": "success",
	}
}
