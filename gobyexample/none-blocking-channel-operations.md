# [Go by Example](../gobyexample.md): None-Blocking Channel Operations

通道上的基本發送和接收被阻斷進程。 然而，我們可以將 *select* 與默認子句一起使用，以實現非阻塞的發送，接收甚至非阻塞的多路選擇。

``` go
package main

import "fmt"

func main() {
    messages := make(chan string)
    signals := make(chan bool)

    select {					// 這是不阻塞的接收。
    case msg := <-messages:			// 如果消息上有一個值，
        fmt.Println("received message", msg)	// 則select將使用帶有該值的<-messages情況
    default:					// 如果不是，它將立即採用默認情況
        fmt.Println("no message received")
    }

    msg := "hi"
    select {					// 非阻塞發送的工作方式與此類似
    case messages <- msg:			// 無法將msg發送到消息通道
        fmt.Println("sent message", msg)	// 因為該通道沒有緩衝區(m := make(chan string, 1)
    default:					// 也沒有接收器。 因此，選擇默認情況。
        fmt.Println("no message sent")
    }

    select {					// 在default子句之上使用多種情況
    case msg := <-messages:			// 用來實現多路非阻塞選擇
        fmt.Println("received message", msg)	// 在>這裡，我們嘗試對消息和信號進行非阻塞接收
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}
```
[執行](http://play.golang.org/p/TFv6-7OVNVq)

``` shell
$ go run non-blocking-channel-operations.go 
no message received
no message sent
no activity
```

接下來我們稍微修改上例註解的給予緩衝的範例:

``` go
package main

import (
    "fmt"
)

func main() {
    // messages := make(chan string)    //如果不加緩存的話，就全部會選擇defalut
    messages := make(chan string, 1)	//加了緩存的話，會選擇對應的
    signals := make(chan bool)

    // messages <- "test"	// 如果在此給通道值的話，後面結果會不同

    select {
    case msg := <-messages:	//因為messages通道沒有值，因此選擇default執行
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

    msg := "hi world"
    select {
    case messages <- msg:	// 上一範例 messages 無緩存，所以收不到訊息
        fmt.Println("sent message", msg) //因為本範例有緩存，所以這裡的msg發送到messages通道，不會被阻塞住
    default:
        fmt.Println("no message sent")
    }

    select {
    case msg := <-messages:
        fmt.Println("received message", msg) //因為messages已經有值了，所以會選擇這個case執行
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}
```

下一範例: [Closing Channels](closing-channels.md)
