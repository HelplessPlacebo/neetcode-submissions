
func topKFrequent(nums []int, k int) []int {
	counts:= make(map[int]int)

	for _, num := range nums {
		counts[num] ++
	}

	groups:= make([][]int, len(nums)+1)
	for num, count := range counts {
		groups[count] = append(groups[count], num)
	}

	res:= make([]int, 0, k)
	for i := len(groups) - 1; i > 0 && len(res) < k; i-- {
		for _, num := range groups[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}