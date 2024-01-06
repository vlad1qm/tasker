package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	t "tasker/internal/task"
)

var SubListCmd = &cobra.Command{
	Use: "sub",
	Short: "list all subtasks",
	Run: func(cmd *cobra.Command, args []string) {
		if taskId == 0{
			cmd.Help()
			os.Exit(1)
		}
		subtasks := t.GetSubTasks(t.IntToString(taskId))
		fmt.Println(subtasks)
	}, 
}