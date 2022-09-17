package main

import (
	"strings"

	"github.com/gizak/termui/v3/widgets"
)

var matrixMessages = []string{
	"Wake up, Neo...",
	"The Matrix has you...",
	"Follow the white rabbit.",
	"Knock, knock, Neo.",
}

type TerminalMessages struct {
	par      *widgets.Paragraph
	interval float64 // how many seconds between each new message
}

func (tm *TerminalMessages) getWidget() interface{} {
	return tm.par
}

func (tm *TerminalMessages) update(seconds float64) {
	index := (int(seconds/tm.interval) % len(matrixMessages)) + 1
	tm.par.Text = strings.Join(matrixMessages[:index], "\n")
}

func newTerminalMessages(interval float64) NeoWidget {
	tm := new(TerminalMessages)
	tm.par = widgets.NewParagraph()
	tm.par.Text = ""
	tm.par.Title = "Morpheus"

	tm.interval = interval

	return tm
}
