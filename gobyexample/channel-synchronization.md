# [Go by Example](../gobyexample.md): channel-synchronization  

我們可以使用通道在 *goroutines* 之間同步執行。 這是一個使用阻塞接收等待 *goroutine* 完成的示例。  
等待多個 *goroutine* 完成時，您可能更喜歡使用 *WaitGroup* 。

``` go
package main

import (
    "fmt"
    "time"
)

// 這是我們將在 *goroutine* 中運行的功能。
// *done通道* 將用於通知另一個 *goroutine* 該功能的工作已完成。
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true	// 傳送值到 done 通道, 用來通知另一 goroutine
}

func main() {

    done := make(chan bool, 1)	// 產生 done通道
    go worker(done)		// 將 done 通道傳給 worker goroutine, 

    <-done	// 當收到 done 通道的值前會阻斷進程
}
```
[執行](http://play.golang.org/p/Nw-1DzIGk5f)

``` shell
$ go run channel-synchronization.go      
working...done                  
```

如果試著移除 <-done 這行的話，程式直接離開，影響是來不及等到 worker goroutine 開始(更別說要它結束)

下一範例: [Channel Directions](channel-directions.md)
