/*
1513. Number of substrings with only ones.



PROBLEM
You are given a binary string s

Return the number of substrings with all characters being 1's

Return it modulo 10^9 + 7
*/


/*
Brute force approach similar to 3234
Check every possible substring
Check if the substring at s[i,j] contains ONLY 1's
*/
package main

import "fmt"

func numSumBruteForce(s string) int {
	n := len(s)
	count := 0

			// Iterates through each possible substring s[i,j]
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			allOnes := true

			// Check if substring s[i,j] contains only 1's
			for k := i; k <= j; k++ {
				if s[k] == '0' {
					allOnes = false
					break
				}
			}

			if allOnes {
				count++
			}
		}
	}
	return count
} 


/*
This next solution is very beautiful it's O(n) complexity
Uses triangular number formula (k(k+1)/2)
This formula works in this case because of combinatorics of consecutive substrings.
Similar to pascals triangle in a way.
*/
func numSumOptimized(s string) int {
	consecutiveOnes := 0
	count := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			consecutiveOnes++
			count = (count + consecutiveOnes) % (1e9 + 7)
		} else {
			consecutiveOnes = 0
		}
	}
	return count
}


func main() {
	fmt.Printf("Test 1: s=\"0110111\" -> got %d want 9\n", numSumBruteForce("0110111"))
	fmt.Printf("Test 1: s=\"101\" -> got %d want 2\n", numSumBruteForce("101"))
	fmt.Printf("Test 1: s=\"111111\" -> got %d want 21\n", numSumBruteForce("111111"))
	fmt.Printf("Test 1: s=\"0110111\" -> got %d want 9\n", numSumOptimized("0110111"))
	fmt.Printf("Test 1: s=\"101\" -> got %d want 2\n", numSumOptimized("101"))
	fmt.Printf("Test 1: s=\"111111\" -> got %d want 21\n", numSumOptimized("111111"))
}
