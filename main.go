package main

import (
	inter "go_reloaded/internal"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		panic("Error: invalid number of arguments")
	}
	input := inter.Read("txt/" + os.Args[1])
	result := inter.CorrectAll(input)
	inter.Write(result, "txt/"+os.Args[2])
}
