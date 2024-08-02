package main

import (
	"strings"
)

func echo(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
}

func echoWithStringJoin(args []string) {
	strings.Join(args[1:], " ")
}
