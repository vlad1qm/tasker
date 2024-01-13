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
		taskId := common.IntToString(taskId)
		field, _ := cmd.Flags().GetString("field")
		field = cases.Title(language.English, cases.Compact).String(field)
		fieldValue := t.GetTaskFieldValue(taskId, field)
		i := common.Input{
			FieldName: field,
		}
			switch {
			case common.IsInSlice(field, t.TaskRowFields):
				i.Data = fieldValue
				result = i.Row()
			case common.IsInSlice(field, t.TaskEditorFields):
				i.Data = fieldValue
				result = i.Text()
			case common.IsInSlice(field, t.TaskChoiceFields):
				i.Choices = t.Choices[field]
				result = i.Choice()
			}
		t.EditTask(taskId, field, result)
	},
}