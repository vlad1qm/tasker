package task

import (
	"gopkg.in/yaml.v2"
	"strconv"
	"fmt"
	"os"
	"io/fs"
	"time"
	"github.com/chzyer/readline"
)

const (
	NoneFilePermissions fs.FileMode = 0000
	DefaultFilePermissions fs.FileMode = 0666
	DefaultFilePath string = "data.yaml"
	DefaultTimeFormat string = "15:04:05 02.01.2006"
)
var FilePath string = DefaultFilePath

type TaskType interface {
	Task | SubTask
	GetId() string
	GetValueOf(string) string
}
type FieldType interface {
	string | bool
}

type Yaml struct {
	Path string
	Permission fs.FileMode
	Data []byte
}

func (y *Yaml)Read()error{
	data, err := os.ReadFile(y.Path)
	if err != nil {
		return err
	}
	y.Data = data
	return nil
}

func (y *Yaml)Write(content []Task)error{
	if y.Permission == NoneFilePermissions {
		y.Permission = DefaultFilePermissions
	}
	tasks := map[string][]Task{"tasks": content}
	data, err := yaml.Marshal(tasks)
	if err != nil{
		return err
	}
	err = os.WriteFile(y.Path, data, y.Permission)
	if err != nil {
		return err
	}
	return nil
}

func (y *Yaml)GetTasks()[]Task{
	var tasks map[string][]Task
	err := y.Read()
	if err != nil{
		fmt.Println(err)
		return []Task{}
	}
	err = yaml.Unmarshal(y.Data, &tasks)
	if err != nil {
		fmt.Println(err)
		return []Task{}
	}
	return tasks["tasks"]
}

func StringToInt(s string)int{
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println()
	}
	return i
}

func IntToString(i int)string{
	return strconv.Itoa(i)
}

func GetCurrentTime()string{
	timeNow := time.Now()
	return timeNow.Format(DefaultTimeFormat)
}

func GetNewId[T TaskType](t []T)int{
	var newId int
	if len(t) == 0{
		return 1
	}
	for _, task := range t{
		taskId := StringToInt(task.GetId())
		if taskId > newId{
			newId = taskId
		}
	}
	return newId + 1
}

func FindIndex[T TaskType](t []T, fieldName string, key string) int{
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

func RowInput(fieldName string, text string)string{
	prefix := fmt.Sprintf("%v >>> ", fieldName)
	input, _ := readline.New(prefix)
	defer input.Close()
	datax := text
	data2 := []byte(datax)
	input.WriteStdin(data2)

	value, _ := input.Readline()
	return value
}