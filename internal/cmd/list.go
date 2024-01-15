package cmd

import (
	"fmt"
	"os"
	c "tasker/internal/config"
	t "tasker/internal/task"

	"github.com/spf13/cobra"
	"github.com/wsxiaoys/terminal/color"
)

var ListCmd = &cobra.Command{
	Use: "list",
	Short: "list all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(c.Configuration)
		 _, tasks, err := t.GetTasks()
		if err != nil{
			fmt.Println(err)
		}
		if len(tasks) == 0{
			color.Println("@rThere are no tasks")
			os.Exit(1)
		}
		tlt := t.TaskListTable[t.Task]{
			Tasks: tasks, 
			ColumnFilterFields: t.Config.ColumnTaskListFilter, 
			RowFilterFields: t.Config.RowTaskListFilter,
			Colorize: true,
		}
		tlt.MakeTaskTable()
	}, 
}