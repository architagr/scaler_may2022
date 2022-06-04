package scaler_product

func FindScalerProduct(A [][]int, B int) [][]int {

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			A[i][j] *= B
		}
	}
	return A
}
