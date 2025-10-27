package task

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

const defaultDataFile = "tasks.csv"

// Storage manages task persistence
type Storage struct {
	filepath string
}

// NewStorage creates a new storage
func NewStorage(customPath string) *Storage {
	if customPath == "" {
		home, _ := os.UserHomeDir()
		customPath = filepath.Join(home, "Code", "go-projects", "ToDo-app", ".tasks", defaultDataFile)
	}

	// Create directory if it doesn't exist
	dir := filepath.Dir(customPath)
	os.MkdirAll(dir, os.ModePerm)

	return &Storage{filepath: customPath}
}

// loadFile opens the file with an exclusive lock
func (s *Storage) loadFile() (*os.File, error) {
	f, err := os.OpenFile(s.filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading: %w", err)
	}

	// Exclusive lock obtained on the file descriptor
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}

// closeFile unlocks and closes the file
func (s *Storage) closeFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}

// LoadTasks loads all tasks from the CSV file
func (s *Storage) LoadTasks() ([]*Task, error) {
	f, err := s.loadFile()
	if err != nil {
		return nil, err
	}
	defer s.closeFile(f)

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	var tasks []*Task

	// If the file is empty or only has a header
	if len(records) <= 1 {
		return tasks, nil
	}

	// Pular cabeçalho
	for _, record := range records[1:] {
		if len(record) != 4 {
			continue
		}

		id, _ := strconv.Atoi(record[0])
		createdAt, _ := time.Parse(time.RFC3339, record[2])
		isComplete, _ := strconv.ParseBool(record[3])

		task := &Task{
			ID:          id,
			Description: record[1],
			CreatedAt:   createdAt,
			IsComplete:  isComplete,
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// SaveTasks saves all tasks to the CSV file
func (s *Storage) SaveTasks(tasks []*Task) error {
	f, err := s.loadFile()
	if err != nil {
		return err
	}
	defer s.closeFile(f)

	// Truncate the file before writing
	if err := f.Truncate(0); err != nil {
		return err
	}
	if _, err := f.Seek(0, 0); err != nil {
		return err
	}

	writer := csv.NewWriter(f)
	defer writer.Flush()

	// Write header
	if err := writer.Write([]string{"ID", "Description", "CreatedAt", "IsComplete"}); err != nil {
		return err
	}

	// Write tasks
	for _, task := range tasks {
		record := []string{
			strconv.Itoa(task.ID),
			task.Description,
			task.CreatedAt.Format(time.RFC3339),
			strconv.FormatBool(task.IsComplete),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

// AddTask adds a new task
func (s *Storage) AddTask(description string) error {
	tasks, err := s.LoadTasks()
	if err != nil {
		return err
	}

	// Calculate next ID
	nextID := 1
	for _, task := range tasks {
		if task.ID >= nextID {
			nextID = task.ID + 1
		}
	}

	newTask := NewTask(nextID, description)
	tasks = append(tasks, newTask)

	return s.SaveTasks(tasks)
}

// CompleteTask sets a task as complete
func (s *Storage) CompleteTask(id int) error {
	tasks, err := s.LoadTasks()
	if err != nil {
		return err
	}

	found := false
	for _, task := range tasks {
		if task.ID == id {
			task.Complete()
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}

	return s.SaveTasks(tasks)
}

// DeleteTask removes a task
func (s *Storage) DeleteTask(id int) error {
	tasks, err := s.LoadTasks()
	if err != nil {
		return err
	}

	var filtered []*Task
	found := false
	for _, task := range tasks {
		if task.ID != id {
			filtered = append(filtered, task)
		} else {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}

	return s.SaveTasks(filtered)
}
