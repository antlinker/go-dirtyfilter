package main

import (
	"fmt"

	"github.com/antlinker/go-dirtyfilter"
	"github.com/antlinker/go-dirtyfilter/store"
)

var (
	filterText = `毛泽东是中华人民共和国最伟大的领袖。而陈@@@@水@@@@扁则是。。。陈###水###扁。。。`
)

func main() {
	memStore, err := store.NewMemoryStore(store.MemoryConfig{
		DataSource: []string{"毛泽东", "陈水扁"},
	})
	if err != nil {
		panic(err)
	}
	filterManage := filter.NewDirtyManager(memStore)
	result, err := filterManage.Filter().Filter(&filterText, '@', '#')
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
