package cmd

import (
	"fmt"
	"os"
	"strings"
	t "tasker/internal/task"
	"tasker/internal/common"

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
		taskId := common.IntToString(taskId)
		subTaskId := common.IntToString(subTaskId)
		field, _ := cmd.Flags().GetString("field")
		field = cases.Title(language.English, cases.Compact).String(field)
		fieldValue := t.GetSubTaskFieldValue(taskId, subTaskId, field)
		i := common.Input{
			FieldName: field,
		}
			switch {
			case common.IsInSlice(field, t.SubTaskRowFields):
				i.Data = fieldValue
				result = i.Row()
			case common.IsInSlice(field, t.SubTaskEditorFields):
				i.Data = fieldValue
				result = i.Text()
			case common.IsInSlice(field, t.SubTaskChoiceFields):
				i.Choices = t.Choices[field]
				result = i.Choice()
			}
		t.EditSubTask(taskId, subTaskId, field, result)
	},
}