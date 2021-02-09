package main

import "fyne.io/fyne/v2"

func main() {
	c := &pcr{}
	c.w.SetContent(c.loadUI(c.application))
	c.prefs()
	c.w.Resize(fyne.NewSize(350, 1))
	c.w.SetFixedSize(true)
	c.w.ShowAndRun()
}
