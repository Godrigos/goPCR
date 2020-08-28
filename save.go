package main

import (
	"fmt"
	"io"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/storage"
	"fyne.io/fyne/widget"
)

func (p *pcr) save(b bool) {
	if b {
		save := dialog.NewFileSave(func(
			write fyne.URIWriteCloser, err error) {
			if write == nil {
				return
			}
			l, err := io.WriteString(write,
				time.Now().Format("Jan 2 2006 - 15:04:05")+"\n"+
					p.str+
					fmt.Sprintf(
						"\n\nAdd %.3f \u00B5L of stock "+
							"DNA to each reaction.\n",
						p.dna),
			)
			if err != nil {
				str := "Could not write to file!"
				fileError := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
					fyne.TextStyle{})
				dialog.ShowCustom("Error", "Ok", fileError, p.w)
			} else {
				str := fmt.Sprintf("File saved successfully!\n"+
					"(%.d bytes)", l)
				done := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
					fyne.TextStyle{})
				dialog.ShowCustom("Done", "Ok", done, p.w)
			}
		}, p.w)
		save.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		save.Show()
	}
}
