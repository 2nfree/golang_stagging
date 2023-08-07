package utils

import (
	"bufio"
	"io"
	"os"
)

// ReadLines 按行读取文件
func ReadLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{""}, err
	}
	defer f.Close()
	var ret []string
	line := bufio.NewReader(f)
	for {
		content, _, err := line.ReadLine()
		if err == io.EOF {
			break
		}
		ret = append(ret, string(content))
	}
	return ret, nil
}
