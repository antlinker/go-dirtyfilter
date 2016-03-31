package store

import (
	"bytes"
	"io"
	"sync/atomic"

	"github.com/antlinker/go-cmap"
)

const (
	// DefaultDelim 默认读取敏感词的分隔符
	DefaultDelim = '\n'
)

// NewMemoryStore 创建敏感词内存存储
func NewMemoryStore(config MemoryConfig) (*MemoryStore, error) {
	memStore := &MemoryStore{
		dataStore: cmap.NewConcurrencyMap(),
	}
	if config.Delim == 0 {
		config.Delim = DefaultDelim
	}
	if dataLen := len(config.DataSource); dataLen > 0 {
		for i := 0; i < dataLen; i++ {
			memStore.dataStore.Set(config.DataSource[i], 1)
		}
	} else if config.Reader != nil {
		buf := new(bytes.Buffer)
		io.Copy(buf, config.Reader)
		buf.WriteByte(config.Delim)
		for {
			line, err := buf.ReadString(config.Delim)
			if err != nil {
				if err == io.EOF {
					break
				}
				return nil, err
			}
			memStore.dataStore.Set(line, 1)
		}
		buf.Reset()
	}
	return memStore, nil
}

// MemoryConfig 敏感词内存存储配置
type MemoryConfig struct {
	// Reader 敏感词数据源
	Reader io.Reader
	// Delim 读取数据的分隔符
	Delim byte
	// DataSource 敏感词数据源
	DataSource []string
}

// MemoryStore 提供内存存储敏感词
type MemoryStore struct {
	version   uint64
	dataStore cmap.ConcurrencyMap
}

// Write Write
func (ms *MemoryStore) Write(words ...string) error {
	if len(words) == 0 {
		return nil
	}
	for i, l := 0, len(words); i < l; i++ {
		ms.dataStore.Set(words[i], 1)
	}
	atomic.AddUint64(&ms.version, 1)
	return nil
}

// Read Read
func (ms *MemoryStore) Read() <-chan string {
	chResult := make(chan string)
	go func() {
		for ele := range ms.dataStore.Elements() {
			chResult <- ele.Key.(string)
		}
		close(chResult)
	}()
	return chResult
}

// ReadAll ReadAll
func (ms *MemoryStore) ReadAll() ([]string, error) {
	dataKeys := ms.dataStore.Keys()
	dataLen := len(dataKeys)
	result := make([]string, dataLen)
	for i := 0; i < dataLen; i++ {
		result[i] = dataKeys[i].(string)
	}
	return result, nil
}

// Remove Remove
func (ms *MemoryStore) Remove(words ...string) error {
	if len(words) == 0 {
		return nil
	}
	for i, l := 0, len(words); i < l; i++ {
		ms.dataStore.Remove(words[i])
	}
	atomic.AddUint64(&ms.version, 1)
	return nil
}

// Version Version
func (ms *MemoryStore) Version() uint64 {
	return ms.version
}
