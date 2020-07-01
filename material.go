package main

type material interface {
	scatter(rayIn *Ray, rec *HitRecord, attenuation *Color, rayOut *Ray) bool
}

type Lambertian struct {
	albedo Color
}

func lambertian(color Color) Lambertian {
	return Lambertian{albedo: color}
}

func (l *Lambertian) scatter(rayIn *Ray, rec *HitRecord, attenuation *Color, rayOut *Ray) bool {
	var scatterDirection = rec.normal.plus(randomUnitVector())
	*rayOut = ray(rec.p, scatterDirection)
	*attenuation = l.albedo
	return true
}

type Metal struct {
	albedo Color
	fuzz   float64
}

func metal(color Color, fuzz float64) Metal {
	if fuzz > 1 {
		fuzz = 1
	}
	return Metal{albedo: color, fuzz: fuzz}
}

func (m *Metal) scatter(rayIn *Ray, rec *HitRecord, attenuation *Color, rayOut *Ray) bool {
	var reflected = reflect(rayIn.direction.unit(), rec.normal)
	*rayOut = ray(rec.p, reflected.plus(randomInUnitSphere().scale(m.fuzz)))
	*attenuation = m.albedo
	return (rayOut.direction.dot(rec.normal) > 0)
}
