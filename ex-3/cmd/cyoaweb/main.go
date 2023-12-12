package main

import (
	"flag"
	"fmt"
	"os"

	cyoa "github.com/Mauricio-3107/gophercises.git/ex-3"
)

func main() {
	filename := flag.String("file", "gopher.json", "A json file to choose you own adventure")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", story)
}
