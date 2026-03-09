// store/tags.go
package store

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

type Store struct {
	path string
	mu   sync.RWMutex
	data map[string][]string // repo full name -> tags
}

func New(path string) *Store {
	s := &Store{
		path: path,
		data: make(map[string][]string),
	}
	s.load()
	return s
}

func DefaultPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "ghboard", "tags.json")
}

func (s *Store) Get(repo string) []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data[repo]
}

func (s *Store) Set(repo string, tags []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[repo] = tags
}

func (s *Store) Remove(repo string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, repo)
}

func (s *Store) Save() error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := os.MkdirAll(filepath.Dir(s.path), 0700); err != nil {
		return err
	}
	data, err := json.MarshalIndent(s.data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, data, 0600)
}

func (s *Store) load() {
	data, err := os.ReadFile(s.path)
	if err != nil {
		return
	}
	json.Unmarshal(data, &s.data)
}
