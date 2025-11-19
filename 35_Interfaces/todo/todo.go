package todos

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
 Text   string    `json:"text"`
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	
	json, err := json.Marshal(todo)
	if err != nil {
		return errors.New("could not convert todo to JSON")
	}
	return os.WriteFile(fileName, json, 0644)

}

func NewNote(text string) (*Todo, error) {
	if text == "" {
		return nil, errors.New("title cannot be empty")
	}
	return &Todo{
		Text:     text,	
	}, nil

}

func (t Todo) Display()  {
	fmt.Println(t.Text)
}


