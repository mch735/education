package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var filepath string

	flag.StringVar(&filepath, "file", "", "file path for analyze words")
	flag.Parse()

	reader, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	list := Words{}

	for scanner.Scan() {
		list.Add(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(list.Tops(5))
}
