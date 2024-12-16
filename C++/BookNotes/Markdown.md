# [Markdown tutorial](https://markdown.com.cn/basic-syntax/headings.html)

## 标题
```
#  -  ###### 对应标题等级  1 - 6  （由大至小）
空格后加标题 兼容性原则
```
optional ：
```
一级标题 文本下标注======（任意数量）
二级标题 文本下标注－－－－（任意数量）
```
## 段落
文本空白行隔开
## 换行
建议 两个空格加换行 或 `<br>`标注 （兼容）
## 强调
```
* italic text *
** bold text **
***italic + bold ***
```
## 引用
```
块引用
>
> .empty line.
>> (嵌套引用)
> -(无序列表)
> ## 标题语法
...
```
## 列表
```
有序列表
1.
2.
3.
4.

1.
1.
<Tab><Tab>+content
<Tab><Tab>+content
1.
1.

1.
8.
3.
5.

无序列表
- 
- 
	-
- 
	-
- 

* 
* 
* 
* 

+
+
+ 
+ 

```

## 代码
```
`单词 短语 `

``...`a`...  ``

<Tab>
<Tab>代码块

围栏代码块---
```语言名称

```(语法高亮)
```
## 分割线
```
单独一行

*** 

--- 

___

确保兼容性 前后加空行
```
## 链接
```
[content](url "title")

<email> <url>

**[]()**
*[]()*
[`code`]()

	PART 1
[content] [label](link)
	PART 2
[label]: <link>/link  "title"/'title'/(title)

url中的空格用%20代替 （兼容性）
```
## 图片
```
![content](图片链接 "Title")

结合链接 [![]()]()
```
## 转义字符
```
反斜杠字符达到转义目的 :
\ ` * _ {} [] () # + - . ! |

< &lt;   & &amp;
```
## 内嵌HTML标签
```
行级内联标签 <span> <cite> <del> 
此外 <a> <img>   <em>
区块标签 <div> <table> <pre> <p> 前后加空行

```

---

## 扩展语法
### 表格
```
|   |   |   |
|---|---|---|
| :-|:-:| -:|  表中的管道符：&#124;
```
### 脚注
```
 aaaaabbbb,[^1] cccc
 [^1]:This is a footnote.
```
### 标题编号
```
### Heading {#custom-id} 

[[Heading IDs](#custom-id)
```
### 定义列表
```
First term 
: DEFINITION
: DEFINITION
: DEFINITION
（空一行）
Second term
: DEFINITION
: DEFINITION
: DEFINITION
```
### 删除线
`~~content~~ content`
### 任务列表
```
- [x] content
- [ ] content
```
### 输入表情
```
复制粘贴或
表情符号简码
:tent:   :joy:
```
[表情符号简码列表](https://gist.github.com/rxaviers/7360908 "github页面")
