package main

import "fmt"

/* factorial of 5: 5*4*3*2*1
factorial of 5 = 5 * factorial of 4(4*3*2*1)
factorial of 4 = 4 * factorial of 3(3*2*1)
factorial of 3 = 3 * factorial of 2(2*1)
factorial of 2 = 2 * factorial of 1(1*1)
*/

func myfactorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * myfactorial(n-1)
}

func main() {
	n := 5
	result := myfactorial(n)
	fmt.Printf("Factorial of %d is %d\n", n, result)
}
