# [Go by Example](../gobyexample.md): Channels

*通道* 是連接並發 *goroutine* 的管道。 您可以將值從一個 *goroutine* 發送到通道，並將這些值接收到另一 *goroutine* 。

``` go
package main

import "fmt"

func main() {
    // 使用 make(chan val-type) 創建一個新通道。 通道通過它們傳達的值來分類。
    messages := make(chan string)

    // 使用 channel<- 語法將值發送到通道。
    // 在這裡，我們從一個新的goroutine向上面創建的消息通道發送 "ping"。
    go func() { messages<- "ping" }()

    // 使用 <-channel 語法從通道接收一個值。 在這裡，我們將收到上面發送的 "ping" 消息並打印出來。
    msg := <-messages
    fmt.Println(msg)
}
```
[執行](http://play.golang.org/p/MaLY7AiAkHM)

``` shell
$ go run channels.go 
ping
```

當我們運行程序時， "ping"消息已通過我們的通道從一個 *goroutine* 成功傳遞到了另一個 *goroutine* 。  

默認情況下，發送和接收阻斷進程，直到發送方和接收方都準備好為止。 此屬性使我們可以在程序結尾等待 "ping" 消息，而不必使用任何其他同步。

有興趣的話可以閱讀[Go 的並發：Goroutine 與 Channel 介紹](https://peterhpchen.github.io/2020/03/08/goroutine-and-channel.html)一文。  

下一範例: [Channel Buffering](channel-buffering.md)
