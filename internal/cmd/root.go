package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var taskId int
var subTaskId int

var rootCmd = &cobra.Command{
	Use: "task",
	Short: "working with tasks",
	Long: "",
	Run: func(cmd *cobra.Command, args []string){
		cmd.Help()
	},
}

func Execute(){
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init(){
	rootCmd.AddCommand(ListCmd)
	rootCmd.AddCommand(AddCmd)
	rootCmd.AddCommand(EditCmd)
	rootCmd.AddCommand(ShowCmd)
	rootCmd.AddCommand(DeleteCmd)

	EditCmd.Flags().VarP(TaskFieldsChoices, "field", "f", TaskFieldChoicesHelp)
	EditCmd.PersistentFlags().IntVarP(&taskId, "id", "i", 0, "task id")
	EditCmd.AddCommand(SubEditCmd)
	SubEditCmd.Flags().VarP(SubTaskFieldsChoices, "field", "f", SubTaskFieldChoicesHelp)
	SubEditCmd.PersistentFlags().IntVarP(&subTaskId, "sid", "s", 0, "sub task id")

	ListCmd.AddCommand(SubListCmd)
	SubListCmd.PersistentFlags().IntVarP(&taskId, "id", "i", 0, "task id")

	AddCmd.AddCommand(SubAddCmd)
	SubAddCmd.PersistentFlags().IntVarP(&taskId, "id", "i", 0, "task id")

	ShowCmd.AddCommand(SubShowCmd)
	ShowCmd.PersistentFlags().IntVarP(&taskId, "id", "i", 0, "task id")
	SubShowCmd.PersistentFlags().IntVarP(&subTaskId, "sid", "s", 0, "sub task id")

	DeleteCmd.AddCommand(SubDeleteCmd)
	DeleteCmd.PersistentFlags().IntVarP(&taskId, "id", "i", 0, "task id")
	SubDeleteCmd.PersistentFlags().IntVarP(&subTaskId, "sid", "s", 0, "sub task id")
}