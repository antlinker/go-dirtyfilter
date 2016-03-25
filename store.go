package filter

// DirtyStore 提供敏感词的读取、写入存储接口
type DirtyStore interface {
	// Write 将敏感词写入存储区，如果写入失败则返回error
	Write(words ...string) error

	// Read 以迭代的方式读取敏感词
	Read() <-chan string

	// ReadAll 获取所有的敏感词数据，如果获取失败则返回error
	ReadAll() ([]string, error)

	// Remove 移除敏感词,如果移除失败则返回error
	Remove(words ...string) error

	// Version 数据存储版本号
	Version() uint64
}
