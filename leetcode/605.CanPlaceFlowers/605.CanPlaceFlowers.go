package canplaceflowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	// n is zero, no need to plant flower
	if n == 0 {
		return true 
	}
	// if n is greater than length of array, we return false 
	if n > len(flowerbed) {
		return false 
	}

	//in case only one spot is available, one case if n is less than 1 and first spot is zero (not filled)
	if len(flowerbed) == 1 {
		return n <= 1 && flowerbed[0] == 0 
	}

    
	for i := 0;i<len(flowerbed);i++ {
		//once we filled in all empty spots, we return true 
        if n == 0 {
            return true 
        }
		//3 cases (we are in first spot we only fill in that spot in case second spot is empty)
        if i == 0 {
            if  flowerbed[i] == 0 && flowerbed[i+1] == 0  {
                flowerbed[i] = 1
                n--
            }
		//in case it is in the last spot and previous spot is zero, we fill in that spot 
        } else if i == len(flowerbed) - 1 {
            if  flowerbed[i] == 0 && flowerbed[i-1] == 0  { 
                flowerbed[i] = 1 
                n--
            }
		// we have current spot is empty and prev/next spot are empty
        } else if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1 
            n--
        } 
    }
    
	//we filled in all spots 
    return n <= 0 
}