# 換行

```
換行是後面要兩個空白  
這個是後面有兩個空白的效果  
所以如果感覺顯示跟自己想的不一樣  
很有可能是後面少了兩個空白  
```

## 效果

換行是後面要兩個空白  
這個是後面有兩個空白的效果  
所以如果感覺顯示跟自己想的不一樣  
很有可能是後面少了兩個空白  

```
這邊示範的是
後面沒有兩個空白的效果
所以如果感覺顯示跟自己想的不一樣
很有可能是後面少了兩個空白

```

## 效果

換行是後面要兩個空白
這個是後面沒有兩個空白的效果
所以如果感覺顯示跟自己想的不一樣
很有可能是後面少了兩個空白

# 分節

```
分節在 markdown 中就是換行，而且前面最好不要有空白，以免誤判。要看分節的效果，最好是像這段話，相當長，長到讓瀏覽器會換行為止，它跟一般的換行不同之處在於中間有空白行 

例如這裡所示就是兩小節  
```

## 效果

分節在 markdown 中就是換行，而且前面最好不要有空白，以免誤判。  
要看分節的效果，最好是像這段話，相當長，長到讓瀏覽器會換行為止，  
它跟一般的換行不同之處在於中間有空白行  

例如這裡所示就是兩小節  

# 標題

```
下列各級標題，在 # 後面要有空白
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
*一個星號是斜體*  
**兩個星號是粗體**  
***三個星號又會是怎樣呢？***  
```

## 效果

*一個星號是斜體*  
**兩個星號是粗體**  
***三個星號又會是怎樣呢？***  

# 清單
```
* 無序列表一  //可以使用星號建立無序清單
  * 無序子列表 1 // 第二階的前置可以是 *+- 之一
  + 無序子列表 1 // 第二階的前置可以是 *+- 之一
  - 無序子列表 2
- 無序列表二  //或是短橫線（負號）
  - 無序列表 1
+ 無序列表三  //使用半形加號也可以
  + 列表1  

1. 有序列表一  
  a. 只要前置不相同即可  
  b. 只要前置不相同即可  
5. 有序列表二, 可以發現數字並沒按順序，只是指引它是清單用的
3. 有序列表三  
  A. 只要前置不相同即可  
    1. 這是第三階  
    2. 如果跟第一階一樣的編號，也是可以的，但是後面仍然要2個空白  
  B. 這邊會變成第四層，自己對齊並沒有幫助
```

## 效果

* 無序列表一  //可以使用星號建立無序清單  
  + 無序子列表 1 // 第二階的前置要跟第一階一樣  
  - 無序子列表 2  
- 無序列表二  //或是短橫線（負號）  
  * 無序列表 1  
+ 無序列表三  //使用半形加號也可以  
  + 無序列表 1  

1. 有序列表一  
  a. 只要前置不相同即可  
  b. 只要前置不相同即可  
5. 有序列表二, 可以發現數字並沒按順序，只是指引它是清單用的
3. 有序列表三  
  a. 只要前置不相同即可  
    1. 這是第三階  
    2. 如果跟第一階一樣的編號，也是可以的，但是後面仍然要2個空白  
  b. 這邊會變成第四層，自己對齊並沒有幫助

# 分隔線

```
擺在這邊講的原因是，它跟無序或粗斜體等等有點像  
分隔線是底下的符號(兩個都可以)，先決條件是前後要有空白行  

***
---
```

## 效果  

***  
---  

# 連結
```
首先，單純要是網址本身變連結，就用 <URL> 的方式，例如  
&lt;https://www.google.com&gt;  
&lt;wade@gmail.com&gt;  

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

首先，單純要是網址本身變連結，就用 <URL> 的方式，例如  
<https://www.google.com>  
<wade@gmail.com>  

[這是一個行內連結](https://www.google.com)

[這是一個帶有標題的行內連結](https://www.google.com "Google's Homepage")
//滑鼠放在連結上會顯示 Google's Homepage

[Arbitrary case-insensitive reference text]

參考標的物也可以使用數字 [1]

直接使用文字對應也可以 [這段文字連到參考項目]

[arbitrary case-insensitive reference text]: https://www.mozilla.org

[1]: http://slashdot.org

[這段文字連到參考項目]: http://www.reddit.com

### 再看一個完整例子  

```
In a hole in the ground there lived a hobbit. Not a nasty, dirty, wet hole, filled with the ends
of worms and an oozy smell, nor yet a dry, bare, sandy hole with nothing in it to sit down on or to
eat: it was a [hobbit-hole](https://en.wikipedia.org/wiki/Hobbit#Lifestyle "Hobbit lifestyles"), and that means comfort.
```

## 效果  

In a hole in the ground there lived a hobbit. Not a nasty, dirty, wet hole, filled with the ends
of worms and an oozy smell, nor yet a dry, bare, sandy hole with nothing in it to sit down on or to
eat: it was a [hobbit-hole](https://en.wikipedia.org/wiki/Hobbit#Lifestyle "Hobbit lifestyles"), and that means comfort.

或底下的，看起來效果一樣，內容是不同的

In a hole in the ground there lived a hobbit. Not a nasty, dirty, wet hole, filled with the ends
of worms and an oozy smell, nor yet a dry, bare, sandy hole with nothing in it to sit down on or to
eat: it was a [hobbit-hole][10], and that means comfort.

[10]: <https://en.wikipedia.org/wiki/Hobbit#Lifestyle> "Hobbit lifestyles"


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

如果圖片在網路上，又要連結別的網址時是怎麼寫呢？  
注意中括號、小括號  

```
[![An old rock in the desert](https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon48.png "Shiprock, New Mexico by Beau Rogers")](https://www.flickr.com/photos/beaurogers/31833779864/in/photolist-Qv3rFw-34mt9F-a9Cmfy-5Ha3Zi-9msKdv-o3hgjr-hWpUte-4WMsJ1-KUQ8N-deshUb-vssBD-6CQci6-8AFCiD-zsJWT-nNfsgB-dPDwZJ-bn9JGn-5HtSXY-6CUhAL-a4UTXB-ugPum-KUPSo-fBLNm-6CUmpy-4WMsc9-8a7D3T-83KJev-6CQ2bK-nNusHJ-a78rQH-nw3NvT-7aq2qf-8wwBso-3nNceh-ugSKP-4mh4kh-bbeeqH-a7biME-q3PtTf-brFpgb-cg38zw-bXMZc-nJPELD-f58Lmo-bXMYG-bz8AAi-bxNtNT-bXMYi-bXMY6-bXMYv)
```

## 效果

[![An old rock in the desert](https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon48.png "Shiprock, New Mexico by Beau Rogers")](https://www.flickr.com/photos/beaurogers/31833779864/in/photolist-Qv3rFw-34mt9F-a9Cmfy-5Ha3Zi-9msKdv-o3hgjr-hWpUte-4WMsJ1-KUQ8N-deshUb-vssBD-6CQci6-8AFCiD-zsJWT-nNfsgB-dPDwZJ-bn9JGn-5HtSXY-6CUhAL-a4UTXB-ugPum-KUPSo-fBLNm-6CUmpy-4WMsc9-8a7D3T-83KJev-6CQ2bK-nNusHJ-a78rQH-nw3NvT-7aq2qf-8wwBso-3nNceh-ugSKP-4mh4kh-bbeeqH-a7biME-q3PtTf-brFpgb-cg38zw-bXMZc-nJPELD-f58Lmo-bXMYG-bz8AAi-bxNtNT-bXMYi-bXMY6-bXMYv)

# Youtube

```
[![Markdown Youtube 教學-英](https://img.youtube.com/vi/HUBNt18RFbo/0.jpg)](https://www.youtube.com/watch?v=HUBNt18RFbo)
```

[![IMAGE ALT TEXT HERE](https://img.youtube.com/vi/HUBNt18RFbo/0.jpg)](https://www.youtube.com/watch?v=HUBNt18RFbo)

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

```
還有一種是開頭4個空白的話也算, 反而如果在 3個反引號裡面會被解釋
    <html>  
    <head>  
      <title>Test</title>  
    </head>  
    <body>  
        Hello  
    </body>  
    </html> 
====== 上面在 markdown 中是跟下面內容一樣 ======  
    &lt;html&gt;  
    &lt;head&gt;  
      &lt;title&gt;Test<lt;/title&gt;  
    &lt;/head&gt;  
    &lt;body&gt;  
        Hello  
    &lt;/body&gt;  
    &lt;/html&gt;  
```

## 效果  

===== 反而下面的內容有4個空白在行首，會被當成程式碼 =====

    <html>
    <head>
      <title>Test</title>
    </head>
    <body>
        Hello
    </body>
    </html>

用3個反引號有語法強調效果, 但是不同釋釋器效果不一樣  
```c  
    int main() {  
        printf ("Hello\n");  
    }  
```

# 如何抑制反引單？

```
三個反引號裡面，不管2個或1個反引號都被抑制  
``這裡就示範 `抑制的方法`, 先看看效果``  
```

## 效果  

三個反引號裡面，不管2個或1個反引號都被抑制  
``這裡就示範 `抑制的方法`, 先看看效果``  

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
表格繪製時，最好養成的習慣每行後面加兩個空白，以免有些時候容易出錯。

Markdown | Less | Pretty
--- | --- | ---:
*Still* | `renders` | **nicely**
1 | 2 | 3
```

| Tables        | Are           | Cool  |  
| ------------- |:-------------:| -----:|  
| col 3 is      | right-aligned | $1600 |  
| col 2 is      | centered      |   $12 |  
| zebra stripes | are neat      |    $1 |  

Markdown | Less | Pretty  
--- | --- | ---:  
*Still* | `renders` | **nicely**  
1 | 2 | 3  

# 引言

```
> 引言（Blockquotes）常常出現在電子郵件中，表示摘錄來信原句並回覆。  
>   
>> 這一行是引言的一部分引言。  
>> 這樣就很清楚了  
>
> 再加一行看看, 上面的空白第一層引言是必要的喔

引言的中斷，必須有一個空白引言

> 這是一段非常長的引言區塊，只要在句首使用了正確的符號與空格，你可以持續不間斷的撰寫，整段文字都還是會被包含在引言區塊中。當然你依舊可以在引言區塊中 *使用* **Markdown** 的行內格式標記語法。
```

## 效果

> 引言（Blockquotes）常常出現在電子郵件中，表示摘錄來信原句並回覆。  
>   
>> 這一行是引言的一部分引言。  
>> 這樣就很清楚了  
>  
> 再加一行看看, 上面的空白第一層引言是必要的喔

引言的中斷，必須有一個空白引言

> 這是一段非常長的引言區塊，只要在句首使用了正確的符號與空格，你可以持續不間斷的撰寫，整段文字都還是會被包含在引言區塊中。當然你依舊可以在引言區塊中 *使用* **Markdown** 的行內格式標記語法。  

# 合併

### 引言有清單

```
> #### The quarterly results look great!  
>  
> - Revenue was off the chart.  
> - Profits were higher than ever.  
>  
>  *Everything* is going according to **plan**.  
```

## 效果  

> #### The quarterly results look great!  
>  
> - Revenue was off the chart.  
> - Profits were higher than ever.  
>  
>  *Everything* is going according to **plan**.  

### 清單中有引言  

```
1 如果清單中有引言呢？  
2 清單本身開頭是 =-*. 而引言是 >  
  > 所以如果用小節的方式, 也就是有個空白行  
3 是不是就可以在清單中有引言呢？是的，但是這邊算是新的清單喔！！！  
```

## 效果  

1 如果清單中有引言呢？  
2 清單本身開頭是 =-*. 而引言是 >  此時並不會變引言  
  
  > 所以如果用小節的方式, 也就是有個空白行  
  
3 是不是就可以在清單中有引言呢？是的，但是這邊算是新的清單喔！！！  

### 無序清單中有引言，效果又如何？
```
*   This is the first list item.
*   Here's the second list item.

    > A blockquote would look great below the second list item.

*   And here's the third list item.
```

## 效果

*   This is the first list item.
*   Here's the second list item.

    > A blockquote would look great below the second list item.

*   And here's the third list item.

# 待辦事項

```
- [ ] 待辦事項
- [x] 已完成事項
  - [ ] 另一條待辦事項
  - [x] 新一條完成事項
```

## 效果

- [ ] 待辦事項
- [x] 已完成事項
  - [ ] 另一條待辦事項
  - [x] 新一條完成事項

