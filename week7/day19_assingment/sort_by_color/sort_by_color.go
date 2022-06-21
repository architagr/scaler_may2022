package sort_by_color

func SortByColor(A []int) []int {
	zero := 0
	two := len(A) - 1
	one := 0

	//one<=two not one<two because two represents next nextIndex where 2 can be stored
	for one <= two {
		if A[one] == 0 {
			swap(A, one, zero)
			zero++
			one++
		} else if A[one] == 2 {
			swap(A, one, two)
			two--
		} else {
			one++
		}
	}
	return A
}
func swap(a []int, m, n int) {
	tem := a[m]
	a[m] = a[n]
	a[n] = tem
}
