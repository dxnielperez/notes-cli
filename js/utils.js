const fs = require("fs");

function loadNotes() {
  const data = fs.readFileSync("notes.json", "utf8");
  return JSON.parse(data);
}

function saveNotes(notes) {
  fs.writeFileSync("notes.json", JSON.stringify(notes, null, 2));
}

function generateUniqueId(existingIds) {
  let id;
  do {
    id = Math.floor(1000 + Math.random() * 8000);
  } while (existingIds.includes(id));
  return id;
}

function addNote(text) {
  const notes = loadNotes();
  const existingIds = notes.map((note) => note.id);
  const uniqueId = generateUniqueId(existingIds);
  const newNote = { id: uniqueId, text };
  notes.push(newNote);
  saveNotes(notes);
  console.log("Note added:", newNote);
}

function listNotes() {
  const notes = loadNotes();
  if (notes.length === 0) {
    console.log("No notes found.");
    return;
  } else {
    notes.forEach((note) => console.log(`${note.id}: ${note.text}`));
  }
}

function deleteNote(id) {
  const notes = loadNotes();
  const updatedNotes = notes.filter((note) => note.id != id);
  if (notes.length === updatedNotes.length) {
    console.log("Note not found");
  } else {
    console.log("Note deleted successfully.");
    saveNotes(updatedNotes);
  }
}

function updateNote(id, newNote) {
  const notes = loadNotes();
  const note = notes.find((note) => note.id == id);

  if (note) {
    note.text = newNote;
    saveNotes(notes);
    console.log("Note successfully updated");
  } else {
    console.log("Note not found");
  }
}

module.exports = { addNote, listNotes, deleteNote, updateNote };
