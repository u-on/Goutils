package Goutils

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
)

func GbkToUtf8(b []byte) ([]byte, error) {
	reader := bytes.NewReader(b)
	decoder := simplifiedchinese.GBK.NewDecoder()
	transformedReader := transform.NewReader(reader, decoder)
	transformedBytes, err := io.ReadAll(transformedReader)
	if err != nil {
		return nil, err
	}
	return transformedBytes, nil
}
