package validanagram

import (
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	//XOR is not good here [in case of similar chars in but different in strings] like "aa" and "bb" or "xaabby" and "xddrry"
	 mp := map[rune]int{}
	 
	 //we add all chars from first string 
	 for _,c := range s {
		 mp[c] = mp[c] + 1
	 }
	 
	 //we check second string and compare with map 
	 for _,c := range t {
		 if v, exists := mp[c]; exists {
			 if v == 1 {
				 //remove (we should get final map empty)
				 delete(mp,c)
			 } else  {
				 mp[c] = v - 1 
			 }
		 } else { //if not found in map, return false 
			 return false 
		 }
	 }
	 
	 return len(mp) == 0 
	 
	 
 }

 ///////

func isAnagram2(s string, t string) bool {
	//sort both strings and compare if they are equal 

	s = sortStr(s)
	t = sortStr(t) 
	
	return s == t 
}


func sortStr(s string) string {
	ss := strings.Split(s, "")
	sort.Strings(ss)

	return strings.Join(ss,"")
}