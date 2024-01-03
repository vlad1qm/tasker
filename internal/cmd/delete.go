package cmd

import (
	"os"

	"github.com/spf13/cobra"
	t "tasker/internal/task"
)

var DeleteCmd = &cobra.Command{
	Use: "delete",
	Short: "delete task",
	Run: func(cmd *cobra.Command, args []string) {
		if taskId == 0{
			cmd.Help()
			os.Exit(1)
		}
		t.DeleteTask(t.IntToString(taskId))
	}, 
}
