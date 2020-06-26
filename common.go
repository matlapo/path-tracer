package main

import "math"

var infinity float64 = math.Inf(1)

const pi float64 = math.Pi

func degreesToRads(degrees float64) float64 {
	return degrees * pi / 180
}

func clamp(x float64, min float64, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
