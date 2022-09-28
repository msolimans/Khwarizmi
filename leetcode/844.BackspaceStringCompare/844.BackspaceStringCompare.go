package backspacestringcompare

func backspaceCompare(s string, t string) bool {
	//we can use stack to remove chars in case we encountered # (use same concept as in 2390.RemovingStarsFromaString.go)
    //use the following for o(1) space 

	i, j :=len(s) - 1, len(t) - 1

	for i >= 0 || j >= 0 {
		//i>=0 and j>=0 are here to make sure we ar still inside array (otherwise no need for comparison anymore, we are waiting for other to complete)
		if i>=0 && j>=0 && s[i] == t[j] && s[i] != '#' && t[j] != '#' { //make sure these are chars 
			i--
			j--
		} else if (i>=0 && s[i] == '#') || (j>=0 && t[j] == '#') {
			
			//consider repititive hashes :) we count them and multiply count o hashes by 2 to calc how many elems to skip 
			if i>= 0 && s[i] == '#' {
				hashes := countHash(s,i)
				i = walk(s,i,hashes)
				//i-=hashes*2 //"a##c" - #a#c ith and jth should go back to -2 -2 in this case 
			}
			if j>=0 && t[j] == '#' {
				hashes := countHash(t,j)
				// j-=hashes*2
				j = walk(t,j,hashes)
			}
		} else { //they are not equal but none of them has hash in this index (not equal)
			return false 
		}
	}

	return  i < 0 && j < 0
	// aa#b#cd
	// add##cd



}

func walk(s string, start int, hashes int ) int {
	start -= hashes //now we need to count chars 
	i := start 
	for hashes > 0 {
		if i >= 0 {
			if s[i] == '#' {//add more steps to it 
				hashes++
				i--
			} else {
				i-- 
				hashes--
			}
		} else {
			i-- 
			hashes--
		} 
	}

	return i 
}
//return new index is better 
func countHash(s string, i int) int { 
	count := 0 
	for ;i>=0;i--{
		if s[i] == '#' {
			count++
		} else {
			break
		}
	}
	return count 
}