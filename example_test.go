// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu (duv123+git@baidu.com)
// Date: 2020/5/16

package lazyfuncs_test

import (
	"fmt"

	"github.com/fsgo/lazyfuncs"
)

func ExampleGroup() {
	group := lazyfuncs.NewGroup()

	group.Register(func() {
		fmt.Println("fn1")
	})

	group.Register(func() {
		fmt.Println("fn2")
	})

	group.RegisterWithOrder(func() {
		fmt.Println("fn3")
	}, -1)

	group.RegisterWithOrder(func() {
		fmt.Println("fn4")
	}, -2)

	group.Execute()
	// OutPut:
	// fn4
	// fn3
	// fn1
	// fn2
}

func ExampleDefault() {
	lazyfuncs.Default = lazyfuncs.NewGroup()

	lazyfuncs.Register(func() {
		fmt.Println("world")
	})
	lazyfuncs.RegisterWithOrder(func() {
		fmt.Println("hello")
	}, -1)

	lazyfuncs.Execute()
	// OutPut:
	// hello
	// world
}
