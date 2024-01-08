package task

import (
	"gopkg.in/yaml.v2"
	"strconv"
	"fmt"
	"os"
	"io/fs"
	"time"
	"github.com/chzyer/readline"
	"path/filepath"
	"os/exec"
)

const (
	NoneFilePermissions fs.FileMode = 0000
	DefaultFilePermissions fs.FileMode = 0666
	DefaultFileWritePermissions fs.FileMode = 0755
	DefaultNewFolderPermissions fs.FileMode = 0777

	DefaultFilePath string = "data.yaml"
	DefaultTmpPath string = "/tmp/tasks/"
	DefaultTextEditor string = "vim"
	DefaultTimeFormat string = "15:04:05 02.01.2006"
	DefaultRowPrompt string = ">>>"
)
var (
	FileWritePermissiong fs.FileMode
	FilePermissions fs.FileMode = DefaultFilePermissions
	FileWritePermissions fs.FileMode = DefaultFileWritePermissions
	NewFolderPermissions fs.FileMode = DefaultNewFolderPermissions

	FilePath string = DefaultFilePath
	TmpPAth string = DefaultTmpPath
	TextEditor string = DefaultTextEditor
	TimeFormat string = DefaultTimeFormat
	RowPrompt string = DefaultRowPrompt
)

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

func FileRead(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}

func FileWrite(path string, data []byte) error {
	err := os.WriteFile(path, data, FileWritePermissions)
	if err != nil {
		fmt.Println(err)
	}
	return err
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
		y.Permission = FilePermissions
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
	return timeNow.Format(TimeFormat)
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

func FindTaskIndex[T TaskType](t []T, fieldName string, key string) int{
	for index, value := range t{
		if key == value.GetValueOf(fieldName) {
			return index
		}
	}
	return -1
}

func FindIndex(i []string, key string) int{
	for index, value := range i{
		if key == value {
			return index
		}
	}
	return -1
}

func MakeInterfaceSlice(s []string) []interface{} {
	new := make([]interface{}, len(s))
	for index, value := range s{
		new[index] = value
	}
	return new
}

func MakeMap(headers []string, row []string) map[string]string{
	m := map[string]string{}
	count := len(headers)
	for i := 0; i < count; i++{
		m[headers[i]] = row[i]
	}
	return m
}

func IsInSlice(name string, s []string)bool{
	for _, el := range s{
		if el == name{
			return true
		}
	}
	return false
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

func DeleteFromSliceByIndex(s []string, indexToDelete int) []string{
	newSlice := []string{}
	for index, value := range s{
		if indexToDelete != index{
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}

func RowInput(fieldName string, text string)string{
	prefix := fmt.Sprintf("%v %v ", fieldName, RowPrompt)
	input, _ := readline.New(prefix)
	defer input.Close()
	datax := text
	data2 := []byte(datax)
	input.WriteStdin(data2)

	value, _ := input.Readline()
	return value
}

func TextInput(fieldName string, text string)string{
	folderPath := TmpPAth
	path := filepath.Join(folderPath, fieldName)
	err := os.MkdirAll(folderPath, NewFolderPermissions)
	if err != nil{
		fmt.Println(err)
	}
	text = fmt.Sprintf("%v\n", text)
	FileWrite(path, []byte(text))
	cmd := exec.Command(TextEditor, path)
	cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    err = cmd.Run()
	if err != nil{
		fmt.Println(err)
	}
	data, err := FileRead(path)
	sData := string(data)
	last := len(sData) - 1
	sData = sData[:last]
	if err != nil{
		fmt.Println(err)
	}
	os.Remove(path)
	return string(sData)
}

func ChoiceInput(fieldName string, choices []string)string{
	message := fmt.Sprintf("Choose value for %v", fieldName)
	fmt.Println(message)
	fmt.Println("Pick a digit")
	for index, choice := range choices{
		message := fmt.Sprintf("Type %v for %v '%v'", index, fieldName, choice)
		fmt.Println(message)
	}
	var chosen int
	for {
		fmt.Scanln(&chosen)
		if (len(choices)) <= chosen || chosen < 0 {
			fmt.Printf("Value must be from 0 to %d\n", len(choices) - 1)
		} else {
			break
		}
	}
	return choices[chosen]
}