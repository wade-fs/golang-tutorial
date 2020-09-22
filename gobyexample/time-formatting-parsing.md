# [Go by Example](../gobyexample.md): Time Formatting and Parsing

Go支持通過基於樣式為主布局進行時間格式化和解析。  
這樣說不容易懂，簡單講就是，我想要的時間格式像是  
 xxx-xxx-xxx yy:yy:yy  
這樣子的

``` go
package main

import (
    "fmt"
    "time"
)

func main() {
    p := fmt.Println

    // 先取得一個時間值：現在時刻
    // 將時間以 RFC3339 規範格式印出
    t := time.Now()
    p(t.Format(time.RFC3339))

    // 另一種寫法，但是時間的不是現在時刻，而是指定的
    t1, e := time.Parse(
        time.RFC3339,
        "2012-11-01T22:08:41+00:00")
    p(t1)	// 印出來的內容是 Go 對時間內定的訊息格式

    // 回到前面的現在時刻 t, 以下面 HH:MMPM 的格式印出來
    p(t.Format("3:04PM"))

    // 同理，將現在時刻 t 以下面格式印出來
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))

    // 格式本身是字串，所以可以用變數來表達
    // 但是這個樣式缺了年月日秒，系統會以現在時刻補足
    form := "3 04 PM"
    t2, e := time.Parse(form, "8 41 PM")
    p(t2)

    // 底下就很簡單，也很麻煩，自行指定格式（含定位數）
    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

    // time.Parse() 當然有機會錯誤，訊息中會解釋原因
    ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    p(e)
}
```
[執行](http://play.golang.org/p/0ooeler0RfR)
``` shell
$ go run time-formatting-parsing.go 
2014-04-15T18:00:15-07:00
2012-11-01 22:08:41 +0000 +0000
6:00PM
Tue Apr 15 18:00:15 2014
2014-04-15T18:00:15.161182-07:00
0000-01-01 20:41:00 +0000 UTC
2014-04-15T18:00:15-00:00
parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": ...
```

亂數常常用時間為種子，因此接下來示例亂數

下一範例: [亂數](random-numbers.md)
