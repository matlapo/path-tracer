package main

import "math"

type Sphere struct {
	center Point3
	radius float64
}

func sphere(cen Point3, r float64) Sphere {
	return Sphere{center: cen, radius: r}
}

func (sphere Sphere) hit(ray Ray, tMin float64, tMax float64, rec HitRecord) bool {
	var oc Vector = ray.origin.minus(sphere.center)
	var a = ray.direction.lengthSquared()
	var halfB = oc.dot(ray.direction)
	var c = oc.lengthSquared() - sphere.radius*sphere.radius
	var discrimnant = halfB*halfB - a*c

	if discrimnant > 0 {
		var root = math.Sqrt(discrimnant)
		var temp = (-halfB - root) / a
		if temp < tMax && temp > tMin {
			rec.t = temp
			rec.p = ray.at(rec.t)
			var outwardNormal = (rec.p.minus(sphere.center)).divide(sphere.radius)
			rec.setFaceNormal(ray, outwardNormal)
			return true
		}
		temp = (-halfB + root) / a
		if temp < tMax && temp > tMin {
			rec.t = temp
			rec.p = ray.at(rec.t)
			var outwardNormal = (rec.p.minus(sphere.center)).divide(sphere.radius)
			rec.setFaceNormal(ray, outwardNormal)
			return true
		}
	}
	return false
}
