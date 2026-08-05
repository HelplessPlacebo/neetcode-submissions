func twoSum(nums []int, target int) []int {
	remains := make(map[int]int)
	result := make([]int, 2)

	for ix, num := range nums {
		vIx, exist := remains[num]
		if exist {
			if vIx > ix {
				result = []int{ix, vIx}
			}
			if ix > vIx {
				result = []int{vIx, ix}
			}
			break
		}

		remain := target - num
		if _, ok := remains[remain]; !ok {
			remains[remain] = ix
		}
	}
	return result
}
