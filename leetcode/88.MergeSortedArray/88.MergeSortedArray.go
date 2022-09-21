package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int)  {
	
	//choose whats largest from right most sides and put in empty spaces (sort of greedy)
	c := len(nums1) - 1
	n1 := m - 1
	n2 := n - 1 // last element in nums2 

	for c >= 0 {
		if n1 >= 0 && n2 >= 0 {
			if  nums1[n1] > nums2[n2] {
				nums1[c] = nums1[n1]
				c--
				n1--
			} else {
				nums1[c] = nums2[n2]
				c--
				n2--
			} 
		} else if n2 >= 0 { //left over from nums2 if any 
			nums1[c] = nums2[n2]
			c-- 
			n1-- 
		} else {
			break
		}
	}	

}