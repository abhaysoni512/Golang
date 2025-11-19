package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/notes"
)

func getUserInput(prompt string)(string){
	fmt.Print(prompt)
    reader :=	bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil{
		fmt.Println("Error reading input :", err)
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func getNoteData()(string,string){
	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")
	return title, content
}


func main() {
	title, content := getNoteData()

	userNote , err := notes.NewNote(title, content)
	if err != nil{
		fmt.Println("Error creating note :", err)
		return
	}
	fmt.Println("Note created successfully with details : \n",userNote.DisplayNote())

	err = userNote.Save()
	if err != nil{
		fmt.Println("Error saving note :", err)
		return
	}
	fmt.Println("Note saved successfully")
	




}


