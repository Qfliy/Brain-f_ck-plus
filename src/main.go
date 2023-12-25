package main

import (
	"bf+/info"
	"bf+/intrepreter"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var i = intrepreter.NewInterpreter("")
var args = os.Args

func ranCode(program string) {
	i.Load(program)
	i.Run()
}

func readAndRanCode(fileName string) {
	var code, err = os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: \n\t the file was not found")
		return
	}

	ranCode(string(code))
}

var commands = map[string]func(param string){
	"str": func(param string) {
		ranCode(param)
		os.Exit(0)
	},
	"run": func(param string) {
		readAndRanCode(param)
		os.Exit(0)
	},
}

func ranShell() {
	reader := bufio.NewReader(os.Stdin)
	var code string

	fmt.Println(`Use "&" for exit`)

	for {
		fmt.Print("?> ")
		code, _ = reader.ReadString('\n')

		if strings.TrimSpace(code) == "&" {
			break
		}

		ranCode(code)
		fmt.Print("\n")
	}

	os.Exit(0)
}

func main() {
	switch len(args) {
	case 1:
		info.Info("Info\n")
	case 2:
		if args[1] == "shell" {
			ranShell()
		}

		fallthrough
	case 3:
		if procedure, ok := commands[args[1]]; ok {
			procedure(args[2])
		}

		fallthrough
	default:
		info.Info("Eroor command\n")
	}
}
