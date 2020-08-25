package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var (
	err error

	application = app.NewWithID("com.github.godrigos.gopcr")
	w           = application.NewWindow("goPCR")

	// Define widget for labels and values for stock tab
	bufferLab     = widget.NewLabel("Buffer:")
	bufferVal     = widget.NewEntry()
	bufferValScrl = widget.NewHScrollContainer(bufferVal)
	bufferUnit    = widget.NewLabel("x")

	dntpsLab     = widget.NewLabel("DNTPs:")
	dntpsVal     = widget.NewEntry()
	dntpsValScrl = widget.NewHScrollContainer(dntpsVal)
	dntpsUnit    = widget.NewLabel("nmol/\u00B5L")

	mgcl2Lab     = widget.NewLabel("MgCl\u2082:")
	mgcl2Val     = widget.NewEntry()
	mgcl2ValScrl = widget.NewHScrollContainer(mgcl2Val)
	mgcl2Unit    = widget.NewLabel("nmol/\u00B5L")

	primer1Lab     = widget.NewLabel("Primer 1:")
	primer1Val     = widget.NewEntry()
	primer1ValScrl = widget.NewHScrollContainer(primer1Val)
	primer1Unit    = widget.NewLabel("pmol/\u00B5L")

	primer2Lab     = widget.NewLabel("Primer 2:")
	primer2Val     = widget.NewEntry()
	primer2ValScrl = widget.NewHScrollContainer(primer2Val)
	primer2Unit    = widget.NewLabel("pmol/\u00B5L")

	primer3Lab     = widget.NewLabel("Primer 3:")
	primer3Val     = widget.NewEntry()
	primer3ValScrl = widget.NewHScrollContainer(primer3Val)
	primer3Unit    = widget.NewLabel("pmol/\u00B5L")

	primer4Lab     = widget.NewLabel("Primer 4:")
	primer4Val     = widget.NewEntry()
	primer4ValScrl = widget.NewHScrollContainer(primer4Val)
	primer4Unit    = widget.NewLabel("pmol/\u00B5L")

	glycerolLab     = widget.NewLabel("Glycerol:")
	glycerolVal     = widget.NewEntry()
	glycerolValScrl = widget.NewHScrollContainer(glycerolVal)
	glycerolUnit    = widget.NewLabel("%")

	dmsoLab     = widget.NewLabel("DMSO:")
	dmsoVal     = widget.NewEntry()
	dmsoValScrl = widget.NewHScrollContainer(dmsoVal)
	dmsoUnit    = widget.NewLabel("%")

	taqLab     = widget.NewLabel("Taq:")
	taqVal     = widget.NewEntry()
	taqValScrl = widget.NewHScrollContainer(taqVal)
	taqUnit    = widget.NewLabel("U/\u00B5L")

	dnacLab     = widget.NewLabel("DNA:")
	dnacVal     = widget.NewEntry()
	dnacValScrl = widget.NewHScrollContainer(dnacVal)
	dnacUnit    = widget.NewLabel("ng/\u00B5L")

	// Define widget for labels and values for PCR tab
	bufferLabMix     = widget.NewLabel("Buffer:")
	bufferValMix     = widget.NewEntry()
	bufferValMixScrl = widget.NewHScrollContainer(bufferValMix)
	bufferUnitMix    = widget.NewLabel("x")

	dntpsLabMix     = widget.NewLabel("DNTPs:")
	dntpsValMix     = widget.NewEntry()
	dntpsValMixScrl = widget.NewHScrollContainer(dntpsValMix)
	dntpsUnitMix    = widget.NewLabel("nmol/\u00B5L")

	mgcl2LabMix     = widget.NewLabel("MgCl\u2082:")
	mgcl2ValMix     = widget.NewEntry()
	mgcl2ValMixScrl = widget.NewHScrollContainer(mgcl2ValMix)
	mgcl2UnitMix    = widget.NewLabel("nmol/\u00B5L")

	primer1LabMix     = widget.NewLabel("Primer 1:")
	primer1ValMix     = widget.NewEntry()
	primer1ValMixScrl = widget.NewHScrollContainer(primer1ValMix)
	primer1UnitMix    = widget.NewLabel("pmol/\u00B5L")

	primer2LabMix     = widget.NewLabel("Primer 2:")
	primer2ValMix     = widget.NewEntry()
	primer2ValMixScrl = widget.NewHScrollContainer(primer2ValMix)
	primer2UnitMix    = widget.NewLabel("pmol/\u00B5L")

	primer3LabMix     = widget.NewLabel("Primer 3:")
	primer3ValMix     = widget.NewEntry()
	primer3ValMixScrl = widget.NewHScrollContainer(primer3ValMix)
	primer3UnitMix    = widget.NewLabel("pmol/\u00B5L")

	primer4LabMix     = widget.NewLabel("Primer 4:")
	primer4ValMix     = widget.NewEntry()
	primer4ValMixScrl = widget.NewHScrollContainer(primer4ValMix)
	primer4UnitMix    = widget.NewLabel("pmol/\u00B5L")

	glycerolLabMix     = widget.NewLabel("Glycerol:")
	glycerolValMix     = widget.NewEntry()
	glycerolValMixScrl = widget.NewHScrollContainer(glycerolValMix)
	glycerolUnitMix    = widget.NewLabel("%")

	dmsoLabMix     = widget.NewLabel("DMSO:")
	dmsoValMix     = widget.NewEntry()
	dmsoValMixScrl = widget.NewHScrollContainer(dmsoValMix)
	dmsoUnitMix    = widget.NewLabel("%")

	taqLabMix     = widget.NewLabel("Taq:")
	taqValMix     = widget.NewEntry()
	taqValMixScrl = widget.NewHScrollContainer(taqValMix)
	taqUnitMix    = widget.NewLabel("U/\u00B5L")

	dnacLabMix     = widget.NewLabel("DNA:")
	dnacValMix     = widget.NewEntry()
	dnacValMixScrl = widget.NewHScrollContainer(dnacValMix)
	dnacUnitMix    = widget.NewLabel("ng")

	// Define widget for labels and values for mix tab
	reactionVolLab     = widget.NewLabel("Volume:")
	reactionVolVal     = widget.NewEntry()
	reactionVolValScrl = widget.NewHScrollContainer(reactionVolVal)
	reactionVolUnit    = widget.NewLabel("\u00B5L")

	reactionNumLab     = widget.NewLabel("Reactions:")
	reactionNumVal     = widget.NewEntry()
	reactionNumValScrl = widget.NewHScrollContainer(reactionNumVal)
	reactionNumUnit    = widget.NewLabel("x")

	// Define results variables
	reactionVol, reactionNum       float64
	bufferPCR, bufferStock, buffer float64
	dntpsPCR, dntpsStock, dntps    float64
	mgcl2PCR, mgcl2Stock, mgcl2    float64
	p1PCR, p1Stock, p1             float64
	p2PCR, p2Stock, p2             float64
	p3PCR, p3Stock, p3             float64
	p4PCR, p4Stock, p4             float64
	glyPCR, glyStock, glycerol     float64
	dmsoPCR, dmsoStock, dmsoFinal  float64
	taqPCR, taqStock, taq          float64
	dnaPCR, dnaStock, dna          float64
	h2o, mixVolume                 float64
	p2p3p4, gly, dmso, str         string
)
