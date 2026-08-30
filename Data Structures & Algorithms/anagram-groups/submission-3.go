type Key26 = [26]int

func makeHashKey(str string) [26]int {
	key := Key26{}
	for i := 0; i < len(str); i++ {
		key[str[i]-'a']++
	}
	return key
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, str := range strs {
		hash := makeHashKey(str)
		groups[hash] = append(groups[hash], str)

	}

	res := make([][]string, 0, len(groups))
	for _, group := range groups {
		res = append(res, group)
	}
	return res
}