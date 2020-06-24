package main

// HitRecord is a abstraction for any hittable surface (e.g a sphere, a list of spheres, etc)
type HitRecord struct {
	p         Point3
	normal    Vector
	t         float64
	frontFace bool // we keep track of which side (front or back) the ray is hitting the object
}

// set the normal vector to the appropriate direction.
// the normal always (by convention) points against the ray.
// we therefore keep track of which side the ray is hitting.
func (rec HitRecord) setFaceNormal(ray Ray, outwardNormal Vector) {
	rec.frontFace = ray.direction.dot(outwardNormal) < 0
	if rec.frontFace {
		rec.normal = outwardNormal
	} else {
		rec.normal = outwardNormal.inv()
	}
}

// Hittable takes a ray as input. Most ray tracers have a valid interval for hits.
type hittable interface {
	hit(ray Ray, tMin float64, tMax float64, rec HitRecord) bool
}
