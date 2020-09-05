package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func (p *pcr) loadUI(application fyne.App) *widget.TabContainer {
	p.application = app.NewWithID("com.github.godrigos.goPCR")
	p.w = p.application.NewWindow("goPCR")
	p.w.SetIcon(icon)
	p.application.Settings().SetTheme(theme.LightTheme())

	// Define action button
	calc := widget.NewButton("Calculate", p.calculate)

	// Define widgets for labels and values of stock tab
	p.bufferLab = widget.NewLabel("Buffer:")
	p.bufferVal = widget.NewEntry()
	p.bufferValScrl = widget.NewHScrollContainer(p.bufferVal)
	p.bufferUnit = widget.NewLabel("x")

	p.dntpsLab = widget.NewLabel("DNTPs:")
	p.dntpsVal = widget.NewEntry()
	p.dntpsValScrl = widget.NewHScrollContainer(p.dntpsVal)
	p.dntpsUnit = widget.NewLabel("nmol/\u00B5L")

	p.mgcl2Lab = widget.NewLabel("MgCl\u2082:")
	p.mgcl2Val = widget.NewEntry()
	p.mgcl2ValScrl = widget.NewHScrollContainer(p.mgcl2Val)
	p.mgcl2Unit = widget.NewLabel("nmol/\u00B5L")

	p.primer1Lab = widget.NewLabel("Primer 1:")
	p.primer1Val = widget.NewEntry()
	p.primer1ValScrl = widget.NewHScrollContainer(p.primer1Val)
	p.primer1Unit = widget.NewLabel("pmol/\u00B5L")

	p.primer2Lab = widget.NewLabel("Primer 2:")
	p.primer2Val = widget.NewEntry()
	p.primer2ValScrl = widget.NewHScrollContainer(p.primer2Val)
	p.primer2Unit = widget.NewLabel("pmol/\u00B5L")

	p.primer3Lab = widget.NewLabel("Primer 3:")
	p.primer3Val = widget.NewEntry()
	p.primer3ValScrl = widget.NewHScrollContainer(p.primer3Val)
	p.primer3Unit = widget.NewLabel("pmol/\u00B5L")

	p.primer4Lab = widget.NewLabel("Primer 4:")
	p.primer4Val = widget.NewEntry()
	p.primer4ValScrl = widget.NewHScrollContainer(p.primer4Val)
	p.primer4Unit = widget.NewLabel("pmol/\u00B5L")

	p.glycerolLab = widget.NewLabel("Glycerol:")
	p.glycerolVal = widget.NewEntry()
	p.glycerolValScrl = widget.NewHScrollContainer(p.glycerolVal)
	p.glycerolUnit = widget.NewLabel("%")

	p.dmsoLab = widget.NewLabel("DMSO:")
	p.dmsoVal = widget.NewEntry()
	p.dmsoValScrl = widget.NewHScrollContainer(p.dmsoVal)
	p.dmsoUnit = widget.NewLabel("%")

	p.taqLab = widget.NewLabel("Taq:")
	p.taqVal = widget.NewEntry()
	p.taqValScrl = widget.NewHScrollContainer(p.taqVal)
	p.taqUnit = widget.NewLabel("U/\u00B5L")

	p.dnacLab = widget.NewLabel("DNA:")
	p.dnacVal = widget.NewEntry()
	p.dnacValScrl = widget.NewHScrollContainer(p.dnacVal)
	p.dnacUnit = widget.NewLabel("ng/\u00B5L")

	stockTab := fyne.NewContainerWithLayout(layout.NewGridLayout(3),
		p.bufferLab, p.bufferValScrl, p.bufferUnit,
		p.dntpsLab, p.dntpsValScrl, p.dntpsUnit,
		p.mgcl2Lab, p.mgcl2ValScrl, p.mgcl2Unit,
		p.primer1Lab, p.primer1ValScrl, p.primer1Unit,
		p.primer2Lab, p.primer2ValScrl, p.primer2Unit,
		p.primer3Lab, p.primer3ValScrl, p.primer3Unit,
		p.primer4Lab, p.primer4ValScrl, p.primer4Unit,
		p.glycerolLab, p.glycerolValScrl, p.glycerolUnit,
		p.dmsoLab, p.dmsoValScrl, p.dmsoUnit,
		p.taqLab, p.taqValScrl, p.taqUnit,
		p.dnacLab, p.dnacValScrl, p.dnacUnit)
	stockGroup := widget.NewTabItem("Stock", stockTab)

	// Define widgets for labels and values of PCR tab
	p.bufferLabMix = widget.NewLabel("Buffer:")
	p.bufferValMix = widget.NewEntry()
	p.bufferValMixScrl = widget.NewHScrollContainer(p.bufferValMix)
	p.bufferUnitMix = widget.NewLabel("x")

	p.dntpsLabMix = widget.NewLabel("DNTPs:")
	p.dntpsValMix = widget.NewEntry()
	p.dntpsValMixScrl = widget.NewHScrollContainer(p.dntpsValMix)
	p.dntpsUnitMix = widget.NewLabel("nmol/\u00B5L")

	p.mgcl2LabMix = widget.NewLabel("MgCl\u2082:")
	p.mgcl2ValMix = widget.NewEntry()
	p.mgcl2ValMixScrl = widget.NewHScrollContainer(p.mgcl2ValMix)
	p.mgcl2UnitMix = widget.NewLabel("nmol/\u00B5L")

	p.primer1LabMix = widget.NewLabel("Primer 1:")
	p.primer1ValMix = widget.NewEntry()
	p.primer1ValMixScrl = widget.NewHScrollContainer(p.primer1ValMix)
	p.primer1UnitMix = widget.NewLabel("pmol/\u00B5L")

	p.primer2LabMix = widget.NewLabel("Primer 2:")
	p.primer2ValMix = widget.NewEntry()
	p.primer2ValMixScrl = widget.NewHScrollContainer(p.primer2ValMix)
	p.primer2UnitMix = widget.NewLabel("pmol/\u00B5L")

	p.primer3LabMix = widget.NewLabel("Primer 3:")
	p.primer3ValMix = widget.NewEntry()
	p.primer3ValMixScrl = widget.NewHScrollContainer(p.primer3ValMix)
	p.primer3UnitMix = widget.NewLabel("pmol/\u00B5L")

	p.primer4LabMix = widget.NewLabel("Primer 4:")
	p.primer4ValMix = widget.NewEntry()
	p.primer4ValMixScrl = widget.NewHScrollContainer(p.primer4ValMix)
	p.primer4UnitMix = widget.NewLabel("pmol/\u00B5L")

	p.glycerolLabMix = widget.NewLabel("Glycerol:")
	p.glycerolValMix = widget.NewEntry()
	p.glycerolValMixScrl = widget.NewHScrollContainer(p.glycerolValMix)
	p.glycerolUnitMix = widget.NewLabel("%")

	p.dmsoLabMix = widget.NewLabel("DMSO:")
	p.dmsoValMix = widget.NewEntry()
	p.dmsoValMixScrl = widget.NewHScrollContainer(p.dmsoValMix)
	p.dmsoUnitMix = widget.NewLabel("%")

	p.taqLabMix = widget.NewLabel("Taq:")
	p.taqValMix = widget.NewEntry()
	p.taqValMixScrl = widget.NewHScrollContainer(p.taqValMix)
	p.taqUnitMix = widget.NewLabel("U/\u00B5L")

	p.dnacLabMix = widget.NewLabel("DNA:")
	p.dnacValMix = widget.NewEntry()
	p.dnacValMixScrl = widget.NewHScrollContainer(p.dnacValMix)
	p.dnacUnitMix = widget.NewLabel("ng")

	pcrTab := fyne.NewContainerWithLayout(layout.NewGridLayout(3),
		p.bufferLabMix, p.bufferValMixScrl, p.bufferUnitMix,
		p.dntpsLabMix, p.dntpsValMixScrl, p.dntpsUnitMix,
		p.mgcl2LabMix, p.mgcl2ValMixScrl, p.mgcl2UnitMix,
		p.primer1LabMix, p.primer1ValMixScrl, p.primer1UnitMix,
		p.primer2LabMix, p.primer2ValMixScrl, p.primer2UnitMix,
		p.primer3LabMix, p.primer3ValMixScrl, p.primer3UnitMix,
		p.primer4LabMix, p.primer4ValMixScrl, p.primer4UnitMix,
		p.glycerolLabMix, p.glycerolValMixScrl, p.glycerolUnitMix,
		p.dmsoLabMix, p.dmsoValMixScrl, p.dmsoUnitMix,
		p.taqLabMix, p.taqValMixScrl, p.taqUnitMix,
		p.dnacLabMix, p.dnacValMixScrl, p.dnacUnitMix)
	pcrGroup := widget.NewTabItem("PCR", pcrTab)

	// Define widgets for labels and values of mix tab
	p.reactionVolLab = widget.NewLabel("Volume:")
	p.reactionVolVal = widget.NewEntry()
	p.reactionVolValScrl = widget.NewHScrollContainer(p.reactionVolVal)
	p.reactionVolUnit = widget.NewLabel("\u00B5L")

	p.reactionNumLab = widget.NewLabel("Reactions:")
	p.reactionNumVal = widget.NewEntry()
	p.reactionNumValScrl = widget.NewHScrollContainer(p.reactionNumVal)
	p.reactionNumUnit = widget.NewLabel("x")

	mix := fyne.NewContainerWithLayout(layout.NewGridLayout(3),
		p.reactionVolLab, p.reactionVolValScrl, p.reactionVolUnit,
		p.reactionNumLab, p.reactionNumValScrl, p.reactionNumUnit)
	mixTab := fyne.NewContainerWithLayout(layout.NewBorderLayout(mix,
		calc, nil, nil), mix, calc)
	mixGroup := widget.NewTabItem("Mix", mixTab)

	// Define widgets for setting tab
	pcrSetLabel := widget.NewLabelWithStyle("PCR Values:",
		fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	p.customPCR = widget.NewSelectEntry([]string{"Buffer",
		"DNTPs", "MgCl₂", "Primer 1", "Primer 2", "Primer 3", "Primer 4",
		"Glycerol", "DMSO", "Taq", "DNA"})
	p.customPCR.Disable()
	p.selectPCRVal = widget.NewEntry()
	p.selectPCRScrl = widget.NewHScrollContainer(p.selectPCRVal)
	p.selectPCRBtn = widget.NewButton("Save", p.saveCustomPCR)
	customPCRSet := fyne.NewContainerWithLayout(layout.NewGridLayout(3),
		p.customPCR, p.selectPCRScrl, p.selectPCRBtn)
	pcrSet := fyne.NewContainerWithLayout(layout.NewBorderLayout(customPCRSet,
		nil, nil, nil), customPCRSet)

	mixSetLabel := widget.NewLabelWithStyle("Mix Values:",
		fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	p.customMix = widget.NewSelectEntry([]string{"Volume",
		"Reactions"})
	p.customMix.Disable()
	p.selectMixVal = widget.NewEntry()
	p.selectMixScrl = widget.NewHScrollContainer(p.selectMixVal)
	p.selectMixBtn = widget.NewButton("Save", p.saveCustomMix)
	customMixSet := fyne.NewContainerWithLayout(layout.NewGridLayout(3),
		p.customMix, p.selectMixScrl, p.selectMixBtn)
	mixSet := fyne.NewContainerWithLayout(layout.NewBorderLayout(customMixSet,
		nil, nil, nil), customMixSet)

	stockSetLabel := widget.NewLabelWithStyle("Stock Values:",
		fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	p.customStock = widget.NewSelectEntry([]string{"Buffer",
		"DNTPs", "MgCl₂", "Primer 1", "Primer 2", "Primer 3", "Primer 4",
		"Glycerol", "DMSO", "Taq", "DNA"})
	p.customStock.Disable()
	p.selectStockVal = widget.NewEntry()
	p.selectStockScrl = widget.NewHScrollContainer(p.selectStockVal)
	p.selectStockBtn = widget.NewButton("Save", p.saveCustomStock)
	customStockSet := fyne.NewContainerWithLayout(layout.NewGridLayout(3),
		p.customStock, p.selectStockScrl, p.selectStockBtn)
	stockSet := fyne.NewContainerWithLayout(layout.NewBorderLayout(
		customStockSet, nil, nil, nil), customStockSet)

	warning := widget.NewLabelWithStyle("All saved changes will take"+
		" effect\nafter application restart!",
		fyne.TextAlignCenter, fyne.TextStyle{})

	set := fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		pcrSetLabel, pcrSet, mixSetLabel, mixSet, stockSetLabel, stockSet,
		warning)
	defaultBtn := widget.NewButton("Restore Defaults", p.restore)

	setGroup := fyne.NewContainerWithLayout(layout.NewBorderLayout(
		set, defaultBtn, nil, nil), set, defaultBtn)

	settingsTab := widget.NewTabItemWithIcon("", theme.SettingsIcon(),
		setGroup)

	// Create tab widget
	tabs := widget.NewTabContainer(pcrGroup, mixGroup, stockGroup, settingsTab)

	return tabs
}
