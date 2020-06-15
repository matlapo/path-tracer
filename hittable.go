package main

type HitRecord struct {
	p      Point3
	normal Vector
	t      float64
}

type hittable interface {
	hit(ray Ray, tMin float64, tMax float64, rec HitRecord) bool
}
