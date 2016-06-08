package filter

import (
	"io"
)

// DirtyFilter 提供敏感词过滤接口
type DirtyFilter interface {
	// Filter 文本过滤函数
	// excludes 表示排除指定的字符
	// 返回文本中出现的敏感词,如果敏感词不存在则返回nil
	// 如果出现异常，则返回error
	Filter(text string, excludes ...rune) ([]string, error)

	// FilterResult 文本过滤函数
	// excludes 表示排除指定的字符
	// 返回文本中出现的敏感词及出现次数,如果敏感词不存在则返回nil
	// 如果出现异常，则返回error
	FilterResult(text string, excludes ...rune) (map[string]int, error)

	// FilterReader 从可读流中过滤敏感词
	// excludes 表示排除指定的字符
	// 返回可读流中出现的敏感词，如果敏感词不存在则返回nil
	// 如果出现异常，则返回error
	FilterReader(reader io.Reader, excludes ...rune) ([]string, error)

	// FilterReaderResult 从可读流中过滤敏感词
	// excludes 表示排除指定的字符
	// 返回可读流中出现的敏感词及出现次数，如果敏感词不存在则返回nil
	// 如果出现异常，则返回error
	FilterReaderResult(reader io.Reader, excludes ...rune) (map[string]int, error)

	// Replace 使用字符替换文本中的敏感词
	// delim 替换的字符
	// 如果出现异常，则返回error
	Replace(text string, delim rune) (string, error)
}
