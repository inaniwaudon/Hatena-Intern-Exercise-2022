package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(p []byte) (int, error) {
	l, ok := reader.r.Read(p)
	if ok == nil {
		for i := 0; i < l; i++ {
			temp := p[i]
			if ('A' <= temp && temp <= 'Z') || ('a' <= temp && temp <= 'z') {
				p[i] += 13
			}
			if ('Z'-13 < temp && temp <= 'Z') || 'z'-13 < temp {
				p[i] -= 26
			}
		}
	}
	return l, ok
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
