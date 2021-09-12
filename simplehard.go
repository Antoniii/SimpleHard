// Go - Recursion
// https://www.tutorialspoint.com/go/go_recursion.htm

package main

import "fmt"
	
func f3nplus1(n , m int) int {
	 
	if n == 1 {
        return m
    } else if n%2 == 0 {
        n = n/2
		m += 1
		return f3nplus1(n,m)
    } else {
        n = n*3+1
		m+=1
		return f3nplus1(n,m)
    }
}

func main() {
	fmt.Println(f3nplus1(27,0))
	fmt.Println(f3nplus1(7,0))
	fmt.Println(f3nplus1(26,0))
	fmt.Println(f3nplus1(94,0))
	fmt.Println(f3nplus1(97,0))
}

// func foo(params ...int) {
//     fmt.Println(len(params))
// }

// func main() {
//     foo()
//     foo(1)
//     foo(1,2,3)
// }
