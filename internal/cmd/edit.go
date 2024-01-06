package cmd

import (
	"fmt"
	"os"
	"strings"
	t "tasker/internal/task"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/spf13/cobra"
)

var (
	TaskFieldsChoices = newChoice(t.LCTaskFields, "")
	TaskFieldChoicesHelp string = fmt.Sprintf("task field to edit:[%v]", strings.Join(t.LCTaskFields, ","))
)

var EditCmd = &cobra.Command{
	Use: "edit",
	Short: "edit task",
	Run: func(cmd *cobra.Command, args []string){
		if taskId == 0{
			cmd.Help()
			os.Exit(1)
		}
		var result string
		taskId := t.IntToString(taskId)
		field, _ := cmd.Flags().GetString("field")
		field = cases.Title(language.English, cases.Compact).String(field)
		fieldValue := t.GetTaskFieldValue(taskId, field)
			switch {
			case t.IsInSlice(field, t.TaskRowFields):
				result = t.RowInput(field, fieldValue)
			case t.IsInSlice(field, t.TaskEditorFields):
				result = t.TextInput(field, fieldValue)
			case t.IsInSlice(field, t.TaskChoiceFields):
				choices := t.Choices[field]
				result = t.ChoiceInput(field, choices)
			}
		t.EditTask(taskId, field, result)
	},
}