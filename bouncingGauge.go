package main

import (
	"neo/wave"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type BouncingGauge struct {
	gauge  *widgets.Gauge
	wave   wave.Wave
	period float64
}

func (bg *BouncingGauge) getWidget() interface{} {
	return bg.gauge
}

func (bg *BouncingGauge) update(seconds float64) {
	bg.gauge.Percent = int(wave.ValueAt(bg.wave, seconds, bg.period, 100))
}

func newBouncingGauge(color ui.Color, wave wave.Wave, period float64) NeoWidget {
	bg := new(BouncingGauge)
	bg.gauge = widgets.NewGauge()
	bg.gauge.Percent = 0
	bg.gauge.BarColor = color

	bg.wave = wave
	bg.period = period

	return bg
}
