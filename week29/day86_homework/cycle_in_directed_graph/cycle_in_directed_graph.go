package cycle_in_directed_graph

var f int

func CycleDirectedGraph(A int, B [][]int) int {
	at := make(map[int]map[int]bool)
	for i := 0; i < len(B); i++ {
		val, ok := at[B[i][0]]
		if !ok {
			val = make(map[int]bool)
		}
		val[B[i][1]] = true
		at[B[i][0]] = val
	}
	visited := make([]int, A+1)
	for i := 0; i <= A; i++ {
		visited[i] = -1
	}
	for i := 1; i <= A; i++ {
		if visited[i] == -1 {
			f = 0
			check(at, visited, i)
			if f == 1 {
				return 1
			}

		}
	}
	return 0
}

func check(g map[int]map[int]bool, vis []int, i int) {
	vis[i] = 1
	for v := range g[i] {
		if vis[v] == -1 {
			check(g, vis, v)
		} else if vis[v] == 1 {
			f = 1
		}
	}
	vis[i] = 2
}
