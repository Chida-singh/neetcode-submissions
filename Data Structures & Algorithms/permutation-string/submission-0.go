func checkInclusion(s1 string, s2 string) bool {

	if len(s2) < len(s1){
		return false
	}

	var s1Count [26]int
	var s2Count [26]int


	for i:= 0; i < len(s1);i++{
		s1Count[s1[i] - 'a']++
		s2Count[s2[i] - 'a']++
	}

	if s1Count == s2Count{
		return true
	}

	for right := len(s1); right < len(s2); right++{
		left := right - len(s1)

		s2Count[s2[left] - 'a']--
		s2Count[s2[right] - 'a']++


		if s1Count == s2Count{
			return true
		}
	}
	return false
}
