package task

import (
	"fmt"
	"os"

	"github.com/fatih/structs"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/wsxiaoys/terminal/color"
)

var TaskListFilter = []string{FieldNote}


type TaskTable[T TaskType] struct {
	Task T
	TaskId string
	TaskTitle string
	TaskDescription string
	TableTitle string
	Headers []string
	Body []string
	Data [][]string
	Colorize bool
}

func (tt *TaskTable[T]) MakeHeaders(){
	headers :=structs.Names(tt.Task)
	tt.Headers = headers
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

func (tt *TaskTable[T]) MakeTitle(){
	title := func(n interface{})string{
		switch n.(type){
		case Task:
			return "Task:"
		case SubTask:
			return fmt.Sprintf("Task: %v\nTask title: %v\nTask description: %v\nSubTask:", 
			tt.TaskId, tt.TaskTitle, tt.TaskDescription)
		default:
			return ""
		}
	}
	tt.TableTitle = title(tt.Task)
}

func (tt *TaskTable[T]) ColorizeTaskTable(row []string)[]string{
	for index, element := range row{
		switch element{
		case FieldPriority:
			rowColor := PriorityColors[row[index+1]]
			msg := fmt.Sprintf("%v %v ", rowColor, row[index + 1])
			row[index + 1] = color.Sprintf(msg)
		case FieldStatus:
			rowColor := StatusColors[row[index+1]]
			msg := fmt.Sprintf("%v %v ", rowColor, row[index + 1])
			row[index + 1] = color.Sprintf(msg)
		case FieldUrl:
			rowColor := ColorUrl
			msg := fmt.Sprintf("%v %v ", rowColor, row[index + 1])
			row[index + 1] = color.Sprintf(msg)
		case FieldChecked:
			var rowColor string
			var newValue string
			if row[index + 1] == "yes"{
				rowColor = ColorCheckedTrue
				newValue = "[X]"
			} else {
				rowColor = ColorCheckedFalse
				newValue = "[O]"
			}
			msg := fmt.Sprintf("%v %v ", rowColor, newValue)
			row[index + 1] = color.Sprintf(msg)
			}
		}
		return row
}

func (tt *TaskTable[T]) MakeTaskTable(){
	tt.MakeHeaders()
	tt.MakeTaskBody()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	tt.MakeTitle()
	t.SetTitle(tt.TableTitle)
	for _, row := range tt.Data {
		if row[1] != ""{
			t.AppendRow(MakeRow(tt.ColorizeTaskTable(row)))
		}
		}
		t.SetStyle(table.StyleColoredBlackOnYellowWhite)
		t.SetColumnConfigs([]table.ColumnConfig{
			{Number: 2, WidthMin: 50},
		})
		t.Render()
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
				count := 0
				subtasks := value.([]SubTask)
				for _, subtask := range subtasks{
					if subtask.Checked == "no"{
						count++
					}
				}
				str = IntToString(count)
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

func (tlt *TaskListTable[T]) MakeTaskTable(){
	tlt.MakeHeaders()
	tlt.MakeListTasksData()
	tlt.Filter()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(MakeRow(tlt.Headers))
	for _, row := range tlt.Data{
		for i, t := range row{
			row[i] = text.WrapSoft(t, 40)
		}
		t.AppendRow(MakeRow(tlt.ColorizeTaskListTable(row)))
	}
	t.SetStyle(table.StyleColoredBlackOnYellowWhite)
	t.SortBy([]table.SortBy{
		{Name: "Status", Mode: table.AscNumeric},
	})
	t.SuppressEmptyColumns()
	t.Render()
}

func (tlt *TaskListTable[T]) ColorizeTaskListTable(row []string)[]string{
	for index, element := range row{
		switch tlt.Headers[index]{
		case FieldPriority:
			rowColor := PriorityColors[element]
			msg := fmt.Sprintf("%v %v ", rowColor, element)
			row[index] = color.Sprintf(msg)
		case FieldStatus:
			rowColor := StatusColors[element]
			msg := fmt.Sprintf("%v %v ", rowColor, element)
			row[index] = color.Sprintf(msg)
		case FieldChecked:
			var rowColor string
			var newValue string
			if row[index] == "yes"{
				rowColor = ColorCheckedTrue
				newValue = "[X]"
			} else {
				rowColor = ColorCheckedFalse
				newValue = "[O]"
			}
			msg := fmt.Sprintf("%v %v ", rowColor, newValue)
			row[index] = color.Sprintf(msg)
		case FieldUrl:
			if element != ""{
				rowColor := ColorUrl
				msg := fmt.Sprintf("%v %v ", rowColor, element)
				row[index] = color.Sprintf(msg)
			}

		}
	
}
	return row
}

func MakeRow(s []string) table.Row {
	new := make(table.Row, len(s))
	for index, value := range s{
		new[index] = value
	}
	return new
}