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
	width := 0
	for _, val := range v {
		if len(val.name) > width {
			width = len(val.name)
		}
	}

	buf := new(bytes.Buffer)
	for _, val := range v {
		if width < len(val.name) {
			panic(width)
		}
		val.width = width
		buf.WriteString(fmt.Sprintf("\t%v\n", val))
	}

	return buf.String()
}

type command struct {
	name        string
	description string
	width       int
}

func (v *command) String() string {
	if v.width < len(v.name) {
		panic(fmt.Sprint(v.width, len(v.name)))
	}
	return fmt.Sprintf("%s: %s%s", v.name, strings.Repeat(" ", v.width-len(v.name)), v.description)
}

func Describe(name, description string) {
	usage = append(usage, &command{name: name, description: description})
}

func Usage() string {
	return usage.String()
}
