package geom3d

// Transform represents a rigid 3D transform:
// p_world = R * p_local + T
type Transform struct {
	R Mat3
	T Vec3
}

// IdentityTransform returns the identity rigid transform.
func IdentityTransform() Transform {
	return Transform{
		R: IdentityMat3(),
		T: Vec3{},
	}
}

// ApplyPoint applies the transform to a point.
func (tf Transform) ApplyPoint(p Vec3) Vec3 {
	return tf.R.MulVec(p).Add(tf.T)
}

// ApplyVector applies the transform to a vector.
//
// Translation is not applied.
func (tf Transform) ApplyVector(v Vec3) Vec3 {
	return tf.R.MulVec(v)
}

// Compose returns the composition tf ∘ other.
//
// The returned transform is equivalent to applying other first,
// then applying tf.
func (tf Transform) Compose(other Transform) Transform {
	return Transform{
		R: tf.R.Mul(other.R),
		T: tf.R.MulVec(other.T).Add(tf.T),
	}
}

// Inverse returns the inverse of a rigid transform.
//
// This assumes tf.R is a proper rotation matrix.
func (tf Transform) Inverse() Transform {
	rt := tf.R.Transpose()
	return Transform{
		R: rt,
		T: rt.MulVec(tf.T).Scale(-1),
	}
}
