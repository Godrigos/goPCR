package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (p *pcr) checkFloat(err error, s string) bool {
	if err != nil {
		str := ("Non numerical value for\n" + s + "!")
		warning := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
			fyne.TextStyle{})
		dialog.ShowCustom("Error", "Ok", warning, p.w)
		return true
	}
	return false
}
