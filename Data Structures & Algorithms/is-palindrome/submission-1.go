func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isLetter(b byte) bool {
	r := rune(b)
	return unicode.IsLower(r) || unicode.IsUpper(r)
}

func isValidChar(b byte) bool {
	return isDigit(b) || isLetter(b)
}

func isPalindrome(s string) bool {
	left:= 0
	right:= len(s) - 1
	for left < right {
		leftChar:= s[left]
		rightChar:= s[right]
		if isValidChar(leftChar) && isValidChar(rightChar) {
			left++
			right --
			ls:= strings.ToLower(string(leftChar))
			rs:= strings.ToLower(string(rightChar))
			if ls != rs {
				return false
			}

		}
		if !isValidChar(leftChar) {
			left++
		}
		if !isValidChar(rightChar)  {
			right--
		}
	}
	return true
}
