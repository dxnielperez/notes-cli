import json
import os

NOTES_FILE = "notes.json"

def load_notes():
    if os.path.exists(NOTES_FILE):
        with open(NOTES_FILE, "r") as file:
            return json.load(file)
    return []

def save_notes(notes):
    with open(NOTES_FILE, "w") as file:
        json.dump(notes, file, indent=4)

def add_note():
    title = input("Enter note title: ")
    content = input("Enter note content: ")
    notes = load_notes()
    notes.append({"title": title, "content": content})
    save_notes(notes)
    print("Note added successfully!\n")

def view_notes():
    notes = load_notes()
    if not notes:
        print("No notes found.\n")
    else:
        for idx, note in enumerate(notes, start=1):
            print(f"{idx}. {note['title']}\n   {note['content']}\n")

def update_note():
    notes = load_notes()
    view_notes()
    if not notes:
        return
    
    try:
        index = int(input("Enter note number to update: ")) - 1
        if 0 <= index < len(notes):
            notes[index]["title"] = input("Enter new title: ")
            notes[index]["content"] = input("Enter new content: ")
            save_notes(notes)
            print("Note updated successfully!\n")
        else:
            print("Invalid note number.\n")
    except ValueError:
        print("Invalid input. Please enter a number.\n")

def delete_note():
    notes = load_notes()
    view_notes()
    if not notes:
        return
    
    try:
        index = int(input("Enter note number to delete: ")) - 1
        if 0 <= index < len(notes):
            del notes[index]
            save_notes(notes)
            print("Note deleted successfully!\n")
        else:
            print("Invalid note number.\n")
    except ValueError:
        print("Invalid input. Please enter a number.\n")

def main():
    while True:
        print("\nNotes CLI - Choose an option:")
        print("1. Add Note")
        print("2. View Notes")
        print("3. Update Note")
        print("4. Delete Note")
        print("5. Exit")
        
        choice = input("Enter your choice: ")
        if choice == "1":
            add_note()
        elif choice == "2":
            view_notes()
        elif choice == "3":
            update_note()
        elif choice == "4":
            delete_note()
        elif choice == "5":
            print("Exiting...\n")
            break
        else:
            print("Invalid choice. Please try again.\n")

if __name__ == "__main__":
    main()
