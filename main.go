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
	// else, ray hits the background.
	var unitDirection Vector = ray.direction.unit()
	// by converting ray.direction to a unit vector, we can have
	// a blend based on the value of y (which we know is now -1 < y < 1)
	// and the 1/2 factor keeps 0 <= t <= 1 required for the lerp.
	t = 0.5 * (unitDirection.y + 1.0)
	return (color(1.0, 1.0, 1.0).scale(1.0 - t)).plus(color(0.5, 0.7, 1.0).scale(t))
}

func hitSphere(center Point3, radius float64, ray Ray) float64 {
	var oc Vector = ray.origin.minus(center)
	var a = ray.direction.lengthSquared()
	var halfB = oc.dot(ray.direction)
	var c = oc.lengthSquared() - radius*radius
	var discrimnant = halfB*halfB - a*c
	if discrimnant < 0 {
		return -1.0
	}
	return (-halfB - math.Sqrt(discrimnant)) / a
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

	// viewport conserves same aspect ratio, but has different units
	const viewportHeight = 2.0
	const viewportWidth = aspectRatio * viewportHeight

	// the eye is 1.0 units away from the viewport
	const focalLength = 1.0

	var origin = point3(0, 0, 0)
	var horizontal = vec3(viewportWidth, 0, 0)
	var vertical = vec3(0, viewportHeight, 0)

	// origin - horizontal/2 - vertical/2 - focal_length == lower left corner of viewport
	// see 3D diagram.
	var lowerLeftCorner = ((origin.minus(horizontal.scale(0.5))).minus(vertical.scale(0.5))).minus(vec3(0, 0, focalLength))

	fmt.Printf("DEBUG %f %f %f\n", lowerLeftCorner.x, lowerLeftCorner.y, lowerLeftCorner.z)

	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			// 0 <= u,v <= 1
			var u = float64(i) / float64(imageWidth-1)
			var v = float64(j) / float64(imageHeight-1)
			// lower_left_corner + u*horizontal + v*vertical - origin
			// u,v = fraction of viewport => u*horizontal = current point on viewport
			var dir Vector = ((lowerLeftCorner.plus(horizontal.scale(u))).plus(vertical.scale(v))).minus(origin)
			// vector from camera's eye to the viewport
			var ray Ray = ray(origin, dir)
			var pixelColor = rayColor(ray)

			writeColor(w, pixelColor)
		}
	}

	w.Flush()
}
