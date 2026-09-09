package funct

import (
	"fmt"
	"math"
)

type Tuple struct {
	X float64
	Y float64
}

func FindDistance(p1, p2 Tuple) float64 {
	return math.Sqrt(math.Pow(p1.X-p2.X, 2) + math.Pow(p2.Y-p1.Y, 2))
}

// func TestBound(center, point Tuple) float64 {
// 	return math.Pow(point.Y-center.Y, 2) + math.Pow(center.X-point.X, 2)
// }

func FindCenter(arr []Tuple) (Tuple, error) {
	t1 := arr[0]
	t2 := arr[1]
	t3 := arr[2]

	var m21 float64 = FindM(t1, t2)
	var m32 float64 = FindM(t2, t3)
	// fmt.Println(m21, m32)

	var ah_m21 float64 = -1 / m21
	var ah_m32 float64 = -1 / m32

	p_21 := FindMiddle(t1, t2)
	p_32 := FindMiddle(t2, t3)
	if ah_m21 == ah_m32 {
		return Tuple{}, fmt.Errorf("NOT FOUND")
	}
	return FindTar(ah_m21, ah_m32, p_21, p_32), nil

}

func FindMiddle(t1, t2 Tuple) Tuple {
	return Tuple{X: (t1.X + t2.X) / 2, Y: (t1.Y + t2.Y) / 2}
}
func FindTar(m1, m2 float64, p1, p2 Tuple) Tuple {
	X_intersect := ((p2.Y - m2*p2.X) - (p1.Y - m1*p1.X)) / (m1 - m2)
	y_intersect := m1*X_intersect + (p1.Y - m1*p1.X)

	return Tuple{
		Y: y_intersect, X: X_intersect,
	}
}

func FindM(t1, t2 Tuple) float64 {
	return (float64(t2.Y) - float64(t1.Y)) / (float64(t2.X) - float64(t1.X))
}

func IsInt(X float64) bool {
	return X == float64(int64(X))
}
