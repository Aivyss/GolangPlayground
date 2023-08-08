package main

import (
	"fmt"
	"math"
)

func main() {
	foo1()

	compareTo := func(num1, num2 int) int {
		result := 0

		if num1 > num2 {
			result = 1
		} else if num1 == num2 {
			result = 0
		} else {
			result = -1
		}

		return result
	}

	fmt.Println(compareTo(2, 1), compareTo(1, 1), compareTo(1, 2))
	func() {
		fmt.Println("you can also invoke anonymous function directly.")
	}()

	testAnonym()
}

func foo1() {
	fmt.Println("foo invoked")
}

type item struct {
	name    string
	price   int
	special bool
}

func testAnonym() {
	i1 := item{
		name:    "item 1",
		price:   1000,
		special: false,
	}
	i2 := item{
		name:    "item 2",
		price:   10000,
		special: false,
	}
	i3 := item{
		name:    "item 3",
		price:   100,
		special: false,
	}
	i4 := item{
		name:    "item 4",
		price:   10,
		special: false,
	}
	i5 := item{
		name:    "item 5",
		price:   10,
		special: true,
	}

	cart10000 := []item{i1, i2}
	cartCnt3 := []item{i1, i3, i4}
	cartSpecial := []item{i1, i5}
	cartCnt3Special := []item{i1, i4, i5}
	noDiscount := []item{i1}

	fmt.Println(discountFunction(cart10000...)(cart10000...))
	fmt.Println(discountFunction(cartCnt3...)(cartCnt3...))
	fmt.Println(discountFunction(cartSpecial...)(cartSpecial...))
	fmt.Println(discountFunction(cartCnt3Special...)(cartCnt3Special...))
	fmt.Println(discountFunction(noDiscount...)(noDiscount...))
}

func discountFunction(items ...item) func(items ...item) int {
	special := false
	cnt3 := len(items) >= 3
	totalPrice := 0

	for _, item := range items {
		totalPrice += item.price

		if item.special {
			special = item.special
		}
	}

	thousand10 := totalPrice > 10_000

	getTotalPrice := func(items ...item) int {
		total := 0

		for _, item := range items {
			total += item.price

			if item.special {
				special = item.special
			}
		}

		return total
	}

	var discount func(items ...item) int
	if thousand10 {
		fmt.Println("thousand10")

		discount = func(items ...item) int {
			summation := getTotalPrice(items...)

			return int(math.Floor(float64(summation) * 0.90))
		}
	} else if cnt3 && special {
		fmt.Println("cnt3 && special")

		discount = func(items ...item) int {
			summation := getTotalPrice(items...)

			var discount int
			if summation < 3000 {
				discount = 0
			} else {
				discount = 3000
			}

			return summation - discount
		}
	} else if cnt3 {
		fmt.Println("cnt3")

		discount = func(items ...item) int {
			summation := getTotalPrice(items...)

			var discount int
			if summation < 1000 {
				discount = 0
			} else {
				discount = 1000
			}

			return summation - discount
		}
	} else {
		fmt.Println("No Discount")

		discount = func(items ...item) int {
			return getTotalPrice(items...)
		}
	}

	return discount
}
