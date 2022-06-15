package excel_column_number

var power26 map[int]int = map[int]int{
	0: 1,
	1: 26,
	2: 676,
	3: 17576,
	4: 456976,
}
var cellnumber map[string]int = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
	"F": 6,
	"G": 7,
	"H": 8,
	"I": 9,
	"J": 10,
	"K": 11,
	"L": 12,
	"M": 13,
	"N": 14,
	"O": 15,
	"P": 16,
	"Q": 17,
	"R": 18,
	"S": 19,
	"T": 20,
	"U": 21,
	"V": 22,
	"W": 23,
	"X": 24,
	"Y": 25,
	"Z": 26,
}

func GetCoumnNumber(A string) int {
	n := len(A)
	ans := 0
	for i := 1; i < n; i++ {
		ans += power26[i]
	}
	for i, x := n-1, 0; i >= 0; i, x = i-1, x+1 {
		if x == 0 {
			ans += cellnumber[string(A[i])]
		} else {
			ans += (cellnumber[string(A[i])] - 1) * power26[x]
		}
	}

	return ans
}
