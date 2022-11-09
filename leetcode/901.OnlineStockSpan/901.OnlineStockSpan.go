package onlinestockspan

type StockSpanner struct {
    arr []int 
    spans []int 
}

//we could also use monotonic stack (that stores stack data in sorted mode)

func Constructor() StockSpanner {
    return StockSpanner{arr: []int{}, spans: []int{}}
}

func (this *StockSpanner) Next(price int) int {
	this.arr = append(this.arr, price)
    //we maintain precalc values inside spans array 
    this.spans = append(this.spans, 1)
    j := len(this.arr) - 1
    res := 0 
	//start from end of the array, loop back until we visit an item that's greater than price argument 
    for j >= 0 && this.arr[j] <= price {
		//depend on old calculated values (stored in spans array) 
        res += this.spans[j]
		//move j pointer back with how many steps we calculated back (downside of this can be o(n^2) in case of 9,8,7,6,5,4,3)
        j -= this.spans[j]
    }
    this.spans[len(this.arr) - 1] = res

    return res 
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor()
 * param_1 := obj.Next(price)
 */