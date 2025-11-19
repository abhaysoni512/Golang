package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"examples/35_interfaces/note"
	todos "examples/35_interfaces/todo"
)

// Interface is a contract that certain values( typically structs) can choose to fulfill by implementing specific methods.
// whatever struct that implements those methods is said to satisfy that interface
type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getTodoData()
	todo, err := todos.NewNote(todoText)
	if err != nil {
		fmt.Println("Error creating todo :", err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	userNote, err := note.NewNote(title, content)
	if err != nil {
		fmt.Println("Error creating note :", err)
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}

}

// func printSomething(value any) {
// 	switch value.(type) {
// 	case string:
// 		fmt.Println("It's a string :", value)
// 	case int:
// 		fmt.Println("It's an integer :", value)
// 	default:
// 		fmt.Println("Unknown type")
// 	}
// }

// interface function that takes any outputtable data type
// func add (a,b interface{}){
// 	aInt, aIsint := a.(int)
// 	bInt, bIsint := b.(int)
	
// 	if aIsint && bIsint{
// 		fmt.Println("Sum is :", aInt + bInt)
// 	}
// }

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}
func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Error saving data :")
		return err

	}
	fmt.Println("Data saved successfully")
	return nil
}
func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input :", err)
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func getTodoData() string {
	text := getUserInput("Todo text: ")
	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")
	return title, content
}
