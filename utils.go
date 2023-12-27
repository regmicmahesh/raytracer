package raytracer

const EPSILON float64 = 0.001

func Equals(a, b float64) bool {

	return (a - b) < EPSILON

}
