# [Go by Example](../gobyexample.md): Tickers

[計時器](timers.md)用於將來要執行某項操作時- *Tickers*(定時器) 用於需要固定週期進行重複操作時。 這是一個自動定期滴答的定時器示例，直到我們停止為止。

``` go
package main

import (
    "fmt"
    "time"
)

func main() {
    // 定時器使用與計時器類似的機制：發送值的通道。
    // 在這裡，我們將使用通道上的內建的 select 來等待值(每500毫秒到達一次)。
    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

    // 休息1.6秒供定時器作業
    time.Sleep(1600 * time.Millisecond)

    // 定時器也跟計時器一樣可以停止，
    // 一旦定時器停止，就再也無法從通道收到值
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
}
```
[執行](http://play.golang.org/p/gs6zoJP-Pl9)

``` shell
$ go run tickers.go
Tick at 2012-09-23 11:29:56.487625 -0700 PDT
Tick at 2012-09-23 11:29:56.988063 -0700 PDT
Tick at 2012-09-23 11:29:57.488076 -0700 PDT
Ticker stopped
```

當我們運行該程序時，定時器只來得及進行三次，1.6 秒後停止。

下一範例: [Worker Pools](worker-pools.md)
