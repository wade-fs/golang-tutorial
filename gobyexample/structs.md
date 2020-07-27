# [Go by Example](../gobyexample.md): 結構

Go 的結構跟 C/C++ 的結構或類別很類似，如果在 Go 使用結構來搭配 (interfaces](interfaces.md) 大概就可以完成常規的 OOP 操作了。

``` go
package main

import "fmt"

type person struct {	// 結構定義跟 C/C++ 類似，但是它就長這樣，沒更多的了
    name string		// 不像 C/C++ 在 `}` 後面還可以(應該)跟著尾巴
    age  int
}

func newPerson(name string) *person {
    p := person{name: name}	// 產生結構並賦與值的方法（之一）
    p.age = 42			// 注意 name: name 就是你想的那樣
    return &p			// 所以不必特別加 this 來修飾
}

func main() {
    fmt.Println(person{"Bob", 20})	// 無名產生結構並賦與初始化
    fmt.Println(person{name: "Alice", age: 30}) // 產生結構時依名稱賦值
    fmt.Println(person{name: "Fred"})	// 如果有欄位少了，也會有該型別的初值
    fmt.Println(&person{name: "Ann", age: 40}) // 指標也沒問題
    fmt.Println(newPerson("Jon"))

    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    sp := &s
    fmt.Println(sp.age)		// 結構指標變數的取值不是 `->`, 而是一樣用 `.`

    sp.age = 51
    fmt.Println(sp.age)
}
```
[執行](http://play.golang.org/p/n7jt1x3iw4Z)

``` shell
$ go run structs.go
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
&{Jon 42}
Sean
50
51
```

結構比較厲害的地方，請參考後面的[Json](json.md), 甚至是用在[資料庫](http://gorm.io/docs/models.html)也很方便，因此設計 Web APP 幾乎到處都看得到結構。  
我會專門找 Web APP 程式設計的書來翻譯，敬請期待。

下一範例: [方法](methods.md)
