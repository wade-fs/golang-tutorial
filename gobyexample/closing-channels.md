# [Go by Example](../gobyexample.md): Closing Channels

*關閉通道* 表示將不再發送任何值。 這對於將完成情況傳達給頻道的接收者很有用。


``` go
package main

import "fmt"

func main() {
    // 在此示例中，我們將使用作業渠道將要完成的工作從
    // main（）goroutine 傳遞給worker goroutine。
    // 當我們沒有該 worker 的更多工作時，我們將關閉工作渠道。
    jobs := make(chan int, 5)
    done := make(chan bool)

    // 這是 worker goroutine。
    // 它反復從具有jobs 的作業中接收 j,more := <-jobs。
    // 在這種特殊的2值接收形式中，如果作業已關閉並且
    // 通道中的所有值均已接收，則 more 的值將為false。
    // 我們用它來通知我們完成所有工作後的 done。
    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    // 這通過 jobs 通道將3個作業發送給 worker，然後將其關閉。
    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    // 我們使用前面看到的同步方法來等待 worker
    <-done
}
```
[執行](http://play.golang.org/p/vCvRjcMq7p3)

``` shell
$ go run closing-channels.go 
sent job 1
received job 1
sent job 2
received job 2
sent job 3
received job 3
sent all jobs
received all jobs
```

封閉渠道的想法自然引出了我們的下一個示例：  

下一範例: [Range over Channels](range-over-channels.md)
