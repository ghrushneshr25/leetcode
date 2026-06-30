/*
@leetcode

name: Hash Map
algorithm: |-
    Initialize an empty hash map to store value → index.
    Iterate through the array from the first element to the last.
    For each element:
          Compute the complement:
          complement = target - current_element
    Check whether the complement exists in the hash map.
    If it exists, return the stored index and the current index.
    Otherwise, insert the current element and its index into the hash map.
    Repeat until the pair is found.
    Return the indices of the two numbers.
complexity:
    time: O(n)
    space: O(n)
notes: I use a hash map to store previously seen numbers and their indices. For each element, I compute its complement (target - num) and check if it already exists in the map. If it does, I've found the pair; otherwise, I store the current number. This achieves O(n) time and O(n) space complexity.
@endleetcode
*/

func twoSum(nums []int, target int) []int {
    hashMap := map[int]int{}
    for index , num := range nums {
        checkVal := target - num 
        if val, ok := hashMap[checkVal]; ok {
            return []int{val, index}
        }
        hashMap[num] = index
    }
    return []int{0,0}
}