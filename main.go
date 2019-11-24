package main

import "fmt"

func main() {

	testUniqueString := "happy"
	fmt.Printf("\nRunning uniqueCharacters on string '%s'\n", testUniqueString)
	fmt.Printf("%t\n", uniqueCharacters(testUniqueString))

	testPerma := "acba"
	testPermb := "a"
	fmt.Printf("\nRunning isPermutation on strings '%s' and '%s'\n", testPerma, testPermb)
	fmt.Printf("%t\n", isPermutation(testPerma, testPermb))
}
