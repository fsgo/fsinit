# lazyfuncs
延迟方法管理，用于解决使用`init`方法后，执行逻辑不可控的问题。  

## 对外接口 
```go
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

// NewGroup 创建一个新组，用于和默认值分开管理
NewGroup() Group
```

## 使用示例

### 案例1

原逻辑：
```go
package xxx
func init(){
    // 加载配置，初始化子模块
}
```

调整为：
```go
package xxx

import (
    "github.com/fsgo/lazyfuncs"
)

func init(){
    // 注册后不会立马执行，测试更方便
    lazyfuncs.Register(Init)
}

func Init(){
    // 加载配置，初始化子模块
    // 异常panic
}
```

```go
package main

import (
    "github.com/fsgo/lazyfuncs"
 
    // 需要将对于的包引入，以执行其init方法，注册延迟执行函数
    // 若确保已在其他相关依赖包中引入，则可不添加
    _ "xzy/xxx"
)

func main(){
    lazyfuncs.Execute()
}
```