package main

import (
	"fmt"
)

type Cat struct {
	name      string
	years     int
	food      string
	voiceWord string
	isSilence bool
}

func (cat Cat) getName() string {
	return cat.name
}

func (cat *Cat) setName(name string) string {
	cat.name = name
	return cat.name
}

func (cat Cat) voice() string {
	if cat.isSilence != false {
		return cat.voiceWord
	}
	return "The cat can't meow!"
}

func main() {
	cat := Cat{"Kitty", 2, "meat", "meow!!", true}
	fmt.Println(cat.getName())
	fmt.Println(cat.setName("Barsik"))
	fmt.Println(cat.getName())
}
