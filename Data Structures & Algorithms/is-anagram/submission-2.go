type Int26 [26]int
func isAnagram(s string, t string) bool {
 if len(s) != len(t) {
	return false
 }

 var int26S Int26
 var int26T Int26
 for i:= 0; i< len(s); i++ {
	int26S[ s[i] - 'a'] ++
 }

  for j:= 0; j< len(t); j++ {
	int26T[t[j] - 'a'] ++
 }

 for ix, v:= range int26S {
	if v != int26T[ix] {
		return false
	}
 }

 return true
}
