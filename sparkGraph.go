package main

import (
	"neo/wave"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type SparkGraph struct {
	sparklineGroup *widgets.SparklineGroup
	data           []float64
	speed          float64
}

func (sg *SparkGraph) getWidget() interface{} {
	return sg.sparklineGroup
}

func (sg *SparkGraph) update(seconds float64) {
	index := int(seconds * sg.speed)
	sg.sparklineGroup.Sparklines[0].Data = append(sg.data[(index)%400:], sg.data[:(index)%400]...)
}

func newSparkGraph(color ui.Color, waveForm wave.Wave, period, amplitude, speed float64) NeoWidget {
	sg := new(SparkGraph)
	sg.data = make([]float64, 400)
	for i := range sg.data {
		sg.data[i] = wave.ValueAt(waveForm, float64(i)/5, period, amplitude)
	}
	sg.speed = speed

	sl := widgets.NewSparkline()
	sl.Data = sg.data
	sl.LineColor = color

	sg.sparklineGroup = widgets.NewSparklineGroup(sl)

	return sg
}
