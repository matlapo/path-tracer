package main

import (
	"bufio"
	"strconv"
)

// Write the translated [0,255] value of each color component.
func writeColor(w *bufio.Writer, pixelColor Color) {
	var ir int = int(255.999 * float64(pixelColor.x))
	var ig int = int(255.999 * float64(pixelColor.y))
	var ib int = int(255.999 * float64(pixelColor.z))

	d1 := []byte(strconv.Itoa(ir) + " " + strconv.Itoa(ig) + " " + strconv.Itoa(ib) + "\n")

	_, err := (*w).Write(d1)
	check(err)
}
