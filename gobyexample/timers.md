# [Go by Example](../gobyexample.md): Timers

我們通常希望在將來的某個時間執行Go代碼，或者在某個時間間隔重複執行。 Go的內置 *timer* 和 *ticker* 功能使這兩項任務變得容易。 我們將首先關注 *timer* ，然後關注 *ticker*。

``` go
package main

import (
    "fmt"
    "time"
)

func main() {
    // 計時器代表將來的單個事件。
    timer1 := time.NewTimer(2 * time.Second) // 該計時器將等待2秒鐘。

    // 上面您告訴計時器您要等待多長時間，此處提供將在該時間通知的頻道。
    <-timer1.C	// 在計時器的通道 C 阻塞進程，直到它發送一個值表示計時器已啟動
    fmt.Println("Timer 1 fired")

    // 如果您只想等待，則可以使用time.Sleep。
    // 計時器可能有用的一個原因是，您可以在計時器啟動之前將其取消。
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 fired")
    }()
    stop2 := timer2.Stop()  // 這是一個在啟動前就停止計時器的例子。
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }

    // 給timer2足夠的時間（如果有的話）觸發，
    // 這邊反而可以表明它實際上已停止。
    time.Sleep(2 * time.Second)
}
```
[執行](http://play.golang.org/p/gF7VLRz3URM)

``` shell
$ go run timers.go
Timer 1 fired
Timer 2 stopped
```
啟動程序後，第一個計時器將觸發約2秒，但第二個計時器將在有機會觸發之前停止運行。  

下一範例: [Timers](timers.md)
