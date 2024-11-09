package views

import (
	"github.com/charmbracelet/huh"
	"github.com/hunterMotko/budgot/internal/data"
	"github.com/hunterMotko/budgot/internal/utils"
)

func addForm(action string, opts []string, rec *data.Record) *huh.Form {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Date").Placeholder("Date Format: YYYY-MM-DD").Value(&rec.Date).Validate(utils.ValidDateStr),
			huh.NewInput().Title("Amount").Placeholder("Amount").Value(&rec.Amount).Validate(utils.ValidAmount),
			huh.NewText().Title("Description").CharLimit(155).Lines(2).Value(&rec.Description),
			huh.NewSelect[string]().
				Title("Category").
				OptionsFunc(func() []huh.Option[string] {
					return huh.NewOptions(opts...)
				}, &action).Value(&rec.Category),
		),
	)
	return form
}

func RunAdd(action string, opts []string) (*data.Record, error) {
  var rec data.Record
	form := addForm(action, opts, &rec)
	err := form.Run()
	if err != nil {
    return nil, err
	}
  return &rec, nil
}
