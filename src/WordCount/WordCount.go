package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var strCache map[string]int = make(map[string]int)
	var strMap map[string]int = make(map[string]int)
	var fields []string = strings.Fields(s)
	for _, fieldKey := range fields {
		if _, ok := strCache[fieldKey]; ok {
			strCache[fieldKey] = strCache[fieldKey] + 1
		} else {
			strCache[fieldKey] = 1
		}
		strMap[fieldKey] = strCache[fieldKey]
	}
	fmt.Printf("strMap: %v", strMap)
	return strMap
}

func main() {
	wc.Test(WordCount)
}
