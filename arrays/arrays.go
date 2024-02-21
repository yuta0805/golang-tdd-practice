package arrays


func Sum(arr []int) int {
	var result int
	for _, v := range arr {
		result += v
	}

	return result
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// 末尾の要素を除いたスライスの合計を返す
func SumAllTails(numbersToSum ...[]int) []int {
    var sums []int
    for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))	
		}
    }

    return sums
}
