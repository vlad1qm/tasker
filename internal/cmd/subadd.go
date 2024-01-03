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
		if taskId == 0{
			cmd.Help()
			os.Exit(1)
		}
		subtask := t.SubTask{}
		for _, field := range t.SubTaskAddFields{
			result := t.RowInput(field, "")
			subtask.Create(field, result)
		}
		subtask.Add(t.IntToString(taskId))
	},
}