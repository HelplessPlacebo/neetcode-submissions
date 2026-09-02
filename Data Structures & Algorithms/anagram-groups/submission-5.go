type Int26 [26]int

func makeGroupKey(str string) Int26 {
	k:= Int26{}
	for i:= 0; i<len(str); i++ {
		k[str[i] - 'a'] ++
	}

	return k
}

func groupAnagrams(strs []string) [][]string {
	groups:= make(map[Int26][]int)

	for ix, str:= range strs {
		k:= makeGroupKey(str)
		groups[k] = append(groups[k], ix)
	}

	res:= make([][]string, 0, len(groups))

	for _, strsI:= range groups {
		s := []string{}
		for _, i:= range strsI{
			s = append(s, strs[i])
		}
		res = append(res, s) 
	}
	return res
}