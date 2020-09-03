# [Go by Example](../gobyexample.md): Atomic Counters

Go中管理狀態的主要機制是通過[通道](channels.md)進行通信。  
例如，我們在[工作池](worker-pools.md)中看到了這一點。  
但是，還有其他一些用於管理狀態的選項。 在這裡，我們將研究對多個 *goroutine* 訪問的 *Atomic Counters*(原子計數器)使用 *sync/atomic* 軟件包。

``` go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {

    var ops uint64

    var wg sync.WaitGroup

    for i := 0; i < 50; i++ {
        wg.Add(1)

        go func() {
            for c := 0; c < 1000; c++ {

                atomic.AddUint64(&ops, 1)
            }
            wg.Done()
        }()
    }

    wg.Wait()

    fmt.Println("ops:", ops)
}
```
[執行](http://play.golang.org/p/j-14agntvEO)

``` shell
$ go run atomic-counters.go
ops: 50000
```


下一範例: [Mutexes](mutexes.md)
