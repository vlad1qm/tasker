package cmd

import (
	"github.com/spf13/cobra"
	t "tasker/internal/task"
)

var ListCmd = &cobra.Command{
	Use: "list",
	Short: "list all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		_, tasks := t.GetTasks()
		tlt := t.TaskListTable[t.Task]{Tasks: tasks}
		tlt.MakeTaskTable()
	}, 
}