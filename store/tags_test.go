// store/tags_test.go
package store

import (
	"path/filepath"
	"testing"
)

func TestAddAndGetTags(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tags.json")
	s := New(path)

	s.Set("owner/repo", []string{"tools", "ai"})
	tags := s.Get("owner/repo")
	if len(tags) != 2 || tags[0] != "tools" || tags[1] != "ai" {
		t.Errorf("unexpected tags: %v", tags)
	}
}

func TestPersistence(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tags.json")
	s1 := New(path)
	s1.Set("owner/repo", []string{"reference"})
	s1.Save()

	s2 := New(path)
	tags := s2.Get("owner/repo")
	if len(tags) != 1 || tags[0] != "reference" {
		t.Errorf("tags not persisted: %v", tags)
	}
}
