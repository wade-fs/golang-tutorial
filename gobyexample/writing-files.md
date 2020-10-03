# [Go by Example](../gobyexample.md): Writing Files

寫入檔案在 Go 遵循著與先前看過的[讀檔範例](reading-files.md)相似的方法。  

``` go
package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
)

// 檢查是否發生錯誤，若有則印出訊息並停止程式
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // 首先展示，如何將程式中的資料([]byte)寫入檔案
    d1 := []byte("hello\ngo\n")		// \n 是 c-型式的換行控制字元
    err := ioutil.WriteFile("/tmp/dat1", d1, 0644) // 0644 是檔案權限
    check(err)	// 檢查是否發生錯誤，若有則印出訊息並停止程式

    // 上例具備簡易寫檔特色，
    // 如果我們要進一步控制寫入檔案的過程，就需要更彈性作法
    // 首先創建新檔案(存在舊檔則先刪除它)
    f, err := os.Create("/tmp/dat2")
    check(err)

    // 要養成習慣關閉所有開啟的物件，例如資料庫或網路連線
    defer f.Close()	// 透過 defer 來保證最後離開程式時自動關閉檔案

    // 這邊一樣示範寫 []byte, 事實上可以寫入任何資料
    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    // 例如，可以直接寫入「字串」
    n3, err := f.WriteString("writes\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n3)

    // 為了保證已寫入儲存體（硬碟/網路等），可以使用 sync 達成
    // 通常 Sync() 是給非 bufio 的讀寫，
    // 因為非 bufio 預設是直接寫入，但是作業系統本身是以 read-through 處理
    // 也就是仍然不會馬上寫入，除非有人想讀它，所以調用 Sync() 迫使系統寫入
    // 註：如果關檔的話，也有 Sync() 相同的功效
    f.Sync()

    // 為了效率的話，尤其對頻繁的小資料讀寫，可以改用 bufio
    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n4)

    // Flush() 與 Sync() 有類似的功效
    // 而 Flush() 則是將緩存中的資料立刻傾印到儲存體中
    w.Flush()

}
```
[執行](http://play.golang.org/p/fQ7sd4gXv0F)
``` shell
$ go run writing-files.go 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered
```
接下來我們會繼續閱讀像 stdin/stdout 這類串流的資料輸出與輸入  

下一範例: [Line Filters](line-filters.md)
