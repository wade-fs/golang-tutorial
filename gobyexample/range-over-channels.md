# [Go by Example](../gobyexample.md): Range over Channels

在[前面的示例](closing-channels.md)中，我們看到了 *for* 和 *range* 如何在基本數據結構上提供迭代。 我們還可以使用此語法來迭代從通道接收的值。


``` go
package main

import "fmt"

func main() {

    queue := make(chan string, 2) // 將在隊列通道中迭代2個值。
    queue <- "one"
    queue <- "two"
    close(queue)

    // 當從隊列中接收到每個元素時, 此 range 將對其進行迭代。
    // 因為我們用 close() 關閉了上面的通道，
    // 所以迭代在接收到2個元素後終止。
    for elem := range queue {
        fmt.Println(elem)
    }
}
```
[執行](http://play.golang.org/p/yQMclmwOYs9)

``` shell
$ go run range-over-channels.go
one
two
```

此示例還表明，關閉非空通道之後，仍然可以接收並處理剩餘的值。

下一範例: [Timers](timers.md)
