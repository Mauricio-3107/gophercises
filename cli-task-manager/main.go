package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Mauricio-3107/gophercises.git/cli-task-manager/cmd"
	"github.com/Mauricio-3107/gophercises.git/cli-task-manager/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks-cli.db")
	must(db.Init(dbPath))

	must(cmd.RootCmd.Execute())

}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
