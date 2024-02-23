package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/go-project/note"
	"example.com/go-project/todo"
)

type Saver interface {
	Save() error
}

type Displayer interface {
	Display()
}

// Embadable interface
type Outputtable interface {
	Saver
	Displayer
}

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)

	if err != nil {
		return
	}

	todoText := getUserInput("Todo text:")

	userTodo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(userTodo)

}

func outputData(data Outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data Saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}

	fmt.Println("Saving the note succeded")

	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""

	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
