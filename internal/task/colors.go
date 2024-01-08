package task

import (
	"github.com/olekukonko/tablewriter"
)

var PriorityColors = map[string]int{
	PriorityUrgent: tablewriter.FgHiMagentaColor,
	PriorityHigh: tablewriter.FgHiRedColor,
	PriorityMedium: tablewriter.FgHiYellowColor,
	PriorityLow: tablewriter.FgHiCyanColor,
}

var StatusColors = map[string]int{
	StatusOpen: tablewriter.FgMagentaColor,
	StatusNew: tablewriter.Normal,
	StatusPause: tablewriter.FgYellowColor,
	StatusClosed: tablewriter.FgGreenColor,
}
var (
	BoldText int = tablewriter.Bold
	NormalColor int = tablewriter.Normal
	UrlColor int = tablewriter.FgHiBlueColor
	TaskListColorType string = "task_list"
	TaskColorType string = "task"
)


type ColorFilter struct {
	Tasks map[string]string
	Headers []string
	RowColors []tablewriter.Colors
	Type string
}

type Stage func()

func (c *ColorFilter)FillNormal(){
	for i := 0; i < len(c.Headers); i++{
		c.RowColors = append(c.RowColors, tablewriter.Colors{})
	}
}

func (c *ColorFilter) Process(){
	c.FillNormal()
	var stages []Stage
	switch c.Type{
	case TaskListColorType:
		stages = []Stage{c.TaskListPriority, c.TaskListStatus}
	case TaskColorType:
		stages = []Stage{c.TaskPriority, c.TaskStatus}
	}
	for _, stage := range stages{
		stage()
	}
}


func (c *ColorFilter)TaskListPriority(){
	fieldName := FieldPriority
	index := FindIndex(c.Headers, fieldName)
	fieldValue := c.Tasks[fieldName]
	colorValue := PriorityColors[fieldValue]
	colorValues := c.RowColors[index]
	if len(colorValues) == 0 {
		c.RowColors[index] = tablewriter.Colors{BoldText, colorValue}
	}
}

func (c *ColorFilter)TaskPriority(){
	fieldName := FieldPriority
	if IsInSlice(fieldName, c.Headers){
		index := FindIndex(c.Headers, fieldName)
		fieldValue := c.Tasks[fieldName]
		colorValue := PriorityColors[fieldValue]
		colorValues := c.RowColors[index]
		if len(colorValues) == 0 {
			c.RowColors[index] = tablewriter.Colors{BoldText, NormalColor}
			c.RowColors = append(c.RowColors, tablewriter.Colors{BoldText, colorValue})
		}
	}
}

func (c *ColorFilter)TaskListStatus(){
	fieldName := FieldStatus
	fieldValue := c.Tasks[fieldName]
	colorValue, exists := StatusColors[fieldValue]
	for index, cell := range c.RowColors{
		if len(cell) == 0 && exists{
			c.RowColors[index]= tablewriter.Colors{BoldText, colorValue}
		}
	}
}

func (c *ColorFilter)TaskStatus(){
	fieldName := FieldStatus
	if IsInSlice(fieldName, c.Headers){
	fieldValue := c.Tasks[fieldName]
	colorValue, exists := StatusColors[fieldValue]
	for index, cell := range c.RowColors{
		if len(cell) == 0 && exists{
			c.RowColors[index]= tablewriter.Colors{BoldText, NormalColor}
			c.RowColors = append(c.RowColors, tablewriter.Colors{BoldText, colorValue})
		}
	}
	}
}
