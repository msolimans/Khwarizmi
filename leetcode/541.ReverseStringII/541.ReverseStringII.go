package reversewordsinastringiii


func reverseStr(s string, k int) string {
    chs := make([]byte, len(s))
    for i := 0; i<len(s); i++ {
        chs[i] = s[i]
    }
    
    reverseStrUsingIndex(chs, k, 0)
	
	return string(chs)
    
}

func reverseStrUsingIndex(chs []byte, k int, at int) {
    //do reverse here 
    //reverse first k and jump k 
    if at > len(chs) {
        return 
    }

    res := reverse(chs, at, at+k-1)
    at += k + k //2k
    reverseStrUsingIndex(res, k, at)
}

func reverse(chs []byte, start, end int) []byte {
    if start > len(chs) {
        return chs
    }
    
    //out of boudary, shrink it 
    if end > len(chs) - 1 {
        end = len(chs) - 1
    }
    
    
    for start < end { 
        temp := chs[start]
        chs[start] = chs[end]
        chs[end] = temp
        start++
        end--
    }
    
    return chs
    
}