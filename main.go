package main

import (
	"log"
	"math"
	"neo/wave"
	"time"

	ui "github.com/gizak/termui/v3"
)

type NeoWidget interface {
	update(float64)
	getWidget() interface{}
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	nwgs := make([]NeoWidget, 7)
	nwgs[0] = newSparkGraph(ui.ColorCyan, wave.Sawtooth, 2*math.Pi, 1, 12)
	nwgs[1] = newLineGraph(ui.ColorYellow, wave.Sin, 2*math.Pi, 1, 12)
	nwgs[2] = newRandomPieChart(.75)
	nwgs[3] = newBouncingGauge(ui.ColorBlue, wave.Sin, 6.0)
	nwgs[4] = newBouncingGauge(ui.ColorRed, wave.Triangle, 5.0)
	nwgs[5] = newBouncingGauge(ui.ColorGreen, wave.Sawtooth, 2.0)
	nwgs[6] = newTerminalMessages(5)

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/2, nwgs[0].getWidget()),
			ui.NewCol(1.0/2, nwgs[1].getWidget()),
		),
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/4, nwgs[2].getWidget()),
			ui.NewCol(1.0/4,
				ui.NewRow(1./3, nwgs[3].getWidget()),
				ui.NewRow(1./3, nwgs[4].getWidget()),
				ui.NewRow(1./3, nwgs[5].getWidget()),
			),
			ui.NewCol(1.0/4, nwgs[6].getWidget()),
		),
	)

	ui.Render(grid)

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Millisecond * 100).C
	startTime := time.Now()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(grid)
			}
		case currTime := <-ticker:
			timeDelta := currTime.Sub(startTime).Seconds()
			for _, wg := range nwgs {
				wg.update(timeDelta)
			}
			ui.Render(grid)
		}
	}
}
