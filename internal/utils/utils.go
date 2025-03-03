package utils

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/hunterMotko/budgot/internal/config"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lucasb-eyer/go-colorful"
)

var (
	configFileErr error = errors.New("Config File Error.\nPlease Use budgot init to configure cli")
	writeFileErr  error = errors.New("FILE WRITE ERROR")
	dateErr       error = errors.New("Please follow correct date format")
	monthErr      error = errors.New("Incorrect Month")
	dayErr        error = errors.New("Incorrect Day")
	yearErr       error = errors.New("Incorrect Year")
)

func ValidDateStr(date string) error {
	if len(date) != 10 {
		return dateErr
	}
  if !strings.Contains(date, "-") {
    return dateErr
  }
	dates := strings.Split(date, "-")
	year, month, day := dates[0], dates[1], dates[2]
	if len(dates) != 3 || len(year) != 4 || len(month) != 2 || len(day) != 2 {
		return dateErr
	}
	err := checkMonth(month)
	if err != nil {
		return err
	}
	err = checkDay(day)
	if err != nil {
		return err
	}
	err = checkYear(year)
	if err != nil {
		return err
	}
	return nil
}

func checkMonth(month string) error {
	i, err := ParseInt(month)
	if err != nil {
		return err
	}
	if i > 12 {
		return monthErr
	}
	return nil
}

func checkDay(day string) error {
	i, err := ParseInt(day)
	if err != nil {
		return err
	}
	if i > 31 {
		return monthErr
	}
	return nil
}

func checkYear(year string) error {
	curYear := time.Now().Year()
	i, err := ParseInt(year)
	if err != nil {
		return err
	}
	if i > curYear {
		return yearErr
	}
	return nil
}

func ValidAmount(amt string) error {
	_, err := ParseAmount(amt)
	return err
}

func ParseInt(n string) (int, error) {
	i, err := strconv.ParseInt(n, 10, 0)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

func ParseAmount(amt string) (float64, error) {
	n, err := strconv.ParseFloat(amt, 64)
	if err != nil {
		return 0, errors.New("Parse amount error")
	}
	return n, nil
}

func CheckDBFileExists(conf *config.Config) bool {
	if _, err := os.Stat(conf.String()); errors.Is(err, os.ErrNotExist) {
    return false
	}
	return true
}

func FormatTimeNow() time.Time {
  str := "2006-01-02" 
  now := time.Now().Format(str)
  res, err := time.Parse(str, now)
  if err != nil {
    log.Fatalf("WTF TIME: %v", err)
  }
  return res
}

func BarPercentages(perc float64) (float64, float64) {
	if perc < 0 {
		return 1.0, math.Abs(perc)
	}
	return perc, 1.0
}
// Generate a blend of colors.
func MakeRampStyles(colorA, colorB string, steps float64) (s []lipgloss.Style) {
	cA, _ := colorful.Hex(colorA)
	cB, _ := colorful.Hex(colorB)
	for i := 0.0; i < steps; i++ {
		c := cA.BlendLuv(cB, i/steps)
		s = append(s, lipgloss.NewStyle().Foreground(lipgloss.Color(colorToHex(c))))
	}
	return
}
// Convert a colorful.Color to a hexadecimal format.
func colorToHex(c colorful.Color) string {
	return fmt.Sprintf("#%s%s%s", colorFloatToHex(c.R), colorFloatToHex(c.G), colorFloatToHex(c.B))
}
// Helper function for converting colors to hex. Assumes a value between 0 and
// 1.
func colorFloatToHex(f float64) (s string) {
	s = strconv.FormatInt(int64(f*255), 16)
	if len(s) == 1 {
		s = "0" + s
	}
	return
}
