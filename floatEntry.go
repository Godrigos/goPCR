package main

import (
	"regexp"
	"strings"

	"fyne.io/fyne/v2/widget"
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
// It filters input so only numbers and a single '.' are allowed with a
// maximum of 4 decimal places.
// Implements: fyne.Focusable
func (e *floatEntry) TypedRune(r rune) {
	reg := regexp.MustCompile(`[0-9]|[.]`)
	if !reg.MatchString(string(r)) {
		return
	}

	reg2 := regexp.MustCompile(`[.]`)
	if reg2.MatchString(e.Text) && r == '.' {
		return
	}

	if len(strings.SplitAfter(e.Text, ".")) > 1 {
		if len([]rune(strings.SplitAfter(e.Text, ".")[1])) > 3 {
			return
		}
	}

	e.Entry.TypedRune(r)
}
