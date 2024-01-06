package cmd

import (
	// "fmt"
	t "tasker/internal/task"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use: "add",
	Short: "add task",
	Run: func(cmd *cobra.Command, args []string){
		var result string
		task := t.Task{}
		for _, field := range t.TaskAddFields{
			switch {
			case t.IsInSlice(field, t.TaskRowFields):
				result = t.RowInput(field, "")
			case t.IsInSlice(field, t.TaskEditorFields):
				result = t.TextInput(field, "")
			case t.IsInSlice(field, t.TaskChoiceFields):
				choices := t.Choices[field]
				result = t.ChoiceInput(field, choices)
			}
			task.Create(field, result)
		}
		task.Add()
	},
}