package main

func main() {

	// Set stock, pcr and mix values from application preferences
	prefs()

	w.SetFixedSize(true)

	w.SetIcon(icon)
	w.SetContent(loadUI())
	w.ShowAndRun()
}
