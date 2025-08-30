package sorting

/*
Given an unsorted array, find the maximum difference between the successive elements in its sorted form.

Return 0 if the array contains less than 2 elements.

Example 1:


Input: [3,6,9,1]
Output: 3
Explanation: The sorted form of the array is [1,3,6,9], either
             (3,6) or (6,9) has the maximum difference 3.

Example 2:


Input: [10]
Output: 0
Explanation: The array contains less than 2 elements, therefore return 0.

Note:

You may assume all elements in the array are non-negative integers and fit in the 32-bit signed integer range.
Try to solve it in linear time/space.

题目大意 #
在数组中找到 2 个数字之间最大的间隔。要求尽量用 O(1) 的时间复杂度和空间复杂度。

解题思路 #
虽然使用排序算法可以 AC 这道题。先排序，然后依次计算数组中两两数字之间的间隔，找到最大的一个间隔输出即可。

这道题满足要求的做法是基数排序。
*/

// quickSort
func maximumGap1(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	quickSort164(nums, 0, n-1)
	res := 0
	for i := 0; i < n-1; i++ {
		res = max(nums[i+1]-nums[i], res)
	}
	return res
}

func quickSort164(a []int, low, high int) {
	if low > high {
		return
	}
	p := partition164(a, low, high)
	quickSort164(a, low, p-1)
	quickSort164(a, p+1, high)
}

func partition164(a []int, low, high int) int {
	pivot := a[high]
	i := low - 1
	for j := low; j < high; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	// pivot and the first e larger switch
	a[i+1], a[high] = a[high], a[i+1]
	return i + 1
}

// Radix Sort - optimized with bit operations for linear time complexity
// Time: O(d*n) where d=4 for 32-bit integers, Space: O(n+k) where k=256
// Example walkthrough with nums = [170, 45, 802]
func maximumGap2(nums []int) int {
	// Base case: need at least 2 elements to have a gap
	if len(nums) < 2 { // omit nil check , because len(nil) is 0.
		return 0
	}

	// Step 1: Find the maximum number to determine how many bits we need to process
	// Example: max(170, 45, 802) = 802
	// 802 in binary = 1100100010 (10 bits, so we need 2 passes of 8 bits each)
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > m {
			m = nums[i]
		}
	}

	// Step 2: Setup radix sort parameters
	R := 256                      // Radix = 2^8 = 256 buckets (process 8 bits at a time)
	mask := R - 1                 // 255 in binary: 11111111 (used to extract 8 bits)
	aux := make([]int, len(nums)) // Auxiliary array for stable sorting

	// Step 3: Process 8 bits at a time, from least significant to most significant
	// Example: First pass (shift=0), second pass (shift=8)
	for shift := 0; m>>shift > 0; shift += 8 {
		// Step 3a: Initialize counting array for current 8-bit digit
		count := make([]int, R) // count[i] = number of elements with digit i

		// Step 3b: Count frequency of each 8-bit digit at current position
		// Example (shift=0): 170&255=170, 45&255=45, 802&255=34
		// count[170]=1, count[45]=1, count[34]=1
		for i := 0; i < len(nums); i++ {
			// Extract 8-bit digit: shift right to get desired position,
			// then AND with mask to get only 8 bits
			// Example: (802 >> 0) & 255 = 802 & 255 = 34 (last 8 bits)
			//          (802 >> 8) & 255 = 3 & 255 = 3 (next 8 bits)
			digit := (nums[i] >> shift) & mask
			count[digit]++
		}

		// Step 3c: Convert counts to starting positions (cumulative sum)
		// Example: count=[0,0,...,1,0,...,1,0,...,1,0,...] (at positions 34,45,170)
		// After cumsum: count[34]=1, count[45]=2, count[170]=3
		// This means: digit 34 starts at pos 0, digit 45 at pos 1, digit 170 at pos 2
		for i := 1; i < R; i++ {
			count[i] += count[i-1]
		}

		// Step 3d: Build sorted array for current digit position
		// Process from right to left to maintain stability
		// Example: 802(digit=34): count[34]=1→0, aux[0]=802
		//          45(digit=45):  count[45]=2→1, aux[1]=45
		//          170(digit=170): count[170]=3→2, aux[2]=170
		// Result: aux=[802, 45, 170]
		for i := len(nums) - 1; i >= 0; i-- {
			digit := (nums[i] >> shift) & mask
			count[digit]--              // Decrement to get correct position
			aux[count[digit]] = nums[i] // Place element in correct position
		}

		// Step 3e: Copy sorted result back to original array
		// After first pass (shift=0): nums=[802, 45, 170] (sorted by last 8 bits)
		// After second pass (shift=8): nums=[45, 170, 802] (fully sorted)
		copy(nums, aux)
	}

	// Step 4: Find maximum gap between consecutive elements in sorted array
	// Example: gaps are [170-45=125, 802-170=632], max = 632
	maxGap := 0
	for i := 1; i < len(nums); i++ {
		gap := nums[i] - nums[i-1]
		if gap > maxGap {
			maxGap = gap
		}
	}

	return maxGap
}

// ASCII Example: Radix Sort Process Visualization
/*
Example: nums = [170, 45, 75, 90, 2, 802, 24, 66]

Step 1: Convert to binary and show 8-bit chunks
170 = 10101010 = [10101010] (8 bits)
45  = 00101101 = [00101101]
75  = 01001011 = [01001011]
90  = 01011010 = [01011010]
2   = 00000010 = [00000010]
802 = 1100100010 = [00000011][00100010] (16 bits, needs 2 chunks)
24  = 00011000 = [00011000]
66  = 01000010 = [01000010]

Processing from Least Significant 8 bits (shift=0):
┌─────────────────────────────────────────────────────────────┐
│                    First 8-bit Pass                        │
└─────────────────────────────────────────────────────────────┘

Original: [170, 45, 75, 90, 2, 802, 24, 66]
LSB 8bits:[170, 45, 75, 90, 2, 34, 24, 66] (802&255=34)

Bucket sort by last 8 bits:
Bucket[2]:   [2]                    (00000010)
Bucket[24]:  [24]                   (00011000)
Bucket[34]:  [802]                  (00100010)
Bucket[45]:  [45]                   (00101101)
Bucket[66]:  [66]                   (01000010)
Bucket[75]:  [75]                   (01001011)
Bucket[90]:  [90]                   (01011010)
Bucket[170]: [170]                  (10101010)

After pass 1: [2, 24, 802, 45, 66, 75, 90, 170]

Processing next 8 bits (shift=8):
┌─────────────────────────────────────────────────────────────┐
│                    Second 8-bit Pass                       │
└─────────────────────────────────────────────────────────────┘

Current: [2, 24, 802, 45, 66, 75, 90, 170]
Next8bit:[0,  0,  3,  0,  0,  0,  0,   0] (values>>8)

Bucket sort by next 8 bits:
Bucket[0]: [2, 24, 45, 66, 75, 90, 170]  (all have 0 in upper bits)
Bucket[3]: [802]                         (802>>8 = 3)

Final sorted: [2, 24, 45, 66, 75, 90, 170, 802]

Visual representation:
Original:  [170, 45, 75, 90, 2, 802, 24, 66]
           ┌─┐ ┌─┐ ┌─┐ ┌─┐ ┌┐ ┌──┐ ┌─┐ ┌─┐
           │?│ │?│ │?│ │?│ │?│ │? │ │?│ │?│
           └─┘ └─┘ └─┘ └─┘ └┘ └──┘ └─┘ └─┘
              ↓ Sort by LSB 8 bits ↓
Pass 1:    [2, 24, 802, 45, 66, 75, 90, 170]
           ┌┐ ┌─┐ ┌──┐ ┌─┐ ┌─┐ ┌─┐ ┌─┐ ┌──┐
           │✓│ │✓│ │? │ │✓│ │✓│ │✓│ │✓│ │✓ │
           └┘ └─┘ └──┘ └─┘ └─┘ └─┘ └─┘ └──┘
              ↓ Sort by next 8 bits ↓
Final:     [2, 24, 45, 66, 75, 90, 170, 802]
           ┌┐ ┌─┐ ┌─┐ ┌─┐ ┌─┐ ┌─┐ ┌──┐ ┌──┐
           │✓│ │✓│ │✓│ │✓│ │✓│ │✓│ │✓ │ │✓ │
           └┘ └─┘ └─┘ └─┘ └─┘ └─┘ └──┘ └──┘

Maximum Gap: 802 - 170 = 632
*/

// Most Significant Digit first - Demo function
func sort2(nums []int) {
	n := len(nums)
	if nums == nil || n < 2 {
		return
	}

	// This is just a placeholder demo function
	// The actual radix sort is implemented in maximumGap2
}
