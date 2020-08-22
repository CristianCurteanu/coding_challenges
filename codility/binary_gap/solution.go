package binary_gap

import (
	"strconv"
	"strings"
)

// func main() {
// 	fmt.Println("--result(1041):", BinaryGap(1041))
// 	fmt.Println("--result(529):", BinaryGap(529))
// 	fmt.Println("--result(32):", BinaryGap(32))
// 	fmt.Println("--result(20):", BinaryGap(20))
// 	fmt.Println("--result(9):", BinaryGap(9))
// }

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
		_, max = MinMax(counters)
	}

	return max
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
