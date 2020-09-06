# [Go by Example](../gobyexample.md): Mutexes

在前面的示例中，我們看到瞭如何使用[原子操作](atomic-counters.md)來管理簡單的計數器狀態。  
對於更複雜的狀態，我們可以使用Mutex(互斥鎖)安全地跨多個 *goroutine* 訪問數據。

``` go
package main

import (
    "fmt"
    "math/rand"
    "sync"
    "sync/atomic"
    "time"
)

func main() {
    // 本範例中 state 是個 map
    var state = make(map[int]int)

    // 互斥鎖 mutex 會同步存取 state
    var mutex = &sync.Mutex{}

    // 我們將跟踪我們進行了多少次讀寫操作。
    var readOps uint64
    var writeOps uint64

    // 在這裡，我們啟動100個goroutine，以對該狀態執行重複讀取，
    // 每個goroutine中每毫秒執行一次。
    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {
                // 對於每次讀取，我們選擇一個要訪問的鍵，
                // Lock()用來鎖住互斥鎖以確保對狀態的獨占訪問，讀取所選鍵的值，
                // Unlock()解開互斥鎖，並增加readOps原子計數值。
                key := rand.Intn(5)
                mutex.Lock()
                total += state[key]
                mutex.Unlock()
                atomic.AddUint64(&readOps, 1)
                // 在讀取之間稍等片刻
                time.Sleep(time.Millisecond)
            }
        }()
    }
    // 我們還將使用與讀取相同的模式，啟動10個goroutine來模擬寫入。
    for w := 0; w < 10; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }
    // 等待1秒來讓10個goroutine在狀態和互斥量上工作一秒鐘
    time.Sleep(time.Second)

    // 取得最後作業計數並報告
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)

    // 最終鎖定狀態，說明它結束值(如何結束)
    mutex.Lock()
    fmt.Println("state:", state)
    mutex.Unlock()
}
```
[執行](http://play.golang.org/p/0WEmOOjoCjp)

``` shell
$ go run mutexes.go
readOps: 83285
writeOps: 8320
state: map[1:97 4:53 0:33 2:15 3:2]
```

運行該程序表明，在互斥鎖同步狀態下，我們總共執行了約90,000次操作。  

接下來，我們將研究僅使用 [goroutine](goroutines.md) 和 [通道](channels.md) 來實現相同的狀態管理任務。

下一範例: [Stateful Goroutines](stateful-goroutines.md)
