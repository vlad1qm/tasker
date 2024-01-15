package config

import (
	"errors"
	"tasker/internal/common"
)

var Configuration Config

const (
	ConfigPath string =  "/Users/v.minkin/.config/tasker/"
	ConfigName string = "config.yaml"
	DefaultFilePath string = "data.yaml"
	DefaultTmpPath string = "/tmp/tasks/"
	DefaultTextEditor string = "vim"
	DefaultTimeFormat string = "15:04:05 02.01.2006"
	DefaultRowPrompt string = ">>>"
)

type Config struct {
	DataPath string `yaml:"data_file_path"`
	DataFileName string `yaml:"data_file_name"`
	TmpPath string `yaml:"tmp_path"`
	TextEditor string `yaml:"text_editor"`
	TimeFormat string `yaml:"time_format"`
	RowPrompt string `yaml:"row_prompt"`
	TableTheme string `yaml:"table_theme"`
	TaskTableSortBy string `yaml:"table_sort_by"`
	TaskTableSortDirection string `yaml:"table_sort_direction"`
	ColumnNumberMinWidth int `yaml:"column_number_min_width"`
	ColumnMinWidth int `yaml:"column_min_width"`
	ColumnTaskListFilter []string `yaml:"column_task_list_filter"`
	RowTaskListFilter map[string]string `yaml:"row_task_list_filter"`
}

func (c *Config)Init()error{
	y := common.Yaml{FilePath: ConfigPath, FileName: ConfigName}
	y.Read()
	decoded := y.Decode(Config{})
	config, ok := decoded.(Config)
	if ok{
		*c = config
		return nil
	}
	return errors.New("couldnt init config file")
}

