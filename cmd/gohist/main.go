package main

import (
	"fmt"
	"github.com/jvzantvoort/history"
	"os"
	"strings"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) < 1 {
		fmt.Println("gohist <app> line ...")
		os.Exit(1)
	}
	appname, arguments := arguments[0], arguments[1:]
	hist := history.NewHistory(appname, ".")
	hist.Add(strings.Join(arguments, " "))
}
