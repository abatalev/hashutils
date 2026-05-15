package hashutils

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestFileInfo struct {
	FileName string
	Content  string
}

func TestCalcHashOfFiles(t *testing.T) {
	variants := []struct {
		files       []TestFileInfo
		resultFiles []string
		result      string
	}{
		{
			files: []TestFileInfo{
				{FileName: "Dockerfile", Content: "FROM alpine:latest\nCOPY *.go /opt/app/"},
				{FileName: "file1.go", Content: "aaa"},
				{FileName: "file2.go", Content: "bbb"},
			},
			resultFiles: []string{"Dockerfile", "file1.go", "file2.go"},
			result:      "3f9974ce06d2aa489d19adcc4c35f23bada79567",
		},
	}
	for n, variant := range variants {
		dirName := filepath.Join(t.TempDir(), "v"+strconv.Itoa(n))
		require.NoError(t, os.Mkdir(dirName, 0750))
		for _, f := range variant.files {
			fileName := filepath.Join(dirName, f.FileName)
			require.NoError(t, os.WriteFile(fileName, []byte(f.Content), 0600))
		}
		hash, err := CalcHashOfFiles(dirName, variant.resultFiles)
		require.NoError(t, err, n)
		require.Equal(t, variant.result, hash, n)
	}
}

func TestCalcHashOfFilesMissingFile(t *testing.T) {
	_, err := CalcHashOfFiles(t.TempDir(), []string{"nonexistent.go"})
	require.Error(t, err)
}

func TestCalcHashOfFilesEmpty(t *testing.T) {
	result, err := CalcHashOfFiles(t.TempDir(), []string{})
	require.NoError(t, err)
	require.Equal(t, "da39a3ee5e6b4b0d3255bfef95601890afd80709", result)
}
