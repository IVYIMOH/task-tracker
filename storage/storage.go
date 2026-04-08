package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"task-tracker/task"
)

const filename = "tasks.json"

type TaskStore struct {
	tasks  map[int]*task.Task
	nextID int
}

func LoadOrCreate() (*TaskStore, error) {
	store := &TaskStore{
		tasks:  make(map[int]*task.Task),
		nextID: 1,
	}

	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Check if file is empty
	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}

	if info.Size() == 0 {
		return store, nil
	}

	// Read and parse existing tasks
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var tasksList []*task.Task
	if err := json.Unmarshal(data, &tasksList); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Populate the map and find max ID
	maxID := 0
	for _, t := range tasksList {
		store.tasks[t.ID] = t
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	store.nextID = maxID + 1

	return store, nil
}

func (s *TaskStore) Save() error {
	// Convert map to slice for JSON serialization
	tasksList := make([]*task.Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		tasksList = append(tasksList, t)
	}

	data, err := json.MarshalIndent(tasksList, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func (s *TaskStore) Add(description string) (*task.Task, error) {
	newTask := task.NewTask(s.nextID, description)
	s.tasks[newTask.ID] = newTask
	s.nextID++
	
	if err := s.Save(); err != nil {
		return nil, err
	}
	
	return newTask, nil
}

func (s *TaskStore) Update(id int, description string) error {
	t, exists := s.tasks[id]
	if !exists {
		return fmt.Errorf("task with ID %d not found", id)
	}
	
	t.UpdateDescription(description)
	return s.Save()
}

func (s *TaskStore) Delete(id int) error {
	if _, exists := s.tasks[id]; !exists {
		return fmt.Errorf("task with ID %d not found", id)
	}
	
	delete(s.tasks, id)
	return s.Save()
}

func (s *TaskStore) MarkStatus(id int, status task.Status) error {
	t, exists := s.tasks[id]
	if !exists {
		return fmt.Errorf("task with ID %d not found", id)
	}
	
	t.UpdateStatus(status)
	return s.Save()
}

func (s *TaskStore) List(filterStatus *task.Status) []*task.Task {
	result := make([]*task.Task, 0, len(s.tasks))
	
	for _, t := range s.tasks {
		if filterStatus == nil || t.Status == *filterStatus {
			result = append(result, t)
		}
	}
	
	// Sort by ID (simple bubble sort since map iteration is unordered)
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i].ID > result[j].ID {
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	
	return result
}