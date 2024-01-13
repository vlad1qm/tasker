package cmd

import (
	t "tasker/internal/task"
	"tasker/internal/common"
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
			i := common.Input{
				FieldName: field,
				Data: "",
			}
			switch {
			case common.IsInSlice(field, t.SubTaskRowFields):
				result = i.Row()
			case common.IsInSlice(field, t.SubTaskEditorFields):
				result = i.Text()
			case common.IsInSlice(field, t.SubTaskChoiceFields):
				i.Choices = t.Choices[field]
				result = i.Choice()
			}
			subtask.Create(field, result)
		}
	
		subtask.Add(common.IntToString(taskId))
	},
}