package cmd

import (
	"fmt"
	"strings"

	"github.com/Mauricio-3107/gophercises.git/cli-task-manager/db"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This command is to add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		if len(task) <= 0 {
			fmt.Println("You dind't type anything.")
			return
		}
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			return
		}

		fmt.Printf("Added the \"%s\" task", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
