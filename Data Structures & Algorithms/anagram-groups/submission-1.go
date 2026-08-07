type key26 = [26]int

func charCount(str string) key26 {
	key := key26{}
	for i := 0; i < len(str); i++ {
		key[str[i]-'a']++
	}
	return key
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[key26][]string)

	for _, str := range strs {
		hash := charCount(str)
		groups[hash] = append(groups[hash], str)

	}

	res := make([][]string, 0, len(groups))
	for _, group := range groups {
		res = append(res, group)
	}
	return res
}
