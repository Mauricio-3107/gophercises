package main

import (
	"fmt"

	secret "github.com/Mauricio-3107/gophercises.git/secrets-api-cli"
)

func main() {
	v := secret.File("my-fake-key", ".secrets")
	err := v.Set("twitter", "my awesome secret")
	if err != nil {
		panic(err)
	}
	plain, err := v.Get("twitter")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain:", plain)
}
