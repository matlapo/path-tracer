package main

// HitRecord is a abstraction for any hittable surface (e.g a sphere, a list of spheres, etc).
// It keeps track of geometry information when a ray hits a surface.
type HitRecord struct {
	p      Point3  // where the ray it the surface.
	normal Vector  // normal to the surface at point p.
	t      float64 // value of t such that P(t) = A + t*B.
	// this implementation of ray tracing makes the normal vector of any surface
	// always points outward. We thus keep track of which side (front or back)
	// the ray is hitting the object.
	frontFace bool
}

// set the normal vector to the appropriate direction.
// the normal always (by convention) points against the ray.
// we therefore keep track of which side the ray is hitting.
func (rec *HitRecord) setFaceNormal(ray Ray, outwardNormal Vector) {
	rec.frontFace = ray.direction.dot(outwardNormal) < 0
	if rec.frontFace {
		rec.normal = outwardNormal
	} else {
		rec.normal = outwardNormal.inv()
	}
}

// Hittable is an interface that all surfaces must implement.
// Most ray tracers have a valid interval for hits.
type hittable interface {
	hit(ray Ray, tMin float64, tMax float64, rec *HitRecord) bool
}
