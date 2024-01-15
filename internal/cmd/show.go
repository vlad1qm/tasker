package cmd

import (
	"fmt"
	"os"

	"tasker/internal/common"
	t "tasker/internal/task"

	"github.com/spf13/cobra"
	"github.com/wsxiaoys/terminal/color"
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
		if task.IsEmpty(){
			color.Printf(fmt.Sprintf("%vTask with id %v was not found\n", t.ColorRed, taskId))
			os.Exit(1)
		}
		tt := t.TaskTable[t.Task]{Task: task, Colorize: true}
		tt.MakeTaskTable()
	}, 
}
