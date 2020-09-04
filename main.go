package main

import "fyne.io/fyne"

func main() {
	c := &pcr{}
	c.w.SetContent(c.loadUI(c.application))
	c.prefs()
	c.w.Resize(fyne.NewSize(550, 1))
	c.w.SetFixedSize(true)
	c.w.ShowAndRun()
}
