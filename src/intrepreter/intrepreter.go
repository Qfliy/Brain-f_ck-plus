package intrepreter

import (
	"fmt"
)

type Interpreter struct {
	code           string
	memory         []byte
	index          int
	programCounter int
	addressLoops   []int
}

func NewInterpreter(input string) *Interpreter {
	var self = &Interpreter{}
	self.Load(input)
	return self
}

func (interpreter *Interpreter) Load(input string) {
	interpreter.code = input
	interpreter.memory = make([]byte, 30000)
	interpreter.programCounter = 0
	interpreter.index = 0
	interpreter.addressLoops = []int{}
}

func (interpreter *Interpreter) Run() {
	var commands = interpreter.getAllCommands()

	for interpreter.programCounter < len(interpreter.code) {
		var command rune = rune(interpreter.code[interpreter.programCounter])

		if procedure, ok := commands[command]; ok {
			procedure()
		}

		interpreter.programCounter++
	}
}

func (interpreter *Interpreter) getAllCommands() map[rune]func() {
	return map[rune]func(){
		'+': func() { interpreter.memory[interpreter.index]++ },
		'-': func() { interpreter.memory[interpreter.index]-- },
		'>': func() { interpreter.index++ },
		'<': func() { interpreter.index-- },
		'.': func() { fmt.Printf("%c", interpreter.memory[interpreter.index]) },
		',': func() { fmt.Scanf("%c", &interpreter.memory[interpreter.index]) },
		'[': interpreter.beginLoop,
		']': interpreter.endLoop,
	}
}

func (interpreter *Interpreter) beginLoop() {
	if interpreter.memory[interpreter.index] == 0 {
		var loopCount int = 1

		for loopCount > 0 {
			interpreter.programCounter++

			switch interpreter.code[interpreter.programCounter] {
			case '[':
				loopCount++
			case ']':
				loopCount--
			}
		}
	} else {
		interpreter.addressLoops = append(interpreter.addressLoops, interpreter.programCounter)
	}
}

func (interpreter *Interpreter) endLoop() {
	if interpreter.memory[interpreter.index] != 0 {
		interpreter.programCounter = interpreter.addressLoops[len(interpreter.addressLoops)-1]
	} else {
		interpreter.addressLoops = interpreter.addressLoops[:len(interpreter.addressLoops)-1]
	}
}
