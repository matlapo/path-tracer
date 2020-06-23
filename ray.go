package main

// Ray representing P(t) = A + t*B
// A is the eye and B is a point on the viewport
type Ray struct {
	origin    Point3
	direction Vector
}

func ray(origin Point3, direction Vector) Ray {
	return Ray{origin: origin, direction: direction}
}

// P(t) = A + t*B
func (r Ray) at(t float64) Point3 {
	return r.origin.plus(r.direction.scale(t))
}
