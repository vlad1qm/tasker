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
	SubTaskFieldsChoices = newChoice(t.LCSubTaskFields, "")
	SubTaskFieldChoicesHelp string = fmt.Sprintf("subtask field to edit:[%v]", strings.Join(t.LCSubTaskFields, ","))
)

var SubEditCmd = &cobra.Command{
	Use: "sub",
	Short: "edit subtask",
	Run: func(cmd *cobra.Command, args []string){
		if taskId == 0 || subTaskId == 0{
			cmd.Help()
			os.Exit(1)
		}
		var result string
		taskId := t.IntToString(taskId)
		subTaskId := t.IntToString(subTaskId)
		field, _ := cmd.Flags().GetString("field")
		field = cases.Title(language.English, cases.Compact).String(field)
		fieldValue := t.GetSubTaskFieldValue(taskId, subTaskId, field)
			switch {
			case t.IsInSlice(field, t.SubTaskRowFields):
				result = t.RowInput(field, fieldValue)
			case t.IsInSlice(field, t.SubTaskEditorFields):
				result = t.TextInput(field, fieldValue)
			case t.IsInSlice(field, t.SubTaskChoiceFields):
				choices := t.Choices[field]
				result = t.ChoiceInput(field, choices)
			}
		t.EditSubTask(taskId, subTaskId, field, result)
	},
}