package config

import (
	"errors"
	"tasker/internal/common"
)

var Configuration Config

const (
	ConfigPath string = "config.yaml"
	DefaultFilePath string = "data.yaml"
	DefaultTmpPath string = "/tmp/tasks/"
	DefaultTextEditor string = "vim"
	DefaultTimeFormat string = "15:04:05 02.01.2006"
	DefaultRowPrompt string = ">>>"
)

type Config struct {
	DataFile string `yaml:"data_file"`
	TmpPath string `yaml:"tmp_path"`
	TextEditor string `yaml:"text_editor"`
	TimeFormat string `yaml:"time_format"`
	RowPrompt string `yaml:"row_prompt"`
}

func (c *Config)Init()error{
	y := common.Yaml{Path: ConfigPath}
	y.Read()
	decoded := y.Decode(Config{})
	config, ok := decoded.(Config)
	if ok{
		*c = config
		return nil
	}
	return errors.New("couldnt init config file")
}

