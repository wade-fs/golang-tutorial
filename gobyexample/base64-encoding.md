# [Go by Example](../gobyexample.md): Base64 Encoding

Go 提供內建的 [base64](http://en.wikipedia.org/wiki/Base64) 支援在 [encoding/base64](https://golang.org/pkg/encoding/base64/)  

``` go
package main

import (
    b64 "encoding/base64" // 這樣匯入的目的在於簡化程式碼
    "fmt"
)

func main() {
    // 要編碼的文本內容, 長度不限
    data := "abc123!?$*&()'-=@~"

    // Go 支援兩類，第一類是標準編碼 - StdEncoding()
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)
    // 解碼標準編碼的方法很類似
    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

    // Go 支援第二類編碼 - URLEncoding()
    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    // 解碼 Url 編碼文本的方法很類似
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}
```
[執行](http://play.golang.org/p/XLftf8Gvj4y)
``` shell
$ go run base64-encoding.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~
YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
```

進階請見 [Cipher 示例](https://golang.org/src/crypto/cipher/example_test.go)  
或是 [使用 Go 的 AES 加密與解密](https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes/)  
完整的示例如[用 Go 加密/解密數據完整示例](https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/)  

註：Java 支援的相當老舊，已被證實可以攻破，Go 則預設支援更新的機制  

下一範例: [Reading Files](reading-files.md)
