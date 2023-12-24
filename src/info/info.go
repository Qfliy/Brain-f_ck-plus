package info

import "fmt"

const command string = `
bf+ - intrepreter brain f*ck

version 2.0

arguments:
	info:
		bf+
	run program:
		bf+ run "path\to\fail.bf"
	ran string:
		bf+ str ">++++++++[<+++++++++>-]<."

comands:
	. - putChar
	, - getChar
	+ - increment address
	- - decrement address
	> - increment pointer
	< - decrement pointer
	[ - start loop
	] - end loop

`

func Info(prePrint string) {
	fmt.Print(prePrint, command)
}
