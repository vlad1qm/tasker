package task

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"tasker/internal/common"
	c "tasker/internal/config"
	"github.com/wsxiaoys/terminal/color"
)

const (
	PriorityDefault string = PriorityMedium
	PriorityLow string = "low"
	PriorityMedium string = "medium"
	PriorityHigh string = "high"
	PriorityUrgent string = "urgent"

	StatusDefault string = StatusNew
	StatusNew string = "new"
	StatusOpen string = "in progress"
	StatusPause string = "pause"
	StatusClosed string = "done"
	StatusWaiting string = "waiting"

	CheckedTrue string = "yes"
	CheckedFalse string = "no"

	FieldId string = "Id"
	FieldTitle string = "Title"
	FieldUrl string = "Url"
	FieldDescription string = "Description"
	FieldNote string = "Note"
	FieldStatus string = "Status"
	FieldPriority string = "Priority"
	FieldSubTasks string = "SubTasks"
	FieldChecked string = "Checked"
)

var (
	Config = c.Config{}

	LCFieldTitle string = strings.ToLower(FieldTitle)
	LCFieldUrl string = strings.ToLower(FieldUrl)
	LCFieldDescription string = strings.ToLower(FieldDescription)
	LCFieldNote string = strings.ToLower(FieldNote)
	LCFieldStatus string = strings.ToLower(FieldStatus)
	LCFieldPriority string = strings.ToLower(FieldPriority)
	LCFieldChecked string = strings.ToLower(FieldChecked)
	LCTaskFields = []string{LCFieldTitle, LCFieldUrl, LCFieldDescription, LCFieldNote, LCFieldStatus, LCFieldPriority}
	LCSubTaskFields = []string{LCFieldTitle, LCFieldDescription, LCFieldNote, LCFieldChecked}

	TaskAddFields = []string{FieldTitle, FieldUrl, FieldDescription, FieldNote, FieldStatus, FieldPriority}
	TaskEditFields = []string{FieldTitle, FieldUrl, FieldDescription, FieldNote, FieldStatus, FieldPriority}
	TaskRowFields = []string{FieldTitle, FieldUrl, FieldDescription}
	TaskEditorFields = []string{FieldNote}
	TaskChoiceFields = []string{FieldStatus, FieldPriority}
	Choices = map[string][]string{
		FieldStatus: {StatusNew, StatusOpen, StatusPause, StatusWaiting, StatusClosed},
		FieldPriority: {PriorityLow, PriorityMedium, PriorityHigh, PriorityUrgent},
		FieldChecked: {CheckedTrue, CheckedFalse},
	}
)

type Task struct {
	Id string `yaml:"id"`
	Status string `yaml:"status"`
	Priority string `yaml:"priority"`
	Title string `yaml:"title"`
	Url string `yaml:"url"`
	Description string `yaml:"description"`
	Note string `yaml:"note"`
	Created string `yaml:"created"`
	Updated string `yaml:"updated"`
	SubTasks []SubTask `yaml:"subtasks"`
}

func (t *Task)IsEmpty()bool{
	return t.Id == ""
}

func (t *Task)Create(fieldName string, fieldData string){
	taskElements := reflect.ValueOf(t).Elem()
	taskField := taskElements.FieldByName(fieldName)
	if taskField.CanSet(){
		taskField.SetString(fieldData)
	}
}

func (t *Task)Add()error{
	y, tasks, err := GetTasks()
	if err != nil{
		color.Printf(fmt.Sprintf("%vError %v was occured\n", ColorRed, err))
		os.Exit(1)
	}
	t.Id = common.IntToString(GetNewId(tasks))
	if t.Status == ""{
		t.Status = StatusDefault
	}
	t.Created = common.GetCurrentTime(Config.TimeFormat)
	if t.Priority == ""{
		t.Priority = PriorityDefault
	}
	tasks = append(tasks, *t)
	y.Write(tasks)
	color.Printf(fmt.Sprintf("%vTask with id %v was created\n", ColorGreen, t.Id))
	return nil
}

func (t Task)GetId()string{
	return t.Id
}

func (t Task)GetValueOf(fieldName string)string{
	task := &t
	elements := reflect.ValueOf(task).Elem()
	field := elements.FieldByName(fieldName)
	return field.String()
}

type TaskType interface {
	Task | SubTask
	GetId() string
	GetValueOf(string) string
}

func GetNewId[T TaskType](t []T)int{
	var newId int
	if len(t) == 0{
		return 1
	}
	for _, task := range t{
		taskId := common.StringToInt(task.GetId())
		if taskId > newId{
			newId = taskId
		}
	}
	return newId + 1
}

func FindTaskIndex[T TaskType](t []T, fieldName string, key string) int{
	for index, value := range t{
		if key == value.GetValueOf(fieldName) {
			return index
		}
	}
	return -1
}

func DeleteFromTasks[T TaskType](t []T, fieldName string, taskId string) []T{
	index := 0
	for _, value := range t{
		if taskId != value.GetValueOf(fieldName){
			t[index] = value
			index++
		}
	}
	return t[:index]
}

func GetTask(taskId string)Task{
	_, tasks, err := GetTasks()
	if err != nil{
		color.Printf(fmt.Sprintf("%vError %v was occured\n", ColorRed, err))
		os.Exit(1)
	}

	for _, task := range tasks{
		if task.Id == taskId{
			return task
		}
	}
	return Task{}
}

func GetTasks()(common.Yaml, []Task, error){
	y := common.Yaml{FilePath: Config.DataPath, FileName: Config.DataFileName}
	y.Read()
	decoded := y.Decode([]Task{})
	tasks, ok := decoded.([]Task)
	if ok{
		return y, tasks, nil
	}
	return y, []Task{}, errors.New("couldnt get tasks")

}

func GetTaskFieldValue(taskId string, fieldName string)string{
	task := GetTask(taskId)
	if task.IsEmpty(){
		color.Printf(fmt.Sprintf("%vTask with id %v was not found\n", ColorRed, taskId))
		os.Exit(1)
	}
	field := task.GetValueOf(fieldName)
	return field
}

func EditTask(taskId string, fieldName string, fieldData string)error{
	y, tasks, err := GetTasks()
	if err != nil{
		color.Printf(fmt.Sprintf("%vError %v was occured\n", ColorRed, err))
		os.Exit(1)
	}
	taskIndex := FindTaskIndex(tasks, FieldId, taskId)

	if taskIndex == -1 {
		color.Printf(fmt.Sprintf("%vTask with id %v was not found\n", ColorRed, taskId))
		os.Exit(1)
	}

	t := &tasks[taskIndex]
	taskElements := reflect.ValueOf(t).Elem()
	taskField := taskElements.FieldByName(fieldName)

	if taskField.CanSet(){
		taskField.SetString(fieldData)
	}
	t.Updated = common.GetCurrentTime(Config.TimeFormat)
	
	y.Write(tasks)
	return nil
}

func DeleteTask(taskId string)error{
	y, tasks, err := GetTasks()
	if err != nil{
		color.Printf(fmt.Sprintf("%vError %v was occured\n", ColorRed, err))
		os.Exit(1)
	}
	taskIndex := FindTaskIndex(tasks, FieldId, taskId)
	if taskIndex == -1 {
		color.Printf(fmt.Sprintf("%vTask with id %v was not found\n", ColorRed, taskId))
		os.Exit(1)
	}
	result := DeleteFromTasks(tasks, FieldId, taskId)
	y.Write(result)
	return nil
}
