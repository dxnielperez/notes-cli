package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func loadNotes(filename string) ([]Note, error) {
	var notes []Note
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Note{}, nil
		}
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	if len(bytes) == 0 {
		return []Note{}, nil
	}
	err = json.Unmarshal(bytes, &notes)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func saveNotes(filename string, notes []Note) error {
	data, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func createNote(filename, title, content string) error {
	notes, err := loadNotes(filename)
	if err != nil {
		return err
	}
	var id int
	if len(notes) > 0 {
		id = notes[len(notes)-1].ID + 1
	} else {
		id = 1
	}
	newNote := Note{ID: id, Title: title, Content: content}
	notes = append(notes, newNote)
	if err := saveNotes(filename, notes); err != nil {
		return err
	}
	fmt.Printf("Note created with ID %d\n", newNote.ID)
	return nil
}

func listNotes(filename string) error {
	notes, err := loadNotes(filename)
	if err != nil {
		return err
	}
	if len(notes) == 0 {
		fmt.Println("No notes found.")
		return nil
	}
	for _, note := range notes {
		fmt.Printf("ID: %d\nTitle: %s\nContent: %s\n\n", note.ID, note.Title, note.Content)
	}
	return nil
}

func updateNote(filename string, id int, title, content string) error {
	notes, err := loadNotes(filename)
	if err != nil {
		return err
	}
	updated := false
	for i, note := range notes {
		if note.ID == id {
			notes[i].Title = title
			notes[i].Content = content
			updated = true
			break
		}
	}
	if !updated {
		return errors.New("note not found")
	}
	if err := saveNotes(filename, notes); err != nil {
		return err
	}
	fmt.Println("Note updated.")
	return nil
}

func deleteNote(filename string, id int) error {
	notes, err := loadNotes(filename)
	if err != nil {
		return err
	}
	index := -1
	for i, note := range notes {
		if note.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("note not found")
	}
	notes = append(notes[:index], notes[index+1:]...)
	if err := saveNotes(filename, notes); err != nil {
		return err
	}
	fmt.Println("Note deleted.")
	return nil
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  add <title> <content>         Create a new note")
	fmt.Println("  list                          List all notes")
	fmt.Println("  update <id> <title> <content> Update an existing note")
	fmt.Println("  delete <id>                   Delete a note")
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}
	filename := "notes.json"
	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Usage: add <title> <content>")
			return
		}
		title := os.Args[2]
		content := os.Args[3]
		if err := createNote(filename, title, content); err != nil {
			fmt.Println("Error:", err)
		}
	case "list":
		if err := listNotes(filename); err != nil {
			fmt.Println("Error:", err)
		}
	case "update":
		if len(os.Args) < 5 {
			fmt.Println("Usage: update <id> <title> <content>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID:", os.Args[2])
			return
		}
		title := os.Args[3]
		content := os.Args[4]
		if err := updateNote(filename, id, title, content); err != nil {
			fmt.Println("Error:", err)
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: delete <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID:", os.Args[2])
			return
		}
		if err := deleteNote(filename, id); err != nil {
			fmt.Println("Error:", err)
		}
	default:
		printUsage()
	}
}
