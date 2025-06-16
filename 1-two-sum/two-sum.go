func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    for i := 0; i< len(nums); i++{
        complements := target - nums[i]
        if j, found := seen[complements]; found{
            return []int{i,j}
        }
        seen[nums[i]] = i
    }
    return nil
}