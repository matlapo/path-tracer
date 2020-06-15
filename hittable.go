package main

type HitRecord struct {
	p         Point3
	normal    Vector
	t         float64
	frontFace bool
}

// this function mutates the HitRecord directly, not sure if this is a good idea
func (rec HitRecord) setFaceNormal(ray Ray, outwardNormal Vector) {
	rec.frontFace = ray.direction.dot(outwardNormal) < 0
	if rec.frontFace {
		rec.normal = outwardNormal
	} else {
		rec.normal = outwardNormal.inv()
	}
}

type hittable interface {
	hit(ray Ray, tMin float64, tMax float64, rec HitRecord) bool
}
