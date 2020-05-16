/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/5/16
 */

package lazyfuncs

import (
	"testing"
)

func TestExecute(t *testing.T) {
	Default = NewGroup()
	num := 0
	Register(func() {
		num++
	})
	Register(func() {
		num += 90
	})
	Execute()
	want := 91
	if num != want {
		t.Errorf("after Execute,num=%d, want=%d", num, want)
	}
}

func TestExecute_panic(t *testing.T) {
	Default = NewGroup()
	num := 0
	Register(func() {
		num += 3
	})

	Execute()
	if num != 3 {
		t.Errorf("got=%d, want=%d", num, 3)
	}

	Func()()
	if num != 6 {
		t.Errorf("got=%d, want=%d", num, 6)
	}

	defer func() {
		if re := recover(); re == nil {
			t.Errorf("want panic")
		}
	}()

	// 已经执行了 Execute 方法了，不允许注册新的方法
	Register(func() {})
}
