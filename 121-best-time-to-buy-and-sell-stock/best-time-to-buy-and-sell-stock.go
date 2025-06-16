func maxProfit(prices []int) int {
    lowest := prices[0]
    // highest := prices[0]
    
    // profit := make(map[int]int)

    // hIndex := 0
    // lIndex :=0

    maxProfit := 0
    for i := 0; i< len(prices); i++{
        // buyState := false
        if prices[i] <lowest{
            lowest = prices[i]
        }
        profit := prices[i] - lowest
        if profit > maxProfit {
            maxProfit = profit
        }
        
    }

    return maxProfit
    
}