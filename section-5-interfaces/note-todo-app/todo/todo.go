package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const fileName = "todo.json"

// The json pachage will only read fields that are public
// Because of this by default it will make the json keys upper case
// If we don't want that behavior we can use metadata tags, defining how these fields should be treated.
type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	jsonData, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}

func New(text string) (*Todo, error) {
	if text == "" {
		return nil, errors.New("invalid input: input cannot be empty")
	}

	return &Todo{
		Text: text,
	}, nil
}
