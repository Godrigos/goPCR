package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func calculate() {
	// Store error evaluation
	var errors []bool

	// Reactions volume and number convertion to float64 from string
	reactionVol, err = strconv.ParseFloat(reactionVolVal.Text, 64)
	errors = append(errors, checkFloat(err, "reaction volume"))
	reactionNum, err = strconv.ParseFloat(reactionNumVal.Text, 64)
	errors = append(errors, checkFloat(err, "reaction number"))

	// Other reagents conversions and calculation
	// Buffer
	bufferPCR, err = strconv.ParseFloat(bufferValMix.Text, 64)
	errors = append(errors, checkFloat(err, "PCR buffer"))
	bufferStock, err = strconv.ParseFloat(bufferVal.Text, 64)
	errors = append(errors, checkFloat(err, "Stock buffer"))
	buffer = (bufferPCR * reactionVol / bufferStock) * reactionNum
	// DNTPs
	dntpsPCR, err = strconv.ParseFloat(dntpsValMix.Text, 64)
	errors = append(errors, checkFloat(err, "PCR DNTPs"))
	dntpsStock, err = strconv.ParseFloat(dntpsVal.Text, 64)
	errors = append(errors, checkFloat(err, "Stock DNTPs"))
	dntps = (dntpsPCR * reactionVol / dntpsStock) * reactionNum
	// MgCl2
	mgcl2PCR, err = strconv.ParseFloat(mgcl2ValMix.Text, 64)
	errors = append(errors, checkFloat(err, "PCR MgCl₂"))
	mgcl2Stock, err = strconv.ParseFloat(mgcl2Val.Text, 64)
	errors = append(errors, checkFloat(err, "Stock MgCl₂"))
	mgcl2 = (mgcl2PCR * reactionVol / mgcl2Stock) * reactionNum
	//Primers
	p1PCR, err = strconv.ParseFloat(primer1ValMix.Text, 64)
	errors = append(errors, checkFloat(err, "PCR primer 1"))
	p1Stock, err = strconv.ParseFloat(primer1Val.Text, 64)
	errors = append(errors, checkFloat(err, "Stock primer 1"))
	p1 = (p1PCR * reactionVol / p1Stock) * reactionNum
	p2PCR, err = strconv.ParseFloat(primer2ValMix.Text, 64)
	errors = append(errors, checkFloat(err, "PCR primer 2"))
	p2Stock, err = strconv.ParseFloat(primer2Val.Text, 64)
	errors = append(errors, checkFloat(err, "Stock primer 2"))
	p2 = (p2PCR * reactionVol / p2Stock) * reactionNum
	p3PCR, err = strconv.ParseFloat(primer3ValMix.Text, 64)
	errors = append(errors, checkFloat(err, "PCR primer 3"))
	p3Stock, err = strconv.ParseFloat(primer3Val.Text, 64)
	errors = append(errors, checkFloat(err, "Stock primer 3"))
	p3 = (p3PCR * reactionVol / p3Stock) * reactionNum
	p4PCR, err = strconv.ParseFloat(primer4ValMix.Text, 64)
	errors = append(errors, checkFloat(err, "PCR primer 4"))
	p4Stock, err = strconv.ParseFloat(primer4Val.Text, 64)
	errors = append(errors, checkFloat(err, "Stock primer 4"))
	p4 = (p4PCR * reactionVol / p4Stock) * reactionNum
	// Glycerol
	glyPCR, err = strconv.ParseFloat(glycerolValMix.Text, 64)
	errors = append(errors, checkFloat(err, "PCR glycerol"))
	glyStock, err = strconv.ParseFloat(glycerolVal.Text, 64)
	errors = append(errors, checkFloat(err, "Stock glycerol"))
	glycerol = (glyPCR * reactionVol / glyStock) * reactionNum
	// DMSO
	dmsoPCR, err = strconv.ParseFloat(dmsoValMix.Text, 64)
	errors = append(errors, checkFloat(err, "PCR DMSO"))
	dmsoStock, err = strconv.ParseFloat(dmsoVal.Text, 64)
	errors = append(errors, checkFloat(err, "Stock DMSO"))
	dmsoFinal = (dmsoPCR * reactionVol / dmsoStock) * reactionNum
	// Taq
	taqPCR, err = strconv.ParseFloat(taqValMix.Text, 64)
	errors = append(errors, checkFloat(err, "PCR DMSO"))
	taqStock, err = strconv.ParseFloat(taqVal.Text, 64)
	errors = append(errors, checkFloat(err, "Stock DMSO"))
	taq = (taqPCR * reactionVol / taqStock) * reactionNum
	//DNA
	dnaPCR, err = strconv.ParseFloat(dnacValMix.Text, 64)
	errors = append(errors, checkFloat(err, "PCR DNA mass"))
	dnaStock, err = strconv.ParseFloat(dnacVal.Text, 64)
	errors = append(errors, checkFloat(err, "Stock DNA concentration"))
	dna = (dnaPCR / dnaStock)
	// H2O
	h2o = ((reactionVol * reactionNum) - buffer - dntps - mgcl2 - p1 -
		p2 - p3 - p4 - taq - (dna * reactionNum) - glycerol - dmsoFinal)
	// Final mix volume without DNA addition
	mixVolume = (buffer + dntps + mgcl2 + p1 + p2 + p3 + p4 + taq +
		glycerol + dmsoFinal + h2o)

	// Show results if there was no errors
	if contains(errors, true) {
	} else if h2o < 0 {
		str = "Your PCR concentrations\n" +
			"have values that are not\n" +
			"possible to achieve in the mix!"
		h2oError := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
			fyne.TextStyle{})
		dialog.ShowCustom("Error", "Ok", h2oError, w)
	} else {
		if p2 > 0 {
			p2p3p4 = fmt.Sprintf("Primer 2: %.3f \u00B5L\n", p2)
		}
		if p3 > 0 {
			p2p3p4 = p2p3p4 + fmt.Sprintf("Primer 3: %.3f \u00B5L\n", p3)
		}
		if p4 > 0 {
			p2p3p4 = p2p3p4 + fmt.Sprintf("Primer 4: %.3f \u00B5L\n", p4)
		}
		if glycerol > 0 {
			gly = fmt.Sprintf("Glycerol: %.3f \u00B5L\n", glycerol)
		}
		if dmsoFinal > 0 {
			dmso = fmt.Sprintf("DMSO: %.3f \u00B5L\n", dmsoFinal)
		}
		str = fmt.Sprintf("%.f reactions of %.2f \u00B5L need:\n\n",
			reactionNum, reactionVol) +
			fmt.Sprintf("Buffer: %.3f \u00B5L\n", buffer) +
			fmt.Sprintf("DNTPs: %.3f \u00B5L\n", dntps) +
			fmt.Sprintf("MgCl₂: %.3f \u00B5L\n", mgcl2) +
			fmt.Sprintf("Primer 1: %.3f \u00B5L\n", p1) +
			p2p3p4 + gly + dmso +
			fmt.Sprintf("Taq: %.3f \u00B5L\n", taq) +
			fmt.Sprintf("H₂O: %.3f \u00B5L\n", h2o) +
			fmt.Sprintf("Volume: %.3f \u00B5L", mixVolume)
		results := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
			fyne.TextStyle{})
		dnaSample := widget.NewLabelWithStyle(
			fmt.Sprintf("Add %.3f \u00B5L of stock\nDNA to each reaction.", dna),
			fyne.TextAlignCenter, fyne.TextStyle{})
		finalDialog := fyne.NewContainerWithLayout(layout.NewBorderLayout(results,
			dnaSample, nil, nil), results, dnaSample)
		dialog.ShowCustomConfirm("PCR Mix", "Save", "Cancel", finalDialog,
			save, w)
	}
}
