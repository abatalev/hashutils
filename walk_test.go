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
	}
	assertions := require.New(t)
	for n, variant := range variants {
		dirName := filepath.Join(t.TempDir(), "v"+strconv.Itoa(n))
		assertions.NoError(os.Mkdir(dirName, 0750))
		for _, f := range variant.files {
			fileName := filepath.Join(dirName, f.FileName)
			assertions.NoError(os.WriteFile(fileName, []byte(f.Content), 0600))
		}
		files := WalkDirWithPatterns(dirName, variant.patterns)
		assertions.ElementsMatch(variant.resultFiles, files, n)
	}
}
