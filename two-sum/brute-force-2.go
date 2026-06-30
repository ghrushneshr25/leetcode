/*
@leetcode

name: Brute force - 2
algorithm: ""
complexity:
    time: ""
    space: ""
notes: ""
@endleetcode
*/

func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i]+nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return nil
}