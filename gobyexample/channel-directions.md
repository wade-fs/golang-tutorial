# [Go by Example](../gobyexample.md): Channel Directions

將通道用作功能參數時，可以指定通道是否僅用於發送或接收值。 這種特異性增加了程序的類型安全性。

``` go
package main

import "fmt"

// 此ping功能僅接受用於發送值的通道。 嘗試在此通道上接收將是編譯時錯誤。
func ping(pings chan<- string, msg string) {
    pings <- msg
}

// 乒乓功能接受一個通道進行接收（ping），第二個通道進行發送（pong）。
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
```
[執行](http://play.golang.org/p/mjNJDHwUH4R)

``` shell
$ go run channel-directions.go
passed message
```

下一範例: [Select](select.md)
