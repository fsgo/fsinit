// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu (duv123+git@baidu.com)
// Date: 2020/5/16

package lazyfuncs

import (
	"sort"
	"sync"
)

// NewGroup 创建一个新的组
func NewGroup() *Group {
	return &Group{}
}

type lazyFn struct {
	fn    func()
	order int
}

type Group struct {
	lazyFns  []*lazyFn
	lock     sync.Mutex
	executed bool
}

func (g *Group) Register(fn func()) {
	g.RegisterWithOrder(fn, 0)
}

func (g *Group) RegisterWithOrder(fn func(), order int) {
	lf := &lazyFn{
		fn:    fn,
		order: order,
	}
	g.lock.Lock()
	defer g.lock.Unlock()

	if g.executed {
		panic("already Executed, cannot register new func")
	}

	g.lazyFns = append(g.lazyFns, lf)
}

func (g *Group) Execute() {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.executed {
		return
	}
	g.executed = true

	sort.SliceStable(g.lazyFns, func(i, j int) bool {
		return g.lazyFns[i].order < g.lazyFns[j].order
	})
	for _, lf := range g.lazyFns {
		lf.fn()
	}
}
