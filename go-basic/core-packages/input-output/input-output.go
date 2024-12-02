package inputoutput

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func FileIOdemo() {
	fmt.Println("Program Name: ", os.Args[0])
	for _, fName := range os.Args[1:] { //[1:] means we're interested in the arguments after the program name.
		fmt.Println("F Name: ", fName)
		fmt.Println("os.Args[1:]: ", os.Args[1:])
		file, err := os.Open(fName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		file.Close()
	}
}

func CalculateFileSize() {
	for _, fName := range os.Args[1:] {
		file, err := os.Open(fName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		data, err := io.ReadAll(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Println("The File has: ", float64(len(data))*0.0001, " KB")
		file.Close()
	}
}

func WordLineAndCharacterCount() {
	var wordCount, lineCount, charCount int
	for _, fName := range os.Args[1:] {
		file, err := os.Open(fName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		scan := bufio.NewScanner(file)
		for scan.Scan(){
			str := scan.Text()
			wordCount += len(strings.Fields(str))
			charCount += len(str)
			lineCount++
		}

		
		fmt.Println("The File has: ", lineCount, " lines, ",wordCount, " words and ", charCount, " characters.")
		file.Close()
	}
}
