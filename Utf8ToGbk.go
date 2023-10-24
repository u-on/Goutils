package Goutils

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
)

// Utf8ToGbk 将 UTF-8 编码的字符串转换为 GBK 编码。
//
// text: 要转换的 UTF-8 字符串。
// 返回值: 转换后的 GBK 编码字符串。
func Utf8ToGbk(text string) string {
	decoder := simplifiedchinese.GBK.NewDecoder()
	content, _ := io.ReadAll(transform.NewReader(bytes.NewReader([]byte(text)), decoder))
	return string(content)
}
