package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	t "tasker/internal/task"
)

var SubShowCmd = &cobra.Command{
	Use: "sub",
	Short: "show subtask",
	Run: func(cmd *cobra.Command, args []string)  {
		if taskId == 0 || subTaskId == 0{
			cmd.Help()
			os.Exit(1)
		}
		subtask := t.GetSubTask(t.IntToString(taskId), t.IntToString(subTaskId))
		fmt.Println(subtask)
	}, 
}
