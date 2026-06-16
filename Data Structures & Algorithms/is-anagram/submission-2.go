func isAnagram(s string, t string) bool {
	
	if len(s) != len(t){
		return false
	}
	array := [26]int{}
	runeS := []rune(s)
	runeT := []rune(t)
		
	for i:= 0; i<len(runeS); i++{
		array[runeS[i] - 'a']++
		array[runeT[i] - 'a']--
	}

	for _,v := range array{
		if v != 0{
			return false
		}
		
	}
	return true

	
}
