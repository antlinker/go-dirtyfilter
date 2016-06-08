package main

import (
	"fmt"

	"github.com/antlinker/go-dirtyfilter"
	"github.com/antlinker/go-dirtyfilter/store"
)

var (
	filterText = `我是需要过滤的内容，内容为：**文*@@件**名，需要过滤。。。`
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
