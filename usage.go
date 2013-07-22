package command

import (
	"bytes"
	"fmt"
	"strings"
)

var usage commands

func init() {
	usage = make(commands, 0)
}

type commands []*command

func (v commands) String() string {
	width1 := 0
	width2 := 0
	for _, val := range v {
		if len(val.name) > width1 {
			width1 = len(val.name)
		}
		if len(val.options) > width2 {
			width2 = len(val.options)
		}
	}

	buf := new(bytes.Buffer)
	for _, val := range v {
		val.width1 = width1
		val.width2 = width2
		buf.WriteString(fmt.Sprintf("\t%v\n", val))
	}

	return buf.String()
}

type command struct {
	name        string
	options     string
	description string
	width1      int
	width2      int
}

func (v *command) String() string {
	s1 := strings.Repeat(" ", v.width1-len(v.name))
	s2 := strings.Repeat(" ", v.width2-len(v.options))
	return fmt.Sprintf("%s:  %s%s  %s%s", v.name, s1, v.options, s2, v.description)
}

func Describe(name, options, description string) {
	usage = append(usage, &command{name: name, options: options, description: description})
}

func Usage() string {
	return usage.String()
}
