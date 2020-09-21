# [Go by Example](../gobyexample.md): Epoch

程序中的一個常見要求是獲取自Unix時代以來的秒數，毫秒或納秒數。  
底下是Go中的操作方法。

``` go
package main

import (
    "fmt"
    "time"
)

func main() {
    // time.Now() 搭配 Unix 或 UnixNano 來取得自 Unix epoch 以來經過的時間
    now := time.Now()
    secs := now.Unix()
    nanos := now.UnixNano()
    fmt.Println(now)

    millis := nanos / 1000000	// 1ms == 10^6 ns
    fmt.Println(secs)		//  1s == 10^3ms
    fmt.Println(millis)
    fmt.Println(nanos)		//  1s == 10^9 ns

    fmt.Println(time.Unix(secs, 0))
    fmt.Println(time.Unix(0, nanos))
}
```
[執行](http://play.golang.org/p/0ooeler0RfR)
``` shell
$ go run epoch.go 
2012-10-31 16:13:58.292387 +0000 UTC
1351700038
1351700038292
1351700038292387000
2012-10-31 16:13:58 +0000 UTC
2012-10-31 16:13:58.292387 +0000 UTC
```
接下來，我們將看另一項與時間相關的任務：時間解析和格式化。

下一範例: [時間解析和格式化](time-formatting-parsing.md)
