package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// birthday_001.txt
	// filename := "birth_day_001.txt"
	// result, err := match(filename, 4)
	// if err != nil {
	// 	fmt.Println("no match")
	// 	os.Exit(1)
	// }
	// fmt.Println(result)
	files, err := os.ReadDir("./sample")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Dir: ", file.Name())
		} else {
			tmp, err := match(file.Name(), 0)
			fmt.Println("match: ", tmp, err)
		}
	}
}

func match(filename string, total int) (string, error) {
	// got:  birthday_001.txt
	// want: Birth.day - 1 of 4.txt
	// 1. Extract the extension
	pieces := strings.Split(filename, ".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")
	// 2. Extract the number
	pieces = strings.Split(tmp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		fmt.Printf("%s didn't match our patter", filename)
		return "", nil
	}
	// 3. Extract the filename (You already have it in tmp)
	return fmt.Sprintf("%s - %d of %d.%s", strings.Title(name), number, total, ext), nil
}
