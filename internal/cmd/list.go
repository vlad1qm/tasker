package cmd

import (
	"os"
	t "tasker/internal/task"

	"github.com/spf13/cobra"
	"github.com/wsxiaoys/terminal/color"
)

var ListCmd = &cobra.Command{
	Use: "list",
	Short: "list all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		_, tasks := t.GetTasks()
		if len(tasks) == 0{
			color.Println("@rThere are no tasks")
			os.Exit(1)
		}
		tlt := t.TaskListTable[t.Task]{
			Tasks: tasks, 
			FilterFields: t.TaskListFilter, 
			Colorize: true,
		}
		tlt.MakeTaskTable()
	}, 
}