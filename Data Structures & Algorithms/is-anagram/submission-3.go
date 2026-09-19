type AlphEl [26]int

func toElphEls(str string) AlphEl {
	res:= AlphEl{}

	for i:=0; i < len(str); i++ {
		res[str[i] - 'a']++
	}

	return res
}

func isAnagram(s string, t string) bool {
	return toElphEls(s) == toElphEls(t)
}
