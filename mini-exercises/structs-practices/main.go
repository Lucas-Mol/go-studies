package main

import (
	"bufio"
	"fmt"
	"github.com/Lucas-Mol/go-studies/mini-exercises/structs-practices/note"
	"github.com/Lucas-Mol/go-studies/mini-exercises/structs-practices/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todoText := getUserInput("To Do text: ")
	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputAndSaveData(userNote)
	if err != nil {
		return
	}
	err = outputAndSaveData(userTodo)
	if err != nil {
		return
	}
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")
	return value
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func outputAndSaveData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Data failed to save:", err)
		return err
	}

	fmt.Println("Saved data successfully")
	return nil
}
