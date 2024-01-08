package task

import (
	// "reflect"
	"fmt"
	"os"

	"github.com/fatih/structs"
	"github.com/olekukonko/tablewriter"
)

var TaskListFilter = []string{FieldNote}

type TaskTable[T TaskType] struct {
	Task T
	Headers []string
	Body []string
	Data [][]string
	Colorize bool
}

func (tt *TaskTable[T]) MakeHeaders(){
	tt.Headers = structs.Names(tt.Task)
}

func (tt *TaskTable[T])ColorizeRow(row []string)[]tablewriter.Colors{
	headers := []string{row[0]}
	r := []string{row[1]}
	m := MakeMap(headers, r)
	c := ColorFilter{
		Tasks: m, 
		Headers: headers, 
		RowColors: []tablewriter.Colors{},
		Type: TaskColorType,
	}
	c.Process()
	return c.RowColors
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
	if tt.Colorize{
		for _, row := range tt.Data{
			colors := tt.ColorizeRow(row)
			table.Rich(row, colors)
		}
	} else {
		table.AppendBulk(tt.Data)
	}
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	// table.SetColMinWidth(1, 50)
	// table.SetAutoWrapText(false)
	table.Render()
}

type TaskListTable[T TaskType] struct {
	Tasks []T
	FilterFields []string
	Headers []string
	Data [][]string
	Colorize bool
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

func (tlt *TaskListTable[T])Filter(){
	var index int
	for _, filter := range tlt.FilterFields{
		index = FindIndex(tlt.Headers, filter)
		tlt.Headers = DeleteFromSliceByIndex(tlt.Headers, index)
		for taskIndex, _ := range tlt.Data{
			tlt.Data[taskIndex] = DeleteFromSliceByIndex(tlt.Data[taskIndex], index)
		}
	}
}

func (tlt *TaskListTable[T])ColorizeRow(row []string)[]tablewriter.Colors{
	m := MakeMap(tlt.Headers, row)
	c := ColorFilter{
		Tasks: m, 
		Headers: tlt.Headers, 
		RowColors: []tablewriter.Colors{},
		Type: TaskListColorType,
	}
	c.Process()
	return c.RowColors
}

func (tlt *TaskListTable[T]) MakeTaskTable(){
	tlt.MakeHeaders()
	tlt.MakeListTasksData()
	tlt.Filter()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(tlt.Headers)
	if tlt.Colorize{
		for _, row := range tlt.Data{
			colors := tlt.ColorizeRow(row)
			table.Rich(row, colors)
		}
	} else {
		table.AppendBulk(tlt.Data)
	}
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()
}