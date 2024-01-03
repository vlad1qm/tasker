package task

import (
	"reflect"
)

const (
	PriorityDefault string = PriorityMedium
	PriorityLow string = "low"
	PriorityMedium string = "medium"
	PriorityHigh string = "high"
	PriorityUrgent string = "urgent"

	StatusDefault string = StatusNew
	StatusNew string = "new"
	StatusOpen string = "open"
	StatusPause string = "pause"
	StatusClosed string = "closed"

	FieldId string = "Id"
)

var TaskAddFields = []string{"Title", "Url", "Description"}

type Task struct {
	Id string `yaml:"id"`
	Title string `yaml:"title"`
	Url string `yaml:"url"`
	Description string `yaml:"description"`
	Note string `yaml:"note"`
	Priority string `yaml:"priority"`
	Created string `yaml:"created"`
	Status string `yaml:"status"`
	Updated string `yaml:"updated"`
	SubTasks []SubTask `yaml:"subtasks"`
}

func (t *Task)Create(fieldName string, fieldData string){
	taskElements := reflect.ValueOf(t).Elem()
	taskField := taskElements.FieldByName(fieldName)
	if taskField.CanSet(){
		taskField.SetString(fieldData)
	}
}

func (t *Task)Add()error{
	y, tasks := GetTasks()
	t.Id = IntToString(GetNewId(tasks))
	if t.Status == ""{
		t.Status = StatusDefault
	}
	t.Created = GetCurrentTime()
	if t.Priority == ""{
		t.Priority = PriorityDefault
	}
	tasks = append(tasks, *t)
	y.Write(tasks)
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

func GetTask(taskId string)Task{
	_, tasks := GetTasks()
	for _, task := range tasks{
		if task.Id == taskId{
			return task
		}
	}
	return Task{}
}

func GetTasks()(Yaml, []Task){
	y := Yaml{Path: FilePath}
	tasks := y.GetTasks()
	return y, tasks
}

func EditTask(taskId string, fieldName string, fieldData string)error{
	y, tasks := GetTasks()
	taskIndex := FindIndex(tasks, FieldId, taskId)

	t := &tasks[taskIndex]
	taskElements := reflect.ValueOf(t).Elem()
	taskField := taskElements.FieldByName(fieldName)

	if taskField.CanSet(){
		taskField.SetString(fieldData)
	}
	
	y.Write(tasks)
	return nil
}

func DeleteTask(taskId string)error{
	y, tasks := GetTasks()
	result := DeleteFromTasks(tasks, FieldId, taskId)
	y.Write(result)
	return nil
}
