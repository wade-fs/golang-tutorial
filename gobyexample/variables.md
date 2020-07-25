# [Go by Example](../gobyexample.md): Variables

在Go中，變數是 `明確聲明，並由編譯器使用`, 例如 `檢查函數調用的類型正確性`。

``` go
package main

import "fmt"

func main() {

    var a = "initial"	// var 用來宣告變數，此處型態由編譯器根據初始值判定
    fmt.Println(a)

    var b, c int = 1, 2	// var 一次可以宣告多個變數，明確宣告型態, 並賦與初始值
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int		// var 宣告變數時，只宣告型態
    fmt.Println(e)

    f := "apple"	// 利用 := 即時宣告變數及初始值，這種用法不會給予型態宣態
			// 正規用法是: var f string = "apple"
    fmt.Println(f)
}
```
[執行](http://play.golang.org/p/iYyAIilyBRf)

``` shell
$ go run variables.go
initial
1 2
true
0
apple
```

對 go 而言，單純用 var 宣告變數時，都會賦與初始值，以上例的整數來說，初始值為 0

下一範例: [常數](constants.md)
