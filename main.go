package main

import (
	// "fmt"
	// "tasker/internal/task"
	// "reflect"
	"tasker/internal/cmd"
)

func main(){
	// y := task.Yaml{Path: "data.yaml"}
	// tasks := y.GetTasks()
	// for taskIndex := range tasks{
	// 	task := &tasks[taskIndex]
	// 	elements := reflect.ValueOf(task).Elem()
	// 	taskId := elements.FieldByName("Id")
	// 	for subTaskIndex := range tasks[taskIndex].SubTasks{
	// 		subTask := &tasks[taskIndex].SubTasks[subTaskIndex]
	// 		subTaskElements := reflect.ValueOf(subTask).Elem()
	// 		subTaskTitle := subTaskElements.FieldByName("Title")
	// 		subTaskTitle.SetString("Test SubTask Title")
	// 	}
	// 	if taskId.String() == "2"{
	// 		description := elements.FieldByName("Description")
	// 		description.SetString("Test Description")
	// 	}
	// }
	// y.Write(tasks)
	// t := task.GetTask("2", tasks)
	// fmt.Println(t)
	// st := task.GetSubTask("2", "2", tasks)
	// fmt.Println(st)
	// t := y.GetTasks()
	// fmt.Println(t)
	// t := task.Task{
	// 	Title: "Заголовок задачи",
	// 	Url: "https:/ya.ru",
	// 	Description: "Описание задачи",
	// 	Note: "Заметка",
	// }
	// t.Add()
	// st := task.SubTask{
	// 	Title: "Заголовок подзадачи",
	// 	Description: "Описание подзадачи",
	// 	Note: "Заметка",
	// }
	// st.Add("2")
	// err := task.EditSubTask("2", "1", "Title", "Отредактированный заголовок подзадачи")
	// if err != nil{
	// 	fmt.Println(err)
	// }
	cmd.Execute()
}