/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/5/16
 */

package lazyfuncs

// Default 默认分组
var Default = NewGroup()

// Register 注册延迟函数
func Register(fn func()) {
	Default.Register(fn)
}

// RegisterWithOrder 注册延迟函数，同时指定执行顺序
// 若有相同的order,则先注册的排前面
func RegisterWithOrder(fn func(), order int) {
	Default.RegisterWithOrder(fn, order)
}

// Execute 执行所有注册的函数
// 执行之后不允许再注册新的函数，否则会panic
func Execute() {
	Default.Execute()
}

// Func 返回一个包含了所有注册函数的方法
func Func() func() {
	return Default.Func()
}
