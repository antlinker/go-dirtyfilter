# Golang Dirty Filter

[![GoDoc](https://godoc.org/github.com/antlinker/go-dirtyfilter?status.svg)](https://godoc.org/github.com/antlinker/go-dirtyfilter)

> 基于DFA算法；
> 支持动态修改敏感词，同时支持特殊字符的筛选；
> 敏感词的存储支持内存存储及MongoDB存储。

## 获取

``` bash
$ go get -v github.com/antlinker/go-dirtyfilter
```

## 使用

``` go
package main

import (
  "fmt"

  "github.com/antlinker/go-dirtyfilter"
  "github.com/antlinker/go-dirtyfilter/store"
)

var (
  filterText = `我是需要过滤的内容，内容为：**文@@件，需要过滤。。。`
)

func main() {
  memStore, err := store.NewMemoryStore(store.MemoryConfig{
    DataSource: []string{"文件"},
  })
  if err != nil {
    panic(err)
  }
  filterManage := filter.NewDirtyManager(memStore)
  result, err := filterManage.Filter().Filter(filterText, '*', '@')
  if err != nil {
    panic(err)
  }
  fmt.Println(result)
}
```

## 输出结果

```
[文件]
```

## License

	Copyright 2016.All rights reserved.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.