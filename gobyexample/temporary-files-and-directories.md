# [Go by Example](../gobyexample.md): Directories

在整個程序執行過程中，我們通常希望創建程序退出後不需要的數據。   
臨時文件和目錄可用於此目的，因為它們不會隨著時間的推移而污染文件系統。

``` go
package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // 創建臨時文件的最簡單方法是調用ioutil.TempFile。  
    // 它創建一個文件並將其打開以進行讀取和寫入。  
    // 我們提供“”作為第一個參數，它是臨時檔案名稱的所在目錄
    // 第二個參數是臨時檔案的前置檔名，以便自己識別
    // 如此, ioutil.TempFile將在操作系統的默認目錄創建文件。
    f, err := ioutil.TempFile("", "sample")
    check(err)

    fmt.Println("Temp file name:", f.Name())
    // 自動移除它
    defer os.Remove(f.Name())

    _, err = f.Write([]byte{1, 2, 3, 4})
    check(err)

    // 同樣的，可以創建臨時目錄, 參數與 TempFile() 一樣
    dname, err := ioutil.TempDir("", "sampledir")
    check(err)
    fmt.Println("Temp dir name:", dname)
    // 因為目錄是多層架構，建議使用 RemoveAll() 移除
    defer os.RemoveAll(dname)

    // 在臨時目錄下，檔名就不必再使用 TempFile(), 
    // 因為用完之後，會整個目錄移除
    fname := filepath.Join(dname, "file1")
    err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
    check(err)
}
```
[執行](http://play.golang.org/p/nMpjCsALS6P)
``` shell
$ go run temporary-files-and-directories.go
Temp file name: /tmp/sample610887201
Temp dir name: /tmp/sampledir898854668
```

下一範例：[Testing](testing.md)
