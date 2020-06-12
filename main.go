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

func main() {
	const imageWidth int = 256
	const imageHeight int = 256

	f, err := os.Create("./image.ppm")
	check(err)
	defer f.Close()

	d1 := []byte("P3\n" + strconv.Itoa(imageWidth) + " " + strconv.Itoa(imageHeight) + "\n255\n")

	w := bufio.NewWriter(f)
	n4, err := w.Write(d1)
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			r := float64(i) / float64(imageWidth-1)
			g := float64(j) / float64(imageHeight-1)
			b := 0.25

			var ir int = int(255.999 * float64(r))
			var ig int = int(255.999 * float64(g))
			var ib int = int(255.999 * float64(b))

			d1 := []byte(strconv.Itoa(ir) + " " + strconv.Itoa(ig) + " " + strconv.Itoa(ib) + "\n")

			_, err := w.Write(d1)
			check(err)
		}
	}

	w.Flush()
}
