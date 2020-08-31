package main

func (p *pcr) prefs() {

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

	p.bufferVal.SetText(p.application.Preferences().String("StockBuffer"))
	p.dntpsVal.SetText(p.application.Preferences().String("StockDNTPs"))
	p.mgcl2Val.SetText(p.application.Preferences().String("StockMgCl2"))
	p.primer1Val.SetText(p.application.Preferences().String("StockP1"))
	p.primer2Val.SetText(p.application.Preferences().String("StockP2"))
	p.primer3Val.SetText(p.application.Preferences().String("StockP3"))
	p.primer4Val.SetText(p.application.Preferences().String("StockP4"))
	p.glycerolVal.SetText(p.application.Preferences().String("StockGly"))
	p.dmsoVal.SetText(p.application.Preferences().String("StockDMSO"))
	p.taqVal.SetText(p.application.Preferences().String("StockTaq"))
	p.dnacVal.SetText(p.application.Preferences().String("StockDNA"))

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

	p.bufferValMix.SetText(p.application.Preferences().String("PCRBuffer"))
	p.dntpsValMix.SetText(p.application.Preferences().String("PCRDNTPs"))
	p.mgcl2ValMix.SetText(p.application.Preferences().String("PCRMgCl2"))
	p.primer1ValMix.SetText(p.application.Preferences().String("PCRP1"))
	p.primer2ValMix.SetText(p.application.Preferences().String("PCRP2"))
	p.primer3ValMix.SetText(p.application.Preferences().String("PCRP3"))
	p.primer4ValMix.SetText(p.application.Preferences().String("PCRP4"))
	p.glycerolValMix.SetText(p.application.Preferences().String("PCRGly"))
	p.dmsoValMix.SetText(p.application.Preferences().String("PCRDMSO"))
	p.taqValMix.SetText(p.application.Preferences().String("PCRTaq"))
	p.dnacValMix.SetText(p.application.Preferences().String("PCRDNA"))

	// Define mix default values
	p.application.Preferences().SetString("MixVolume", "25")
	p.application.Preferences().SetString("MixReactNum", "1")

	p.reactionVolVal.SetText(p.application.Preferences().String("MixVolume"))
	p.reactionNumVal.SetText(p.application.Preferences().String("MixReactNum"))
}
