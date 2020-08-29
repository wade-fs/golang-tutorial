# [Go by Example](../gobyexample.md): Select

將通道用作功能參數時，可以指定通道是否僅用於發送或接收值。 這種特異性增加了程序的類型安全性。

``` go
package main

import (
    "fmt"
    "time"
)

func main() {
    // 使用 select 將橫跨兩個通道
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)	// 遞延一段時間，用來模擬網路通訊的阻斷
        c1 <- "one"	// 將通道用在 goroutine 中
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    // 因為會有兩個通道，所以 for 2次
    for i := 0; i < 2; i++ {
        select {	// 使用 select 可以在並發進行透過通道傳送訊息時進行處理
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
```
[執行](http://play.golang.org/p/FzONhs4-tae)

``` shell
$ time go run select.go 
received one
received two

real    0m2.245s
```

注意到，因為 goroutine 並發特性，整體時間只有2秒，而不是1+2秒

下一範例: [Timeouts](timeouts.md)
