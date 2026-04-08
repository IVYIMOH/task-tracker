# 📝 Task Tracker CLI (Go)

A simple and efficient **command-line task manager** built with Go.
Track your tasks, update progress, and stay organized — all from your terminal.

---

## 🚀 Features

* Add new tasks
* Update task descriptions
* Delete tasks
* Mark tasks as:

  * Todo ⭕
  * In Progress 🔄
  * Done ✅
* List tasks with optional filters
* Persistent storage using JSON (`tasks.json`)

---

## 📁 Project Structure

```
task-tracker/
├── main.go              # CLI entry point
├── go.mod              # Go module file
├── tasks.json          # Persistent task storage
├── task-cli            # Compiled binary (if built)
├── storage/
│   └── storage.go      # Task storage logic
└── task/
    └── task.go         # Task model and methods
```

---

## ⚙️ Installation & Setup

### 1. Clone the repository

```bash
git clone https://github.com/ivyimoh/task-tracker.git
cd task-tracker
```

### 2. Build the CLI

```bash
go build -o task-cli
```

### 3. Run the application

```bash
./task-cli
```

---

## 📌 Usage

### ➕ Add a Task

```bash
task-cli add "Buy groceries"
```

### ✏️ Update a Task

```bash
task-cli update 1 "Buy milk and eggs"
```

### ❌ Delete a Task

```bash
task-cli delete 1
```

### 🔄 Mark Task as In Progress

```bash
task-cli mark-in-progress 1
```

### ✅ Mark Task as Done

```bash
task-cli mark-done 1
```

### 📋 List Tasks

#### All Tasks

```bash
task-cli list
```

#### Completed Tasks

```bash
task-cli list done
```

#### Pending Tasks

```bash
task-cli list todo
```

#### In-Progress Tasks

```bash
task-cli list in-progress
```

---

## 🧠 How It Works

* Tasks are stored in a local `tasks.json` file.
* On startup, the app:

  * Loads existing tasks
  * Creates the file if it doesn’t exist
* Each task contains:

  * ID
  * Description
  * Status
  * Created timestamp
  * Updated timestamp
* Data is automatically saved after every operation.

---

## ⚡ Design Highlights

* **Modular architecture**

  * `task/` handles task structure and logic
  * `storage/` manages persistence
* **Simple CLI interface** using `os.Args`
* **Lightweight** — no external dependencies
* **Deterministic ordering** of tasks (sorted by ID)

---

## ⚠️ Limitations

* No concurrency handling (single-user CLI)
* Uses simple sorting (bubble sort) — fine for small datasets
* No advanced filtering or search (yet)

---

## 🔮 Future Improvements

* Add due dates and priorities
* Improve sorting (use Go’s `sort` package)
* Add search functionality
* Export/import tasks
* Interactive CLI (e.g., with prompts)
* Multi-user or cloud sync support

---

## 💡 Example Output

```
📋 All Tasks
-----------
⭕ [todo] ID: 1 | Buy groceries
🔄 [in-progress] ID: 2 | Write code
✅ [done] ID: 3 | Read a book
```

---

## 🧑‍💻 Author

**Ivy Imoh**
Building practical tools, one step at a time.

---

## 📜 License

This project is open-source and available under the MIT License.

---
PROJECT URL: https://roadmap.sh/projects/task-tracker
