package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// LimitedReader struct
type LimitedReader struct {
	r io.Reader
	n int64
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.n <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.n {
		p = p[0:l.n]
	}
	n, err = l.r.Read(p)
	l.n -= int64(n)
	return
}

// LimitReader is limited reader that returns n
func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}

func main() {
	// r := LimitReader(os.Stdin, 5)
	r := LimitReader(os.Stdin, 5)
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	fmt.Printf("%v", string(buf))
}
