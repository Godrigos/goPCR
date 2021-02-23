package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (p *pcr) loadUI(application fyne.App) *container.AppTabs {
	p.application = app.NewWithID("com.github.godrigos.goPCR")
	p.w = p.application.NewWindow("goPCR")
	p.w.SetIcon(resourceIconPng)
	p.setTheme()

	// Define action button
	calc := widget.NewButton("Calculate", p.calculate)

	// Define widgets for labels and values of stock tab
	p.bufferLab = widget.NewLabel("Buffer:")
	p.bufferVal = newFloatEntry()
	p.bufferUnit = widget.NewLabel("x")

	p.dntpsLab = widget.NewLabel("dNTPs:")
	p.dntpsVal = newFloatEntry()
	p.dntpsUnit = widget.NewLabel("nmol/\u00B5L")

	p.mgcl2Lab = widget.NewLabel("MgCl\u2082:")
	p.mgcl2Val = newFloatEntry()
	p.mgcl2Unit = widget.NewLabel("nmol/\u00B5L")

	p.primer1Lab = widget.NewLabel("Primer 1:")
	p.primer1Val = newFloatEntry()
	p.primer1Unit = widget.NewLabel("pmol/\u00B5L")

	p.primer2Lab = widget.NewLabel("Primer 2:")
	p.primer2Val = newFloatEntry()
	p.primer2Unit = widget.NewLabel("pmol/\u00B5L")

	p.primer3Lab = widget.NewLabel("Primer 3:")
	p.primer3Val = newFloatEntry()
	p.primer3Unit = widget.NewLabel("pmol/\u00B5L")

	p.primer4Lab = widget.NewLabel("Primer 4:")
	p.primer4Val = newFloatEntry()
	p.primer4Unit = widget.NewLabel("pmol/\u00B5L")

	p.glycerolLab = widget.NewLabel("Glycerol:")
	p.glycerolVal = newFloatEntry()
	p.glycerolUnit = widget.NewLabel("%")

	p.dmsoLab = widget.NewLabel("DMSO:")
	p.dmsoVal = newFloatEntry()
	p.dmsoUnit = widget.NewLabel("%")

	p.taqLab = widget.NewLabel("Taq:")
	p.taqVal = newFloatEntry()
	p.taqUnit = widget.NewLabel("U/\u00B5L")

	p.dnacLab = widget.NewLabel("DNA:")
	p.dnacVal = newFloatEntry()
	p.dnacUnit = widget.NewLabel("ng/\u00B5L")

	stockTab := fyne.NewContainerWithLayout(layout.NewGridLayout(3),
		p.bufferLab, p.bufferVal, p.bufferUnit,
		p.dntpsLab, p.dntpsVal, p.dntpsUnit,
		p.mgcl2Lab, p.mgcl2Val, p.mgcl2Unit,
		p.primer1Lab, p.primer1Val, p.primer1Unit,
		p.primer2Lab, p.primer2Val, p.primer2Unit,
		p.primer3Lab, p.primer3Val, p.primer3Unit,
		p.primer4Lab, p.primer4Val, p.primer4Unit,
		p.glycerolLab, p.glycerolVal, p.glycerolUnit,
		p.dmsoLab, p.dmsoVal, p.dmsoUnit,
		p.taqLab, p.taqVal, p.taqUnit,
		p.dnacLab, p.dnacVal, p.dnacUnit)
	stockGroup := container.NewTabItem("Stock", stockTab)

	// Define widgets for labels and values of PCR tab
	p.bufferLabMix = widget.NewLabel("Buffer:")
	p.bufferValMix = newFloatEntry()
	p.bufferUnitMix = widget.NewLabel("x")

	p.dntpsLabMix = widget.NewLabel("dNTPs:")
	p.dntpsValMix = newFloatEntry()
	p.dntpsUnitMix = widget.NewLabel("nmol/\u00B5L")

	p.mgcl2LabMix = widget.NewLabel("MgCl\u2082:")
	p.mgcl2ValMix = newFloatEntry()
	p.mgcl2UnitMix = widget.NewLabel("nmol/\u00B5L")

	p.primer1LabMix = widget.NewLabel("Primer 1:")
	p.primer1ValMix = newFloatEntry()
	p.primer1UnitMix = widget.NewLabel("pmol/\u00B5L")

	p.primer2LabMix = widget.NewLabel("Primer 2:")
	p.primer2ValMix = newFloatEntry()
	p.primer2UnitMix = widget.NewLabel("pmol/\u00B5L")

	p.primer3LabMix = widget.NewLabel("Primer 3:")
	p.primer3ValMix = newFloatEntry()
	p.primer3UnitMix = widget.NewLabel("pmol/\u00B5L")

	p.primer4LabMix = widget.NewLabel("Primer 4:")
	p.primer4ValMix = newFloatEntry()
	p.primer4UnitMix = widget.NewLabel("pmol/\u00B5L")

	p.glycerolLabMix = widget.NewLabel("Glycerol:")
	p.glycerolValMix = newFloatEntry()
	p.glycerolUnitMix = widget.NewLabel("%")

	p.dmsoLabMix = widget.NewLabel("DMSO:")
	p.dmsoValMix = newFloatEntry()
	p.dmsoUnitMix = widget.NewLabel("%")

	p.taqLabMix = widget.NewLabel("Taq:")
	p.taqValMix = newFloatEntry()
	p.taqUnitMix = widget.NewLabel("U/\u00B5L")

	p.dnacLabMix = widget.NewLabel("DNA:")
	p.dnacValMix = newFloatEntry()
	p.dnacUnitMix = widget.NewLabel("ng")

	pcrTab := fyne.NewContainerWithLayout(layout.NewGridLayout(3),
		p.bufferLabMix, p.bufferValMix, p.bufferUnitMix,
		p.dntpsLabMix, p.dntpsValMix, p.dntpsUnitMix,
		p.mgcl2LabMix, p.mgcl2ValMix, p.mgcl2UnitMix,
		p.primer1LabMix, p.primer1ValMix, p.primer1UnitMix,
		p.primer2LabMix, p.primer2ValMix, p.primer2UnitMix,
		p.primer3LabMix, p.primer3ValMix, p.primer3UnitMix,
		p.primer4LabMix, p.primer4ValMix, p.primer4UnitMix,
		p.glycerolLabMix, p.glycerolValMix, p.glycerolUnitMix,
		p.dmsoLabMix, p.dmsoValMix, p.dmsoUnitMix,
		p.taqLabMix, p.taqValMix, p.taqUnitMix,
		p.dnacLabMix, p.dnacValMix, p.dnacUnitMix)
	pcrGroup := container.NewTabItem("PCR", pcrTab)

	// Define widgets for labels and values of mix tab
	p.reactionVolLab = widget.NewLabel("Volume:")
	p.reactionVolVal = newFloatEntry()
	p.reactionVolUnit = widget.NewLabel("\u00B5L")

	p.reactionNumLab = widget.NewLabel("Reactions:")
	p.reactionNumVal = newFloatEntry()
	p.reactionNumUnit = widget.NewLabel("x")

	mix := fyne.NewContainerWithLayout(layout.NewGridLayout(3),
		p.reactionVolLab, p.reactionVolVal, p.reactionVolUnit,
		p.reactionNumLab, p.reactionNumVal, p.reactionNumUnit)
	mixTab := fyne.NewContainerWithLayout(layout.NewBorderLayout(mix,
		calc, nil, nil), mix, calc)
	mixGroup := container.NewTabItem("Mix", mixTab)

	// Define widgets for setting tab
	pcrSetLabel := widget.NewLabelWithStyle("PCR Values:",
		fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	p.customPCR = widget.NewSelectEntry([]string{"Buffer",
		"DNTPs", "MgCl₂", "Primer 1", "Primer 2", "Primer 3", "Primer 4",
		"Glycerol", "DMSO", "Taq", "DNA"})
	p.selectPCRVal = newFloatEntry()
	p.selectPCRBtn = widget.NewButton("Save", p.saveCustomPCR)
	customPCRSet := fyne.NewContainerWithLayout(layout.NewGridLayout(3),
		p.customPCR, p.selectPCRVal, p.selectPCRBtn)
	pcrSet := fyne.NewContainerWithLayout(layout.NewBorderLayout(customPCRSet,
		nil, nil, nil), customPCRSet)

	mixSetLabel := widget.NewLabelWithStyle("Mix Values:",
		fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	p.customMix = widget.NewSelectEntry([]string{"Volume",
		"Reactions"})
	p.selectMixVal = newFloatEntry()
	p.selectMixBtn = widget.NewButton("Save", p.saveCustomMix)
	customMixSet := fyne.NewContainerWithLayout(layout.NewGridLayout(3),
		p.customMix, p.selectMixVal, p.selectMixBtn)
	mixSet := fyne.NewContainerWithLayout(layout.NewBorderLayout(customMixSet,
		nil, nil, nil), customMixSet)

	stockSetLabel := widget.NewLabelWithStyle("Stock Values:",
		fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	p.customStock = widget.NewSelectEntry([]string{"Buffer",
		"DNTPs", "MgCl₂", "Primer 1", "Primer 2", "Primer 3", "Primer 4",
		"Glycerol", "DMSO", "Taq", "DNA"})
	p.selectStockVal = newFloatEntry()
	p.selectStockBtn = widget.NewButton("Save", p.saveCustomStock)
	customStockSet := fyne.NewContainerWithLayout(layout.NewGridLayout(3),
		p.customStock, p.selectStockVal, p.selectStockBtn)
	stockSet := fyne.NewContainerWithLayout(layout.NewBorderLayout(
		customStockSet, nil, nil, nil), customStockSet)

	warning := widget.NewLabelWithStyle("All saved changes will take"+
		" effect\nafter application restart!",
		fyne.TextAlignCenter, fyne.TextStyle{})

	set := fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		pcrSetLabel, pcrSet, mixSetLabel, mixSet, stockSetLabel, stockSet,
		warning)
	defaultBtn := widget.NewButton("Restore Defaults", p.restore)

	darkThemeBtn := widget.NewButton("Dark", p.darkTheme)
	lightThemeBtn := widget.NewButton("Light", p.lightTheme)
	themeBtn := fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		darkThemeBtn, lightThemeBtn)

	btns := fyne.NewContainerWithLayout(layout.NewBorderLayout(
		defaultBtn, themeBtn, nil, nil), defaultBtn, themeBtn)

	setGroup := fyne.NewContainerWithLayout(layout.NewBorderLayout(
		set, btns, nil, nil), set, btns)

	settingsTab := container.NewTabItemWithIcon("", theme.SettingsIcon(),
		setGroup)

	// Create tab widget
	tabs := container.NewAppTabs(pcrGroup, mixGroup, stockGroup, settingsTab)

	return tabs
}
