package notes

import (
	"errors"
	"os"
	"time"
	"strings"
	"encoding/json"
)

type Note struct {
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func NewNote(title string, content string) ( *Note, error){
	if title == ""{
		return nil, errors.New("title cannot be empty")
	}
	return &Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil

}

func (n Note) DisplayNote () string{
	return  " Title: " + n.Title + " Content: " + n.Content + " Created At: " + n.CreatedAt.String()
}

func (n Note) Save() error{
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName)+".json"
    json, err := json.Marshal(n)
	if err != nil{
		return errors.New("could not convert note to JSON")
	}
	return os.WriteFile(fileName, json, 0644)

}