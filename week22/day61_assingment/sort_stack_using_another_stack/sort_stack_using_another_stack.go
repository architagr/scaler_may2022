package sort_stack_using_another_stack

/**
 * @input A : Integer array
 *
 * @Output Integer array.
 */
func SortStack(A []int) []int {
	altStackTop := -1
	alt := make([]int, len(A))
	for i := len(A) - 1; i >= 0; i-- {
		x := A[i]
		for altStackTop > -1 && alt[altStackTop] < x {
			A[i] = alt[altStackTop]
			i++
			altStackTop--
		}
		altStackTop++
		alt[altStackTop] = x
	}
	l, r := 0, len(A)-1
	for l <= r {
		alt[l], alt[r] = alt[r], alt[l]
		l++
		r--
	}
	return alt
}
