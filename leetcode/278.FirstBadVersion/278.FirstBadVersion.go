package firstbadversion

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

 func firstBadVersion(n int) int {
    //binary search 
    s := 1
    e := n 
    bad := 0
    
    for s <= e {
        mid := (s+e)/2 
        if !isBadVersion(mid) { //not bad meaning all previous versions are good 
            s = mid + 1
        } else { //mid is bad means we need to look back and see if we can get a nother bad one previous 
            bad = mid 
            e = mid - 1
        }
        
    }
    
    return bad 
    
}



//[0,1,2]
//[1,2,4]
///################
//used only for mocking isBadVersion 
func isBadVersion(n int) bool {
	//bad versions are 12 
	
	ns := []int{1,2,3,4,5,6,7,8,9,10,12,13,14,15,16}

	for i := 0;i<len(ns) - 1;i++ {
		if ns[i] - i > 1 {
			return ns[i] == n 
		}
	}

	return false
}