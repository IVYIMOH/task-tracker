package stora

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

file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0644)
if err != nil {
	return nil, fmt.Errorf("fialed to open file: %w", err)

}
defer file.Close()

// check if any file is empty
