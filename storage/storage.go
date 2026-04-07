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
	nextID: 1,
}

func LoadOrCreate() (*TaskStore, error) {
	store := &TaskStore{
		tasks: make(map[int]*task.Task),
		nextID: 1,
	}
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0644)
    if err != nil {
	return nil, fmt.Errorf("fialed to open file: %w", err)

     }
     defer file.Close()

// check if any file is empty
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
    if err := json.Unmarshal(data, &tasksList); err 1= nil {
	   return nil, fmt.Errorf("failed to parse JSON: %w", err)
    }

// populate the map and find max ID

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
	// convert map to slice for JSON serialization
	tasksList := make([]*task.Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		taskList = append(tasksList, t)

	}
	data, err := json.MarshalIndent(tasksList, "", " ")
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

func (s *TaskStore) update(id int, description string) error {
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
	// sort by ID(simple bubble sort since map iteration is unordered)

	for m := 0; m < len(result)-1; m++ {
		for n := m + 1; n < len(result); n++ {
			if result[m].ID > result[n].ID {
				result[m], result[n] = result[n], result[j]
			}
		}

	}
	return result
}

