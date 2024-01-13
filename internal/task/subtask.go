package task

import (
	"reflect"
	"os"
	"fmt"
	"github.com/wsxiaoys/terminal/color"
	"tasker/internal/common"
)

var (
	SubTaskAddFields = []string{FieldTitle, FieldDescription, FieldNote, FieldChecked}
	SubTaskEditFields = []string{FieldTitle, FieldDescription, FieldNote}
	SubTaskRowFields = []string{FieldTitle, FieldDescription}
	SubTaskEditorFields = []string{FieldNote}
	SubTaskChoiceFields = []string{FieldChecked}
)


type SubTask struct {
	Id string `yaml:"id"`
	Checked string `yaml:"checked"`
	Title string `yaml:"title"`
	Description string `yaml:"description"`
	Created string `yaml:"created"`
	Updated string `yaml:"updated"`
	Note string `yaml:"note"`
}

func (st *SubTask)Create(fieldName string, fieldData string){
	taskElements := reflect.ValueOf(st).Elem()
	taskField := taskElements.FieldByName(fieldName)
	if taskField.CanSet(){
		taskField.SetString(fieldData)
	}
}

func (st *SubTask)Add(taskId string)error{
	y, tasks, err := GetTasks()
	if err != nil{
		fmt.Println(err)
	}
	taskIndex := FindTaskIndex(tasks, FieldId, taskId)
	subtasks := tasks[taskIndex].SubTasks
	st.Id = common.IntToString(GetNewId(subtasks))
	st.Created = common.GetCurrentTime(common.TimeFormat)
	subtasks = append(subtasks, *st)
	tasks[taskIndex].SubTasks = subtasks
	y.Write(tasks)
	color.Printf("@gSubTask with id %v was created within Task with id %v\n", st.Id, taskId)
	return nil
}

func (st SubTask)GetId()string{
	return st.Id
}

func (st SubTask)GetValueOf(fieldName string)string{
	subtask := &st
	elements := reflect.ValueOf(subtask).Elem()
	field := elements.FieldByName(fieldName)
	return field.String()
}

func GetSubTask(taskId string, subTaskId string)SubTask{
	_, tasks, err := GetTasks()
	if err != nil{
		fmt.Println(err)
	}
	for _, task := range tasks{
		if task.Id == taskId{
			for _, subtask := range task.SubTasks{
				if subtask.Id == subTaskId{
					return subtask
				}
			}
		}
	}
	return SubTask{}
}

func GetSubTasks(taskId string)[]SubTask{
	_, tasks, err := GetTasks()
	if err != nil{
		fmt.Println(err)
	}
	if len(tasks) == 0{
		color.Println("@rThere are no tasks")
		os.Exit(1)
	}
	for _, task := range tasks{
		if task.Id == taskId{
			return task.SubTasks
		}
	}
	return []SubTask{}
}

func GetSubTaskFieldValue(taskId string, subTaskId string, fieldName string)string{
	subtask := GetSubTask(taskId, subTaskId)
	field := subtask.GetValueOf(fieldName)
	return field
}

func EditSubTask(taskId string, subTaskId string, fieldName string, fieldData string)error{
	fieldType := reflect.ValueOf(fieldData)
	y, tasks, err := GetTasks()
	if err != nil{
		fmt.Println(err)
	}

	taskIndex := FindTaskIndex(tasks, FieldId, taskId)
	subtasks := tasks[taskIndex].SubTasks
	subTaskIndex := FindTaskIndex(subtasks, FieldId, subTaskId)

	st := &subtasks[subTaskIndex]
	taskElements := reflect.ValueOf(st).Elem()
	taskField := taskElements.FieldByName(fieldName)
	
	switch fieldType.Kind() {
	case reflect.String:
		if taskField.CanSet(){
			taskField.SetString(fieldType.String())
		}
	case reflect.Bool:
		if taskField.CanSet(){
			taskField.SetBool(fieldType.Bool())
		}
	}
	st.Updated = common.GetCurrentTime(common.TimeFormat)
	y.Write(tasks)
	return nil
}

func DeleteSubTask(taskId string, subTaskId string)error{
	y, tasks, err := GetTasks()
	if err != nil{
		fmt.Println(err)
	}
	taskIndex := FindTaskIndex(tasks, FieldId, taskId)
	subtasks := tasks[taskIndex].SubTasks
	result := DeleteFromTasks(subtasks, FieldId, subTaskId)
	tasks[taskIndex].SubTasks = result
	y.Write(tasks)
	return nil
}