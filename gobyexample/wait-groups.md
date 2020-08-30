# [Go by Example](../gobyexample.md): WaitGroups

要等待多個 *goroutine* 完成，我們可以使用一個 *wait groups*。

``` go
package main

import (
    "fmt"
    "sync"
    "time"
)

// 這是我們將在每個 *goroutine* 中運行的功能。
// 請注意，必須通過指標將 *WaitGroup* 傳遞給函數。
func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()	// 記得將 WaitGroup 關閉來表示已完成

    fmt.Printf("Worker %d starting\n", id)

    time.Sleep(time.Second)	// 模擬工作負載
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    // wg(WaitGroup) 將用來等待在此處啟動的所有 goroutines 完成工作
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)		// WatiGroup 必須擴充計數器(增加空間)
        go worker(i, &wg)	// 啟動 goroutine
    }

    wg.Wait()			// 然後進程被阻斷，直至計數器被清為 0
}				// 表示所有工作都已完成
```
[執行](http://play.golang.org/p/7mWXl0yVe6I)

``` shell
$ go run waitgroups.go
Worker 5 starting
Worker 3 starting
Worker 4 starting
Worker 1 starting
Worker 2 starting
Worker 4 done
Worker 1 done
Worker 2 done
Worker 5 done
Worker 3 done
```

每次調用時，工作啟動和結束的順序可能會有所不同, 這是 goroutine 並發的特色。

下一範例: [Rate Limiting](rate-limiting.md)
