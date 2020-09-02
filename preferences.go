package main

func (p *pcr) prefs() {

	// Define default values for stock solutions
	p.bufferVal.SetText(p.application.Preferences().
		StringWithFallback("StockBuffer", "10"))
	p.dntpsVal.SetText(p.application.Preferences().
		StringWithFallback("StockDNTPs", "2.5"))
	p.mgcl2Val.SetText(p.application.Preferences().
		StringWithFallback("StockMgCl2", "50"))
	p.primer1Val.SetText(p.application.Preferences().
		StringWithFallback("StockP1", "10"))
	p.primer2Val.SetText(p.application.Preferences().
		StringWithFallback("StockP2", "10"))
	p.primer3Val.SetText(p.application.Preferences().
		StringWithFallback("StockP3", "10"))
	p.primer4Val.SetText(p.application.Preferences().
		StringWithFallback("StockP4", "10"))
	p.glycerolVal.SetText(p.application.Preferences().
		StringWithFallback("StockGly", "100"))
	p.dmsoVal.SetText(p.application.Preferences().
		StringWithFallback("StockDMSO", "100"))
	p.taqVal.SetText(p.application.Preferences().
		StringWithFallback("StockTaq", "5"))
	p.dnacVal.SetText(p.application.Preferences().
		StringWithFallback("StockDNA", "10"))

	// Define default values for PCR concentrations
	p.bufferValMix.SetText(p.application.Preferences().
		StringWithFallback("PCRBuffer", "1"))
	p.dntpsValMix.SetText(p.application.Preferences().
		StringWithFallback("PCRDNTPs", "0.2"))
	p.mgcl2ValMix.SetText(p.application.Preferences().
		StringWithFallback("PCRMgCl2", "1.5"))
	p.primer1ValMix.SetText(p.application.Preferences().
		StringWithFallback("PCRP1", "0.2"))
	p.primer2ValMix.SetText(p.application.Preferences().
		StringWithFallback("PCRP2", "0.2"))
	p.primer3ValMix.SetText(p.application.Preferences().
		StringWithFallback("PCRP3", "0"))
	p.primer4ValMix.SetText(p.application.Preferences().
		StringWithFallback("PCRP4", "0"))
	p.glycerolValMix.SetText(p.application.Preferences().
		StringWithFallback("PCRGly", "0"))
	p.dmsoValMix.SetText(p.application.Preferences().
		StringWithFallback("PCRDMSO", "0"))
	p.taqValMix.SetText(p.application.Preferences().
		StringWithFallback("PCRTaq", "0.05"))
	p.dnacValMix.SetText(p.application.Preferences().
		StringWithFallback("PCRDNA", "5"))

	// Define mix default values
	p.reactionVolVal.SetText(p.application.Preferences().
		StringWithFallback("MixVolume", "25"))
	p.reactionNumVal.SetText(p.application.Preferences().
		StringWithFallback("MixReactNum", "1"))
}
