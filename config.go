package main

import (
	"os"
	"path"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var warnPosition = fyne.NewPos(10, 200)
var err error
var myDir, _ = os.Executable()
var execDir = path.Dir(myDir)

var application = app.NewWithID("com.godrigos.pcrcalc")
var w = application.NewWindow("PCRCalc")

// Define widget for labels and values for stock tab
var bufferLab = widget.NewLabel("Buffer:")
var bufferVal = widget.NewEntry()
var bufferValScrl = widget.NewHScrollContainer(bufferVal)
var bufferUnit = widget.NewLabel("x")

var dntpsLab = widget.NewLabel("DNTPs:")
var dntpsVal = widget.NewEntry()
var dntpsValScrl = widget.NewHScrollContainer(dntpsVal)
var dntpsUnit = widget.NewLabel("nmol/\u00B5L")

var mgcl2Lab = widget.NewLabel("MgCl\u2082:")
var mgcl2Val = widget.NewEntry()
var mgcl2ValScrl = widget.NewHScrollContainer(mgcl2Val)
var mgcl2Unit = widget.NewLabel("nmol/\u00B5L")

var primer1Lab = widget.NewLabel("Primer 1:")
var primer1Val = widget.NewEntry()
var primer1ValScrl = widget.NewHScrollContainer(primer1Val)
var primer1Unit = widget.NewLabel("pmol/\u00B5L")

var primer2Lab = widget.NewLabel("Primer 2:")
var primer2Val = widget.NewEntry()
var primer2ValScrl = widget.NewHScrollContainer(primer2Val)
var primer2Unit = widget.NewLabel("pmol/\u00B5L")

var primer3Lab = widget.NewLabel("Primer 3:")
var primer3Val = widget.NewEntry()
var primer3ValScrl = widget.NewHScrollContainer(primer3Val)
var primer3Unit = widget.NewLabel("pmol/\u00B5L")

var primer4Lab = widget.NewLabel("Primer 4:")
var primer4Val = widget.NewEntry()
var primer4ValScrl = widget.NewHScrollContainer(primer4Val)
var primer4Unit = widget.NewLabel("pmol/\u00B5L")

var glycerolLab = widget.NewLabel("Glycerol:")
var glycerolVal = widget.NewEntry()
var glycerolValScrl = widget.NewHScrollContainer(glycerolVal)
var glycerolUnit = widget.NewLabel("%")

var dmsoLab = widget.NewLabel("DMSO:")
var dmsoVal = widget.NewEntry()
var dmsoValScrl = widget.NewHScrollContainer(dmsoVal)
var dmsoUnit = widget.NewLabel("%")

var taqLab = widget.NewLabel("Taq:")
var taqVal = widget.NewEntry()
var taqValScrl = widget.NewHScrollContainer(taqVal)
var taqUnit = widget.NewLabel("U/\u00B5L")

var dnacLab = widget.NewLabel("DNA:")
var dnacVal = widget.NewEntry()
var dnacValScrl = widget.NewHScrollContainer(dnacVal)
var dnacUnit = widget.NewLabel("ng/\u00B5L")

// Define widget for labels and values for PCR tab
var bufferLabMix = widget.NewLabel("Buffer:")
var bufferValMix = widget.NewEntry()
var bufferValMixScrl = widget.NewHScrollContainer(bufferValMix)
var bufferUnitMix = widget.NewLabel("x")

var dntpsLabMix = widget.NewLabel("DNTPs:")
var dntpsValMix = widget.NewEntry()
var dntpsValMixScrl = widget.NewHScrollContainer(dntpsValMix)
var dntpsUnitMix = widget.NewLabel("nmol/\u00B5L")

var mgcl2LabMix = widget.NewLabel("MgCl\u2082:")
var mgcl2ValMix = widget.NewEntry()
var mgcl2ValMixScrl = widget.NewHScrollContainer(mgcl2ValMix)
var mgcl2UnitMix = widget.NewLabel("nmol/\u00B5L")

var primer1LabMix = widget.NewLabel("Primer 1:")
var primer1ValMix = widget.NewEntry()
var primer1ValMixScrl = widget.NewHScrollContainer(primer1ValMix)
var primer1UnitMix = widget.NewLabel("pmol/\u00B5L")

var primer2LabMix = widget.NewLabel("Primer 2:")
var primer2ValMix = widget.NewEntry()
var primer2ValMixScrl = widget.NewHScrollContainer(primer2ValMix)
var primer2UnitMix = widget.NewLabel("pmol/\u00B5L")

var primer3LabMix = widget.NewLabel("Primer 3:")
var primer3ValMix = widget.NewEntry()
var primer3ValMixScrl = widget.NewHScrollContainer(primer3ValMix)
var primer3UnitMix = widget.NewLabel("pmol/\u00B5L")

var primer4LabMix = widget.NewLabel("Primer 4:")
var primer4ValMix = widget.NewEntry()
var primer4ValMixScrl = widget.NewHScrollContainer(primer4ValMix)
var primer4UnitMix = widget.NewLabel("pmol/\u00B5L")

var glycerolLabMix = widget.NewLabel("Glycerol:")
var glycerolValMix = widget.NewEntry()
var glycerolValMixScrl = widget.NewHScrollContainer(glycerolValMix)
var glycerolUnitMix = widget.NewLabel("%")

var dmsoLabMix = widget.NewLabel("DMSO:")
var dmsoValMix = widget.NewEntry()
var dmsoValMixScrl = widget.NewHScrollContainer(dmsoValMix)
var dmsoUnitMix = widget.NewLabel("%")

var taqLabMix = widget.NewLabel("Taq:")
var taqValMix = widget.NewEntry()
var taqValMixScrl = widget.NewHScrollContainer(taqValMix)
var taqUnitMix = widget.NewLabel("U/\u00B5L")

var dnacLabMix = widget.NewLabel("DNA:")
var dnacValMix = widget.NewEntry()
var dnacValMixScrl = widget.NewHScrollContainer(dnacValMix)
var dnacUnitMix = widget.NewLabel("ng")

// Define widget for labels and values for mix tab
var reactionVolLab = widget.NewLabel("Volume:")
var reactionVolVal = widget.NewEntry()
var reactionVolValScrl = widget.NewHScrollContainer(reactionVolVal)
var reactionVolUnit = widget.NewLabel("\u00B5L")

var reactionNumLab = widget.NewLabel("Reactions:")
var reactionNumVal = widget.NewEntry()
var reactionNumValScrl = widget.NewHScrollContainer(reactionNumVal)
var reactionNumUnit = widget.NewLabel("x")

// Define results variables
var reactionVol, reactionNum float64
var bufferPCR, bufferStock, buffer float64
var dntpsPCR, dntpsStock, dntps float64
var mgcl2PCR, mgcl2Stock, mgcl2 float64
var p1PCR, p1Stock, p1 float64
var p2PCR, p2Stock, p2 float64
var p3PCR, p3Stock, p3 float64
var p4PCR, p4Stock, p4 float64
var glyPCR, glyStock, glycerol float64
var dmsoPCR, dmsoStock, dmsoFinal float64
var taqPCR, taqStock, taq float64
var dnaPCR, dnaStock, dna float64
var h2o, mixVolume float64
var p2p3p4, gly, dmso, str string

// Define action button
var calc = widget.NewButton("Calculate", calculate)

// Solutions stores the default values of stock solutions
type Solutions struct {
	Buffer  string `json:"buffer"`
	DMSO    string `json:"dmso"`
	DnaC    string `json:"dna_c"`
	DNTPs   string `json:"dntps"`
	Gly     string `json:"gly"`
	MgCl2   string `json:"mgcl2"`
	Primer1 string `json:"primer1"`
	Primer2 string `json:"primer2"`
	Primer3 string `json:"primer3"`
	Primer4 string `json:"primer4"`
	Taq     string `json:"taq"`
}

// Mix stores the default values of mix tabS
type Mix struct {
	Volume    string `json:"volume"`
	Reactions string `json:"reactions"`
}
