package main

import (
	"fmt"
)

func main() {
	// ACM Question :
	// give a len n binary string array {s1, s2, ..., sn}
	// len si all same m
	//
	// Can do operations:
	//     var i, j, x int , 1 <= i, j<=n, 1<=x<=m.
	//     eg. si ... sj ... sn
	// switch si[x] and sj[x]
	//
	// after any times operations , array极差 is array's max value - min value.
	// Output the max极差 十进制值。 due to answer maybe big Number , need ans = ans % (1000000000 + 7).

	const MOD = 1000000007

	var n, m int
	fmt.Scan(&n, &m)

	// Read the binary strings
	strings := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&strings[i])
	}

	// Count number of 1s at each position
	ones := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if strings[i][j] == '1' {
				ones[j]++
			}
		}
	}

	// Calculate maximum possible value and minimum possible value
	// We want to maximize (max_string - min_string)
	// This is equivalent to making one string as large as possible and another as small as possible

	var maxVal, minVal int64

	for pos := 0; pos < m; pos++ {
		bitValue := int64(1) << (m - 1 - pos)
		bitValue %= MOD

		onesCount := int64(ones[pos])

		// For maximum string: if there are any 1s at this position, we can put one in our max string
		if onesCount > 0 {
			maxVal = (maxVal + bitValue) % MOD
		}

		// For minimum string: we put a 1 here only if we have more 1s than we can avoid
		// If onesCount >= n, then every string must have a 1 at this position
		if onesCount >= int64(n) {
			minVal = (minVal + bitValue) % MOD
		}
	}

	ans := (maxVal - minVal + MOD) % MOD
	fmt.Printf("%d\n", ans)
}
