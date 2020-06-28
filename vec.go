package main

import (
	"math"
	"math/rand"
)

// Vector to represent point or color
type Vector struct {
	x, y, z float64
}

// Point3 blah
type Point3 = Vector

// Color bla
type Color = Vector

func color(x float64, y float64, z float64) Color {
	return Color{x: x, y: y, z: z}
}

func point3(x float64, y float64, z float64) Point3 {
	return Point3{x: x, y: y, z: z}
}

func vec3(x float64, y float64, z float64) Vector {
	return Vector{x: x, y: y, z: z}
}

func (u Vector) length() float64 {
	return math.Sqrt(u.lengthSquared())
}

func (u Vector) lengthSquared() float64 {
	return u.x*u.x + u.y*u.y + u.z*u.z
}

func (u Vector) inv() Vector {
	return Vector{x: -u.x, y: -u.y, z: -u.z}
}

func (u Vector) plus(v Vector) Vector {
	return Vector{x: u.x + v.x, y: u.y + v.y, z: u.z + v.z}
}

func (u Vector) minus(v Vector) Vector {
	return Vector{x: u.x - v.x, y: u.y - v.y, z: u.z - v.z}
}

func (u Vector) times(v Vector) Vector {
	return Vector{x: u.x * v.x, y: u.y * v.y, z: u.z * v.z}
}

func (u Vector) scale(t float64) Vector {
	return Vector{x: t * u.x, y: t * u.y, z: t * u.z}
}

func (u Vector) divide(t float64) Vector {
	return u.scale(1 / t)
}

func (u Vector) dot(v Vector) float64 {
	return u.x*v.x + u.y*v.y + u.z*v.z
}

func (u Vector) cross(v Vector) Vector {
	return Vector{
		x: u.y*v.z - u.z*v.y,
		y: u.z*v.x - u.x*v.z,
		z: u.x*v.y - u.y*v.x,
	}
}

func (u Vector) unit() Vector {
	return u.divide(u.length())
}

func random() Vector {
	return vec3(rand.Float64(), rand.Float64(), rand.Float64())
}

func randomInRange(min float64, max float64) Vector {
	x := min + rand.Float64()*(max-min)
	y := min + rand.Float64()*(max-min)
	z := min + rand.Float64()*(max-min)
	return vec3(x, y, z)
}

// pick a random point inside a sphere.
// works by picking a random point in a unit
// cube and reject until within a sphere.
func randomInUnitSphere() Vector {
	for {
		var point = randomInRange(-1, 1)
		if point.lengthSquared() >= 1 {
			continue
		}
		return point
	}
}
