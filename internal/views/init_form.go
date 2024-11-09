package views

import (
	"github.com/charmbracelet/huh"
	"github.com/hunterMotko/budgot/internal/data"
	"github.com/hunterMotko/budgot/internal/utils"
)

func initForm(data *data.InitData) *huh.Form {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Starting Balance").Placeholder("$0.00").Value(&data.Balance).Validate(utils.ValidAmount),
		),
		huh.NewGroup(
			huh.NewInput().Title("Food").Placeholder("0.00").Value(&data.Food).Validate(validiteNumbers),
			huh.NewInput().Title("Gifts").Placeholder("0.00").Value(&data.Gifts).Validate(validiteNumbers),
			huh.NewInput().Title("Medical").Placeholder("0.00").Value(&data.Medical).Validate(validiteNumbers),
			huh.NewInput().Title("Home").Placeholder("0.00").Value(&data.Home).Validate(validiteNumbers),
			huh.NewInput().Title("Transportation").Placeholder("0.00").Value(&data.Transportation).Validate(validiteNumbers),
			huh.NewInput().Title("Personal").Placeholder("0.00").Value(&data.Personal).Validate(validiteNumbers),
			huh.NewInput().Title("Pets").Placeholder("0.00").Value(&data.Pets).Validate(validiteNumbers),
			huh.NewInput().Title("Utilies").Placeholder("0.00").Value(&data.Utilities).Validate(validiteNumbers),
			huh.NewInput().Title("Travel").Placeholder("0.00").Value(&data.Travel).Validate(validiteNumbers),
			huh.NewInput().Title("Debt").Placeholder("0.00").Value(&data.Debt).Validate(validiteNumbers),
			huh.NewInput().Title("Other").Placeholder("0.00").Value(&data.Ex_other).Validate(validiteNumbers),
		),
		huh.NewGroup(
			huh.NewInput().Title("Savings").Placeholder("0.00").Value(&data.Savings).Validate(validiteNumbers),
			huh.NewInput().Title("Paycheck").Placeholder("0.00").Value(&data.Paycheck).Validate(validiteNumbers),
			huh.NewInput().Title("Bonus").Placeholder("0.00").Value(&data.Bonus).Validate(validiteNumbers),
			huh.NewInput().Title("Interest").Placeholder("0.00").Value(&data.Interest).Validate(validiteNumbers),
			huh.NewInput().Title("Other").Placeholder("0.00").Value(&data.In_other).Validate(validiteNumbers),
		),
	)
	return form
}

func RunInit() (map[int]float64, error) {
  var data data.InitData
  err := initForm(&data).Run()
  if err != nil {
    return nil, err
  }
  return data.Process(), nil
}

func validiteNumbers(str string) error {
  if str == "" {
    return nil
  }
  return utils.ValidAmount(str)
}
