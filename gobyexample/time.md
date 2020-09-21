# [Go by Example](../gobyexample.md): Time

Go為時間和持續時間提供廣泛的支持；  
跟 C/C++ 不同的是，它簡化並統一時間相關的運算及更進一步的暫停功能  

``` go
package main

import (
    "fmt"
    "time"
)

func main() {
    p := fmt.Println

    // 得到現在的時間
    now := time.Now()
    p(now)

    // 十分彈性與人性化的自訂時間
    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

    // 由時間取得年/月/日/....都相當直覺，更可貴的是精度相當細緻
    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())
    // 這夠簡單了吧，直接取得星期幾，不必再計算了
    p(then.Weekday())
    // 兩個時間的比較也清楚明白
    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))
    // 計算差值也很容易
    diff := now.Sub(then)
    p(diff)
    // 差值也就是『一段時間』或是『持續時間』
    // 事實上它可以很容易跟『時間』轉換，至少相當直覺
    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())
    // 上方示例是減法，下方示例是加法
    p(then.Add(diff))
    p(then.Add(-diff))
}
```
[執行](https://play.golang.org/p/YAM3s1KPc8c)
``` shell
2009-11-10 23:00:00 +0000 UTC m=+0.000000001
2009-11-17 20:34:58.651387237 +0000 UTC
2009
November
17
20
34
58
651387237
UTC
Tuesday
false
true
false
-165h34m58.651387237s
-165.58295871867693
-9934.977523120617
-596098.651387237
-596098651387237
2009-11-10 23:00:00 +0000 UTC
2009-11-24 18:09:57.302774474 +0000 UTC
```
接下來，我們將研究與Unix epoch 相關的時間觀念。
所謂的 epoch 是指硬體上，如果沒有 RTC 記住時間的話，  
開機後就是從 epoch 開始計時，可以稱之為『機器誕生』時間

下一範例: [Epoch](epoch.md)
