package reversevowelsofastring

func reverseVowels(s string) string {
	i, j := 0, len(s) - 1 
	vowels := map[byte]bool{'a':true, 'e': true, 'i': true, 'o': true, 'u': true, 'A' : true, 'E': true, 'I': true, 'O': true, 'U': true}

	res := make([]byte, len(s))
	
	for i <= j {
		 _, iexist := vowels[s[i]]
			 _, jexist := vowels[s[j]]
//hello 
// i j
//holle	
		if !iexist && !jexist { //replace 
			//put them as is 
			res[i] = s[i]
			res[j] = s[j]
			i++
			j--
		} else if iexist && jexist {
			//replace 
			res[i], res[j] = s[j], s[i]
			i++ 
			j--
		} else if iexist { 
			res[i] = s[i]//just in case we have a character that doesn't have a counterpart (it might get replaced in future with j or not)
			res[j] = s[j]
			j--
		} else { //j exists 
			res[i] = s[i]
			res[j] = s[j]//just in case we have a character that doesn't have a counterpart (it might get replaced in future with i or not)
			i++
		}

	}
	return string(res)    
}

 