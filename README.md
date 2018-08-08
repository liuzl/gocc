# gocc - Golang version OpenCC
[![GoDoc](https://godoc.org/github.com/liuzl/gocc?status.svg)](https://godoc.org/github.com/liuzl/gocc)[![Go Report Card](https://goreportcard.com/badge/github.com/liuzl/gocc)](https://goreportcard.com/report/github.com/liuzl/gocc)
## Introduction 介紹
gocc is a golang port of OpenCC([Open Chinese Convert 開放中文轉換](https://github.com/BYVoid/OpenCC/)) which is a project for conversion between Traditional and Simplified Chinese developed by [BYVoid](https://www.byvoid.com/).

gocc stands for "**Go**lang version Open**CC**", it is a total rewrite version of OpenCC in Go. It just borrows the dict files and config files of OpenCC, so it may not produce the same output with the original OpenCC.

## Installation 安裝
### 1, golang package
```sh
go get github.com/liuzl/gocc
```
### 2, Command Line
```sh
git clone https://github.com/liuzl/gocc
cd gocc/cmd
make install
gocc --help
echo "我们是工农子弟兵" | gocc
#我們是工農子弟兵
```

## Usage 使用
```go
package main

import (
    "fmt"
    "log"
    
    "github.com/liuzl/gocc"
)

func main() {
    s2t, err := gocc.New("s2t")
    if err != nil {
        log.Fatal(err)
    }
    in := `自然语言处理是人工智能领域中的一个重要方向。`
    out, err := s2t.Convert(in)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s\n%s\n", in, out)
    //自然语言处理是人工智能领域中的一个重要方向。
    //自然語言處理是人工智能領域中的一個重要方向。
}
```
## Conversions
* `s2t` Simplified Chinese to Traditional Chinese
* `t2s` Traditional Chinese to Simplified Chinese
* `s2tw` Simplified Chinese to Traditional Chinese (Taiwan Standard)
* `tw2s` Traditional Chinese (Taiwan Standard) to Simplified Chinese
* `s2hk` Simplified Chinese to Traditional Chinese (Hong Kong Standard)
* `hk2s` Traditional Chinese (Hong Kong Standard) to Simplified Chinese
* `s2twp` Simplified Chinese to Traditional Chinese (Taiwan Standard) with Taiwanese idiom
* `tw2sp` Traditional Chinese (Taiwan Standard) to Simplified Chinese with Mainland Chinese idiom
* `t2tw` Traditional Chinese (OpenCC Standard) to Taiwan Standard
* `t2hk` Traditional Chinese (OpenCC Standard) to Hong Kong Standard
