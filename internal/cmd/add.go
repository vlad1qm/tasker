package cmd

import (
	t "tasker/internal/task"
	"tasker/internal/common"
	"github.com/spf13/cobra"
)


var AddCmd = &cobra.Command{
	Use: "add",
	Short: "add task",
	Run: func(cmd *cobra.Command, args []string){
		var result string
		task := t.Task{}
		for _, field := range t.TaskAddFields{

			i := common.Input{
				FieldName: field, 
				Data: "",
				Prompt: t.Config.RowPrompt,
				TmpPath: t.Config.TmpPath,
				TextEditor: t.Config.TextEditor,
				NewFolderPermissions: common.NewFolderPermissions,
			}
			
			switch {
			case common.IsInSlice(field, t.TaskRowFields):
				result = i.Row()
			case common.IsInSlice(field, t.TaskEditorFields):
				result = i.Text()
			case common.IsInSlice(field, t.TaskChoiceFields):
				i.Choices = t.Choices[field]
				result = i.Choice()
			}
			task.Create(field, result)
		}
		task.Add()
	},
}