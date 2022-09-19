package another_sequence_problem

var data = []int{1, 1, 2}

func GetValueAnotherSeq(A int) int {

	if A < len(data) {
		return data[A]
	}

	return GetValueAnotherSeq(A-1) + GetValueAnotherSeq(A-2) + GetValueAnotherSeq(A-3) + A
}
