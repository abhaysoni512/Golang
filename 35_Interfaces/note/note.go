package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func NewNote(title string, content string) (*Note, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil

}

func (n Note) Display()  {
	fmt.Println("Title:", n.Title)
	fmt.Println("Content:", n.Content)
	fmt.Println("Created At:", n.CreatedAt.String())

}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(n)
	if err != nil {
		return errors.New("could not convert note to JSON")
	}
	return os.WriteFile(fileName, json, 0644)

}
