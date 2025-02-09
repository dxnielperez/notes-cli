const { addNote, listNotes, deleteNote, updateNote } = require("./utils");

const [, , command, ...args] = process.argv;

switch (command) {
  case "add": {
    const noteText = args.join(" ");
    if (!noteText.trim()) {
      console.log("Error: Note text cannot be empty.");
      break;
    }
    addNote(noteText);
    break;
  }
  case "list": {
    listNotes();
    break;
  }
  case "update": {
    const id = args[0];
    const newNote = args.slice(1).join(" ");
    if (!id || !newNote) {
      console.log("Error: Please provide both the note ID and new text.");
      break;
    }
    updateNote(id, newNote);
    break;
  }
  case "delete": {
    const id = args;
    if (!id) {
      console.log("Error: Please provide the note ID to delete.");
      break;
    }
    deleteNote(id);
    break;
  }
  default:
    console.log(`
    Usage:
      node index.js add "Note content"     - Add a new note
      node index.js list                   - List all notes
      node index.js delete <id>            - Delete a note by ID
      node index.js update <id> "new text" - Update a note by ID
    `);
}
