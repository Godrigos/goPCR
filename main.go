package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {

	// Import stock, pcr and mix values from json files in presets folder
	stock := stock()
	pcr := pcr()
	mix := mix()

	w.SetFixedSize(true)

	// Load icon
	icon, err := fyne.LoadResourceFromPath("./images/icon.png")
	if err != nil {
		panic(err)
	}

	// Set defaul values for stock solutions
	bufferVal.SetText(stock.Buffer)
	dntpsVal.SetText(stock.DNTPs)
	mgcl2Val.SetText(stock.MgCl2)
	primer1Val.SetText(stock.Primer1)
	primer2Val.SetText(stock.Primer2)
	primer3Val.SetText(stock.Primer3)
	primer4Val.SetText(stock.Primer4)
	glycerolVal.SetText(stock.Gly)
	dmsoVal.SetText(stock.DMSO)
	taqVal.SetText(stock.Taq)
	dnacVal.SetText(stock.DnaC)

	// Set defaul values for PCR
	bufferValMix.SetText(pcr.Buffer)
	dntpsValMix.SetText(pcr.DNTPs)
	mgcl2ValMix.SetText(pcr.MgCl2)
	primer1ValMix.SetText(pcr.Primer1)
	primer2ValMix.SetText(pcr.Primer2)
	primer3ValMix.SetText(pcr.Primer3)
	primer4ValMix.SetText(pcr.Primer4)
	glycerolValMix.SetText(pcr.Gly)
	dmsoValMix.SetText(pcr.DMSO)
	taqValMix.SetText(pcr.Taq)
	dnacValMix.SetText(pcr.DnaC)

	// Set defaul values for Mix
	reactionVolVal.SetText(mix.Volume)
	reactionNumVal.SetText(mix.Reactions)

	stockTab := fyne.NewContainerWithLayout(layout.NewGridLayout(3),
		bufferLab, bufferValScrl, bufferUnit,
		dntpsLab, dntpsValScrl, dntpsUnit,
		mgcl2Lab, mgcl2ValScrl, mgcl2Unit,
		primer1Lab, primer1ValScrl, primer1Unit,
		primer2Lab, primer2ValScrl, primer2Unit,
		primer3Lab, primer3ValScrl, primer3Unit,
		primer4Lab, primer4ValScrl, primer4Unit,
		glycerolLab, glycerolValScrl, glycerolUnit,
		dmsoLab, dmsoValScrl, dmsoUnit,
		taqLab, taqValScrl, taqUnit,
		dnacLab, dnacValScrl, dnacUnit)
	stockGroup := widget.NewTabItem("Stock", stockTab)

	mixTab := fyne.NewContainerWithLayout(layout.NewGridLayout(3),
		bufferLabMix, bufferValMixScrl, bufferUnitMix,
		dntpsLabMix, dntpsValMixScrl, dntpsUnitMix,
		mgcl2LabMix, mgcl2ValMixScrl, mgcl2UnitMix,
		primer1LabMix, primer1ValMixScrl, primer1UnitMix,
		primer2LabMix, primer2ValMixScrl, primer2UnitMix,
		primer3LabMix, primer3ValMixScrl, primer3UnitMix,
		primer4LabMix, primer4ValMixScrl, primer4UnitMix,
		glycerolLabMix, glycerolValMixScrl, glycerolUnitMix,
		dmsoLabMix, dmsoValMixScrl, dmsoUnitMix,
		taqLabMix, taqValMixScrl, taqUnitMix,
		dnacLabMix, dnacValMixScrl, dnacUnitMix)
	mixGroup := widget.NewTabItem("PCR", mixTab)

	final := fyne.NewContainerWithLayout(layout.NewGridLayout(3),
		reactionVolLab, reactionVolValScrl, reactionVolUnit,
		reactionNumLab, reactionNumValScrl, reactionNumUnit)
	finalTab := fyne.NewContainerWithLayout(layout.NewBorderLayout(final,
		calc, nil, nil), final, calc)
	finalGroup := widget.NewTabItem("Mix", finalTab)

	// Create tab widget
	var tabs = widget.NewTabContainer(stockGroup, mixGroup, finalGroup)

	w.SetIcon(icon)
	w.SetContent(tabs)
	w.ShowAndRun()
}
