package main

import "fyne.io/fyne/theme"

func prefs() {
	// Set the theme
	application.Settings().SetTheme(theme.LightTheme())

	// Define default values for stock solutions
	application.Preferences().SetString("StockBuffer", "10")
	application.Preferences().SetString("StockDNTPs", "2.5")
	application.Preferences().SetString("StockMgCl2", "50")
	application.Preferences().SetString("StockP1", "10")
	application.Preferences().SetString("StockP2", "10")
	application.Preferences().SetString("StockP3", "10")
	application.Preferences().SetString("StockP4", "10")
	application.Preferences().SetString("StockGly", "100")
	application.Preferences().SetString("StockDMSO", "100")
	application.Preferences().SetString("StockTaq", "5")
	application.Preferences().SetString("StockDNA", "10")

	bufferVal.SetText(application.Preferences().String("StockBuffer"))
	dntpsVal.SetText(application.Preferences().String("StockDNTPs"))
	mgcl2Val.SetText(application.Preferences().String("StockMgCl2"))
	primer1Val.SetText(application.Preferences().String("StockP1"))
	primer2Val.SetText(application.Preferences().String("StockP2"))
	primer3Val.SetText(application.Preferences().String("StockP3"))
	primer4Val.SetText(application.Preferences().String("StockP4"))
	glycerolVal.SetText(application.Preferences().String("StockGly"))
	dmsoVal.SetText(application.Preferences().String("StockDMSO"))
	taqVal.SetText(application.Preferences().String("StockTaq"))
	dnacVal.SetText(application.Preferences().String("StockDNA"))

	// Define default values for PCR concentrations
	application.Preferences().SetString("PCRBuffer", "1")
	application.Preferences().SetString("PCRDNTPs", "0.2")
	application.Preferences().SetString("PCRMgCl2", "1.5")
	application.Preferences().SetString("PCRP1", "0.2")
	application.Preferences().SetString("PCRP2", "0.2")
	application.Preferences().SetString("PCRP3", "0")
	application.Preferences().SetString("PCRP4", "0")
	application.Preferences().SetString("PCRGly", "0")
	application.Preferences().SetString("PCRDMSO", "0")
	application.Preferences().SetString("PCRTaq", "0.05")
	application.Preferences().SetString("PCRDNA", "5")

	bufferValMix.SetText(application.Preferences().String("PCRBuffer"))
	dntpsValMix.SetText(application.Preferences().String("PCRDNTPs"))
	mgcl2ValMix.SetText(application.Preferences().String("PCRMgCl2"))
	primer1ValMix.SetText(application.Preferences().String("PCRP1"))
	primer2ValMix.SetText(application.Preferences().String("PCRP2"))
	primer3ValMix.SetText(application.Preferences().String("PCRP3"))
	primer4ValMix.SetText(application.Preferences().String("PCRP4"))
	glycerolValMix.SetText(application.Preferences().String("PCRGly"))
	dmsoValMix.SetText(application.Preferences().String("PCRDMSO"))
	taqValMix.SetText(application.Preferences().String("PCRTaq"))
	dnacValMix.SetText(application.Preferences().String("PCRDNA"))

	// Define mix default values
	application.Preferences().SetString("MixVolume", "25")
	application.Preferences().SetString("MixReactNum", "1")

	reactionVolVal.SetText(application.Preferences().String("MixVolume"))
	reactionNumVal.SetText(application.Preferences().String("MixReactNum"))
}
