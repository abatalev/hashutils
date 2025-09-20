package hashutils

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar"
)

func WalkDirWithPatterns(workDir string, patters []string) []string {
	if _, err := os.Lstat(workDir); err != nil {
		return []string{}
	}

	if !strings.HasSuffix(workDir, "/") {
		workDir += "/"
	}
	files := make([]string, 0)
	err := filepath.Walk(workDir, func(path string, info fs.FileInfo, err error) error {
		// log.Println("@@@@", path, workDir)
		filename := strings.TrimPrefix(path, workDir)
		if info.IsDir() {
			return nil
		}
		if walkFunc(filename, patters) {
			files = append(files, filename)
		}
		return nil
	})
	if err != nil {
		// panic("!!!") // TODO remove panic
		log.Println("ERROR", err)
		return []string{}
	}
	return files
}

func walkFunc(filename string, patters []string) bool {
	for _, p := range patters {
		if x, _ := doublestar.Match(p, filename); x {
			return true
		}
	}
	return false
}
