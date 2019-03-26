// Unit Tests for go-world.go functions

package main

// In GO any function prefixed with Example will be excuted as test and its output
// is asserted against the Output comment in the function. Its an useful option to
// test the functions which just prints
func ExamplePrintWelcomeMsg() {
	// calling the function to be tested
	printWelcomeMsg()
	// In Go, to test the functions which simply prints,
	// we can use the below output comment which acts as assertion within Example functions

	// Output: Welcome to GO world!
}
