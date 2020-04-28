package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	a := []int{1}
	b := []int{1}
	carry := 0

	for count := 2; count < 1000000; count++ {
		carry = 0
		index := 0
		if len(a) >= len(b) {
			index = len(a)
		} else {
			index = len(b)
		}
		c := make([]int, 0, 0)

		for i := 0; i < index; i++ {
			value1 := 0
			if len(a) > i {
				value1 = a[i]
			}

			value2 := 0
			if len(b) > i {
				value2 = b[i]
			}

			sum := value1 + value2 + carry
			if sum >= 10 {
				carry = 1
				sum -= 10
			} else {
				carry = 0
			}

			c = append(c, sum)
		}

		if carry > 0 {
			c = append(c, carry)
		}

		a = b
		b = c

		if count%10000 == 0 {
			fmt.Println(count + 1)
		}

		if len(b) >= 10 {
			//reverseしている
			if check(b[0:9]) && check(b[len(b)-9:]) {
				fmt.Println("index => ", count+1)
				break
			}
		}
	}
	end := time.Now()
	fmt.Printf("%d ms\n", (end.Sub(start)).Milliseconds())
}

func check(source []int) bool {
	var bits [9]int

	for _, n := range source {
		if n-1 >= 0 {
			bits[n-1] = 1
		}
	}

	sum := 0
	for i := 0; i < len(bits); i++ {
		sum += bits[i]
	}

	return sum == 9
}
