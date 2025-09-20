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

func TestCalcHash(t *testing.T) {
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
	assertions := require.New(t)
	for n, variant := range variants {
		dirName := filepath.Join(t.TempDir(), "v"+strconv.Itoa(n))
		assertions.NoError(os.Mkdir(dirName, 0750))
		for _, f := range variant.files {
			fileName := filepath.Join(dirName, f.FileName)
			assertions.NoError(os.WriteFile(fileName, []byte(f.Content), 0600))
		}
		assertions.Equal(variant.result, CalcHashOfFiles(dirName, variant.resultFiles), n)
	}
}
