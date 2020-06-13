package main

// Ray representing P(t) = A + b*t
type Ray struct {
	origin    Point3
	direction Vector
}

func ray(origin Point3, direction Vector) Ray {
	return Ray{origin: origin, direction: direction}
}

func (r Ray) at(t float64) Point3 {
	return r.origin.plus(r.direction.scale(t))
}
