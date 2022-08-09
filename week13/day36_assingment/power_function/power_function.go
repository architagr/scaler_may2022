package power_function

func Pow(A, B, C int) int {
	if A == 0 {
		return 0
	}
	value := getPow(A, B, C)
	if value >= 0 {
		return value
	} else {
		return value + C
	}

}

func getPow(A, B, C int) int {
	if B == 0 {
		return 1
	}
	A = A % C
	if B%2 == 0 {
		return getPow((A*A)%C, B/2, C) % C
	} else {
		return (A * getPow((A*A)%C, B/2, C)) % C
	}
}
