package pra

import "fmt"

var globalVariable int = 30

func Gvar(g int ) (int){
	fmt.Println("******* start gvar *******")
	fmt.Println("bafore assigning a value to globalVariable =", globalVariable) 

	globalVariable = g

	fmt.Println("after assigning a value to globalVariable =", globalVariable) 
	fmt.Println("****** end gvar ******")
	return globalVariable
}
