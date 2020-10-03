# [Go by Example](../gobyexample.md): File Paths


``` go
package main

import (
    "fmt"
    // filepath 包提供分析與建構檔案路徑的功能, 
    // 此包是可攜至其他作業系統的，例如 Linux dir/file, Windows dir\file
    "path/filepath"
    "strings"
)

func main() {
    // Join 用來建構路徑的方式比較特別, 會處理不同作業系統的分隔字元
    // 我們鼓勵大家使用 Join() 來建構路徑，以免受 / \ 的困擾
    p := filepath.Join("dir1", "dir2", "filename")
    fmt.Println("p:", p)

    // Join() 的另一強大優勢，會自動處理不洽當的路徑分隔元，
    fmt.Println(filepath.Join("dir1//", "filename"))
    fmt.Println(filepath.Join("dir1/../dir1", "filename"))

    // Dir() 取出 dirname, 而 Base() 取出 basename
    fmt.Println("Dir(p):", filepath.Dir(p))
    fmt.Println("Base(p):", filepath.Base(p))

    // 判斷路徑是絕對路徑或是相對的
    fmt.Println(filepath.IsAbs("dir/file"))
    fmt.Println(filepath.IsAbs("/dir/file"))

    filename := "config.json" // 上面 Base() 的結果就類此這裡的值

    // 可以取出副檔名
    ext := filepath.Ext(filename)
    fmt.Println(ext)

    // 底下介紹去掉副檔名的方法
    fmt.Println(strings.TrimSuffix(filename, ext))

    // 去掉前面部份路徑的方法，也可以稱之為『相對前面』的路徑
    rel, err := filepath.Rel("a/b", "a/b/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)

    // 再看看所謂相對路徑
    rel, err = filepath.Rel("a/b", "a/c/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)
}
```
[執行](http://play.golang.org/p/5h3lUytvmyO)
``` shell
$ go run file-paths.go
p: dir1/dir2/filename
dir1/filename
dir1/filename
Dir(p): dir1/dir2
Base(p): filename
false
true
.json
config
t/file
../c/t/file
```

下一範例：[Directories](directories.md)
