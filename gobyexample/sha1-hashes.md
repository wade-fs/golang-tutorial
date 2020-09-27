# [Go by Example](../gobyexample.md): Sha1 Hashes

接下來探討一些編碼的示例，甚至是加密解密運算  
SHA1哈希通常用於計算二進製或文本Blob的短標識。  
例如，git版本控制系統廣泛使用SHA1來識別版本化的文件和目錄。  
以下是在Go中計算SHA1哈希的方法。


``` go
package main

import (
    "crypto/sha1" // Go 實作了常見的哈布功能，sha1 是其一
    "fmt"
)

func main() {
    s := "sha1 this string" // 文本內容，可以很長，不限

    h := sha1.New()

    h.Write([]byte(s)) // 將文本以 []byte 的方式寫入即可

    bs := h.Sum(nil)	// 取其計算完的 sha1 哈希值

    fmt.Println(s)
    fmt.Printf("%x\n", bs)	// 其結果以16進制表達，跟 git commit 一致
}
```
[執行](http://play.golang.org/p/XLftf8Gvj4y)
``` shell
$ go run sha1-hashes.go
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94
```

可以計算其他哈希值，方法都類似，例如 MD5 也很常見，  
有興趣的人可以進階參閱[哈希強度](hash-strength.md)

下一範例: [Base64 編碼](base64-encoding.md)
