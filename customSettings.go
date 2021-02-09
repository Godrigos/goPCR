package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (p *pcr) saveCustomStock() {
	_, err := strconv.ParseFloat(p.selectStockVal.Text, 64)
	if err != nil {
		str := ("Non numerical value for\n" + " stock " + p.customStock.Text +
			"!" + "\nNot applying changes!")
		warning := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
			fyne.TextStyle{})
		dialog.ShowCustom("Error", "Ok", warning, p.w)
		p.customStock.SetText("")
		p.selectStockVal.SetText("")
	} else if p.customStock.Text == "" {
		str := ("Select a value to change!")
		warning := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
			fyne.TextStyle{})
		dialog.ShowCustom("Error", "Ok", warning, p.w)
		p.customStock.SetText("")
		p.selectStockVal.SetText("")
	} else {
		msg := true
		switch p.customStock.Text {
		case "Buffer":
			p.application.Preferences().SetString("StockBuffer",
				p.selectStockVal.Text)
		case "DNTPs":
			p.application.Preferences().SetString("StockDNTPs",
				p.selectStockVal.Text)
		case "MgCl₂":
			p.application.Preferences().SetString("StockMgCl2",
				p.selectStockVal.Text)
		case "Primer 1":
			p.application.Preferences().SetString("StockP1", p.selectStockVal.Text)
		case "Primer 2":
			p.application.Preferences().SetString("StockP2", p.selectStockVal.Text)
		case "Primer 3":
			p.application.Preferences().SetString("StockP3", p.selectStockVal.Text)
		case "Primer 4":
			p.application.Preferences().SetString("StockP4", p.selectStockVal.Text)
		case "Glycerol":
			p.application.Preferences().SetString("StockGly",
				p.selectStockVal.Text)
		case "DMSO":
			p.application.Preferences().SetString("StockDMSO",
				p.selectStockVal.Text)
		case "Taq":
			p.application.Preferences().SetString("StockTaq", p.selectStockVal.Text)
		case "DNA":
			p.application.Preferences().SetString("StockDNA", p.selectStockVal.Text)
		default:
			str := ("Wrong field name, select from menu!")
			warning := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
				fyne.TextStyle{})
			dialog.ShowCustom("Error", "Ok", warning, p.w)
			p.customStock.SetText("")
			p.selectStockVal.SetText("")
			msg = false
		}
		if msg {
			str := ("Default value for\n" + " stock " + p.customStock.Text +
				"\nchanged successfully!")
			warning := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
				fyne.TextStyle{})
			dialog.ShowCustom("Done", "Ok", warning, p.w)
			p.customStock.SetText("")
			p.selectStockVal.SetText("")
		}
	}
}

func (p *pcr) saveCustomPCR() {
	_, err := strconv.ParseFloat(p.selectPCRVal.Text, 64)
	if err != nil {
		str := ("Non numerical value for\n" + " PCR " + p.customPCR.Text +
			"!" + "\nNot applying changes!")
		warning := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
			fyne.TextStyle{})
		dialog.ShowCustom("Error", "Ok", warning, p.w)
		p.customPCR.SetText("")
		p.selectPCRVal.SetText("")
	} else if p.customPCR.Text == "" {
		str := ("Select a value to change!")
		warning := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
			fyne.TextStyle{})
		dialog.ShowCustom("Error", "Ok", warning, p.w)
		p.customPCR.SetText("")
		p.selectPCRVal.SetText("")
	} else {
		msg := true
		switch p.customPCR.Text {
		case "Buffer":
			p.application.Preferences().SetString("PCRBuffer", p.selectPCRVal.Text)
		case "DNTPs":
			p.application.Preferences().SetString("PCRDNTPs", p.selectPCRVal.Text)
		case "MgCl₂":
			p.application.Preferences().SetString("PCRMgCl2", p.selectPCRVal.Text)
		case "Primer 1":
			p.application.Preferences().SetString("PCRP1", p.selectPCRVal.Text)
		case "Primer 2":
			p.application.Preferences().SetString("PCRP2", p.selectPCRVal.Text)
		case "Primer 3":
			p.application.Preferences().SetString("PCRP3", p.selectPCRVal.Text)
		case "Primer 4":
			p.application.Preferences().SetString("PCRP4", p.selectPCRVal.Text)
		case "Glycerol":
			p.application.Preferences().SetString("PCRGly", p.selectPCRVal.Text)
		case "DMSO":
			p.application.Preferences().SetString("PCRDMSO", p.selectPCRVal.Text)
		case "Taq":
			p.application.Preferences().SetString("PCRTaq", p.selectPCRVal.Text)
		case "DNA":
			p.application.Preferences().SetString("PCRDNA", p.selectPCRVal.Text)
		default:
			str := ("Wrong field name, select from menu!")
			warning := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
				fyne.TextStyle{})
			dialog.ShowCustom("Error", "Ok", warning, p.w)
			p.customPCR.SetText("")
			p.selectPCRVal.SetText("")
			msg = false
		}
		if msg {
			str := ("Default value for\n" + " PCR " + p.customPCR.Text +
				"\nchanged successfully!")
			warning := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
				fyne.TextStyle{})
			dialog.ShowCustom("Done", "Ok", warning, p.w)
			p.customPCR.SetText("")
			p.selectPCRVal.SetText("")
		}
	}
}

func (p *pcr) saveCustomMix() {
	_, err := strconv.ParseFloat(p.selectMixVal.Text, 64)
	if err != nil {
		str := ("Non numerical value for\n" + " mix " + p.customMix.Text +
			"!" + "\nNot applying changes!")
		warning := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
			fyne.TextStyle{})
		dialog.ShowCustom("Error", "Ok", warning, p.w)
		p.customMix.SetText("")
		p.selectMixVal.SetText("")
	} else if p.customMix.Text == "" {
		str := ("Select a value to change!")
		warning := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
			fyne.TextStyle{})
		dialog.ShowCustom("Error", "Ok", warning, p.w)
		p.customMix.SetText("")
		p.selectMixVal.SetText("")
	} else {
		msg := true
		switch p.customMix.Text {
		case "Volume":
			p.application.Preferences().SetString("MixVolume", p.selectMixVal.Text)
		case "Reactions":
			p.application.Preferences().SetString("MixReactNum", p.selectMixVal.Text)
		default:
			str := ("Wrong field name, select from menu!")
			warning := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
				fyne.TextStyle{})
			dialog.ShowCustom("Error", "Ok", warning, p.w)
			p.customMix.SetText("")
			p.selectMixVal.SetText("")
			msg = false
		}
		if msg {
			str := ("Default value for\n" + " mix " + p.customMix.Text +
				"\nchanged successfully!")
			warning := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
				fyne.TextStyle{})
			dialog.ShowCustom("Done", "Ok", warning, p.w)
			p.customMix.SetText("")
			p.selectMixVal.SetText("")
		}
	}
}
