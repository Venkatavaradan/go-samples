// Concurrent computation of pi.
// See http://goo.gl/ZuTZM.
//
// This demonstrates Go's ability to handle
// large numbers of concurrent processes.
// It is an unreasonable way to calculate pi.
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("res",pi(5000000,50000))
}

// pi launches n goroutines to compute an
// approximation of pi.
func pi(n int, numthreads int) float64 {
    f := 0.0

	ch := make(chan float64);
	iterations := n/numthreads;
	for k:= 0; k <iterations; k++ {
	   	
		for i:=0; i< numthreads; i++{
			
			go term(ch, float64((k*numthreads)+i))
		
		}
		for i:=0; i< numthreads; i++ {
		f += <-ch
		}
		fmt.Println(" Iteration:", k, " value:",f);
	}
	return f
}

func term(ch chan float64, k float64) {
	
	val := 4 * math.Pow(-1, k) / (2*k + 1)
	//fmt.Println(k, val);
	ch <- val;
	
}

