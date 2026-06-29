func twoSumTwoPass(nums []int, target int) []int {
    m := make(map[int]int)

    // First pass: store all elements
    for i, num := range nums {
        m[num] = i
    }

    // Second pass: find complement
    for i, num := range nums {
        complement := target - num
        if j, found := m[complement]; found && i != j {
            return []int{i, j}
        }
    }

    return nil
}