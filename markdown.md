# 換行
```
換行是後面要兩個空白  
這個是後面有兩個空白的效果
```

## 效果
換行是後面要兩個空白  
這個是後面有兩個空白的效果

```
這邊示範的是
後面沒有兩個空白的效果
```

## 效果
換行是後面要兩個空白
這個是後面有兩個空白的效果

# 標題

```
# H1
## H2
### H3
#### H4
##### H5
###### H6

Alt-H1
======

Alt-H2
------
```
## 效果

# H1
## H2
### H3
#### H4
##### H5
###### H6

Alt-H1
======

Alt-H2
------

# 強調語法
```
*asterisks*  //斜體
**asterisks** //粗體
```

## 效果

*asterisks*  //斜體

**asterisks** //粗體

# 清單
```
* 无序列表一  //可以使用星號建立無序清單
- 无序列表二  //或是短橫線（負號）
+ 无序列表三  //使用半形加號也可以
1. 有序列表一
2. 有序列表二
3. 有序列表三
```

## 效果

* 无序列表一  //可以使用星號建立無序清單
- 无序列表二  //或是短橫線（負號）
+ 无序列表三  //使用半形加號也可以

1. 有序列表一
2. 有序列表二
3. 有序列表三

# 連結
```
[這是一個行內連結](https://www.google.com)
[這是一個帶有標題的行內連結](https://www.google.com "Google's Homepage")
//滑鼠放在連結上會顯示 Google's Homepage
[Arbitrary case-insensitive reference text]
參考標的物也可以使用數字[1]
直接使用文字對應也可以 [這段文字連到參考項目]
[arbitrary case-insensitive reference text]: https://www.mozilla.org
[1]: http://slashdot.org
[這段文字連到參考項目]: http://www.reddit.com
```

## 效果

[這是一個行內連結](https://www.google.com)

[這是一個帶有標題的行內連結](https://www.google.com "Google's Homepage")
//滑鼠放在連結上會顯示 Google's Homepage

[Arbitrary case-insensitive reference text]

參考標的物也可以使用數字 [1]

直接使用文字對應也可以 [這段文字連到參考項目]

[arbitrary case-insensitive reference text]: https://www.mozilla.org

[1]: http://slashdot.org

[這段文字連到參考項目]: http://www.reddit.com

# 圖片

```
行內格式：

![alt text](https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon48.png "Logo 標題文字範例一")

參考連結格式：

![alt text][logo]

[logo]: https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon48.png "Logo 標題文字範例二"
```

## 效果

行內格式：

![alt text](https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon48.png "Logo 標題文字範例一")

參考連結格式：

![alt text][logo]

[logo]: https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon48.png "Logo 標題文字範例二"


# 程式代碼與語法高亮標示

```
行內 `code` 必須使用 `back-ticks` 將文字包起來（一般鍵盤左上方的第一個鍵）。
```

用3個反單引號將文字包起來，如果是程式語言，請指定它，例如：  
\```javascript  
var s = "JavaScript syntax highlighting";  
alert(s);  
\```

## 效果
```javascript  
var s = "JavaScript syntax highlighting";  
alert(s);  
```

\```python  
s = "Python syntax highlighting"  
print s  
\```

## 效果
```python
s = "Python syntax highlighting"
print s
```


```
沒有指定程式語言的話，就沒有語法高亮顯示  
但是，仍然可以用&lt;b&gt; <i><b>標籤</b></i> &lt;/b&gt;來加強內容
```

# 表格  

```
冒號（Colons）是用來對齊的（擺左齊左、擺右齊右，都擺就置中）。
| Tables        | Are           | Cool  |
| ------------- |:-------------:| -----:|
| col 3 is      | right-aligned | $1600 |
| col 2 is      | centered      |   $12 |
| zebra stripes | are neat      |    $1 |

最外圍的豎線（|）不是絕對需要，在原始文檔中你可以不要太在意美觀，
實際轉成網頁或電子書時會呈現得很好。你也可以在表格內使用行內格式。
表格繪製時，最好養成每行後面加兩個空白，以免有些時候容易出錯。

Markdown | Less | Pretty
--- | --- | ---
*Still* | `renders` | **nicely**
1 | 2 | 3
```

| Tables        | Are           | Cool  |  
| ------------- |:-------------:| -----:|  
| col 3 is      | right-aligned | $1600 |  
| col 2 is      | centered      |   $12 |  
| zebra stripes | are neat      |    $1 |  

Markdown | Less | Pretty  
--- | --- | ---  
*Still* | `renders` | **nicely**  
1 | 2 | 3  


