package gray_code

import "math"

func GetGrayCode(n int) []int {
	count := int(math.Pow(2, float64(n))) - 1
	result := make([]int, 0)
	return genrateCode(count, result)
}
func genrateCode(i int, ans []int) []int {
	if i == 0 {
		return append(ans, 0)
	}
	ans = genrateCode(i-1, ans)
	ans = append(ans, i^(i>>1))
	return ans
}
