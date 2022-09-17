package main

import (
	"neo/wave"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type LineGraph struct {
	plot  *widgets.Plot
	data  []float64
	speed float64
}

func (lg *LineGraph) getWidget() interface{} {
	return lg.plot
}

func (lg *LineGraph) update(seconds float64) {
	index := int(seconds * lg.speed)
	lg.plot.Data[0] = append(lg.data[(index)%400:], lg.data[:(index)%400]...)
}

func newLineGraph(color ui.Color, waveForm wave.Wave, period, amplitude, speed float64) NeoWidget {
	lg := new(LineGraph)
	lg.data = make([]float64, 400)
	for i := range lg.data {
		lg.data[i] = wave.ValueAt(waveForm, float64(i)/5, period, amplitude)
	}
	lg.speed = speed

	lg.plot = widgets.NewPlot()
	lg.plot.Data = append(lg.plot.Data, lg.data)
	lg.plot.AxesColor = ui.ColorWhite
	lg.plot.LineColors[0] = color

	return lg
}
