func topKFrequent(nums []int, k int) []int {

	counts := make(map[int]int)
	fr := make([][]int, len(nums)+1)

	for _, num := range nums {
		counts[num]++
	}

	for num, count := range counts {
		fr[count] = append(fr[count], num)
	}

	res := make([]int, 0, k)
	for i := len(fr) - 1; i > 0 && len(res) < k; i-- {
		for _, num := range fr[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}