# [Go by Example](../gobyexample.md): URL Parsing

URL提供了一種統一的資源定位方式。 以下是在Go中解析網址的方法。


``` go
package main

import (
    "fmt"
    "net"
    "net/url"
)

func main() {
    // 我們將解析此示例URL，其中包括:
    // 方案(postgres)，身份驗證信息(user:pass)，主機(host.com)，端口(5432)，
    // 路徑(path)，查詢參數(k=v)和查詢片段(#f)。
    s := "postgres://user:pass@host.com:5432/path?k=v#f"

    // 解析 URL 與處理錯誤
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    // 印出方案
    fmt.Println(u.Scheme)
    // 顯示身份驗證信息
    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    // 密碼因為有驗證需求，有可能發出錯誤
    p, _ := u.User.Password()
    fmt.Println(p)

    // 主機資訊一併處理端口
    fmt.Println(u.Host)
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)

    // 查詢路徑與查詢片段(在 # 後面的部份)應一併處理
    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

    // 至於查詢參數可以多組，以 k=v 的形成 []map[string]string
    fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])
}
```
[執行](http://play.golang.org/p/fHTQn9X7l1B)
``` shell
$ go run url-parsing.go 
postgres
user:pass
user
pass
host.com:5432
host.com
5432
/path
f
k=v
map[k:[v]]
v
```

下一範例: [SHA1 Hashes](sha1-hashes.md)
