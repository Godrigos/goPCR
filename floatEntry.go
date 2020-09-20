package main

import (
	"regexp"

	"fyne.io/fyne/widget"
)

type floatEntry struct {
	widget.Entry
}

func newFloatEntry() *floatEntry {
	entry := &floatEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

// TypedRune receives text input events when the Entry widget is focused.
// Implements: fyne.Focusable
func (e *floatEntry) TypedRune(r rune) {
	reg := regexp.MustCompile(`[0-9]|[.]`)
	if !reg.MatchString(string(r)) {
		return
	}
	e.Entry.TypedRune(r)
}
