package cmap

import (
	"fmt"
	"testing"
)

func TestCMap(t *testing.T) {
	cmap := NewConcurrencyMap()
	cmap.Set("Foo", "bar")
	cmap.Set("Foo1", "bar1")
	cmap.Set("Foo2", "bar2")
	cmap.SetIfAbsent("Foo", "bar2")
	foo, err := cmap.Get("Foo")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Foo:", foo)
	t.Log("Keys:", cmap.Keys())
	t.Log("Values:", cmap.Values())
	t.Log("Map:", cmap.ToMap(), ",Len:", cmap.Len())
	foo2, err := cmap.Remove("Foo2")
	t.Log("Remove value:", foo2)
	t.Log("Map:", cmap.ToMap(), ",Len:", cmap.Len())
}

func TestCMapElements(t *testing.T) {
	cmap := NewConcurrencyMap()
	for i := 0; i < 10; i++ {
		err := cmap.Set(fmt.Sprintf("Foo_%d", i), i)
		if err != nil {
			t.Error(err)
			return
		}
	}
	for element := range cmap.Elements() {
		t.Log("Key:", element.Key, ",Value:", element.Value)
	}
}
