package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
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
func rayColor(r Ray, world hittable, depth int) Color {
	var hitRecord HitRecord

	if depth <= 0 {
		// ray has exceeded max number of bounces, no more light gathered.
		return color(0, 0, 0)
	}

	if world.hit(r, 0.001, infinity, &hitRecord) {
		var scattered Ray
		var attenuation Color
		if hitRecord.material.scatter(&r, &hitRecord, &attenuation, &scattered) {
			return rayColor(scattered, world, depth-1).times(attenuation)
		}
		return color(0, 0, 0)
	}
	// else, ray hits the background.
	var unitDirection Vector = r.direction.unit()
	// by converting ray.direction to a unit vector, we can have
	// a blend based on the value of y (which we know is now -1 < y < 1)
	// and the 1/2 factor keeps 0 <= t <= 1 required for the lerp.
	var t = 0.5 * (unitDirection.y + 1.0)
	return (color(1.0, 1.0, 1.0).scale(1.0 - t)).plus(color(0.5, 0.7, 1.0).scale(t))
}

func hitSphere(center Point3, radius float64, ray Ray) float64 {
	// [P(t) - C]^2 = R^2 where P(t) = A + t*B
	// This gives a polynomial of order 2 (t is the unknown).
	// t here is a value that stretches the ray to any length.
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
	const samplesPerPixel = 100
	const maxDepth = 50

	f, err := os.Create("./image.ppm")
	check(err)
	defer f.Close()

	d1 := []byte("P3\n" + strconv.Itoa(imageWidth) + " " + strconv.Itoa(imageHeight) + "\n255\n")

	w := bufio.NewWriter(f)
	n4, err := w.Write(d1)
	check(err)
	fmt.Printf("wrote header: %d bytes\n", n4)

	var world HittableList
	var dielectric1 = dielectric(1.4)
	var lambertian2 = lambertian(color(0.8, 0.8, 0.0))
	var sphere1 = sphere(point3(0, 0, -1), 0.5, &dielectric1)
	var sphere2 = sphere(point3(0, -100.5, -1), 100, &lambertian2)
	var metal1 = metal(color(0.8, 0.6, 0.2), 0.3)
	var dielectric2 = dielectric(1.4)
	var sphere3 = sphere(point3(1, 0, -1), 0.5, &metal1)
	var sphere4 = sphere(point3(-1, 0, -1), 0.5, &dielectric2)
	world.add(&sphere1)
	world.add(&sphere2)
	world.add(&sphere3)
	world.add(&sphere4)

	var cam Camera = camera()

	// for every pixel in the image.
	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			var pixelColor Color = color(0, 0, 0)
			for s := 0; s < samplesPerPixel; s++ {
				// for the given pixel (i,j), send samplesPerPixel rays into it.
				// If there was no randomness, the ray would always hit the same
				// color in the scene, so we would have no antialiasing.
				var u = (float64(i) + rand.Float64()) / float64(imageWidth-1)
				var v = (float64(j) + rand.Float64()) / float64(imageHeight-1)
				// u,v represents ratios of the image, the same ratio can be applied
				// to the viewport.
				var r Ray = cam.getRay(u, v)
				var rayColor = rayColor(r, &world, maxDepth)
				pixelColor = pixelColor.plus(rayColor)
			}
			writeColor(w, pixelColor, samplesPerPixel)
		}
	}

	w.Flush()
}
