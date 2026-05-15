package hashutils

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWalkDirWithPatterns(t *testing.T) {
	variants := []struct {
		files       []TestFileInfo
		patterns    []string
		resultFiles []string
	}{
		{
			files: []TestFileInfo{
				{FileName: "Dockerfile", Content: "FROM alpine:latest"},
			},
			patterns:    []string{"Docker*"},
			resultFiles: []string{"Dockerfile"},
		},
		{
			files: []TestFileInfo{
				{FileName: "foo.go", Content: "package foo"},
				{FileName: "bar.txt", Content: "hello"},
				{FileName: "baz.go", Content: "package baz"},
			},
			patterns:    []string{"*.go"},
			resultFiles: []string{"foo.go", "baz.go"},
		},
		{
			files: []TestFileInfo{
				{FileName: "a.go", Content: "package a"},
			},
			patterns:    []string{"*.rs"},
			resultFiles: []string{},
		},
	}
	for n, variant := range variants {
		dirName := filepath.Join(t.TempDir(), "v"+strconv.Itoa(n))
		require.NoError(t, os.Mkdir(dirName, 0750))
		for _, f := range variant.files {
			fileName := filepath.Join(dirName, f.FileName)
			require.NoError(t, os.WriteFile(fileName, []byte(f.Content), 0600))
		}
		files, err := WalkDirWithPatterns(dirName, variant.patterns)
		require.NoError(t, err, n)
		require.ElementsMatch(t, variant.resultFiles, files, n)
	}
}

func TestWalkDirWithPatternsNotExists(t *testing.T) {
	_, err := WalkDirWithPatterns("/nonexistent/path", []string{"*"})
	require.Error(t, err)
}

func TestWalkDirWithPatternsEmptyDir(t *testing.T) {
	dir := t.TempDir()
	files, err := WalkDirWithPatterns(dir, []string{"*"})
	require.NoError(t, err)
	require.Empty(t, files)
}
