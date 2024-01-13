package cmd

import (
	"os"

	"github.com/spf13/cobra"
	t "tasker/internal/task"
	"tasker/internal/common"
)

var ShowCmd = &cobra.Command{
	Use: "show",
	Short: "show task",
	Run: func(cmd *cobra.Command, args []string) {
		if taskId == 0{
			cmd.Help()
			os.Exit(1)
		}
		task := t.GetTask(common.IntToString(taskId))
		tt := t.TaskTable[t.Task]{Task: task, Colorize: true}
		tt.MakeTaskTable()
	}, 
}
