package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// some terminology:
// basic idea: trace a path from imaginary eye (the camera)
// through each pixel of the screen and calculating the color for each pixel.

// behind the screen is a scene, which is a mathematical representation of
// a 3D environment.

// (1) calculate the ray from the eye to the pixel
// (2) determine which objects the ray intersects
// (3) compute a color for that intersection point.

// the camera is located at (0,0,0), the screen is towards the negative
// z-axis.

// what's a viewport? basically the screen
// focal length = 1 unit (check the z-axis on the diagram)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// returns the color of the background
// blendValue = (1 - t)*startValue + t*endValue (0 <= t <= 1)
func rayColor(ray Ray) Color {
	var t = hitSphere(point3(0, 0, -1), 0.5, ray)
	if t > 0.0 {
		var N Vector = (ray.at(t)).minus(vec3(0, 0, -1))
		return color(N.x+1, N.y+1, N.z+1).scale(0.5)
	}
	var unitDirection Vector = ray.direction.unit()
	t = 0.5 * (unitDirection.y + 1.0)
	return (color(1.0, 1.0, 1.0).scale(1.0 - t)).plus(color(0.5, 0.7, 1.0).scale(t))
}

// vec3 oc = r.origin() - center;
// auto a = r.direction().length_squared();
// auto half_b = dot(oc, r.direction());
// auto c = oc.length_squared() - radius*radius;
// auto discriminant = half_b*half_b - a*c;

// if (discriminant < 0) {
//     return -1.0;
// } else {
//     return (-half_b - sqrt(discriminant) ) / a;
// }

func hitSphere(center Point3, radius float64, ray Ray) float64 {
	var oc Vector = ray.origin.minus(center)
	var a = ray.direction.lengthSquared()
	var halfB = oc.dot(ray.direction)
	var c = oc.lengthSquared() - radius*radius
	var discrimnant = halfB*halfB - a*c
	if discrimnant < 0 {
		return -1.0
	} else {
		return (-halfB - math.Sqrt(discrimnant)) / a
	}
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
