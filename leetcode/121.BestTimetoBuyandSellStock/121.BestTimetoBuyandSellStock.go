package besttimetobuyandsellstock

func maxProfit(prices []int) int {

    least := prices[0]
	max := 0 

	for i := 1;i<len(prices);i++ {
		if prices[i] < least {
			least = prices[i]
		} else {
			//greater than least (visited earlier) check max now and update if needed
			if prices[i] - least > max {
				max = prices[i] - least 
			}
		}
	}

	return max 


}