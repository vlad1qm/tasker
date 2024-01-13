package main

import (
	"tasker/internal/cmd"
	c "tasker/internal/config"
)

func main(){
	c.Configuration = c.Config{}
	c.Configuration.Init()
	cmd.Execute()
}