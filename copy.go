package main

import (
	"fmt"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

func (p *pcr) copyClip(b bool) {
	if b {
		p.clip = time.Now().Format("Jan 2 2006 - 15:04:05") + "\n" +
			p.str +
			fmt.Sprintf(
				"\n\nAdd %.3f \u00B5L of stock "+
					"DNA to each reaction.\n", p.dna)
		p.w.Clipboard().SetContent(p.clip)
		done := widget.NewLabelWithStyle("Results copied to clipboard!",
			fyne.TextAlignCenter,
			fyne.TextStyle{})
		dialog.ShowCustom("Done", "Ok", done, p.w)
	}
}
