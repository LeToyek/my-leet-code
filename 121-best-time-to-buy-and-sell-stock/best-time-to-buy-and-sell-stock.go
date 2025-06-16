func maxProfit(prices []int) int {
    minPrice, maxProfit := prices[0], 0

    for _, price := range prices {
        if price < minPrice {
            minPrice = price
        } else if price - minPrice > maxProfit {
            maxProfit = price - minPrice
        }
    }

    return maxProfit
}
