package main

import (
	"fmt"
	"strconv"
)

// uniqueCharacters will check if an input string consists of all unique characters (assume ascii)
func uniqueCharacters(input string) bool {

	// ascii only has 128 unique characters. if the string is longer it cannot consist of unique characters
	// NOTE this is only applicable if the string is guarenteed to be ascii
	if len(input) > 128 {
		return false
	}

	// make a map of rune(the ascii value of a character) to bool to keep track of what characters we have already seen
	hash := make(map[rune]bool)

	// ignoring the index of the string, ranging over the string will give us the rune
	for _, c := range input {

		// check our map for the existance of this rune (ok will be true if it exists)
		if _, ok := hash[c]; ok {
			fmt.Printf("Found duplicate character \nrune value: %d\nstring value: %s", c, strconv.QuoteRune(c))
			// we have encountered a duplicate character, so we can end iteration immediatly and return false
			return false
		}

		// the character is unique so far, so add it to our map for future iterations
		hash[c] = true
	}

	// we have made it through the entire string and not found any duplicates, every character is unique!
	return true
}
