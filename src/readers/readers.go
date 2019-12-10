package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (myReader MyReader) Read(a []byte) (int, error) {
	a = a[:cap(a)]
	for i := range a {
		a[i] = 'A'
	}
	// fmt.Println(a)
	return cap(a), nil
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
