# [Go by Example](../gobyexample.md): Interfaces

`界面` 是方法集合的別稱  
它的主要作用在於類似『物件』的使用，我們還是直接看範例吧。  

``` go
package main

import (
    "fmt"
    "math"
)

type geometry interface {	// 可以想成.....幾何類別
    area() float64		// 按照上面的原文，又可以說是『方法集合』
    perim() float64		// 也就是說，在界面中，只需要定義『方法』即可
}

type rect struct {		// 幾何形狀之一：矩形
    width, height float64
}
type circle struct {		// 幾何形狀之一：圓
    radius float64
}

func (r rect) area() float64 {	// 幾何形狀的方法：面積
    return r.width * r.height
}
func (r rect) perim() float64 {	// 幾何形狀的方法：周長
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 { // 幾何形狀的方法：面積
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {// 幾何形狀的方法：周長
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {	// 普通方法，利用共同界面 geometry
    fmt.Println(g)		// 
    fmt.Println(g.area())	// 就可以不管原來結構，呼叫界面中的方法
    fmt.Println(g.perim())	// 這作法很像 OOP 中的繼承
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)			// 矩形跟圓形有共同的界面
    measure(c)			// 所以可以透過 measure() 呼叫界面中的方法

    g := make([]geometry, 0)	// 『幾何』陣列
    g = append(g, r)		// 可以把矩形丟入幾何陣列
    g = append(g, c)		// 再把圓形丟入幾何陣列中，
				// 界面讓 Go 的陣列可以不必被限制統一結構
    for _, s := range g {	// 這時候可以呼叫界面中定義的方法
        fmt.Println(s.area(), s.perim())
    }
}
```
[執行](http://play.golang.org/p/4wmDCAydC1e)

``` shell
$ go run interfaces.go
{3 4}
12
14
{5}
78.53981633974483
31.41592653589793
```

Go 語言的界面是相當強大的特性，可以說寫個稍微不那麼小的程式都會用到  
有興趣的話，請參考[如何使用 Go 語言的界面](https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go)一文  

這邊再回憶一下前面提到的[結構](structs.md)及[方法](methods.md), 這兩個 Go 語言特性使得它具有物件導向的基礎，加上界面的使用，使得 Go 語言效能與結構化設計能同時具備。  

以本範例來思考，界面可以想成是類別中最小的共同方法集合。  
事實上，界面讓 Go 語言具備繼承的概念：界面中含有界面成員：

``` go  
type IGeometry interface {
    geometry
    name() string
    setName(string)
}
.... 參考上例 ....
    r := rect{width: 3, height: 4}
    r.SetName("table")
    c := circle{radius:5}
    c.SetName("wheel")
.... 略, 參考上例 ....
    for _, s := range g {
        fmt.Println(s.Name(), s.area(), s.perim())
    }
```

以上摘要來自[用介面 (Interface) 實踐繼承和多型](https://michaelchen.tech/golang-programming/interface/)  
底下先插個話，界面感覺是物件的一部份，而 Go 語言的多型不常見，只有 String() 最特別，我們將在後面[錯誤處理](error-handling-and-go.md)看到類似多型的用法。如果結構定義了 String() 方法的話，就可以將該物件透過 fmt.Println(v) 印出訊息。  

下一範例: [Errors](errors.md)
