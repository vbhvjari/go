package main

func main() {
	var a int = 10
	var b int
	b = a
	println("Value of b using assign operator is: ", b)

	var x, y int = 20, 10
	x += y
	println("Value of x using add and assign operator is: ", x)
	x -= y
	println("Value of x using subtract and assign operator is: ", x)
	x *= y
	println("Value of x using multiply and assign operator is: ", x)
	x /= y
	println("Value of x using divide and assign operator is: ", x)
	x %= y
	println("Value of x using modulus and assign operator is: ", x)
}
