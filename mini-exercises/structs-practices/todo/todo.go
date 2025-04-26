package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{
		Text: content,
	}, nil
}

func (t Todo) Display() {
	fmt.Printf(t.Text)
}

func (t Todo) Save() error {
	filename := "todo.json"

	jsonData, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonData, 0644)
}
