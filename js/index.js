const { addNote, listNotes, deleteNote, updateNote } = require("./utils");

const [, , command, ...args] = process.argv;

if (command === "add") {
  addNote(args.join(" "));
} else if (command === "list") {
  listNotes();
} else if (command === "delete") {
  deleteNote(args);
} else if (command === "update") {
  const id = args[0];
  const newNote = args.slice(1).join(" ");
  updateNote(id, newNote);
} else {
  console.log("Usage: node index.js [add|list]");
}
