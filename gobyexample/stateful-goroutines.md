# [Go by Example](../gobyexample.md): Stateful Goroutines


在前面的示例中，我們使用了具有[互斥鎖](mutexes.md)的顯式鎖定，以跨多個goroutines同步對共享狀態的訪問。   
另一種選擇是使用goroutine和通道的內置同步功能來實現相同的結果。 這種基於通道的方法與Go共享內存的想法保持了一致，該想法是通過通信並使每個數據完全由1個goroutine擁有。


``` go
package main

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

type readOp struct {	// 在此示例中，我們的狀態將歸於一個goroutine。
    key  int 		// 這將確保並發訪問不會破壞數據。
    resp chan int 	// 為了讀取或寫入該狀態，
}			// 其他goroutine將向擁有的goroutine發送消息並接收相應的回復。
type writeOp struct { 	// 這些readOp和writeOp結構封裝了
    key  int		// 這些請求以及擁有的goroutine進行響應的方式。
    val  int
    resp chan bool
}

func main() {

    var readOps uint64	// 如前面的示例, 
    var writeOps uint64	// 我們需要計算執行的操作數

    reads := make(chan readOp)	// 讀寫通道將其他 goroutines 
    writes := make(chan writeOp)// 用於發出讀取和寫入請求

    go func() {					// 這是有狀態的 goroutine(數量1個)
        var state = make(map[int]int)		// 基本上跟上一例一樣資料結構
        for {					// 但是現在是有狀態 goroutine內的資料
            select {				// 此 goroutine 不斷從讀寫通道中選擇
            case read := <-reads:		// 據此響應讀取
                read.resp <- state[read.key]	//   在響應通道 resp 發送值來指示成功
            case write := <-writes:		// 與寫入請求的到達
                state[write.key] = write.val	//   在響應通道 resp 發送值來指示成功
                write.resp <- true		// 讀取的 resp 響應為其期望值
            }
        }
    }()

    for r := 0; r < 100; r++ {			// 這將啟動100個goroutine，
        go func() {
            for {
                read := readOp{			// 以通過reads通道向擁有狀態的goroutine發出讀取。
                    key:  rand.Intn(5),		// 每次讀取都需要構造一個readOp，
                    resp: make(chan int)}	//
                reads <- read			// 通過讀取通道發送它，
                <-read.resp			// 並通過提供的resp通道接收結果。
                atomic.AddUint64(&readOps, 1)	// 以原子計數器確保數量沒有因為並發消失
                time.Sleep(time.Millisecond)	// 短暫的睡眠用來感覺有在工作
            }
        }()
    }

    for w := 0; w < 10; w++ {			// 同上，這邊是寫入
        go func() {				// 對於讀/寫我們是分別對待
            for {
                write := writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
                    resp: make(chan bool)}
                writes <- write
                <-write.resp
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    time.Sleep(time.Second)			// 同理，假裝系統忙碌了一秒鐘

    readOpsFinal := atomic.LoadUint64(&readOps)	// 最後捕捉並報告作業數
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
}
```
[執行](http://play.golang.org/p/5mf_P9xqBzk)

``` shell
$ go run stateful-goroutines.go
readOps: 71708			// 每個並發讀取作業約 1ms, 所以 100 個在一秒內應該有 10萬個
writeOps: 7177			// 寫入 10 個並發，應該有 1萬個
```				// 但是系統總是還在忙別的事，所以實際數量會較少

對於這種特殊情況，基於goroutine的方法比基於互斥體的方法要複雜得多。  
但是，在某些情況下它可能很有用:
例如，當您涉及其他通道時，或者在管理多個此類互斥鎖時，都容易出錯。   
您應該使用最自然的方法，尤其是在理解程序的正確性方面。

因為並發是重要課題，我們可以轉往[進階並發](advanced-parallel.md)或是新課題，[排序](sorting.md)

下一範例: [排序](sorting.md)
