
type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteString("$")
		sb.WriteString(str)
	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
		res := []string{}
	i := 0
	for i < len(encoded) {
		j := i
		for j < len(encoded) && encoded[j] >= '0' && encoded[j] <= '9' {
			j++
		}
		count, err := strconv.Atoi(encoded[i:j])
		if err != nil {
			break
		}
		from := j + 1
		to := from + count
		res = append(res, encoded[from:to])
		i = to
	}
	return res
}
