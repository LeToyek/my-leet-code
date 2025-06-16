func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    for i:= 0; i < len(nums);i++{
        complement := target- nums[i] 
        if j,found := seen[complement];found{
            return []int{i,j}
        }
        seen[nums[i]] = i
    }
    return nil
}