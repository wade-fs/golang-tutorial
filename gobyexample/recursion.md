# [Go by Example](../gobyexample.md): 遞迴

遞回是相當傳統的課題，Go 與其他語言並沒有不同。  
遞迴的精神是 各個擊破, 底下是經典的範例，
本來是計算 n, 在算答案前，先算 n-1 ... 等最後有結果了，答案就會是 n * f(n-1)  

``` go
package main

import "fmt"

func fact(n int) int {
    if n == 0 {			// 遞迴最主要就是要有『必定會執行到的』終止條件
        return 1		// 如果這個條件時有時無，就會發生無法停止的狀況
    }				// 通常這種情況就是系統資源被吃光光
    return n * fact(n-1)	// 接下來就是調用自己
}

func main() {
    fmt.Println(fact(7))
}
```
[執行](http://play.golang.org/p/smWim1q9ofu)

``` shell
$ go run recursion.go 
5040
```

下一範例: [指標](pointers.md)
