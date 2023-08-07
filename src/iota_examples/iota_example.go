package main

import "fmt"

const (
	a = iota
	b
	c
)

const (
	d = iota
	e
	f
)

const (
	g = iota + 1
	h
	i
)

const (
	j = iota + 1
	k = iota
	l = iota
)

const (
	m = iota
	n
	o = iota
)

const (
	p = iota + 2 // 0 + 2
	q
	r
	s = iota
)

func main() {
	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)
	fmt.Printf("c = %v\n", c)
	fmt.Printf("d = %v\n", d)
	fmt.Printf("e = %v\n", e)
	fmt.Printf("f = %v\n", f)
	fmt.Printf("g = %v\n", g)
	fmt.Printf("h = %v\n", h)
	fmt.Printf("i = %v\n", i)
	fmt.Printf("j = %v\n", j)
	fmt.Printf("k = %v\n", k)
	fmt.Printf("l = %v\n", l)
	fmt.Printf("m = %v\n", m)
	fmt.Printf("n = %v\n", n)
	fmt.Printf("o = %v\n", o)
	fmt.Printf("p = %v\n", p)
	fmt.Printf("q = %v\n", q)
	fmt.Printf("r = %v\n", r)
	fmt.Printf("s = %v\n", s)
}
