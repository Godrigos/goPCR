package main

import (
	"fmt"
	"math"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func (p *pcr) calculate() {
	// Store error evaluation
	var errors []bool

	// Reactions volume and number conversion to float64 from string
	reactionVol, err := strconv.ParseFloat(p.reactionVolVal.Text, 64)
	errors = append(errors, p.checkFloat(err, "reaction volume"))
	reactionVol = math.Abs(reactionVol)
	p.reactionNum, err = strconv.ParseFloat(p.reactionNumVal.Text, 64)
	errors = append(errors, p.checkFloat(err, "reaction number"))
	p.reactionNum = math.Abs(p.reactionNum)

	// Other reagents conversions and calculation
	// Buffer
	p.bufferPCR, err = strconv.ParseFloat(p.bufferValMix.Text, 64)
	errors = append(errors, p.checkFloat(err, "PCR buffer"))
	p.bufferPCR = math.Abs(p.bufferPCR)
	p.bufferStock, err = strconv.ParseFloat(p.bufferVal.Text, 64)
	errors = append(errors, p.checkFloat(err, "Stock buffer"))
	p.bufferStock = math.Abs(p.bufferStock)
	p.buffer = (p.bufferPCR * reactionVol / p.bufferStock) * p.reactionNum

	// DNTPs
	p.dntpsPCR, err = strconv.ParseFloat(p.dntpsValMix.Text, 64)
	errors = append(errors, p.checkFloat(err, "PCR DNTPs"))
	p.dntpsPCR = math.Abs(p.dntpsPCR)
	p.dntpsStock, err = strconv.ParseFloat(p.dntpsVal.Text, 64)
	errors = append(errors, p.checkFloat(err, "Stock DNTPs"))
	p.dntpsStock = math.Abs(p.dntpsStock)
	p.dntps = (p.dntpsPCR * reactionVol / p.dntpsStock) * p.reactionNum

	// MgCl2
	p.mgcl2PCR, err = strconv.ParseFloat(p.mgcl2ValMix.Text, 64)
	errors = append(errors, p.checkFloat(err, "PCR MgCl₂"))
	p.mgcl2PCR = math.Abs(p.mgcl2PCR)
	p.mgcl2Stock, err = strconv.ParseFloat(p.mgcl2Val.Text, 64)
	errors = append(errors, p.checkFloat(err, "Stock MgCl₂"))
	p.mgcl2Stock = math.Abs(p.mgcl2Stock)
	p.mgcl2 = (p.mgcl2PCR * reactionVol / p.mgcl2Stock) * p.reactionNum

	//Primers
	p.p1PCR, err = strconv.ParseFloat(p.primer1ValMix.Text, 64)
	errors = append(errors, p.checkFloat(err, "PCR primer 1"))
	p.p1PCR = math.Abs(p.p1PCR)
	p.p1Stock, err = strconv.ParseFloat(p.primer1Val.Text, 64)
	errors = append(errors, p.checkFloat(err, "Stock primer 1"))
	p.p1Stock = math.Abs(p.p1Stock)
	p.p1 = (p.p1PCR * reactionVol / p.p1Stock) * p.reactionNum

	p.p2PCR, err = strconv.ParseFloat(p.primer2ValMix.Text, 64)
	errors = append(errors, p.checkFloat(err, "PCR primer 2"))
	p.p2PCR = math.Abs(p.p2PCR)
	p.p2Stock, err = strconv.ParseFloat(p.primer2Val.Text, 64)
	errors = append(errors, p.checkFloat(err, "Stock primer 2"))
	p.p2Stock = math.Abs(p.p2Stock)
	p.p2 = (p.p2PCR * reactionVol / p.p2Stock) * p.reactionNum

	p.p3PCR, err = strconv.ParseFloat(p.primer3ValMix.Text, 64)
	errors = append(errors, p.checkFloat(err, "PCR primer 3"))
	p.p3PCR = math.Abs(p.p3PCR)
	p.p3Stock, err = strconv.ParseFloat(p.primer3Val.Text, 64)
	errors = append(errors, p.checkFloat(err, "Stock primer 3"))
	p.p3Stock = math.Abs(p.p3Stock)
	p.p3 = (p.p3PCR * reactionVol / p.p3Stock) * p.reactionNum

	p.p4PCR, err = strconv.ParseFloat(p.primer4ValMix.Text, 64)
	errors = append(errors, p.checkFloat(err, "PCR primer 4"))
	p.p4PCR = math.Abs(p.p4PCR)
	p.p4Stock, err = strconv.ParseFloat(p.primer4Val.Text, 64)
	errors = append(errors, p.checkFloat(err, "Stock primer 4"))
	p.p4Stock = math.Abs(p.p4Stock)
	p.p4 = (p.p4PCR * reactionVol / p.p4Stock) * p.reactionNum

	// Glycerol
	p.glyPCR, err = strconv.ParseFloat(p.glycerolValMix.Text, 64)
	errors = append(errors, p.checkFloat(err, "PCR glycerol"))
	p.glyPCR = math.Abs(p.glyPCR)
	p.glyStock, err = strconv.ParseFloat(p.glycerolVal.Text, 64)
	errors = append(errors, p.checkFloat(err, "Stock glycerol"))
	p.glyStock = math.Abs(p.glyStock)
	p.glycerol = (p.glyPCR * reactionVol / p.glyStock) * p.reactionNum

	// DMSO
	p.dmsoPCR, err = strconv.ParseFloat(p.dmsoValMix.Text, 64)
	errors = append(errors, p.checkFloat(err, "PCR DMSO"))
	p.dmsoPCR = math.Abs(p.dmsoPCR)
	p.dmsoStock, err = strconv.ParseFloat(p.dmsoVal.Text, 64)
	errors = append(errors, p.checkFloat(err, "Stock DMSO"))
	p.dmsoStock = math.Abs(p.dmsoStock)
	p.dmsoFinal = (p.dmsoPCR * reactionVol / p.dmsoStock) * p.reactionNum

	// Taq
	p.taqPCR, err = strconv.ParseFloat(p.taqValMix.Text, 64)
	errors = append(errors, p.checkFloat(err, "PCR DMSO"))
	p.taqPCR = math.Abs(p.taqPCR)
	p.taqStock, err = strconv.ParseFloat(p.taqVal.Text, 64)
	errors = append(errors, p.checkFloat(err, "Stock DMSO"))
	p.taqStock = math.Abs(p.taqStock)
	p.taq = (p.taqPCR * reactionVol / p.taqStock) * p.reactionNum

	//DNA
	p.dnaPCR, err = strconv.ParseFloat(p.dnacValMix.Text, 64)
	errors = append(errors, p.checkFloat(err, "PCR DNA mass"))
	p.dnaPCR = math.Abs(p.dnaPCR)
	p.dnaStock, err = strconv.ParseFloat(p.dnacVal.Text, 64)
	errors = append(errors, p.checkFloat(err, "Stock DNA concentration"))
	p.dnaStock = math.Abs(p.dnaStock)
	p.dna = (p.dnaPCR / p.dnaStock)

	// H2O
	p.h2o = ((reactionVol * p.reactionNum) - p.buffer - p.dntps - p.mgcl2 - p.p1 -
		p.p2 - p.p3 - p.p4 - p.taq - (p.dna * p.reactionNum) - p.glycerol - p.dmsoFinal)

	// Final mix volume without DNA addition
	p.mixVolume = (p.buffer + p.dntps + p.mgcl2 + p.p1 + p.p2 + p.p3 + p.p4 + p.taq +
		p.glycerol + p.dmsoFinal + p.h2o)

	// Show results if there was no errors
	if contains(errors, true) {
	} else if p.h2o < 0 {
		p.str = "Your PCR concentrations\n" +
			"have values that are not\n" +
			"possible to achieve in the mix!"
		h2oError := widget.NewLabelWithStyle(p.str, fyne.TextAlignCenter,
			fyne.TextStyle{})
		dialog.ShowCustom("Error", "Ok", h2oError, p.w)
	} else {
		if p.p2 > 0 {
			p.p2p3p4 = fmt.Sprintf("Primer 2: %.3f \u00B5L\n", p.p2)
		} else {
			p.p2p3p4 = ""
		}
		if p.p3 > 0 {
			p.p2p3p4 = p.p2p3p4 + fmt.Sprintf("Primer 3: %.3f \u00B5L\n", p.p3)
		} else {
			p.p2p3p4 = p.p2p3p4 + ""
		}
		if p.p4 > 0 {
			p.p2p3p4 = p.p2p3p4 + fmt.Sprintf("Primer 4: %.3f \u00B5L\n", p.p4)
		} else {
			p.p2p3p4 = p.p2p3p4 + ""
		}
		if p.glycerol > 0 {
			p.gly = fmt.Sprintf("Glycerol: %.3f \u00B5L\n", p.glycerol)
		} else {
			p.gly = ""
		}
		if p.dmsoFinal > 0 {
			p.dmso = fmt.Sprintf("DMSO: %.3f \u00B5L\n", p.dmsoFinal)
		} else {
			p.dmso = ""
		}
		p.str = fmt.Sprintf("%.f reactions of %.2f \u00B5L need:\n\n",
			p.reactionNum, reactionVol) +
			fmt.Sprintf("Buffer: %.3f \u00B5L\n", p.buffer) +
			fmt.Sprintf("DNTPs: %.3f \u00B5L\n", p.dntps) +
			fmt.Sprintf("MgCl₂: %.3f \u00B5L\n", p.mgcl2) +
			fmt.Sprintf("Primer 1: %.3f \u00B5L\n", p.p1) +
			p.p2p3p4 + p.gly + p.dmso +
			fmt.Sprintf("Taq: %.3f \u00B5L\n", p.taq) +
			fmt.Sprintf("H₂O: %.3f \u00B5L\n", p.h2o) +
			fmt.Sprintf("Volume: %.3f \u00B5L", p.mixVolume)
		results := widget.NewLabelWithStyle(p.str, fyne.TextAlignCenter,
			fyne.TextStyle{})
		dnaSample := widget.NewLabelWithStyle(
			fmt.Sprintf("Add %.3f \u00B5L of stock\nDNA to each reaction.", p.dna),
			fyne.TextAlignCenter, fyne.TextStyle{})
		finalDialog := fyne.NewContainerWithLayout(layout.NewBorderLayout(results,
			dnaSample, nil, nil), results, dnaSample)
		dialog.ShowCustomConfirm("PCR Mix", "Copy", "Cancel", finalDialog,
			p.copyClip, p.w)
	}
}
