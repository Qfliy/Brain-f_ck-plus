package main

import (
	"bf+/info"
	"bf+/intrepreter"
	"fmt"
	"os"
)

func ranCode(program string) {
	var interpreter = intrepreter.NewInterpreter(program)
	interpreter.Run()
}

func readAndRanCode(fileName string) {
	var code, err = os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: \n\t the file was not found")
		return
	}

	ranCode(string(code))
}

func main() {
	args := os.Args

	switch len(args) {
	case 1:
		info.Info("Info\n")
	case 3:
		switch args[1] {
		case "str":
			ranCode(args[2])
		case "run":
			readAndRanCode(args[2])
		default:
			info.Info("Eroor command\n")
		}
	default:
		info.Info("Eroor command\n")
	}
}
