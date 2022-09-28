// Initialise package called greetings by declaring a greetings package to collect related functions 

package greetings

import "fmt"

// Hello returns a greeting for the named person
func Hello (name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

/* 
As above, Function name is func "Hello", 
the Parameter type is (name "string"),
the Return type "string"

In Golang, the ":=" operator is a shortcut for declaring and initialisinga variable
in one line (Go uses the value on the right to determine the variable's type).
Ref: [https://go.dev/doc/tutorial/create-module]

Using fmt package's Sprintf function to create a greeting message. the First argument
is a format string, and the "Sprintf" substitutes the "name" parameter's value for the
%v format verb. Inserting the value of the "name" parameter completes the greeting text.

*/