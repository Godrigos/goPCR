package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

func checkFloat(err error, s string) bool {
	if err != nil {
		str := ("Non numerical value for\n" + s + "!")
		warning := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
			fyne.TextStyle{})
		dialog.ShowCustom("Error", "Ok", warning, w)

		err = nil

		return true
	}
	return false
}
