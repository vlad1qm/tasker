package cmd

import (
	"os"

	"github.com/spf13/cobra"
	t "tasker/internal/task"
	"tasker/internal/common"
)

var SubDeleteCmd = &cobra.Command{
	Use: "sub",
	Short: "delete subtask",
	Run: func(cmd *cobra.Command, args []string) {
		if taskId == 0 || subTaskId == 0{
			cmd.Help()
			os.Exit(1)
		}
		t.DeleteSubTask(common.IntToString(taskId), common.IntToString(subTaskId))
	}, 
}
