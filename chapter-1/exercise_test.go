package main

import (
	"fmt"
	"math"
	"testing"
)

func TestButtonAndLampBruteForce(t *testing.T) {
	n := 25
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	if count%2 == 0 {
		fmt.Println("lampu mati")
	} else {
		fmt.Println("lampu hidup")
	}
}

func TestButtonAndLampEfficient(t *testing.T) {
	n := 25
	num := n
	divisorCount := 1
	// you only need to find it until square root of N
	for i := 2; i*i <= num; i++ {
		factorCount := 0
		for num%i == 0 {
			factorCount++
			num = num / i
		}
		divisorCount = divisorCount * (1 + factorCount)
	}

	if num > 1 {
		// to handle, prime number, which only has 2 factor
		divisorCount = divisorCount * 2
	}

	if divisorCount%2 == 0 {
		fmt.Println("lampu mati")
	} else {
		fmt.Println("lampu hidup")
	}
}

func TestButtonAndLampMoreOptimal(t *testing.T) {
	// this problem was actually find the squared number
	// if found it, then it is the answer
	n := 24
	result := math.Sqrt(float64(n))
	if int(result)*int(result) != n {
		fmt.Println("lampu mati")
	}else {
		fmt.Println("lampu hidup")
	}
}
