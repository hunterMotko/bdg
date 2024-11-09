package data

import (
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hunterMotko/budgot/internal/utils"
)

const (
	// Expenses
	Food int = iota + 1
	Gifts
	Medical
	Home
	Transportation
	Personal
	Pets
	Utilities
	Travel
	Debt
	Ex_other
	// Income
	Savings
	Paycheck
	Bonus
	Interest
	In_other
)

type InitData struct {
	Balance string
	// Expenses
	Food           string
	Gifts          string
	Medical        string
	Home           string
	Transportation string
	Personal       string
	Pets           string
	Utilities      string
	Travel         string
	Debt           string
	Ex_other       string
	// income
	Savings  string
	Paycheck string
	Bonus    string
	Interest string
	In_other string
}

func (id *InitData) Process() map[int]float64 {
	res := make(map[int]float64)
	vals := reflect.ValueOf(*id)
	tp := vals.Type()
	for i := 0; i < vals.NumField(); i++ {
		key := strings.ToLower(tp.Field(i).Name)
		val := vals.Field(i).String()
		if key == "" {
			continue
		}
		if val == "" {
			val = "0"
		}
		flt, err := utils.ParseAmount(val)
		if err != nil {
			log.Fatalf("PROCESS PARSE ERROR: %v\n", err)
		}
		res[categoryInt(key)] = flt
	}
	return res
}

type Category struct {
	Id      int
	Name    string
	Planned float64
}

type Record struct {
	Date        string
	Amount      string
	Description string
	Category    string
}

func (r *Record) GetDateTime() time.Time {
	dt, err := time.Parse("2006-01-02", r.Date)
	if err != nil {
		log.Fatalf("TIME PARSE ERROR: %v", err)
	}
	return dt
}

func (r *Record) GetAmount() float64 {
	n, err := utils.ParseAmount(r.Amount)
	if err != nil {
		log.Fatalf("AMOUNT PARSE ERROR: %v", err)
	}
	return n
}

func (r *Record) GetCategory() int {
	return categoryInt(r.Category)
}

func categoryInt(str string) int {
	switch str {
	case "food":
		return Food
	case "gifts":
		return Gifts
	case "medical":
		return Medical
	case "home":
		return Home
	case "transportation":
		return Transportation
	case "personal":
		return Personal
	case "pets":
		return Pets
	case "utilities":
		return Utilities
	case "travel":
		return Travel
	case "debt":
		return Debt
	case "ex_other":
		return Ex_other
	case "savings":
		return Savings
	case "paycheck":
		return Paycheck
	case "bonus":
		return Bonus
	case "interest":
		return Interest
	case "in_other":
		return In_other
	}
	return 0
}

func CategoryString(n int) string {
	switch n {
	case 1:
		return "food"
	case 2:
		return "gifts"
	case 3:
		return "medical"
	case 4:
		return "home"
	case 5:
		return "transportation"
	case 6:
		return "personal"
	case 7:
		return "pets"
	case 8:
		return "utilities"
	case 9:
		return "travel"
	case 10:
		return "debt"
	case 11:
		return "ex_other"
	case 12:
		return "savings"
	case 13:
		return "paycheck"
	case 14:
		return "bonus"
	case 15:
		return "interest"
	case 16:
		return "in_other"
	}
	return ""
}

type Sums struct {
	Start          float64
	PlannedExpense float64
	PlannedIncome  float64
	TotalExpense   float64
	TotalIncome    float64
}

func (s *Sums) CalcPercChange(end float64) float64 {
	return float64((end / s.Start) - 1)
}
func (s *Sums) ExpensePerc() float64 {
  return float64((s.TotalExpense/s.PlannedExpense)-1)
}
func (s *Sums) IncomePerc() float64 {
  return float64((s.TotalIncome/s.PlannedIncome)-1)
}
func (s *Sums) CalcEndBal() float64 {
	return s.Start + (s.TotalIncome - s.TotalExpense)
}
func (s *Sums) Saved(end float64) float64 {
	return end - s.Start
}
