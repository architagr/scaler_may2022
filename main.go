package main

import (
	"bufio"
	"fmt"
	"os"
	"scaler-may-22/common"
	"scaler-may-22/week24/day68_assingment/auto_complete"
)

func main() {
	// YOUR CODE GOES HERE
	// Please take input and print output to standard input/output (stdin/stdout)
	// E.g. 'fmt.Scanf' for input & 'fmt.Printf' for output

	/*
		4
		6 3
		abcd aecd abaa abef acdcc acbcc
		2 1 3 4 6 5
		ab abc a
		--ans
		abef abaa abcd
		abcd
		acdcc acbcc abef abaa abcd

		5 3
		aaaa aacd abaa abaa aadcc
		3 4 1 2 6
		aa aba abc
		---ans
		aadcc aacd aaaa
		abaa abaa
		-1

		2 3
		aeedd eaa
		3 11
		aee aeed aee
		--ans
		aeedd
		aeedd
		aeedd

		14 8
		bbd cbbdec eaedbcdd bbddc ccced ed aacdbac dbaae a bddedee abbbdee dcddceb aede c
		35 8 54 12 21 46 27 19 66 76 32 58 69 36
		a eaedbcd eaedb cbb ccced cbbd a cc
		-- ans
		aede a abbbdee aacdbac
		eaedbcdd
		eaedbcdd
		cbbdec
		ccced
		cbbdec
		aede a abbbdee aacdbac
		ccced
	*/

	in := bufio.NewReader(os.Stdin)
	t := common.ReadInt(in)
	result := make([][]string, 0)
	for i := 0; i < t; i++ {
		common.ReadArrInt(in)
		words := common.ReadStrArr(in)
		weight := common.ReadArrInt(in)
		prefix := common.ReadStrArr(in)
		root := auto_complete.CreateDictonary(words, weight)

		for _, pre := range prefix {
			res := auto_complete.Search(pre, root)
			if len(res) == 0 {
				result = append(result, []string{"-1"})
			} else {
				result = append(result, res)
			}
		}

	}
	for i := 0; i < len(result); i++ {
		ans := ""
		for j := 0; j < len(result[i]); j++ {
			ans += fmt.Sprintf("%s ", result[i][j])
		}
		fmt.Println(ans)
	}
}
