package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	t "tasker/internal/task"
)

var ListCmd = &cobra.Command{
	Use: "list",
	Short: "list all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		_, tasks := t.GetTasks()
		fmt.Println(tasks)
	}, 
}