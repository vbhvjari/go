// fmt : format package

/* Format Specifiers
%v		default format
$d		decimal integers (int or uint)
%f		floating (float32)
%.2f	floating (float64)
%s		Plain string (string)
%c		character (string)
%q		quoted character/string (string)
%t		boolean - true or false (bool)
%T		type of the value
*/

/*
Print value to console

fmt.print()	    : print value as plain text without formatting(no new line at end)
fmt.println()   : print value as plain text with new line at end
fmt.printf()	: print value with formatting (need to specify format specifiers), no new line at end; use "\n" for new line
*/

package main

import "fmt"

func main() {
	var s1 string = "vaibhav"
	var s2 string
	s2 = "jariwala"
	var s3, s4 string = "surat", "gujarat"
	var (
		s string  = "india"
		i int     = 5
		f float32 = 3.55
		b bool    = true
	)

	fmt.Print(s1)
	fmt.Print("\n")
	fmt.Print("Second value: ", s2)
	fmt.Print("\n")

	fmt.Println("3rd value and 4th value:", s3, s4)

	fmt.Printf("String: %s, Integer: %d, Float: %f and bool: %t", s, i, f, b)
}
