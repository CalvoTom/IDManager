package idmanager

import (
	"github.com/rivo/tview"
)

func Drawer() *Person {
	app := tview.NewApplication()
	form := tview.NewForm().
		AddDropDown("Title", []string{"Mr.", "Ms.", "Mrs.", "Dr."}, 0, nil).
		AddInputField("First name", "", 20, nil, nil).
		AddInputField("Last name", "", 20, nil, nil).
		AddInputField("Age", "", 20, nil, nil).
		AddPasswordField("Password", "", 10, '*', nil).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("Enter personal data").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	return Save(form)
}
