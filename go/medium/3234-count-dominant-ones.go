/*
3234. Count the number of substrings with dominant ones


PROBLEM
You are given a binary string s.

Return the number of substrings with dominant ones.

A string has dominant ones if the number of ones in the string is greater than or equal to the square of the number of zeros in the string.

*/


/*
Brute force approach by checking all substrings
This approach wont work on leetcode since with O(n^3) complexity it exceeds the time limit
*/
package main 

import "fmt"

// Function taking two inputs and should count every single substring in string s and determine the substrings with dominant ones
func countSubstringsBruteForce(s string) int {
	n := len(s)
	count := 0

	// Creating every possible substring ranging from [i, j] where 0 <= i <= j < n
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			ones, zeros := 0, 0

			// Checking each created pair in range [i, j] checking if they contain a 1 or not.
			for l := i; l <= j; l++ {
				if s[l] == '1' {
					ones ++
				} else {
					zeros ++
				}
			}

			// Checking if substring meets condition (ones >= (zeroes)^2)
			if ones >= zeros * zeros {
				count ++
			}
		}
	}
	return count
}


/*
This solution is really similar to the first just an optimization feature with the prefix sums
Still fails in leetcode, time limit exceeds.
*/
func countSubStringsOptimized(s string) int {
	n := len(s)
	count := 0


	// Creating a prefix sum array	
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i]
		if s[i] == '1' {
			prefix[i+1]++
		}
	}

	for i := 0; i < n; i++ {
		zeros := 0

		for j := i; j < n; j++ {
			if s[j] == '0' {
				zeros ++
			}
			ones := prefix[j+1] - prefix[i]

			if ones >= zeros * zeros {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Printf("Test 1: s=\"00011\" -> got %d, want 5\n", countSubstringsBruteForce("00011"))
	fmt.Printf("Test 1: s=\"101101\" -> got %d, want 16\n", countSubstringsBruteForce("101101"))
	fmt.Printf("Test 1: s=\"00011\" -> got %d, want 5\n", countSubStringsOptimized("00011"))
	fmt.Printf("Test 1: s=\"101101\" -> got %d, want 16\n", countSubStringsOptimized("101101"))
	
	
}
