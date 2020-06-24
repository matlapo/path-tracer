package main

type HittableList struct {
	objects []hittable
}

func (hittableList *HittableList) clear() {
	hittableList.objects = nil // TODO this frees up the objects, might want to use list[:0] instead
}

func (hittableList *HittableList) add(h hittable) {
	hittableList.objects = append(hittableList.objects, h)
}

func (hittableList *HittableList) hit(ray Ray, tMin float64, tMax float64, rec HitRecord) bool {
	var tempRecord HitRecord
	var hitAnything = false
	var closestSoFar = tMax

	for _, object := range hittableList.objects {
		if object.hit(ray, tMin, closestSoFar, tempRecord) {
			hitAnything = true
			closestSoFar = tempRecord.t
			// warning: mutates the argument
			rec = tempRecord
		}
	}

	return hitAnything
}
