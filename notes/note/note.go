package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Your note title: %v has the following content: %v\n\n\n\n", note.title, note.content)
}
