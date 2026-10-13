package main

import (
	"fmt"
	"io"
	"strings"
)

// bufferedReader 是一个自定义的缓冲读取器，它实现了 io.Reader 接口。它在读取数据时使用一个内部缓冲区来提高读取效率。

type bufferedReader struct {
	io.Reader
	buffer []byte
}

func newBufferedReader(r io.Reader) *bufferedReader {
	return &bufferedReader{
		Reader: r,
		buffer: make([]byte, 1024),
	}
}

func (br *bufferedReader) Read(p []byte) (n int, err error) {
	n, err = br.Reader.Read(p)

	if err != nil && err != io.EOF {
		return n, err
	}
	return n, err
}

func main() {
	strReader := strings.NewReader("Hello, Go!")

	br := newBufferedReader(strReader)

	buf := make([]byte, 10)

	n, err := br.Read(buf)

	if err != nil && err != io.EOF {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("Read %d bytes: %s \n", n, string(buf[:n]))
}
