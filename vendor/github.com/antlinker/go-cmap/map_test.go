package cmap

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

type TestMap struct {
	gomap       map[interface{}]interface{}
	cmap        ConcurrencyMap
	gomapnolock map[interface{}]interface{}
	sync.RWMutex
}

func NewTestMap() *TestMap {
	return &TestMap{gomap: make(map[interface{}]interface{}), cmap: NewConcurrencyMap(), gomapnolock: make(map[interface{}]interface{}, 1024*1024)}
}
func (m *TestMap) GomapNolockGetSet(key interface{}, value interface{}) bool {
	m.gomap[key] = value
	newvalue := m.gomap[key]
	return newvalue == value

}
func (m *TestMap) GomapGetSet(key interface{}, value interface{}) bool {
	m.Lock()
	m.gomap[key] = value
	m.Unlock()
	m.RLock()
	newvalue := m.gomap[key]
	m.RUnlock()
	return newvalue == value

}
func (m *TestMap) ConcurrencymapGetSet(key interface{}, value interface{}) bool {
	err := m.cmap.Set(key, value)
	if err != nil {
		return false
	}
	newvalue, _ := m.cmap.Get(key)
	return newvalue == value
}

func BenchmarkGoMap(b *testing.B) {
	b.StopTimer()
	testmap := NewTestMap()
	b.StartTimer()
	var i int64 = 0
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			n := atomic.AddInt64(&i, 1)
			var key = fmt.Sprintf("foo_%d", n)
			result := testmap.GomapGetSet(key, n)
			if !result {
				b.Error("执行错误错误结果")
			}
		}
	})
}

func BenchmarkNolockGoMap(b *testing.B) {
	b.StopTimer()
	testmap := NewTestMap()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var key = fmt.Sprintf("foo_%d", i)
		result := testmap.GomapNolockGetSet(key, i)
		if !result {
			b.Error("执行错误错误结果")
		}
	}
}

func BenchmarkConcurrencyMap(b *testing.B) {
	b.StopTimer()
	testmap := NewTestMap()
	b.StartTimer()
	var i int64 = 0
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			n := atomic.AddInt64(&i, 1)
			var key = fmt.Sprintf("foo_%d", n)
			result := testmap.ConcurrencymapGetSet(key, n)
			if !result {
				b.Error("执行错误错误结果")
			}
		}
	})

}
