package main

/*

R.I.P.S.S

Repeat
Inputs
Pseudo
Simple
Solve

0 1 1 2 3 5 8 13 21 34 ...

*/

import "fmt"

var results = make(map[uint]uint)

func main() {
	fmt.Println("Fib(5)", fib(5))
}

// fib returns the nth number in the Fibinocci sequence
//
// fib(2) = 1
func fib(n uint) uint {
	// b/e of uint, we know that all numbers are >=0
	r, ok := results[n]
	if ok {
		return r
	}

	if n == 0 || n == 1 {
		r = n
	} else {
		r = fib(n-1) + fib(n-2)
	}
	results[n] = r
	return r
}
