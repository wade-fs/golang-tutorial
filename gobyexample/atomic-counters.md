# [Go by Example](../gobyexample.md): Atomic Counters

Go中管理狀態的主要機制是通過[通道](channels.md)進行通信。  
例如，我們在[工作池](worker-pools.md)中看到了這一點。  
但是，還有其他一些用於管理狀態的選項。 在這裡，我們將研究對多個 *goroutine* 訪問的 *Atomic Counters*(原子計數器)使用 *sync/atomic* 軟件包。
*譯註* 在計算機領域的 Atomic 通常是指，一件事無法被插斷，或是說會被立即執行完畢之意  
*譯註* 如果是指資料結構的話，其義通常是指小的不能再小，例如整數、Boolean、字元 等等

``` go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    // 我們將使用無符號整數表示我們的（始終為正）計數器。
    var ops uint64
    // WaitGroup將幫助我們等待所有goroutine完成他們的工作。
    var wg sync.WaitGroup

    // 我們將啟動50個goroutine，每個goroutine會將計數器精確地增加1000次。
    for i := 0; i < 50; i++ {
        wg.Add(1)

        // 為了原子地增加計數器，我們使用AddUint64，
        // 並使用＆語法為其提供ops計數器的內存地址。
        go func() {
            for c := 0; c < 1000; c++ {
                atomic.AddUint64(&ops, 1)
            }
            wg.Done()
        }()
    }
    // 等待所有goroutine完成。
    wg.Wait()

    // 現在可以安全訪問 ops，
    // 因為我們知道沒有其他goroutine正在對其進行寫入。
    // 使用atomic.LoadUint64之類的功能，還可以在原子更新時安全地讀取原子。
    fmt.Println("ops:", ops)
}
```
[執行](http://play.golang.org/p/j-14agntvEO)

``` shell
$ go run atomic-counters.go
ops: 50000
```

我們預計將獲得50,000次作業。   
如果我們使用非原子的ops ++來增加計數器，我們可能會得到一個不同的數字，在運行之間會有所變化，  
因為goroutine會相互干擾。   

也就是把 atomic.AddUint64(&ops, 1) 替換成 ops++  
這樣的執行結果每次都有可能不同，而且可以發現都小於 50000  

此外，使用-race標誌運行時，我們會遇到數據爭用失敗的情況。

接下來，我們將討論互斥鎖，這是另一個用於管理狀態的工具。  

下一範例: [Mutexes](mutexes.md)
