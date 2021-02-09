package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (p *pcr) setTheme() {
	switch p.application.Preferences().String("Theme") {
	case "light":
		p.application.Settings().SetTheme(theme.LightTheme())
	case "dark":
		p.application.Settings().SetTheme(theme.DarkTheme())
	default:
		p.application.Settings().SetTheme(theme.LightTheme())
	}
}

func (p *pcr) darkTheme() {
	p.application.Preferences().SetString("Theme", "dark")
	str := ("Theme will be set to dark!")
	warning := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
		fyne.TextStyle{})
	dialog.ShowCustom("Done", "Ok", warning, p.w)
}

func (p *pcr) lightTheme() {
	p.application.Preferences().SetString("Theme", "light")
	str := ("Theme will be set to light!")
	warning := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
		fyne.TextStyle{})
	dialog.ShowCustom("Done", "Ok", warning, p.w)
}
