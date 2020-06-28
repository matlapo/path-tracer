package main

// Camera holds all info about the source of rays.
type Camera struct {
	aspectRatio     float64
	viewportHeight  float64
	viewportWidth   float64
	focalLength     float64
	origin          Point3
	horizontal      Vector
	vertical        Vector
	lowerLeftCorner Point3
}

func camera() Camera {
	var aspectRatio = 16.0 / 9.0
	var viewportHeight = 2.0
	var viewportWidth = aspectRatio * viewportHeight
	var focalLength = 1.0
	var origin = point3(0, 0, 0)
	// vector that points from origin to width, used for computing lower left corner.
	var horizontal = vec3(viewportWidth, 0.0, 0.0)
	var vertical = vec3(0.0, viewportHeight, 0.0)
	// vector that points from origin to lower left corner of the viewport.
	// i.e origin - horizontal/2 - vertical/2 - vec3(0, 0, focal_length)
	var lowerLeftCorner = ((origin.minus(horizontal.scale(0.5))).minus(vertical.scale(0.5))).minus(vec3(0, 0, focalLength))
	return Camera{
		aspectRatio:     aspectRatio,
		viewportHeight:  viewportHeight,
		viewportWidth:   viewportWidth,
		focalLength:     focalLength,
		origin:          origin,
		horizontal:      horizontal,
		vertical:        vertical,
		lowerLeftCorner: lowerLeftCorner,
	}
}

func (camera Camera) getRay(u float64, v float64) Ray {
	// returns a vector from origin to a point on the viewport.
	// i.e origin to (lower_left_corner + u*horizontal + v*vertical - origin).
	// u,v are [0,1) so they represent a fraction of the viewport.
	// note: - origin because lower_left_corner + u*horizontal + v*vertical is not a point on the viewport.
	var dir Vector = ((camera.lowerLeftCorner.plus(camera.horizontal.scale(u))).plus(camera.vertical.scale(v))).minus(camera.origin)
	return ray(camera.origin, dir)
}
