package main

import (
	"fmt"
	"os"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

func save(b bool) {
	if b {
		fileName := fmt.Sprintf(execDir+"/saved/PCRCalc_%s.txt",
			time.Now().Format("Jan-2-2006-15.04.05"))

		f, err := os.Create(fileName)
		if err != nil {
			str := "Could not create the file!"
			fileError := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
				fyne.TextStyle{})
			dialog.ShowCustom("Error", "Ok", fileError, w)
		}
		defer f.Close()

		l, err := f.WriteString(
			time.Now().Format("Jan 2 2006 - 15:04:05") + "\n" +
				str +
				fmt.Sprintf(
					"\n\nAdd %.3f \u00B5L of stock\n"+
						"DNA to each reaction.",
					dna),
		)
		if err != nil {
			str := "Could not write to file!"
			fileError := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
				fyne.TextStyle{})
			dialog.ShowCustom("Error", "Ok", fileError, w)
		} else {
			str := fmt.Sprintf("File saved successfully!\n"+
				"(%.d bytes)", l)
			done := widget.NewLabelWithStyle(str, fyne.TextAlignCenter,
				fyne.TextStyle{})
			dialog.ShowCustom("Done", "Ok", done, w)
		}
		f.Sync()
	}
}
