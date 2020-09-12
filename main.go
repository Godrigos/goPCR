package main

func main() {
	c := &pcr{}
	c.w.SetContent(c.loadUI(c.application))
	c.prefs()
	c.w.ShowAndRun()
}
