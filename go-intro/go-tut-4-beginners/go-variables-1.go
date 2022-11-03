/* Variables - Introduction

Variables can be displayed in multiple ways with GoLang, per varied variables, see below:

	
	var i int (written "variables" of the value "i" as an type "int" Integer)

	once it's set, we can set the i as a value with (i = 42)

	OR
	
	using the same but shorter, we can also use:

	var i int = 42

	or 
	
	i := 42


Naming Variables come with two key #1 Simple naming for readability and 
#2 is the what the convention is for. [Short and concise]
(ie. )
*/ 

package main

import (
	"fmt"
)

func main() {
	var i int
	i = 42

	var j int = 41

	k := 40

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}