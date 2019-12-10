package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Rot13Reader struct {
	r io.Reader
}

func (rot Rot13Reader) Read(p []byte) (x int, err error) {
	x, err = rot.r.Read(p)

	for i := 0; i < len(p); i++ {
		if p[i] >= 'A' && p[i] < 'Z' {
			p[i] = 65 + (((p[i] - 65) + 13) % 26)
		} else if p[i] >= 'a' && p[i] <= 'z' {
			p[i] = 97 + (((p[i] - 97) + 13) % 26)
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	fmt.Println(s)
	r := Rot13Reader{s}
	io.Copy(os.Stdout, &r)

}
