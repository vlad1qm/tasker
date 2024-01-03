package cmd

import (
	t "tasker/internal/task"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use: "add",
	Short: "add task",
	Run: func(cmd *cobra.Command, args []string){
		task := t.Task{}
		for _, field := range t.TaskAddFields{
			result := t.RowInput(field, "")
			task.Create(field, result)
		}
		task.Add()
	},
}