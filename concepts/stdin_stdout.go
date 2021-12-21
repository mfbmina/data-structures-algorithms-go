package concepts

import "fmt"

// Sometimes we need to get or send information from the users.
// In basic programs like ours, we do that by reading information from the terminal and printing it back there.
// This is what we call "reading from the `stdin`" and "printing to the `stdout`"
func StdinAndStdout() {
	var name string

	// We are printing a question to the stdout
	fmt.Println("What is your name?")

	// We are reading the awnser from the user and storing it in a variable called name.
	fmt.Scanln(&name)

	fmt.Println("Hello", name)
}
