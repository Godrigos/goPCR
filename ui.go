package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func loadUI() *widget.TabContainer {

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
	var tabs = widget.NewTabContainer(mixGroup, finalGroup, stockGroup)

	return tabs
}
