package another_coin_problem

func AnotherCoinProblem(A int) int {

	ans := 0
	for A > 0 {
		ans += A % 5
		A /= 5
	}
	return ans
}
