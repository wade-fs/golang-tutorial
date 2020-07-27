# [Go by Example](../gobyexample.md): Funcions

`函式` 是任何程式語言的中樞，這邊只列出簡單說明.  

``` go
package main

import "fmt"

func plus(a int, b int) int {		// 傳回值型態放在後面
    return a + b			// 一旦函式定義有傳回值，就需要明確的返回值以滿足定義
}

func plusPlus(a, b, c int) int {	// 跟上例一樣有多個相同型態的參數，可以精簡宣告方式
    return a + b + c
}

func main() {

    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)		// 順帶說明一下，res 前面已經有宣告，此處不能再用 :=
    fmt.Println("1+2+3 =", res)
}
```
[執行](http://play.golang.org/p/-o49-dQfGbK)

``` shell
$ go run functions.go 
1+2 = 3
1+2+3 = 6
```

下一範例: [多傳回值](multiple-return-values.md)
