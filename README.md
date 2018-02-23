# gocc - Golang port of OpenCC

## Introduction 介紹
gocc is a golang port of OpenCC([Open Chinese Convert 開放中文轉換](https://github.com/BYVoid/OpenCC/)) which is a project for conversion between Traditional and Simplified Chinese.

gocc stands for "**Go** verson Open**CC**", it is a total rewrite version of OpenCC by Go. It just borrows the dict files and config files of OpenCC, so it may not  produces the same output with the original OpenCC.

## Installation 安裝
### 1, golang package
```sh
go get github.com/liuzl/gocc
```
### 2, Command Line
```sh
git clone https://github.com/liuzl/gocc
cd gocc/tools/opencc
make install
gocc --help
```

## Usage 使用
```go
package main

import (
    "fmt"
    "github.com/liuzl/gocc"
    "log"
)

func main() {
    s2t, err := opencc.New("s2t")
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
