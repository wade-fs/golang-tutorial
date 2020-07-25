# [Go by Example](../gobyexample.md): Constants

Go支持字符，字串，布林值和數字值的常數。

``` go
package main

import (
    "fmt"
    "math"
)

const s string = "constant"	// const 宣告常數, 跟 var 一樣賦與型態+初始值

func main() {
    fmt.Println(s)

    const n = 500000000		// const 宣告常數, 跟 var 一樣僅賦與初始值，自動型態判定

    const d = 3e20 / n		// const 宣告常數，跟 var 一樣可以根據『算式值』賦值
    fmt.Println(d)

    fmt.Println(int64(d))	// 這邊將常數(或變數也行)的型態，強迫轉換

    fmt.Println(math.Sin(n))	// 前面未賦與型態的常數，在使用時會自動判定
}
```
[執行](http://play.golang.org/p/Vw-pXSfo9_b)

``` shell
$ go run constant.go 
constant
6e+11
600000000000
-0.28470407323754404
```

對 go 而言，常數一旦判定，其值不能更動，但是其型態卻是可以進行某種轉換，多多使用就可以明白

下一範例: [For](for.md)
