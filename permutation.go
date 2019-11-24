package main

import "fmt"

// isPermutation will check if a and b are permutations of each other
func isPermutation(a, b string) bool {

	// first we can check the lengths. if they are not the same, then they cannot be permutations of each other
	if len(a) != len(b) {
		fmt.Printf("%s is not a permutation of %s\n", a, b)
		return false
	}

	// create a map of runes to integers to keep track of the number of times each character appears in our string
	hash := make(map[rune]int)

	// count each time a character shows up in the first string
	for _, r := range a {
		hash[r]++
	}

	// for each character in the second string, subtract from the count
	for _, r := range b {
		hash[r]--

		// if the count dips below 0, we know that b had at least one extra of that character than a, therefore not a permutation
		if hash[r] < 0 {
			fmt.Printf("%s is not a permutation of %s\n", a, b)
			return false
		}
	}

	// we have made it through the second loop without detecting any extra or missing characters, therefore it is a permutation
	fmt.Printf("%s is a permutation of %s\n", a, b)
	return true
}

// INCORRECT SOLUTION
/*
a failure case of us attempting to be clever...
we initially thought that we could skip any difficult or expensive iteration by adding up the int values of each of the characters in both string
we realized at the end that this this does not actually work
for example:
the ascii value of "2" = 50
the ascii value of "d" = 100
by our original logic, the string "22" would be a permutation of the string "d" (obviusly incorrect)
*/

// INCORRECTisPermutation will check if a and b are permutations of each other
func INCORRECTisPermutation(a, b string) bool {

	// first we can check the lengths. if they are not the same, then they cannot be permutations of each other
	if len(a) != len(b) {
		return false
	}

	// use an int to keep track of the numerical representation of our strings
	x := 0

	// iterate over the first string, adding all of the int rune values to x
	for _, r := range a {
		x += int(r)
	}

	// while iterating over the second string, subtract the int rune values from x
	for _, r := range b {
		x -= int(r)
	}

	// at the end, we can verify that the strings contain the same letters if x is 0 (THIS IS ACTUALLY INCORRECT)
	if x != 0 {
		fmt.Printf("%s is not a permutation of %s\n", a, b)
		return false
	}

	// if is is 0, they are a permutation of each other (ALSO INCORRECT)
	fmt.Printf("%s is a permutation of %s\n", a, b)
	return true
}
