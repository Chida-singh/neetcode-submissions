func characterReplacement(s string, k int) int {
	var l, r,maxfreq,res int
	count := make(map[byte] int)

	for r < len(s){
		count[s[r]]++

		maxfreq = max(maxfreq, count[s[r]])

		for (r - l  + 1 ) - maxfreq > k{
			count[s[l]]--
			l++
		}

		res = max(maxfreq, r - l + 1)
		r++
	}

	return res	
}
