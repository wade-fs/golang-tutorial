# [Go by Example](../gobyexample.md): 方法

`方法` 一般是指物件導向程式設計中屬於某物件的函式，Go 結合 C/C++ 的結構與類別，簡化學習曲線。

``` go
package main

import "fmt"

type rect struct {
    width, height int
}

func (r *rect) area() int {	// 屬於 type rect struct 的方法 area()
    return r.width * r.height	// 先看實際用例 r.area()
}

func (r rect) perim() int {	// 屬於 type rect struct 的方法 perim()
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    rp := &r
    fmt.Println("area: ", rp.area())	// 前面在講指標時有說過，指構在取用時也是 `.`
    fmt.Println("perim:", rp.perim())	// 所以不管是就屬性與方法，在 Go 裡是一致的
}
```
[執行](http://play.golang.org/p/4wmDCAydC1e)

``` shell
$ go run methods.go 
area:  50
perim: 30
area:  50
perim: 30
```

Go 雖然在使用指標時一律使用 `.` 對 C/C++ 的慣用者或許會覺得怪怪的，但是它的一致性又讓人不必理會複雜的指標、指標的指標、物件的指標、或是指標物件的指標等這些複雜事物。

有一篇名稱[不要被 Golang 中的指標及非指標方法接收器咬傷](https://nathanleclaire.com/blog/2014/08/09/dont-get-bitten-by-pointer-vs-non-pointer-method-receivers-in-golang/)，我們一樣透過範例來學結構、指標、方法這跟 C/C++ 差異性相當不同的 Go 特色。

Go 在定義方法時，還可以分成『傳值方法』及『指標方法』：  

``` go
package main

import "fmt"

type Mutatable struct {
    a int
    b int
}

func (m Mutatable) StayTheSame() { // 傳值方法，引用時，是複製一份再呼叫
    m.a = 5
    m.b = 7
}

func (m *Mutatable) Mutate() {	// 指標方法，引用時就是指標的概念
    m.a = 5			// 跟上一方法最大差異就是傳值或是傳參考
    m.b = 7
}

func main() {
    m := &Mutatable{0, 0}
    fmt.Println(m)
    m.StayTheSame()	// m 看起來是指標，在引用傳值方法時，其實是複製一份結構
    fmt.Println(m)
    m.Mutate()		// 在引用指標方法時，用的是相同內容，因為定義時它是用指標
    fmt.Println(m)
}
```
[執行](https://play.golang.org/p/6ylFuHcIpJV)

``` shell
&{0 0}		// 本身是結構指標
&{0 0}		// 呼叫傳值方法時，先先複製一份結構
&{5 7}		// 所以在 Go 的物件方法定義，全部都是採用指標方法
```

上面註解已說的夠清楚了，這邊補充一下...  
上面程式片山如果想再定義 `func (m Mutatable) Mutate()` 或  
`func (m *Mutatable) StayTheSame()` 都會出錯，因為兩個都是結構 Mutatable 的方法，如果再定義就會發生重複定義。其實站在 Go 的角度來思考就知道， Go 將指標簡化成一律使用 `.` 來引用屬性及方法，是分不出來『傳值方法』與『指標方法』兩個的。

下一範例: [Interfaces](interfaces.md)
