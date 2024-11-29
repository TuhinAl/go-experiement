package datatypeoperator

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/*
 String in Go is essentially a read-only slice of bytes. and its structure can be represented as:
	1. Pointer to the data: A reference to the underlying byte array that holds the string data.
	2. Length: The number of bytes in the string.
*/

func SortingDemo() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		words[scan.Text()]++
	}
	fmt.Println(len(words), "unique words")

	type KeyValue struct {
		key   string
		value int
	}
	var slice1 []KeyValue
	for key, value := range words {
		slice1 = append(slice1, KeyValue{key, value})
	}
	sort.Slice(slice1, func(i, j int) bool {
		return slice1[i].value > slice1[j].value
	})
	for _, s := range slice1 {
		fmt.Println(s.key, " appears", s.value, " times")
	}
}
