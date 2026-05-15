package hashutils

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar"
)

func WalkDirWithPatterns(workDir string, patterns []string) ([]string, error) {
	if _, err := os.Lstat(workDir); err != nil {
		return nil, err
	}

	if !strings.HasSuffix(workDir, "/") {
		workDir += "/"
	}
	files := make([]string, 0)
	err := filepath.Walk(workDir, func(path string, info fs.FileInfo, err error) error {
		filename := strings.TrimPrefix(path, workDir)
		if info.IsDir() {
			return nil
		}
		if walkFunc(filename, patterns) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

func walkFunc(filename string, patterns []string) bool {
	for _, p := range patterns {
		if x, _ := doublestar.Match(p, filename); x {
			return true
		}
	}
	return false
}
