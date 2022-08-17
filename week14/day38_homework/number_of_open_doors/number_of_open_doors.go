package number_of_open_doors

import "math"

func NumberOfOpenDoors(A int) int {
	return int(math.Sqrt(float64(A)))
}
