package idmanager

import (
	"strings"

	"github.com/rivo/tview"
)

func Display(oldperson *Person) {
	app := tview.NewApplication()
	table := tview.NewTable().
		SetBorders(true)
	lorem := strings.Split("Title FirstName LastName Age Password "+oldperson.Title+" "+oldperson.FirstName+" "+oldperson.LastName+" "+oldperson.Age+" "+oldperson.Password+" ", " ")
	cols, rows := 5, 2
	word := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			table.SetCell(r, c,
				tview.NewTableCell(lorem[word]).
					SetAlign(tview.AlignCenter))
			word = (word + 1) % len(lorem)
		}
	}
	if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}
}
