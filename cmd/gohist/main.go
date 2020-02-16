package main

import (
"fmt"
"github.com/jvzantvoort/history"
)

func main() {
	fmt.Println("vim-go")
	hist := history.NewHistory("foo", ".")
	hist.Add("update")
}
