// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu (duv123+git@baidu.com)
// Date: 2020/5/16

package lazyfuncs

import (
	"sort"
	"sync"
)

// NewGroup 创建一个新的组
func NewGroup() Group {
	return &group{}
}

// Group 一个init 分组
type Group interface {
	// Register 注册延迟函数
	Register(fn func())

	// RegisterWithOrder 注册延迟函数，同时指定执行顺序
	// 若有相同的order,则先注册的排前面
	RegisterWithOrder(fn func(), order int)

	// Execute 执行所有注册的函数
	// 执行之后不允许再注册新的函数，否则会panic
	Execute()

	// Func 返回一个包含了所有注册函数的方法
	Func() func()
}

type lazyFn struct {
	fn    func()
	order int
}

type group struct {
	lazyFns  []*lazyFn
	lock     sync.Mutex
	executed bool
}

func (g *group) Register(fn func()) {
	g.RegisterWithOrder(fn, 0)
}

func (g *group) RegisterWithOrder(fn func(), order int) {
	lazyFn := &lazyFn{
		fn:    fn,
		order: order,
	}
	g.lock.Lock()
	defer g.lock.Unlock()

	if g.executed {
		panic("already Executed, cannot register new func")
	}

	g.lazyFns = append(g.lazyFns, lazyFn)

	sort.SliceStable(g.lazyFns, func(i, j int) bool {
		return g.lazyFns[i].order < g.lazyFns[j].order
	})
}

func (g *group) Execute() {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.executed = true

	g.Func()()
}

func (g *group) Func() func() {
	return func() {
		for _, lazyFn := range g.lazyFns {
			lazyFn.fn()
		}
	}
}

var _ Group = (*group)(nil)
