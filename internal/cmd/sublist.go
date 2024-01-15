package cmd

import (
	"fmt"
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
		subtasks, err := t.GetSubTasks(common.IntToString(taskId))
		if err != nil{
			fmt.Println(err)
		}
		tlt := t.TaskListTable[t.SubTask]{
			Tasks: subtasks, 
			ColumnFilterFields: t.ColumnTaskListFilter,
			Colorize: false,
		}
		tlt.MakeTaskTable()
	}, 
}