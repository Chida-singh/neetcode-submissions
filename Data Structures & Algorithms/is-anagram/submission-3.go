func isAnagram(s string, t string) bool {
	
	if len(s) != len(t){
		return false
	}
	array := [26]int{}
	
		
	for i:= 0; i<len(s); i++{
		array[s[i] - 'a']++
		array[t[i] - 'a']--
	}

	for _,v := range array{
		if v != 0{
			return false
		}
		
	}
	return true

	
}
