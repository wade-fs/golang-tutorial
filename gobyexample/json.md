# [Go by Example](../gobyexample.md): JSON

Go提供了對JSON編碼和解碼的內置支持，包括往返於內置和自定義數據類型的支持。  
有幾個常見相關性極強的議題可以探討：
1) 結構
2) JSON(RESTful API)
3) 資料庫
4) xml

``` go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

// 此示例會用底下兩個結構展示對特定型別的資料編碼與解碼
// 需要特別注意的是，成員屬性的名稱按 Go 的慣用法，要以大寫開頭
type response1 struct {
    Page   int
    Fruits []string
}

// 底下的用法 `json:"page"` 用來匯出/匯入成 JSON 的欄位
type response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {
    // 先來轉換基本型別真假值(布林值)
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))		// true
    // 也可以轉換整數或...
    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))		// 1
    // 轉換浮點數
    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))		// 2.34
    // 轉換字串
    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))		// "gopher" // JSON 表示法，字串有雙引號

    // 轉換字串切片(陣列) 
    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))		// ["apple","peach","pear"] // JSON 表示法

    // 轉換映射, 對 JSON 來說就是常見的物件格式
    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))		// {"apple":5,"lettuce":7}  // JSON 物件

    // 轉換結構為 JSON 的物件
    res1D := &response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))		// {"Page":1,"Fruits":["apple","peach","pear"]}

    // 同上例
    res2D := &response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

    // 底下開始反向解碼 JSON 字串
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

    // 解碼的型別是 map[string]interface{}，後面在取值時要轉換型別
    var dat map[string]interface{}
    // 因為解碼比較容易發生錯誤，主要來自 JSON 字串要求嚴格的格式
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

    // 因為轉換完的型別是 interface{}, 所以需要型別轉換
    num := dat["num"].(float64)
    fmt.Println(num)
    // 底下是取出切片元素的用法
    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)	// 仍然需要型別轉換
    fmt.Println(str1)

    // 底下解碼的型別是明確宣告的結構，取值時就很自然
    // 也就是前面結構宣告時有加 `json:xxxx`
    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

    // 在上面的示例中，我們始終使用
    // 字節和字符串作為數據和JSON表示之間的中介，並輸出到標準輸出。  
    // 我們還可以將JSON編碼直接流式傳輸到os.Writer之類的
    // os.Stdout甚至HTTP響應主體(自訂 http api)。
    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}
```
[執行](http://play.golang.org/p/JOQpRGJWAxR)
``` shell
$ go run json.go
true
1
2.34
"gopher"
["apple","peach","pear"]
{"apple":5,"lettuce":7}
{"Page":1,"Fruits":["apple","peach","pear"]}
{"page":1,"fruits":["apple","peach","pear"]}
map[num:6.13 strs:[a b]]
6.13
a
{1 [apple peach]}
apple
{"apple":5,"lettuce":7}
```

進階可閱讀[JSON 與 Go](http://blog.golang.org/2011/01/json-and-go.html)及[JSON 套件](http://golang.org/pkg/encoding/json/)

下一範例: [XML](xml.md)
