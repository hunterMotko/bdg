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
	for i := range vals.NumField() {
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
	var dm = map[string]int{
		"food": Food, "gifts": Gifts, "medical": Medical, "home": Home,
		"transportation": Transportation, "personal": Personal, "pets": Pets, "utilities": Utilities,
		"travel": Travel, "debt": Debt, "ex_other": Ex_other, "savings": Savings,
		"paycheck": Paycheck, "bonus": Bonus, "interest": Interest, "in_other": In_other,
	}
	v, ok := dm[str]
	if !ok {
		return 0
	}
	return v
}

func CategoryString(n int) string {
	var dm = map[int]string{
		1: "food", 2: "gifts", 3: "medical", 4: "home",
		5: "transportation", 6: "personal", 7: "pets", 8: "utilities",
		9: "travel", 10: "debt", 11: "ex_other", 12: "savings",
		13: "paycheck", 14: "bonus", 15: "interest", 16: "in_other",
	}
	v, ok := dm[n]
	if !ok {
		return ""
	}
	return v
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
	return float64((s.TotalExpense / s.PlannedExpense) - 1)
}

func (s *Sums) IncomePerc() float64 {
	return float64((s.TotalIncome / s.PlannedIncome) - 1)
}

func (s *Sums) CalcEndBal() float64 {
	return s.Start + (s.TotalIncome - s.TotalExpense)
}

func (s *Sums) Saved(end float64) float64 {
	return end - s.Start
}
