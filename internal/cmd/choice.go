package cmd

import (
	"strings"
	"fmt"
	// "github.com/wsxiaoys/terminal/color"
)

// https://github.com/spf13/pflag/issues/236#issuecomment-931600452

type choiсe struct {
	Allowed []string
	Value string
}

func newChoice(allowed []string, value string) *choiсe {
	return &choiсe{
		Allowed: allowed,
		Value:   value,
	}
}

func (c choiсe) String() string {
	return c.Value
}

func (c *choiсe) Set(currentValue string) error {
	isIncluded := func(opts []string, val string) bool {
		for _, opt := range opts {
			if val == opt {
				return true
			}
		}
		return false
	}
	if !isIncluded(c.Allowed, currentValue) {
		return fmt.Errorf("%s is not included in %s", currentValue, strings.Join(c.Allowed, ","))
	}
	c.Value = currentValue
	return nil
}

func (c *choiсe) Type() string {
	return "string"
}