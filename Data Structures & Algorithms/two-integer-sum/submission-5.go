func twoSum(nums []int, target int) []int {
	remains := make(map[int]int)

	for ix, num := range nums {
		if i, ok := remains[num]; ok {
			return []int{i, ix}
		}
		remains[target-num] = ix

	}

	return []int{}
}