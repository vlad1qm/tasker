package cmd

import (
	// "fmt"
	"os"
	"github.com/spf13/cobra"
	t "tasker/internal/task"
	"tasker/internal/common"
)

var SubShowCmd = &cobra.Command{
	Use: "sub",
	Short: "show subtask",
	Run: func(cmd *cobra.Command, args []string)  {
		if taskId == 0 || subTaskId == 0{
			cmd.Help()
			os.Exit(1)
		}
		task := t.GetTask(common.IntToString(taskId))
		subtask := t.GetSubTask(common.IntToString(taskId), common.IntToString(subTaskId))
		tt := t.TaskTable[t.SubTask]{
			Task: subtask,
			TaskId: task.Id,
			TaskTitle: task.Title,
			TaskDescription: task.Description,
		}
		tt.MakeTaskTable()
	}, 
}
