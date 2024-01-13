package cmd

import (
	"os"
	"github.com/spf13/cobra"
	t "tasker/internal/task"
	"tasker/internal/common"
)

var SubListCmd = &cobra.Command{
	Use: "sub",
	Short: "list all subtasks",
	Run: func(cmd *cobra.Command, args []string) {
		if taskId == 0{
			cmd.Help()
			os.Exit(1)
		}
		subtasks := t.GetSubTasks(common.IntToString(taskId))
		tlt := t.TaskListTable[t.SubTask]{
			Tasks: subtasks, 
			FilterFields: t.TaskListFilter,
			Colorize: false,
		}
		tlt.MakeTaskTable()
	}, 
}