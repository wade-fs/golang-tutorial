# [Go by Example](../gobyexample.md): Directories

Go具有一些有用的功能，可用於處理文件系統中的目錄。  
透過這些標準的方法，就以讓程式碼可攜性變高.   

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
    // mkdir
    err := os.Mkdir("subdir", 0755)
    check(err)

    // rm -rf
    defer os.RemoveAll("subdir")

    // touch 新檔案，可以處理多層路徑
    createEmptyFile := func(name string) {
        d := []byte("")
        check(ioutil.WriteFile(name, d, 0644))
    }

    createEmptyFile("subdir/file1")

    // mkdir -p
    err = os.MkdirAll("subdir/parent/child", 0755)
    check(err)

    createEmptyFile("subdir/parent/file2")
    createEmptyFile("subdir/parent/file3")
    createEmptyFile("subdir/parent/child/file4")

    // find subdir/parent
    c, err := ioutil.ReadDir("subdir/parent")
    check(err)

    fmt.Println("Listing subdir/parent")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    // cd
    err = os.Chdir("subdir/parent/child")
    check(err)

    // find .
    // 搭配前面有 cd, 所以是 find subdir/parent/child
    c, err = ioutil.ReadDir(".")
    check(err)

    fmt.Println("Listing subdir/parent/child")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    err = os.Chdir("../../..")
    check(err)

    // 可以『先深後廣』找遍所有檔案
    fmt.Println("Visiting subdir")
    err = filepath.Walk("subdir", visit)
}

func visit(p string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(" ", p, info.IsDir())
    return nil
}
```
[執行](https://play.golang.org/p/UaeLMS5VQVR)
``` shell
$ go run directories.go
Listing subdir/parent
  child true
  file2 false
  file3 false
Listing subdir/parent/child
  file4 false
Visiting subdir
  subdir true
  subdir/file1 false
  subdir/parent true
  subdir/parent/child true
  subdir/parent/child/file4 false
  subdir/parent/file2 false
  subdir/parent/file3 false
```

下一範例：[Temporary Files and Directories](temporary-files-and-directories.md)
