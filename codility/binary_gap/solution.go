package binary_gap

import (
	"strconv"
	"strings"

	"github.com/CristianCurteanu/coding_challenges/utils"
)

func BinaryGap(N int) int {
	n := int64(N)
	arr := strings.Split(strconv.FormatInt(n, 2), "")

	var counters []int
	var counter int

	for i, e := range arr {
		if e == "0" && (arr[i-1] == "1" || counter != 0) {
			counter++
		}

		if e == "0" && i == len(arr)-1 {
			break
		}

		if e == "0" && arr[i+1] == "1" {
			counters = append(counters, counter)
			counter = 0
		}
	}

	var max int
	if len(counters) != 0 {
		_, max = utils.MinMax(counters)
	}

	return max
}
