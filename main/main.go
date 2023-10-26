package main

import (
	"fmt"
	"idmanager"
)

func main() {
	newperson := idmanager.Drawer()
	fmt.Print(newperson.Age)
}
