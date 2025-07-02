# Intuition 
<!-- Describe the first thought you've had when approaching the problem. -->
- The first thought was the two pointer approch altough I don't think I'd do it the way it's supposed to be done.

# Approach
<!-- Describe your approach to solve this problem. -->
1. The edge case of the slice being too short to use the two pointers.
2. Then the next logic applies:
   1. Pointers are moving synchronously when they are pointing at nums that are different, eg `[1 2 3s 4f 5 6] --> [1 2 3 4s 5f 6]`
   2. When the duplicates are encountered, eg `[1 2s 2f 2 3]` the **fast** pointer is incremented by the loop and **nothing else** is happening at the time, the slow one stays in its place, eg `[1 2s 2f 2 3] --> [1 2s 2 2f 3]`
   3. When they are different again, the slow pointer moves one position and `nums[slow] = nums[fast]` happens, now we have a `[1 2 3s 2 3f]`, the `3` is then returned 


# Complexity
- Time complexity: basically iterating trough the whole slice.
$$O(n)$$
<!-- Add time complexity here and explain why its the case here -->

- Space complexity: since the changes are all in place
$$O(1)$$
<!-- Add space complexity-->


# Code
```go []
package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
```