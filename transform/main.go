package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("main", strings.Fields("-i tmp/jump.jpg -o out.jpg -n 100")...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
