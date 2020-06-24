package main

import "math"

var infinity float64 = math.Inf(1)

const pi float64 = math.Pi

func degreesToRads(degrees float64) float64 {
	return degrees * pi / 180
}
