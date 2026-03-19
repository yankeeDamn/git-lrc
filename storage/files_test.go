package storage

import (
	"os"
	"path/filepath"
	"testing"
)

func TestChmodSecretFileMode(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "secret.txt")
	if err := os.WriteFile(path, []byte("secret"), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	if err := Chmod(path, 0600); err != nil {
		t.Fatalf("failed to chmod file: %v", err)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("failed to stat file: %v", err)
	}
	if got := info.Mode().Perm(); got != 0600 {
		t.Fatalf("unexpected mode: got %o want %o", got, os.FileMode(0600))
	}
}

func TestChmodExecutableMode(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "tool.sh")
	if err := os.WriteFile(path, []byte("#!/bin/sh\necho ok\n"), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	if err := Chmod(path, 0755); err != nil {
		t.Fatalf("failed to chmod file: %v", err)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("failed to stat file: %v", err)
	}
	if got := info.Mode().Perm(); got != 0755 {
		t.Fatalf("unexpected mode: got %o want %o", got, os.FileMode(0755))
	}
}

func TestWriteFileAtomicallyCreatesMissingParentDirectory(t *testing.T) {
	tmpDir := t.TempDir()
	targetDir := filepath.Join(tmpDir, "nested", "config")
	targetPath := filepath.Join(targetDir, ".lrc.toml")
	content := []byte("api_key = \"test\"\n")

	if err := WriteFileAtomically(targetPath, content, 0600); err != nil {
		t.Fatalf("write atomically with missing parent: %v", err)
	}

	got, err := os.ReadFile(targetPath)
	if err != nil {
		t.Fatalf("read target file: %v", err)
	}
	if string(got) != string(content) {
		t.Fatalf("unexpected file content: got %q want %q", string(got), string(content))
	}
}
