package permutation

import (
	"unicode"
)

func  permutateWord(word, charSet string) []string {

	mpSet := map[byte]bool {}

	//record those in char set 
	for i:=0;i<len(charSet);i++ {
		mpSet[charSet[i]] = true 
	}

	//find indexes of those charSet 
	indices := []int{}
	for i := 0;i<len(word);i++ {
		if _,exists := mpSet[word[i]]; exists{
			indices = append(indices, i)
		}
	}

	return permutate(indices, word, 0)



}

func permutate(indices []int, word string, j int) []string {

	//medium-one | mdo
	//m d  m o 
	//M d  m o
	//m D  m o 
	//m d  M o
	//m d  m O
	//M D  m o 
	//m D  M o
	//m d  M O
	//M D  M o 
	//m D  M O 
	//M D  M O
	
	res := []string{} 
	res = append(res, word)

	for i := 0;i<len(indices);i++ {
		
		for j := i;j<len(indices);j++ {
			w := extractWord([]byte(word), indices, i, j)
			res = append(res, w)
		}
	}


	return res 
}

func extractWord(word []byte, indices []int, s,e int)  string {
	//capitalize only those in range s ... e 
	for i := s; i <= e; i++ {
		targetIndex := indices[i]
		word[targetIndex] = byte(unicode.ToUpper(rune(word[targetIndex])))
	}

	return string(word) 
} 

