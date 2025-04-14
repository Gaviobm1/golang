package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

type Saver interface {
	Save() error
}

type Displayer interface {
	Display()
}

type Outputtable interface {
	Saver
	Displayer
}

func main() {
	title, content := getNoteData()
	todoText := getTodoData()

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = outputData(todo)

	if err != nil {
		return
	}

	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	outputData(note)
}

func outputData(data Outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data Saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Save failed")
		return err
	}
	fmt.Println("Save succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func getTodoData() string {
	text := getUserInput("Todo: ")
	return text
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
