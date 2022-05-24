package main

import "fmt"

type Test struct {
	v int
}

func (test *Test) testFunc() {
	test.v++
	fmt.Println(test.v)
}
