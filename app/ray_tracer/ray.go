package ray_tracer

type SceneObject interface{}

// TODO: Can I do new types Point and Vector inherited from Tuple?
type Ray struct {
	origin    Tuple
	direction Tuple
}

func NewRay(origin, direction Tuple) Ray {
	if !origin.IsPoint() {
		panic("Origin must be a point!")
	}
	if !direction.IsVector() {
		panic("Direction must be a vector!")
	}
	return Ray{origin, direction}
}

func (r *Ray) CalcPosition(time float64) Tuple {
	return r.origin.Add(r.direction.Mul(time))
}

func (r *Ray) ApplyTransform(m *Matrix) Ray {
	return NewRay(m.MulTuple(r.origin), m.MulTuple(r.direction))
}
