package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// The json pachage will only read fields that are public
// Because of this by default it will make the json keys upper case
// If we don't want that behavior we can use metadata tags, defining how these fields should be treated.
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content: \n\n%v\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	jsonData, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("invalid input: input cannot be empty")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
