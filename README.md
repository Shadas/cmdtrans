# cmdtrans

## It's a cmd translate tool.

Modify the AppId and SecretKey in config/config.go file. Run `make` to build the binary file.

I put it under /usr/local/bin and use it as follow:

```
ShadasdeMacBook-Pro:cmdtrans shadas$ fy
fy Command-line Interface v0.1.0

Usage:
  fy [flags]

Flags:
  -h, --help            help for fy
  -l, --langcodes       Show the language codes suppored.
  -s, --source string   Choose what language you want to translate from, default as auto. (default "auto")
  -t, --target string   Choose what language you want to translate to, default as zh. (default "zh")
```

```
ShadasdeMacBook-Pro:cmdtrans shadas$ fy translate
翻译

```

```
ShadasdeMacBook-Pro:cmdtrans shadas$ fy 苹果 -t en
Apple
```

Run `fy -l` to check the language supported.
```
ShadasdeMacBook-Pro:cmdtrans shadas$ fy -l
auto | 自动检测
zh | 中文
en | 英语
yue | 粤语
wyw | 文言文
jp | 日语
kor | 韩语
fra | 法语
spa | 西班牙语
th | 泰语
ara | 阿拉伯语
ru | 俄语
pt | 葡萄牙语
de | 德语
it | 意大利语
el | 希腊语
nl | 荷兰语
pl | 波兰语
bul | 保加利亚语
est | 爱沙尼亚语
dan | 丹麦语
fin | 芬兰语
cs | 捷克语
rom | 罗马尼亚语
slo | 斯洛文尼亚语
swe | 瑞典语
hu | 匈牙利语
cht | 繁体中文
vie | 越南语
```

Translate multiple words at one time:
```
ShadasdeMacBook-Pro:cmdtrans shadas$ fy apple boy
苹果 男孩
```

Use pipe to make it as your translate model:
```
ShadasdeMacBook-Pro:cmdtrans shadas$ fy apple | grep "苹"
苹果
```
