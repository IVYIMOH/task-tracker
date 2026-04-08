package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"task-tracker/storage"
	"task-tracker/task"
)

func printHelp() {
	fmt.Println(`Task Tracker CLI - Manage your tasks

Usage:
  task-cli add <description>              Add a new task
  task-cli update <id> <description>      Update task description
  task-cli delete <id>                    Delete a task
  task-cli mark-in-progress <id>          Mark task as in progress
  task-cli mark-done <id>                 Mark task as done
  task-cli list                           List all tasks
  task-cli list done                      List completed tasks
  task-cli list todo                      List tasks not done
  task-cli list in-progress               List tasks in progress

Examples:
  task-cli add "Buy groceries"
  task-cli update 1 "Buy milk and eggs"
  task-cli mark-done 1
  task-cli list done`)
}

func printTasks(tasks []*task.Task, title string) {
	if title != "" {
		fmt.Printf("\n%s\n", title)
		fmt.Println(strings.Repeat("-", len(title)))
	}
	
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	
	for _, t := range tasks {
		statusIcon := "📋"
		switch t.Status {
		case task.Todo:
			statusIcon = "⭕"
		case task.InProgress:
			statusIcon = "🔄"
		case task.Done:
			statusIcon = "✅"
		}
		
		fmt.Printf("%s [%s] ID: %d | %s\n", 
			statusIcon, 
			t.Status, 
			t.ID, 
			t.Description)
	}
	fmt.Println()
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	// Load or create task storage
	store, err := storage.LoadOrCreate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing storage: %v\n", err)
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Error: Missing task description")
			fmt.Println("Usage: task-cli add <description>")
			os.Exit(1)
		}
		description := strings.Join(os.Args[2:], " ")
		newTask, err := store.Add(description)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error adding task: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)

	case "update":
		if len(os.Args) < 4 {
			fmt.Fprintln(os.Stderr, "Error: Missing ID or description")
			fmt.Println("Usage: task-cli update <id> <description>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid ID '%s'\n", os.Args[2])
			os.Exit(1)
		}
		description := strings.Join(os.Args[3:], " ")
		if err := store.Update(id, description); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Task %d updated successfully\n", id)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Error: Missing task ID")
			fmt.Println("Usage: task-cli delete <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid ID '%s'\n", os.Args[2])
			os.Exit(1)
		}
		if err := store.Delete(id); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Task %d deleted successfully\n", id)

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Error: Missing task ID")
			fmt.Println("Usage: task-cli mark-in-progress <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid ID '%s'\n", os.Args[2])
			os.Exit(1)
		}
		if err := store.MarkStatus(id, task.InProgress); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Task %d marked as in progress\n", id)

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Error: Missing task ID")
			fmt.Println("Usage: task-cli mark-done <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid ID '%s'\n", os.Args[2])
			os.Exit(1)
		}
		if err := store.MarkStatus(id, task.Done); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Task %d marked as done\n", id)

	case "list":
		var filterStatus *task.Status
		var title string
		
		if len(os.Args) >= 3 {
			switch os.Args[2] {
			case "done":
				status := task.Done
				filterStatus = &status
				title = "✅ Completed Tasks"
			case "todo":
				status := task.Todo
				filterStatus = &status
				title = "⭕ Pending Tasks"
			case "in-progress":
				status := task.InProgress
				filterStatus = &status
				title = "🔄 Tasks In Progress"
			default:
				fmt.Fprintf(os.Stderr, "Error: Unknown filter '%s'\n", os.Args[2])
				fmt.Println("Valid filters: done, todo, in-progress")
				os.Exit(1)
			}
		} else {
			title = "📋 All Tasks"
		}
		
		tasks := store.List(filterStatus)
		printTasks(tasks, title)

	default:
		fmt.Fprintf(os.Stderr, "Error: Unknown command '%s'\n", command)
		printHelp()
		os.Exit(1)
	}
}