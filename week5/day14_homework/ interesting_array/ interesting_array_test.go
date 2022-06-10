package interesting_array

type TestCase struct {
	input    []int
	expected string
}

func TestFindIntrestingArray(t *testing.T) {
	testcases := make([]TestCase, 0)
	testcases = append(testcases, TestCase{
		input:    []int{9, 17},
		expected: "Yes",
	})

	testcases = append(testcases, TestCase{
		input:    []int{1},
		expected: "no",
	})

	for _, testCase := range testcases {
		t.Run(fmt.Sprintf("finding for %d", testCase.input), func(tb *testing.T) {
			got := FindIntrestingArray(testCase.input)

			if got != testCase.expected {
				tb.Errorf("Tested for %d, expected %d but got %d", testCase.input, testCase.expected, got)
			}
		})
	}
}
