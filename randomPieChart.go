package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/gizak/termui/v3/widgets"
)

type RotatingPieChart struct {
	piechart           *widgets.PieChart
	rotationsPerSecond float64 // how many rotations per second
}

func (pc *RotatingPieChart) getWidget() interface{} {
	return pc.piechart
}

func (pc *RotatingPieChart) update(seconds float64) {
	pc.piechart.AngleOffset = 2.0 * math.Pi * math.Mod(seconds*pc.rotationsPerSecond, 1.0)
}

func randomDataAndOffset() (data []float64, offset float64) {
	rand.Seed(time.Now().UnixMilli())
	noSlices := 2 + rand.Intn(5)
	data = make([]float64, noSlices)
	for i := range data {
		data[i] = rand.Float64()
	}
	offset = 2.0 * math.Pi * rand.Float64()
	return
}

func newRandomPieChart(rotationsPerSecond float64) NeoWidget {
	pc := new(RotatingPieChart)
	pc.piechart = widgets.NewPieChart()
	pc.piechart.Data, pc.piechart.AngleOffset = randomDataAndOffset()
	pc.piechart.LabelFormatter = func(i int, v float64) string {
		return fmt.Sprintf("%.02f", v)
	}
	pc.rotationsPerSecond = rotationsPerSecond

	return pc
}
