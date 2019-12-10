package main

import "fmt"

type TInterface interface {
	message()
}

type TStruct struct {
}

func (tStruct *TStruct) message() {
	if tStruct == nil {
		fmt.Println("<nil>")
	}
	// fmt.Println(tStruct.s)
}

func main() {
	var tInterface TInterface
	var tStruct *TStruct
	tInterface = tStruct

	tInterface = &TStruct{}
	tInterface.message()
}
