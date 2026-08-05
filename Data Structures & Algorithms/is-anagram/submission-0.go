import (
	"slices"
)

func isAnagram(s string, t string) bool {
    if (len(s) != len(t)) {
        return false
    }
	sRunes := []rune(s)
	tRunes := []rune(t)
	slices.Sort(sRunes)
	slices.Sort(tRunes)

	for ix, r := range sRunes {
		if tRunes[ix] != r {
			return false
		}
	}
	return true
}
