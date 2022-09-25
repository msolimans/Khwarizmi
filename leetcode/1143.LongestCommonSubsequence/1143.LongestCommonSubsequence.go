package longestcommonsubsequence

import "math"

//the use of dynamic programming to get longest common subsequence

//      a   b   c   d   e
//      0   1   2   3   4
//-----------------------
//a 0 | 1   1   1   1   1
//    |       \
//c 1 | 1   1   2   2   2
//    |               \
//e 2 | 1   1   2   2   3

//follow diagonals to construct longest common subsequence if needed
func longestCommonSubsequence(text1 string, text2 string) int {
	
	if len(text1) > len(text2) {
		return longestCommonSubsequence(text2, text1)
	}

	//longest common prefix - get 1 string -> match it with all other strings -> get mask then move on to next 
	
	// res := make([][]int, len(text1) * len(text2))
	res := make([][]int, len(text1))

	//O(n^2) 
	for i:=0;i<len(text1);i++{
		//init cols 
		res[i] = make([]int, len(text2))

		for j:=0;j<len(text2);j++ {
			
			//they are similar 
			if text1[i] == text2[j] {
				// diagonal + 1
				diagonal := 0 
				if i != 0 && j != 0 {
					diagonal = res[i-1][j-1]
				}
				res[i][j] = diagonal + 1

			} else {
				// max between left and top 
				top, left := 0,0
				
				if i != 0 {
					top = res[i-1][j]
				}
				if j != 0 {
					left = res[i][j-1]
				}

				res[i][j] = int(math.Max(float64(top), float64(left)))
			}

		}
	}

	//the last element in the array defines the length of the longest common subsequence
	return res[len(text1) - 1][len(text2) - 1]   
}
