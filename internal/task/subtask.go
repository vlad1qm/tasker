package task

import (
	"reflect"
)

var SubTaskAddFields = []string{"Title", "Description"}

type SubTask struct {
	Id string `yaml:"id"`
	Checked bool `yaml:"checked"`
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
	y, tasks := GetTasks()
	taskIndex := FindIndex(tasks, FieldId, taskId)
	subtasks := tasks[taskIndex].SubTasks
	st.Id = IntToString(GetNewId(subtasks))
	st.Created = GetCurrentTime()
	subtasks = append(subtasks, *st)
	tasks[taskIndex].SubTasks = subtasks
	y.Write(tasks)
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
	_, tasks := GetTasks()
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

func EditSubTask[F FieldType](taskId string, subTaskId string, fieldName string, fieldData F)error{
	fieldType := reflect.ValueOf(fieldData)
	y, tasks := GetTasks()

	taskIndex := FindIndex(tasks, FieldId, taskId)
	subtasks := tasks[taskIndex].SubTasks
	subTaskIndex := FindIndex(subtasks, FieldId, subTaskId)

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
	y.Write(tasks)
	return nil
}

func DeleteSubTask(taskId string, subTaskId string)error{
	y, tasks := GetTasks()
	taskIndex := FindIndex(tasks, FieldId, taskId)
	subtasks := tasks[taskIndex].SubTasks
	result := DeleteFromTasks(subtasks, FieldId, subTaskId)
	tasks[taskIndex].SubTasks = result
	y.Write(tasks)
	return nil
}