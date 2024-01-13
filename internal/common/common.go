package common

import (
	"fmt"
	"io/fs"
	"os"

	"os/exec"
	"path/filepath"
	"github.com/mitchellh/mapstructure"
	"github.com/chzyer/readline"
	"gopkg.in/yaml.v2"
	"time"
	"strconv"
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

type Yaml struct {
	Path string
	RawData []byte
	Data interface{}
}

func (y *Yaml)Decode(i interface{})interface{}{
	config := mapstructure.DecoderConfig{
		TagName: "yaml",
	}
	config.Result = &i
	decoder, _ := mapstructure.NewDecoder(&config)
	decoder.Decode(y.Data)
	return i
}

func (y *Yaml)Read()error{
	data, err := FileRead(y.Path)
	if err != nil {
		return err
	}
	y.RawData = data
	var yamlData interface{}
	err = yaml.Unmarshal(y.RawData, &yamlData)
	if err != nil{
		return err
	}
	y.Data = yamlData
	return nil
}

func (y *Yaml)Write(content interface{})error{
	data, err := yaml.Marshal(content)
	if err != nil{
		return err
	}
	err = FileWrite(y.Path, data)
	if err != nil {
		return err
	}
	return nil
}

type Input struct {
	FieldName string
	Choices []string
	Data string
	Prompt string
	TmpPath string
	TextEditor string
	NewFolderPermissions fs.FileMode
}

func (i *Input)Row()string{
	prefix := fmt.Sprintf("%v %v ", i.FieldName, i.Prompt)
	input, err := readline.New(prefix)
	if err != nil{
		fmt.Println(err)
	}
	defer input.Close()
	sourceData := i.Data
	inputData := []byte(sourceData)
	input.WriteStdin(inputData)
	destinationData, err := input.Readline()
	if err != nil {
		fmt.Println(err)
	}
	return destinationData
}

func (i *Input)Text()string{
	path := filepath.Join(i.TmpPath, i.FieldName)
	err := os.MkdirAll(i.TmpPath, i.NewFolderPermissions)
	if err != nil{
		fmt.Println(err)
	}
	text := fmt.Sprintf("%v\n", i.Data)
	FileWrite(path, []byte(text))
	cmd := exec.Command(i.TextEditor, path)
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

func (i *Input)Choice()string{
	message := fmt.Sprintf("Choose value for %v", i.FieldName)
	fmt.Println(message)
	fmt.Println("Pick a digit")
	for index, choice := range i.Choices{
		message := fmt.Sprintf("Type %v for %v '%v'", index, i.FieldName, choice)
		fmt.Println(message)
	}
	var chosen int
	for {
		fmt.Scanln(&chosen)
		if (len(i.Choices)) <= chosen || chosen < 0 {
			fmt.Printf("Value must be from 0 to %d\n", len(i.Choices) - 1)
		} else {
			break
		}
	}
	return i.Choices[chosen]
}

func GetCurrentTime(timeFormat string)string{
	timeNow := time.Now()
	return timeNow.Format(timeFormat)
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

func FindIndex(i []string, key string) int{
	for index, value := range i{
		if key == value {
			return index
		}
	}
	return -1
}

func IsInSlice(name string, s []string)bool{
	for _, el := range s{
		if el == name{
			return true
		}
	}
	return false
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