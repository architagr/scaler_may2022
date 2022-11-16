package finish_maximum_jobs

import "sort"

type Task struct {
	start, end int
}

func MaxJob(A, B []int) int {
	jobs := make([]Task, len(A))
	for i := 0; i < len(A); i++ {
		jobs[i] = Task{
			start: A[i],
			end:   B[i],
		}
	}
	sort.Slice(jobs, func(i, j int) bool {
		if jobs[i].end < jobs[j].end {
			return true
		} else if jobs[i].end > jobs[j].end {
			return false
		} else if jobs[i].start < jobs[j].start {
			return true
		}
		return false
	})

	count := 1
	end := jobs[0].end
	for i := 1; i < len(jobs); i++ {
		if end <= jobs[i].start {
			count++
			end = jobs[i].end
		}
	}
	return count
}
