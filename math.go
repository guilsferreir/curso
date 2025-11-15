package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2))
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}

func mod(a int, b int) int {
	return a % b
}

func sqrt(a int) int {
	return int(math.Sqrt(float64(a)))
}

func cbrt(a int) int {
	return int(math.Cbrt(float64(a)))
}

func log(a int) int {
	return int(math.Log(float64(a)))
}