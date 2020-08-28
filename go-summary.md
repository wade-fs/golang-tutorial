# [](#前言 "前言")前言

[Go 程式語言](https://en.wikipedia.org/wiki/Go_(programming_language)) 是 Google 推出的靜態程式語言，其特色在於核心語法和關鍵字非常精簡（全部只有 25 個關鍵字！）並擷取了靜態語言的效能和動態語言的開發效率的優點，具備垃圾回收、快速編譯等特性，且針對平行化程式設計在使用上非常方便。接下來的文章我們將透過 Golang Web 程式設計來學習 Go 這個程式語言。 

# [](#Go-環境建置 "Go 環境建置")Go 環境建置

安裝方式： 
1. [官網套裝安裝](https://golang.org/)，線上也有官方提供的 [playground](https://play.golang.org/) 線上執行環境可以使用 
2. 使用套件管理工具 homebrew（mac）、apt-get（linux） 等進行安裝  
例如，在 linux 上安裝：

``` shell $ sudo apt-get install golang-go
```

或是主動指定版本

``` shell
$ sudo apt-get install golang-1.8-go
```

在 mac 上安裝：

``` shell
$ brew install go
```

3. 使用 gcc、MinGW（windows）等編譯器編譯原始 Go 檔案  
 對於一般初學者來說，可以使用官網或是套件管理工具來安裝，可以留意 _環境變數_ 是否有設定成功。   
 若有成功安裝，可以打開終端機執行 _go version_ 會出現版本相關訊息，以下是 linux 的範例：

``` shell
$ go version
```

在 Go 中也提供了許多方便指令，方便我們編譯、測試和執行程式：  

## 編譯檔案

``` shell
$ go build
```

## 執行單元測試（unit testing）  
Go 一開始就內建了測試的機制，執行 go test 的話，會自動讀取套件目錄中的 `*_test.go` 來進行編譯並進行測試

``` shell
$ go test
```

## 讓程式可以格式化，符合 go 的語言慣例

``` shell
$ go fmt
```

## 安裝套件  

``` shell  
$ go get  
```

## 靜態分析潛在 bug

``` shell  
$ go vet
```

## 可以快速 build 並執行程式  

``` shell  
$ go run
```

## 展示 go 相關文件

``` shell  
$ godoc
``` 

## 重新命名變數和函式

``` shell  
$ gorename
```

## 產生程式碼

``` shell
$ go generate
```

# [](#你的第一個-Go-程式 "你的第一個 Go 程式")你的第一個 Go 程式  
 在 Go 中程式運行的入口是套件 main。在這個程式中使用並導入了套件 “fmt”，在 Go 程式中程式執行入口是 _func main_ 。  
 若成功在終端機執行 `$ go run`，則會在終端機印出 `Hello, World!` 的 Go 程式。  
 恭喜完成你的第一個 Go 程式了！

``` go  
// 宣告程式屬於哪個 package
package main

// 引入套件
import (
    "fmt"
)

// 程式執行入口
func main() {
    // 使用 := 簡化較正規的變數宣告:
    word := "Hello, World!" // 相當於 var word string = "Hello, World!"
    // 使用 fmt 套件印出字串 word
    fmt.Println(word)
}
```

另外一個範例是引入了 `math/rand` 套件產生隨機整數（由於每次執行時環境中 seed 相同所以會印出同樣值）

``` go  
package main

import (
  "fmt"
  "math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

Go 有支援許多網路程式開發的套件，可以用很簡單的幾行就完成網路伺服器的建置：

``` go
package main

import (
    "io"
    "net/http"
)

// 處理 request 和 response 的函式
func hello(w http.ResponseWriter, r *http.Request) {
    // 印出 hello world
    io.WriteString(w, "Hello world!")
}

func main() {
    // router 讓網址可以對應處理函式
    http.HandleFunc("/", hello)
    // 監聽 8000 port
    http.ListenAndServe(":8000", nil)
}
```

# Modules

自 golang 1.11 版本開始，支援模組化的套件相依管理方法，這邊就簡單介紹一下。  

## 目錄結構

先來說說 GOPATH 是可以依套件而變的，通常在 ${GOPATH}/ 會有 src/ 及 pkg/
一般預設值會是 $HOME/go 

如果為了讓各套件具備更加獨立且完整性來看，應該在將 GOPATH 設為套件的目錄。  
也就是說我們的套件目錄等於 ${GOPATH}, 而源碼會放在 src/，其相依性則放在 pkg/  

但是這個並不是天生自然的結構，需要自己操作:  
- export GOPATH=$(pwd)  # 當然要先切換目錄至你想要的套件目錄處
- mkdir -p src pkg
- .....  DO SOMETHING .....  
  將源碼放在 src/，此套件若還要細分子目錄，千萬不要使用 src/ 的目錄名，這是 golang 的不成文規定  
  此時的工作目錄應該切換至 src/  
- 要特別注意的是， golang src/ 中的源碼，其 package 一樣採用 `package main` 即可  
- 進行編譯測試之前，要先建立 mod 資訊(工作目錄在 ${GOPATH}/src/):  
  go mod init 套件名   # 例如 go mod init utils
- 編譯可採用 go build 或 go get  
  可以在 ${GOPATH}/bin 目錄下看到編譯好的執行檔(utils)

以上要特別注意執行 go mod init 的目錄是在 $GOPATH/src/,   
編譯出來的檔案則在 $GOPATH/bin/
