func findMin(nums []int) int {
	var l,r int
	l = 0
	r = len(nums) - 1

	for l < r{
		m := l+ ((r - l)/2)

		if nums[m] > nums[r]{  // middle value greater than right, which means smallest value present on the right 
			l = m+1 // searching in the right side 
		} else {
			r = m // search in the left side
		} 
		//condition will fail when l is on the smallest value
	}

	return nums[l]
}
