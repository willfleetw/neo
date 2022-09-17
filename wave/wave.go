package wave

import (
	"math"
	"math/rand"
)

type Wave int8

const (
	Undefined Wave = iota
	Triangle
	Sin
	Sawtooth
)

func GetRandomWave() Wave {
	return Wave(rand.Intn(3) + 1)
}

/*
Calculates the value for a given wave function with the given period and amplitude at the given point x
*/
func ValueAt(wave Wave, x, period, amplitude float64) (value float64) {
	switch wave {
	case Triangle:
		value = 2 * math.Abs((x/period)-math.Floor((x/period)+.5))
	case Sin:
		value = math.Sin((2*math.Pi*x)/period-(math.Pi/2))/2 + 0.5
	case Sawtooth:
		value = x/period - math.Floor(1+x/period) + 1
	}

	value *= amplitude
	return
}
