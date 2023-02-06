package batches

func countBatches(A int, B []int, C [][]int, D int) int {
	count := 0
	visited := make([]bool, A+1)
	adjencyList := getAdjencyList(C)
	for i := 1; i <= A; i++ {
		if !visited[i] {
			visited[i] = true
			sum := getStrength(adjencyList, visited, B, i, B[i-1])
			if sum >= D {
				count++
			}
		}
	}
	return count
}

func getStrength(adjencyList map[int][]int, visited []bool, strength []int, node int, sum int) int {
	related := adjencyList[node]
	for _, n := range related {
		if !visited[n] {
			visited[n] = true
			sum = getStrength(adjencyList, visited, strength, n, sum+strength[n-1])
		}
	}
	return sum
}
func getAdjencyList(relations [][]int) map[int][]int {
	adjencyList := make(map[int][]int)
	for _, relation := range relations {
		a, b := relation[0], relation[1]
		nodeA, okA := adjencyList[a]
		nodeB, okB := adjencyList[b]
		if !okA {
			nodeA = make([]int, 0)
		}
		if !okB {
			nodeB = make([]int, 0)
		}
		nodeA = append(nodeA, b)
		nodeB = append(nodeB, a)
		adjencyList[a] = nodeA
		adjencyList[b] = nodeB
	}

	return adjencyList
}
