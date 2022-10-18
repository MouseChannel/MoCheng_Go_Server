package mnet

import (
	"fmt"

	"github.com/MouseChannel/MoCheng_Go_Server/face"
)

type TestStruct struct {
	value int
}

func (test *TestStruct) Test() {
	test.value = 11
}
func (test *TestStruct) Print() {
	fmt.Println(test.value)
}
func NewTestStruct() face.ITest {
	return &TestStruct{
		value: 2,
	}
}
