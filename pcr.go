package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type pcr struct {
	// Application and main windows fields
	application fyne.App
	w           fyne.Window

	// Define fields for labels and values of stock tab
	bufferLab, dntpsLab, mgcl2Lab, primer1Lab       *widget.Label
	primer2Lab, primer3Lab, primer4Lab, glycerolLab *widget.Label
	dmsoLab, taqLab, dnacLab                        *widget.Label

	bufferVal, dntpsVal, mgcl2Val, primer1Val       *widget.Entry
	primer2Val, primer3Val, primer4Val, glycerolVal *widget.Entry
	dmsoVal, taqVal, dnacVal                        *widget.Entry

	bufferValScrl, dntpsValScrl, mgcl2ValScrl      *widget.ScrollContainer
	primer1ValScrl, primer2ValScrl, primer3ValScrl *widget.ScrollContainer
	primer4ValScrl, glycerolValScrl, dmsoValScrl   *widget.ScrollContainer
	taqValScrl, dnacValScrl                        *widget.ScrollContainer

	bufferUnit, dntpsUnit, mgcl2Unit, primer1Unit       *widget.Label
	primer2Unit, primer3Unit, primer4Unit, glycerolUnit *widget.Label
	dmsoUnit, taqUnit, dnacUnit                         *widget.Label

	// Define fields for labels and values of mix tab
	reactionVolLab, reactionNumLab         *widget.Label
	reactionVolVal, reactionNumVal         *widget.Entry
	reactionVolValScrl, reactionNumValScrl *widget.ScrollContainer
	reactionVolUnit, reactionNumUnit       *widget.Label

	// Define fields for labels and values of PCR tab
	bufferLabMix, dntpsLabMix, mgcl2LabMix, primer1LabMix       *widget.Label
	primer2LabMix, primer3LabMix, primer4LabMix, glycerolLabMix *widget.Label
	dmsoLabMix, taqLabMix, dnacLabMix                           *widget.Label

	bufferValMix, dntpsValMix, mgcl2ValMix, primer1ValMix       *widget.Entry
	primer2ValMix, primer3ValMix, primer4ValMix, glycerolValMix *widget.Entry
	dmsoValMix, taqValMix, dnacValMix                           *widget.Entry

	bufferValMixScrl, dntpsValMixScrl, mgcl2ValMixScrl *widget.ScrollContainer
	primer1ValMixScrl, primer2ValMixScrl               *widget.ScrollContainer
	primer3ValMixScrl, primer4ValMixScrl               *widget.ScrollContainer
	glycerolValMixScrl, dmsoValMixScrl                 *widget.ScrollContainer
	taqValMixScrl, dnacValMixScrl                      *widget.ScrollContainer

	bufferUnitMix, dntpsUnitMix, mgcl2UnitMix, primer1UnitMix *widget.Label
	primer2UnitMix, primer3UnitMix, primer4UnitMix            *widget.Label
	glycerolUnitMix, dmsoUnitMix, taqUnitMix, dnacUnitMix     *widget.Label

	// Define results fields
	reactionNum                    float64
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

	// Custom values fields

	customStock                                   *widget.SelectEntry
	customPCR                                     *widget.SelectEntry
	customMix                                     *widget.SelectEntry
	selectStockVal, selectPCRVal, selectMixVal    *widget.Entry
	selectStockScrl, selectPCRScrl, selectMixScrl *widget.ScrollContainer
	selectStockBtn, selectPCRBtn, selectMixBtn    *widget.Button
}
