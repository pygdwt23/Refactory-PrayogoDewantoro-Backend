// package main

// import (
// 	"testing"
// 	"fmt"
// )

// type testCase struct{
// 	description string
// 	input []rune
// 	expected rune
// }

// func missingLetters(chars []rune) rune {
// 	var last int

// 	for _, r := range chars {
// 		if last != 0 && int(r)-last > 1 {
// 			return rune(r - 1)
// 		}
// 		last = int(r)
// 	}
// 	return rune(last)
// }

// func findMissingLetters(t *testing.T) {
// 	testCases := []testCase{
// 		{
// 			"list_letters_first",
// 			[]rune{'c' ,'d' ,'e' ,'g' ,'h'},
// 			'f',
// 		},
// 		{
// 			"list_letters_second",
// 			[]rune{'X', 'Z'},
// 			'Y',
// 		},
// 	}

// 	for _, test := range testCases {
// 		if result := missingLetters(test.input); result != test.expected {
// 			t.Fatalf("FAIL: %s - missingLetters(%+v): %v - expected %v", test.description, test.input, result, test.expected)
// 		}
// 		t.Logf("PASS: %s", test.description)
// 	}
// }

// func main() {
// 	fmt.Println()
// }