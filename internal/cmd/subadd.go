package cmd

import (
	t "tasker/internal/task"
	"os"
	"github.com/spf13/cobra"
)

var SubAddCmd = &cobra.Command{
	Use: "sub",
	Short: "add subtask",
	Run: func(cmd *cobra.Command, args []string){
		var result string
		if taskId == 0{
			cmd.Help()
			os.Exit(1)
		}
		subtask := t.SubTask{}
		for _, field := range t.SubTaskAddFields{
			switch {
			case t.IsInSlice(field, t.SubTaskRowFields):
				result = t.RowInput(field, "")
			case t.IsInSlice(field, t.SubTaskEditorFields):
				result = t.TextInput(field, "")
			case t.IsInSlice(field, t.SubTaskChoiceFields):
				choices := t.Choices[field]
				result = t.ChoiceInput(field, choices)
			}
			subtask.Create(field, result)
		}
	
		subtask.Add(t.IntToString(taskId))
	},
}