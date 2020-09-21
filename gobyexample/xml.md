# [Go by Example](../gobyexample.md): XML

Go通過encoding.xml包提供了對XML和類似XML格式的內置支持。

``` go
package main

import (
    "encoding/xml"
    "fmt"
)

// Plant 將被映射到XML。 
// 後面 `xml:...` 所指示的內容，與 JSON 類似，主要是因為大小寫的要求
// 如果您不需要特別轉換大小寫，事實上也可以不要用 `...` 來指明編解碼間轉換要求
// 在這裡，我們使用XML包的一些特殊功能：
// xml.Name 型別指示此結構的XML元素的名稱； 	// 結果的 <plant ...>
// id，attr表示Id字段是XML屬性而不是嵌套元素。	// 結果的 <... id="27">
type Plant struct {
    XMLName xml.Name `xml:"plant"`
    Id      int      `xml:"id,attr"`
    Name    string   `xml:"name"`
    Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
    return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
        p.Id, p.Name, p.Origin)
}

func main() {
    coffee := &Plant{Id: 27, Name: "Coffee"}
    coffee.Origin = []string{"Ethiopia", "Brazil"}

    // 發出表示我們 Plant 的XML； 使用MarshalIndent產生具縮排的輸出
    out, _ := xml.MarshalIndent(coffee, " ", "  ")
    fmt.Println(string(out))

    // 輸出時，加入 xml 檔頭 // <?xml version="1.0" encoding="UTF-8"?>
    fmt.Println(xml.Header + string(out))

    // 使用Unmarhshal將具有XML的字節流解析為數據結構。
    // 如果XML格式不正確或無法映射到Plant，將返回描述性錯誤。
    var p Plant
    if err := xml.Unmarshal(out, &p); err != nil {
        panic(err)
    }
    fmt.Println(p)

    tomato := &Plant{Id: 81, Name: "Tomato"}
    tomato.Origin = []string{"Mexico", "California"}

    // parent> child> plant 欄位標籤告訴編碼器
    // 將所有Plant嵌套在<parent> <child>下。
    type Nesting struct {
        XMLName xml.Name `xml:"nesting"`
        Plants  []*Plant `xml:"parent>child>plant"`
    }

    nesting := &Nesting{}
    nesting.Plants = []*Plant{coffee, tomato}

    out, _ = xml.MarshalIndent(nesting, " ", "  ")
    fmt.Println(string(out))
}
```
[執行](http://play.golang.org/p/RsNLhMBazeX)
``` shell
$ go run xml.go
 <plant id="27">
   <name>Coffee</name>
   <origin>Ethiopia</origin>
   <origin>Brazil</origin>
 </plant>
<?xml version="1.0" encoding="UTF-8"?>
 <plant id="27">
   <name>Coffee</name>
   <origin>Ethiopia</origin>
   <origin>Brazil</origin>
 </plant>
Plant id=27, name=Coffee, origin=[Ethiopia Brazil]
 <nesting>
   <parent>
     <child>
       <plant id="27">
         <name>Coffee</name>
         <origin>Ethiopia</origin>
         <origin>Brazil</origin>
       </plant>
       <plant id="81">
         <name>Tomato</name>
         <origin>Mexico</origin>
         <origin>California</origin>
       </plant>
     </child>
   </parent>
 </nesting>
```

JSON 與 XML 被證明具有資料表示的等效性(非資料部份當然 JSON更有彈性),  
因此譯者更喜歡用的是 JSON 格式來交換資料  

下一範例: [Time](time.md)
