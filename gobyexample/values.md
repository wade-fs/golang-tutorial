# [Go by Example](../gobyexample.md): Values

Go具有各種值類型，包括字符串，整數，浮點數，布林值等。這是一些基本示例。

``` go
package main

import "fmt"

func main() {

    fmt.Println("go" + "lang")  	// 字串可以透過 '+' 接在一起

    fmt.Println("1+1 =", 1+1)		// 整數, 加法
    fmt.Println("7.0/3.0 =", 7.0/3.0)	// 浮點數, 除法

    fmt.Println(true && false)		// 布林值，且(&&) 操作
    fmt.Println(true || false)		// 布林值，或(||) 操作
    fmt.Println(!true)			// 布林值，否(!) 操作
}
```
[執行](http://play.golang.org/p/YnVS3LZr8pk)

``` shell
$ go run values.go
golang
1+1 = 2
7.0/3.0 = 2.3333333333333335
false
true
false
```

上面示範的都算是 `常數值` 的使用例，後面還會有更複雜的 [Arrays](arrays.md), [Maps](maps.md), [Slices](slices.md), [Channel](channel.md) 等介紹.  
底下就介紹其他的內建資料型態:  

# 布林:

預定義型態是具有名稱的型態，布林型態名稱為 bool，只有兩個預先定義的常數 true 與 false，由於只有兩個值，因此在 Go 的規格書 中，並沒有明確提及 bool 的占用記憶體大小。  
在 Go 官方網站的 The Go Playground 執行以下程式碼，會告訴你 bool 大小是 1：
``` go
package main

import (
    "fmt"
    "unsafe"
)

func main() {
    fmt.Println(unsafe.Sizeof(true))
    fmt.Println(unsafe.Sizeof(false))
}
```
[執行](https://play.golang.org/p/62gvzeKce-p)

``` shell
1
1
```

# 數字

數字型態為整數與浮點數的集合，整數部份支援無號與有號整數，名稱分別為 uint 與 int，int 長度會與 uint 相同，而 uint 長度視平台實作而異，可能是 32 位元或是 64 位元。

如果想要長度固定，無號整數的型態名稱為 uint8、uint16、uint32、uint64，顧名思義，使用的長度分別為 8 位元、16 位元、32 位元與 64 位元，舉例來說，uint8 可儲存的整數範圍為 0 到 255，這也是開發者熟悉的位元組型態，而在 Go 中，byte 正是　uint8 的別名。

有號整數的型態名稱為 int8、int16、int32、int64，顧名思義，使用的長度分別為 8 位元、16 位元、32 位元與 64 位元，舉例來說，int32 可儲存的整數範圍為 -2147483648 到 2147483647，而 rune 為 int32 的別名，可用來儲存 Unicode 碼點（code point）。

如果直接寫下一個整數實字（literal），例如 10，在沒有程式上下文（context）的情況下，10 是未定型態（Untyped），未定義型態整數的預設型態（Default type）為 int 型態，在必須得到一個型態而程式上下文未提供時（例如變數宣告與賦值要進行型態推斷時），就會使用預設型態。

寫下 10 這樣的整數，預設是 10 進位制；可以在數字前加上 0，Go 1.13 後可使用 0o 來表示八進位制，加上 0x 表示 16 進位制，此時 a-f and A-F 都可以用來表示 10 到 15，例如 0xBadFace，G○ 1.13 後可以使用 0x1.0p-1021 來表示浮點數。

Go 1.13 後，可以使用 0b 來定義二進位數字，例如 0b00101101；數字分隔底線在 Go 1.13 後可以使用，例如 1_000_000、0b_1010_0110 或 3.1415_9265。

math 模組上定義了一些常數，可以讓你得知各整數型態的儲存範圍，例如以下程式顯示了各整數型態的儲存範圍：

``` go
package main

import (
    "fmt"
    "math"
    "reflect"
)

func main() {
    fmt.Printf("uint8  : 0 ~ %d\n", math.MaxUint8)
    fmt.Printf("uint16 : 0 ~ %d\n", math.MaxUint16)
    fmt.Printf("uint32 : 0 ~ %d\n", math.MaxUint32)
    fmt.Printf("uint64 : 0 ~ %d\n", uint64(math.MaxUint64))
    fmt.Printf("int8   : %d ~ %d\n", math.MinInt8, math.MaxInt8)
    fmt.Printf("int16  : %d ~ %d\n", math.MinInt16, math.MaxInt16)
    fmt.Printf("int32  : %d ~ %d\n", math.MinInt32, math.MaxInt32)
    fmt.Printf("int64  : %d ~ %d\n", math.MinInt64, math.MaxInt64)
    fmt.Printf("整數預設型態: %s\n", reflect.TypeOf(1))
}
```
[執行](https://play.golang.org/p/dqT0WHrnvrs)

``` shell
$ go run int.go
uint8  : 0 ~ 255
uint16 : 0 ~ 65535
uint32 : 0 ~ 4294967295
uint64 : 0 ~ 18446744073709551615
int8   : -128 ~ 127
int16  : -32768 ~ 32767
int32  : -2147483648 ~ 2147483647
int64  : -9223372036854775808 ~ 9223372036854775807
整數預設型態: int
```

注意到，程式中使用了 uint64 函式，對 math 的一些常數做了明確的型態轉換（Type conversion），這是因為在 Go 中，常數可以是未定型態（Untyped），實際型態會視當時程式環境而定，如果沒有可參考的環境資訊，會使用預設型態。

在這邊的例子中，若是拿掉 uint64，math.MaxUint64 就會採用 int 型態而發生 overflow 的錯誤，使用 uint64 函式進行型態轉換，讓常數有明確的環境資訊可以參考，就不會產生這個錯誤。

在 Go 中，不同型態之間也無法直接進行運算，就算都是整數也不行，例如以下會發生 mismatched types 錯誤：

``` go
package main

import "fmt"

func main() {
    var x int32 = 10
    var y int16 = 20
    fmt.Println(x + y) // mismatched types error
}
```
[執行](https://play.golang.org/p/i7S8vcpYWZg)

``` shell
$ go run prog.go
./prog.go:8:19: invalid operation: x + y (mismatched types int32 and int16)
```

上例想要避免錯誤，你必須明確進行型態轉換，例如寫為 x + int32(y)（或者是 int16(x) + y）。

# Unicode 碼, rune

在 Go 中並沒有字元對應的型態，只有碼點的概念，int32 或其別名 rune 可用來儲存 Unicode 碼點，你可以將單一文字包在單引號之中，例如 '林'，這會以 int32 儲存為 26519，例如 fmt.Println('林') 會顯示 26519，若想顯示為文字，則要使用 fmt.Println(string('林'))。

浮點數的名稱為 float32、float64，分別為 IEEE-754 32 位元與 64 位元浮點數，如果直接寫下一個浮點數實字，預設型態是 float64 型態，可使用科學記號，例如 1.e+0、6.67428e-11 等，常數 math.MaxFloat32、math.MaxFloat64 分別代表著浮點數的最大儲存範圍。

Go 還有複數（Complex number），其中 complex64、complex128，可由一個實部數字，加上一個虛部數字與 i 來表示複數，例如 1 + 2i，寫下一個複數實字，預設型態為 complex128，虛數的部份，在 Go 1.13 後，之前談到的數字表示法都可以使用，有三個函式可以用來處理複數，即 complex、real 與 imag，可參考〈Manipulating complex numbers〉。

Go 還有個 uintptr，可以用來儲存指標值，這之後有機會再來談。
 
# 字串

Go 的字串在實作上使用 UTF-8，就目前必須先知道的是，當使用雙引號包裹一系列文字，會產生字串型態，預設型態為 string，例如，"Justin" 會建立一個字串。

如果對字串使用 len 函式，傳回的會是位元組數量，而不是 Unicode 碼點的數量；如果使用 [] 搭配索引，取得特定索引位置的值，那麼傳回的會是 byte（uint8）型態。

在 Go 中，可以對字串使用切片（slice）操作，傳回的型態會是 string 型態，例如，"Justin"[0:2] 會傳回字串 "Ju"，不過，這是取得索引 0、1 處的位元組，再建立 string 傳回，因此，對於 "語言" 這個字串，如果想用切片操作取得 "語" 這個子字串，必須使用 "語言"[0:3]，因為 Go 的字串在實作上使用 UTF-8，一個中文字基本上佔三個位元組。

## 多行字串

上面使用 ".." 來使用字串時，並不能多行，此時有兩個方法解決，
- 方法一  
``` go
text := `這是第一行  
這邊是第二行  
  
這邊是第四行  
`

- 另一個方法是像這樣:   
``` go
text := "這是第一行\n這是第二行\n\n這是第四行\n"
```


## 特殊字元

字串的中文是 Unicode, 每個中文占3個位元，底下可以很清楚:  
`text := "\x47\x6f\xe8\xaa\x9e\xe8\xa8\x80"`  相當於 "Go語言"

-    \a：U+0007，警示或響鈴
-    \b：U+0008，倒退（backspace）
-    \f：U+000C，饋頁（form feed）
-    \n：U+000A，換行（newline）
-    \r：U+000D，歸位（carriage return）
-    \t：U+0009，水平 tab
-    \v：U+000b，垂直 tab
-    \\：U+005c，反斜線（backslash）
-    \"：U+0022，雙引號
-    \ooo：位元組表示，o 為八進位數字
-    \xhh：位元組表示，h 為十六進位數字
-    \uhhhh：Unicode 點點表示，使用四個 16 進位數字
-    \Uhhhhhhhh：Unicode 點點表示，使用八個 16 進位數字

## 字串 與 []byte 的轉換

``` go
package main

import "fmt"

func main() {
    text1 := "Go語言"
    bs := []byte(text1)
    bs[0] = 103
    text2 := string(bs)
    fmt.Println(text1) // Go語言
    fmt.Println(text2) // go語言
}
```
[執行](https://play.golang.org/p/tLeRXTD-LIu)

``` shell
$ go run prog.go
Go語言
go語言
```

# Rune

這邊特別獨立出來一個議題，如果要像我們想像的 `Go語言` 是 4個字的話，它的單位叫 Rune  
Rune 實作上是 int32 的別名，專門用來儲存 Unicode 的碼點(code point), 使用 fmt 時是 %c  

``` go
package main

import "fmt"

func main() {
    text := "Go語言"
    cs := []rune(text)
    fmt.Printf("%c\n", cs[2]) // 語
    fmt.Println(len(cs))      // 4
}
```
[執行](https://play.golang.org/p/TaUAgfyrZrr)

``` shell
$ go run prog.go
語
4
```

## Rune 與 range

``` go
package main

import "fmt"

func main() {
    text := "Go語言"
    for index, runeValue := range text {
        fmt.Printf("%#U 位元起始位置 %d\n", runeValue, index)
    }
}
```
[執行](https://play.golang.org/p/IEnHa-lHvQo)

``` shell
$ go run prog.go
U+0047 'G' 位元起始位置 0
U+006F 'o' 位元起始位置 1
U+8A9E '語' 位元起始位置 2
U+8A00 '言' 位元起始位置 5
```


下一範例: [變數](variables.md)
