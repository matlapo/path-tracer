package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func rayColor(ray Ray) Color {
	var unitDirection Vector = ray.direction.unit()
	t := 0.5 * (unitDirection.y + 1.0)
	return (color(1.0, 1.0, 1.0).scale(1.0 - t)).plus(color(0.5, 0.7, 1.0).scale(t))
}

func main() {
	const aspectRatio = 16.0 / 9.0
	const imageWidth = 384
	const imageHeight = int(imageWidth / aspectRatio)

	f, err := os.Create("./image.ppm")
	check(err)
	defer f.Close()

	d1 := []byte("P3\n" + strconv.Itoa(imageWidth) + " " + strconv.Itoa(imageHeight) + "\n255\n")

	w := bufio.NewWriter(f)
	n4, err := w.Write(d1)
	check(err)
	fmt.Printf("wrote header: %d bytes\n", n4)

	const viewportHeight = 2.0
	const viewportWidth = aspectRatio * viewportHeight
	const focalLength = 1.0

	var origin = point3(0, 0, 0)
	var horizontal = vec3(viewportWidth, 0, 0)
	var vertical = vec3(0, viewportHeight, 0)
	var lowerLeftCorner = ((origin.minus(horizontal.scale(0.5))).minus(vertical.scale(0.5))).minus(vec3(0, 0, focalLength))

	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			var u = float64(i) / float64(imageWidth-1)
			var v = float64(j) / float64(imageHeight-1)
			var dir Vector = ((lowerLeftCorner.plus(horizontal.scale(u))).plus(vertical.scale(v))).minus(origin)
			var ray Ray = ray(origin, dir)
			var pixelColor = rayColor(ray)

			writeColor(w, pixelColor)
		}
	}

	w.Flush()
}
