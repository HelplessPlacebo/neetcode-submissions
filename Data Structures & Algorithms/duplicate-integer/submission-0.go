func hasDuplicate(nums []int) bool {
	alreadyUsed := make(map[int]int)

	for _, num := range nums {
		_, exist := alreadyUsed[num]
		if exist {
			return true
		}
		alreadyUsed[num] = num
	}
	return false
}

