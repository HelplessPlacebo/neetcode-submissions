
type Int26 [26]int

func hash(str string) Int26 {
	var h Int26 
	for i:= 0; i < len(str); i++ {
		h[str[i] - 'a'] ++ 
	}
	return h 
}

// n - len of strsm , m - longest str in strs
// time : O(n*m + n) -> O(n*m)
// memory: O(n * m)
func groupAnagrams(strs []string) [][]string {
	res:= [][]string{}
	m:= make(map[Int26][]string)

	for i:= 0; i < len(strs); i++ {
		hashed:= hash(strs[i])

		m[hashed] = append(m[hashed], strs[i])
	}

	for _, v:= range m {
		res = append(res, v)
	}
	return res
}
