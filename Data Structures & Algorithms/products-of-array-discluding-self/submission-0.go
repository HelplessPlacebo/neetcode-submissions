func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	prev := 1
	for i := range nums {
		res[i] = prev
		prev *= nums[i]
	}

	next := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= next
		next *= nums[i]
	}

	return res
}