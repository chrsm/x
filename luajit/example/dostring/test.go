package main

import (
	"fmt"

	"github.com/chrsm/gluajit"
)

func main() {
	fmt.Println("testing..")

	l := gluajit.NewState()
	if err := l.DoString("print(\"hi\")\nos.exit(2)"); err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("bye!")
}
