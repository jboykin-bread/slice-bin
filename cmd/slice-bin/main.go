package main

import (
	"fmt"
)

var Version = "development"
var GitCommit string

func main() {
	fmt.Println("Version:", Version)
	fmt.Println("Git Commit:", GitCommit)
	fmt.Println("hello there i'm slice-bin! 💅")
}
