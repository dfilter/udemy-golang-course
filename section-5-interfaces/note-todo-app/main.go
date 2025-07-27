package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/structs-practice/note"
	"example.com/structs-practice/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputable interface {
	saver
	Display()
}

// type outputable interface {
// 	Save() error
// 	Display()
// }

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(note)
}

// "any" is an alias for interface{}
// func printSomething(value interface{}) {
// Type of any value can be narrowed by using ".(type)"
// switch value.(type) {
// case int:
// 	fmt.Println("Integer: ", value)
// case float32:
// case float64:
// 	fmt.Println("Float: ", value)
// case string:
// 	fmt.Println("String: ", value)
// }
// Alternitivly types can be checked like this:
// typedVal, ok := value.(float64)
// if ok is true the value is of the provided type false and an empty value if false
// }

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving failed.")
		return err
	}
	fmt.Println("Saved successfully!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	// when specifying single bytes one must use single quotes '
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
