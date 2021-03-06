package main

import (
	"bufio"
	"math"
	"strconv"
)

// Write the translated [0,255] value of each color component.
func writeColor(w *bufio.Writer, pixelColor Color, samplesPerPixel int32) {
	var r = pixelColor.x
	var g = pixelColor.y
	var b = pixelColor.z

	// Average the color by total number of samples and gamma-correct for gamma = 2.
	var scale = 1.0 / float64(samplesPerPixel)
	r = math.Sqrt(r * scale)
	g = math.Sqrt(g * scale)
	b = math.Sqrt(b * scale)

	var ir = int(256.0 * clamp(r, 0.0, 0.999))
	var ig = int(256.0 * clamp(g, 0.0, 0.999))
	var bg = int(256.0 * clamp(b, 0.0, 0.999))

	d1 := []byte(strconv.Itoa(ir) + " " + strconv.Itoa(ig) + " " + strconv.Itoa(bg) + "\n")

	_, err := (*w).Write(d1)
	check(err)
}
