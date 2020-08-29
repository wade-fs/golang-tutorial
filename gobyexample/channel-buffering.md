# [Go by Example](../gobyexample.md): Channel Buffering

默認情況下，通道是*無緩衝*的，這意味著它們僅在有相應的接收（<-chan）準備接收發送的值時才接受發送（chan<-）。 *緩衝*通道接受數量有限的值，而這些值沒有相應的接收器。

``` go
package main

import "fmt"

func main() {
    // 在這裡，我們創建了一個字符串通道，最多可緩衝2個值。
    // 但是它並非陣列的概念，後面會看到用法
    messages := make(chan string, 2)

    messages <- "buffered"	// 可以丟訊息到通道緩衝
    messages <- "channel"	// 總共設置2筆

    fmt.Println(<-messages)	// 緩衝通道並非陣列的概念
    fmt.Println(<-messages)	// 而是 FIFO 的概念
}
```
[執行](http://play.golang.org/p/3BRCdRnRszb)

``` shell
$ go run channel-buffering.go 
buffered
channel
```

下一範例: [Channel Synchronization](channel-synchronization.md)
