package task

import "fmt"

type FooTask struct{}

func (f *FooTask) Spec() string {
	return "@every 3s"
}

func (f *FooTask) Fn() func() {
	return func() {
		fmt.Println("foo task executed")
	}
}
