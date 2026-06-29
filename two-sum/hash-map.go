func twoSum(nums []int, target int) []int {
    indexMap := make(map[int]int)
    for i, num := range nums {
        complement := target - num
        if idx, exists := indexMap[complement]; exists {
            return []int{idx, i}
        }
        indexMap[num] = i
    }
    return []int{}
}