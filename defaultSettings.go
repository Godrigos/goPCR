package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

func (p *pcr) restore() {
	str := ("This will restore all values back to their defaults!\n" +
		"Are You Sure?")
	warning := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
		fyne.TextStyle{})
	dialog.ShowCustomConfirm("Warning", "Ok", "Cancel", warning,
		p.defaults, p.w)
}

func (p *pcr) defaults(b bool) {
	if b {
		// Define default values for stock solutions
		p.application.Preferences().SetString("StockBuffer", "10")
		p.application.Preferences().SetString("StockDNTPs", "2.5")
		p.application.Preferences().SetString("StockMgCl2", "50")
		p.application.Preferences().SetString("StockP1", "10")
		p.application.Preferences().SetString("StockP2", "10")
		p.application.Preferences().SetString("StockP3", "10")
		p.application.Preferences().SetString("StockP4", "10")
		p.application.Preferences().SetString("StockGly", "100")
		p.application.Preferences().SetString("StockDMSO", "100")
		p.application.Preferences().SetString("StockTaq", "5")
		p.application.Preferences().SetString("StockDNA", "10")

		// Define default values for PCR concentrations
		p.application.Preferences().SetString("PCRBuffer", "1")
		p.application.Preferences().SetString("PCRDNTPs", "0.2")
		p.application.Preferences().SetString("PCRMgCl2", "1.5")
		p.application.Preferences().SetString("PCRP1", "0.2")
		p.application.Preferences().SetString("PCRP2", "0.2")
		p.application.Preferences().SetString("PCRP3", "0")
		p.application.Preferences().SetString("PCRP4", "0")
		p.application.Preferences().SetString("PCRGly", "0")
		p.application.Preferences().SetString("PCRDMSO", "0")
		p.application.Preferences().SetString("PCRTaq", "0.05")
		p.application.Preferences().SetString("PCRDNA", "5")

		// Define mix default values
		p.application.Preferences().SetString("MixVolume", "25")
		p.application.Preferences().SetString("MixReactNum", "1")
	}
}
