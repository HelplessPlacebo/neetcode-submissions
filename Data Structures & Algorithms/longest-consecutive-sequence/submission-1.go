
func longestConsecutive(nums []int) int {
	m:= make(map[int]struct{}, len(nums))
	longest:= 0

	for _, num:= range nums {
		m[num] = struct{}{}
	}

	for _, num:= range nums {
		if _, ok:= m[num - 1]; !ok {
			seqLen:= 0
			for {
				if _, ok:= m[num + seqLen]; !ok {
					break
				}
				seqLen++
			}
			longest = max(longest, seqLen)
		}
	}

	return longest
}
