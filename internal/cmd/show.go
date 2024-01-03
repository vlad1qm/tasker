package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	t "tasker/internal/task"
)

var ShowCmd = &cobra.Command{
	Use: "show",
	Short: "show task",
	Run: func(cmd *cobra.Command, args []string) {
		if taskId == 0{
			cmd.Help()
			os.Exit(1)
		}
		task := t.GetTask(t.IntToString(taskId))
		fmt.Println(task)
	}, 
}
