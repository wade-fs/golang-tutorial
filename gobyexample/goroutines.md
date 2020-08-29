# [Go by Example](../gobyexample.md): Goroutines

*goroutine* 是輕量級的執行線程。

``` go
package main

import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

    f("direct")			// 先從普通函式調用看起，它是同步阻斷式的
				// 也就是行程需要等它執行完才會往下
    go f("goroutine")		// 函式調用前方加上 go 就稱之為 goroutine
				// 它是非同步的，在它還沒有結束前行程立即往下
    go func(msg string) {	// 常見這種自定無名函式的 goroutine
        fmt.Println(msg)
    }("going")

    time.Sleep(time.Second)	// 為了等前面的 goroutine 做完，休息一下
    fmt.Println("done")		// 這種做法並不保證給予 goroutine 足夠時間
}				// 比較好的作法是用 [WaitGroup](waitgroups.md)
```
[執行](http://play.golang.org/p/I7scqRijEJt)

``` shell
$ go run goroutines.go
direct : 0		// 執行時，前方同步阻斷式的先執行完
direct : 1
direct : 2		// 之後才會往下走
goroutine : 0		// 而 goroutine 執行時，非阻斷，非同步的方式，
going			// 就會看到新的行程出現
goroutine : 1		// 當然 goroutine 會持續進行
goroutine : 2		// 這個時間外部正在 time.Sleep()
done
```

要發揮 goroutine 的強大功能，通常搭配 [WaitGroups](waitgroups.md) 或[Channels](channels.md)  
進階文章可以參考[撰寫共時性程式](https://michaelchen.tech/golang-programming/concurrency/)  

下一範例: [Channels](channels.md)
