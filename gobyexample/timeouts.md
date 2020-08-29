# [Go by Example](../gobyexample.md): Timeouts

超時對於連接到外部資源或需要限制執行時間的程序很重要。 借助頻道和選擇功能，可以輕鬆，優雅地在Go中實現超時。

``` go
package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
    }()

    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout 1")
    }

    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2")
    }
}
```
[執行](http://play.golang.org/p/4oOz0j29MJ6)

``` shell
$ go run timeouts.go 
timeout 1
result 2
```


下一範例: [Non-Blocking Channel Operations](none-blocking-channel-operations.md)
