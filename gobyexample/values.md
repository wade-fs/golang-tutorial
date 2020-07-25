# [Go by Example](../gobyexample.md): Values

Go具有各種值類型，包括字符串，整數，浮點數，布林值等。這是一些基本示例。

``` go
package main

import "fmt"

func main() {

    fmt.Println("go" + "lang")  	// 字串可以透過 '+' 接在一起

    fmt.Println("1+1 =", 1+1)		// 整數, 加法
    fmt.Println("7.0/3.0 =", 7.0/3.0)	// 浮點數, 除法

    fmt.Println(true && false)		// 布林值，且(&&) 操作
    fmt.Println(true || false)		// 布林值，或(||) 操作
    fmt.Println(!true)			// 布林值，否(!) 操作
}
```
[執行](http://play.golang.org/p/YnVS3LZr8pk)

這邊示範的都算是 `常數值` 的使用例，後面還會有更複雜的 Map, Slice, Channel 等介紹

下一範例: [變數](variables.md)
