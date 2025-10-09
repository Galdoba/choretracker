package js

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"

	"github.com/Galdoba/choretracker/internal/core/domain"
)

// AI generated comment:
// JsonStore represents a thread-safe storage for chores using JSON file as persistence
type JsonStore struct {
	filepath string
	Chores   map[int64]domain.Chore `json:"store"`
	mutex    sync.RWMutex
}

// AI generated comment:
// New creates a new JsonStore instance
// Initializes storage from existing file or creates new storage if file doesn't exist
// Uses path argument to locate or create the JSON storage file
func New(path string) (*JsonStore, error) {
	js := JsonStore{filepath: path, Chores: make(map[int64]domain.Chore)}
	switch fileExist(path) {
	case false:
		os.MkdirAll(filepath.Dir(path), 0755)
		f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0664)
		if err != nil {
			return nil, fmt.Errorf("failed to create JsonStore file: %v", err)
		}
		defer f.Close()
		emptyData, err := json.Marshal(&js)
		if err != nil {
			return nil, fmt.Errorf("JsonStore template creation failed: %v", err)
		}
		if _, err := f.Write(emptyData); err != nil {
			return nil, fmt.Errorf("failed to save JsonStore template: %v", err)
		}
		return &js, nil
	case true:
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to read JsonStore file: %v", err)
		}
		if err := json.Unmarshal(data, &js); err != nil {
			return nil, fmt.Errorf("failed to ummarshal JsonStore data: %v", err)
		}
	}
	return &js, nil

}

// AI generated comment:
// save writes the current state of JsonStore to the associated JSON file
func (js *JsonStore) save() error {
	data, err := json.MarshalIndent(js, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JsonStore: %v", err)
	}
	if err := os.WriteFile(js.filepath, data, 0644); err != nil {
		return fmt.Errorf("failed to save JsonStore file: %v", err)
	}
	return nil
}

// AI generated comment:
// load reads the JSON file and populates the JsonStore instance with its contents
func (js *JsonStore) load() error {
	data, err := os.ReadFile(js.filepath)
	if err != nil {
		return fmt.Errorf("failed to read JsonStore file: %v", err)
	}
	if err := json.Unmarshal(data, &js); err != nil {
		return fmt.Errorf("failed to unmarshal JsonStore: %v", err)
	}
	return nil
}

// AI generated comment:
// Create adds a new chore to the storage
// Accepts a chore.Chore pointer as argument and returns error if chore already exists
func (js *JsonStore) Create(ch domain.Chore) error {
	js.mutex.Lock()
	defer js.mutex.Unlock()
	if err := js.load(); err != nil {
		return fmt.Errorf("failed to load JsonStore: %v", err)
	}

	if _, ok := js.Chores[ch.ID]; ok {
		return fmt.Errorf("chore %v already created", ch.ID)
	}

	js.Chores[ch.ID] = ch

	if err := js.save(); err != nil {
		return fmt.Errorf("failed to save JsonStore: %v", err)
	}
	return nil
}

// AI generated comment:
// Update modifies an existing chore in the storage
// Accepts a chore.Chore pointer as argument and returns error if chore doesn't exist
func (js *JsonStore) Update(ch domain.Chore) error {
	js.mutex.Lock()
	defer js.mutex.Unlock()
	if err := js.load(); err != nil {
		return fmt.Errorf("failed to load JsonStore: %v", err)
	}

	if _, ok := js.Chores[ch.ID]; !ok {
		return fmt.Errorf("chore %v does not exist", ch.ID)
	}
	js.Chores[ch.ID] = ch

	if err := js.save(); err != nil {
		return fmt.Errorf("failed to save JsonStore: %v", err)
	}
	return nil
}

// AI generated comment:
// Read retrieves a chore from storage by its ID
// Accepts integer ID as argument and returns chore pointer or error if not found
func (js *JsonStore) Read(id int64) (domain.Chore, error) {
	js.mutex.Lock()
	defer js.mutex.Unlock()
	if err := js.load(); err != nil {
		return domain.Chore{}, fmt.Errorf("failed to load JsonStore: %v", err)
	}

	ch, ok := js.Chores[id]
	switch ok {
	default:
		return domain.Chore{}, fmt.Errorf("chore %v does not exist", id)
	case true:
		return ch, nil
	}
}

// AI generated comment:
// Delete removes a chore from storage by its ID
// Accepts integer ID as argument and returns error if chore doesn't exist
func (js *JsonStore) Delete(id int64) error {
	js.mutex.Lock()
	defer js.mutex.Unlock()
	if err := js.load(); err != nil {
		return fmt.Errorf("failed to load JsonStore: %v", err)
	}

	if _, ok := js.Chores[id]; !ok {
		return fmt.Errorf("chore %v does not exist", id)
	}

	delete(js.Chores, id)
	if err := js.save(); err != nil {
		return fmt.Errorf("failed to save JsonStore: %v", err)
	}
	return nil
}

// AI generated comment:
// GetAll returns all chores from the storage as a slice of chore pointers
func (js *JsonStore) GetAll() ([]domain.Chore, error) {
	js.mutex.Lock()
	defer js.mutex.Unlock()
	if err := js.load(); err != nil {
		return nil, fmt.Errorf("failed to load JsonStore: %v", err)
	}
	chores := []domain.Chore{}
	for _, v := range js.Chores {
		chores = append(chores, v)
	}
	sortChoresByID(chores)
	return chores, nil
}

func sortChoresByID(chores []domain.Chore) {
	sort.Slice(chores, func(i, j int) bool {
		return chores[i].ID < chores[j].ID
	})
}

// AI generated comment:
// fileExist checks if a file exists at the given path
// Accepts string path as argument and returns boolean indicating existence
func fileExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}
