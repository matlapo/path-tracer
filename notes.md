Vector from point P to Q is Q - P

intersection with sphere: we are testing for any t, which means that we
need to check that t < 0 such that the sphere center is at z = +1 will
render a sphere even if *it is behind the camera's eye* so we
need to eventually check for that.

For a sphere, the normal points from the center C to the hit point P, so the normal vector is P - C. A common way of visualizing normals is to make the normal vector n a unit vector (so components are -1 < x,y,z < 1) and map them to [0,1] and then map x,y,z to r,g,b.

One design decision is whether we compute all the normals before hitting anything, which is wasteful is we hit an object close to the camera.

Another design decision is whether the normal always points outwards or not. If it always point outwards, then when the ray comes from inside the object, both the ray and normal will point in the same direction. We need to choose because knowing which side the ray is coming from matters for object made of glass for example.

note: the normal vector is independent of the ray's direction!

if the ray always point out, then when we color, we can take the dot product of the ray and normal vector to tell whether which side we are on. Otherwise, we need to keep track of that information.

It comes down to determining this either at coloring time (dot product) or geometry time (extra state).

antialiasing: blending the color arounds edges. Normal cameras get this for free.

Question: what determines if one pixel is the background color, but the one besides is the objects color?

There are multiples rays that go through the same pixel.

The image may be discrete (pixels) but the scene is continuous!

Diffuse materials
-----------------

point3 target = rec.p + rec.normal + random_in_unit_sphere();

rec.p is a 3D point on the sphere. rec.normal is a vector representing the direction the normal, but it originates from (0,0,0). rec.p + rec.normal = a vector from (0,0,0) to where the normal points when positioned on the surface.Adding a random point within a unit sphere to this gives a random direction for the child ray.

Lambertian: the less light that is reflected towards the normal (and thus less towards the camera) the darker the object will be.
