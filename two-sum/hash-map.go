
@leetcode

name: Hash Map
algorithm: Iterate once through the array while storing visited values in a hash map. For each number, check if its complement exists before inserting the current value.
complexity:
    time: O(n)
    space: O(n)
notes: Classic single-pass hash map solution.
@endleetcode


func twoSum(nums []int, target int) []int {
	seen := map[int]int{}
	for i, num := range nums {
		if j, ok := seen[target-num]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}
	return nil
}