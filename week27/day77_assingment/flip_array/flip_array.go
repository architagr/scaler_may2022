package flip_array

func FlipArray(A []int) int {
	n := len(A)

	total := 0
	for i := 0; i < n; i++ {
		total += A[i]
	}
	total /= 2
	dp := make([][]int, 2)
	dp[0] = make([]int, total+1)
	dp[1] = make([]int, total+1)

	for j := 1; j <= total; j++ {
		dp[0][j] = n + 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= total; j++ {

			if j >= A[i-1] {
				dp[i%2][j] = minValue(dp[(i-1)%2][j], 1+dp[(i-1)%2][j-A[i-1]])
			} else {
				dp[i%2][j] = dp[(i-1)%2][j]
			}

		}
	}
	for col := total; col >= 0; col-- {
		minCoins := dp[n%2][col]
		if minCoins >= 1 && minCoins <= n {
			return minCoins
		}
	}

	return 0
}
func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func getTotalSum(A []int) int {
// 	sum := 0
// 	for i := 0; i < len(A); i++ {
// 		sum += A[i]
// 	}
// 	return sum
// }
// func FlipArray(A []int) int {
// 	coins := len(A)
// 	totalSum := getTotalSum(A)
// 	halfSum := totalSum / 2

// 	dp := make([][]int, coins+1)
// 	for i := 0; i <= coins; i++ {
// 		dp[i] = make([]int, halfSum+1)
// 	}

// 	// dp[i][j] => min #coins to get a sum=j, till i coins.

// 	// base case - I: 0 coins => dp[0][j] => n+1

// 	// no answer can be greater than `n` coins i.e., there cannot be > `n` coins to pick, and so `n+1` acts as a decent INT_MAX.

// 	for j := 1; j <= halfSum; j++ {
// 		dp[0][j] = coins + 1
// 	}

// 	// base case -II: 0 sum   => dp[i][0] => 0

// 	// This we don't need to fill since static int[]'s are all 0s by default.

// 	// This is Normal Min #Coins to Pick a given Sum's solution

// 	for i := 1; i <= coins; i++ {
// 		for j := 1; j <= halfSum; j++ {
// 			if j >= A[i-1] {
// 				dp[i][j] = minValue(1+dp[i-1][j-A[i-1]], dp[i-1][j])
// 			} else {
// 				dp[i][j] = dp[i-1][j]
// 			}
// 		}
// 	}
// 	// Iterate through the last row to get the min #coins (i.e., min #elements to negate, for the required min. sum)
// 	// Last Row => We want to consider all the `n` elements, and also, each column in a row represents a sum.
// 	// Check the difference of [(TS/2+1) - (TS/2-1)], [(TS/2+2) - (TS/2-2)], [(TS/2+2) - (TS/2-2)], and so on...
// 	// from min possible difference to maximum possible difference.
// 	// If there's any number >= 1 AND <= totalCoins (A.length),

// 	// that numbe is the min #elements to make negative,

// 	// to  get minimum possible sum of ((TS/2)+k) - ((TS/2)-k)

// 	for col := halfSum; col >= 0; col-- {

// 		minCoins := dp[coins][col]
// 		if minCoins >= 1 && minCoins <= coins {
// 			return minCoins
// 		}
// 	}
// 	return 0
// }

// func minValue(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
