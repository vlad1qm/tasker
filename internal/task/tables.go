package task

import (
	// "reflect"
	"fmt"
	"os"

	"github.com/fatih/structs"
	"github.com/olekukonko/tablewriter"
)

type TaskTable[T TaskType] struct {
	Task T
	Headers []string
	Body []string
	Data [][]string
}

func (tt *TaskTable[T]) MakeHeaders(){
	tt.Headers = structs.Names(tt.Task)
}

func (tt *TaskTable[T]) MakeTaskBody(){
	rawValues := structs.Values(tt.Task)
	for index, value := range rawValues {
		var str string
		if tt.Headers[index] == FieldSubTasks{
			subtasks := value.([]SubTask)
			str = fmt.Sprint(len(subtasks))
		}else{
			str = fmt.Sprint(value)
		}
		tt.Body = append(tt.Body, str)
	}
	for headerIndex, header := range tt.Headers{
		row := []string{header, tt.Body[headerIndex]}
		tt.Data = append(tt.Data, row)
	}
}

func (tt *TaskTable[T]) MakeTaskTable(){
	tt.MakeHeaders()
	tt.MakeTaskBody()
	table := tablewriter.NewWriter(os.Stdout)
	table.AppendBulk(tt.Data)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()
}

type TaskListTable[T TaskType] struct {
	Tasks []T
	Headers []string
	Data [][]string
}

func (tlt *TaskListTable[T]) MakeHeaders(){
	tlt.Headers = structs.Names(tlt.Tasks[0])
}

func (tlt *TaskListTable[T])MakeListTasksData(){
	for _, task := range tlt.Tasks{
		rawValues := structs.Values(task)
		var body = []string{}
		for index, value := range rawValues {
			
			var str string
			if tlt.Headers[index] == FieldSubTasks{
				subtasks := value.([]SubTask)
				str = fmt.Sprint(len(subtasks))
			}else{
				str = fmt.Sprint(value)
			}
			body = append(body, str)
		}
		tlt.Data = append(tlt.Data, body)
	}
}

func (tlt *TaskListTable[T]) MakeTaskTable(){
	tlt.MakeHeaders()
	tlt.MakeListTasksData()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(tlt.Headers)
	table.AppendBulk(tlt.Data)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()
}