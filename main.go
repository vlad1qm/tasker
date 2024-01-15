package main

import (
	"tasker/internal/cmd"
	t "tasker/internal/task"
)

func main(){
	t.Config.Init()
	cmd.Execute()
}