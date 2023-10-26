package idmanager

import (
	"github.com/rivo/tview"
)

type Person struct {
	Title     string
	FirstName string
	LastName  string
	Age       string
	Password  string
}

func Save(form *tview.Form) *Person {
	var newperson Person

	_, newperson.Title = form.GetFormItemByLabel("Title").(*tview.DropDown).GetCurrentOption()
	//newperson.FirstName = form.GetFormItemByLabel("First Name").(*tview.InputField).GetText()
	newperson.LastName = form.GetFormItemByLabel("Last Name").(*tview.InputField).GetText()
	newperson.Age = form.GetFormItemByLabel("Age").(*tview.InputField).GetText()
	newperson.Password = form.GetFormItemByLabel("Password").(*tview.InputField).GetText()

	return &newperson
}
